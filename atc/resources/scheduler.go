package resources

import (
	"context"
	"fmt"
	"time"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
)

type SchedulerResource interface {
	StepConfig() atc.StepConfig
	Resources() db.SchedulerResources
	ResourceTypes() atc.VersionedResourceTypes

	RefreshResourceConfig() (db.ResourceConfigScope, error)
}

type Factory interface {
	ResourcesToSchedule(time.Duration, time.Duration) ([]SchedulerResource, error)
}

type Planner interface {
	Create(
		atc.StepConfig,
		db.SchedulerResources,
		atc.VersionedResourceTypes,
		[]db.BuildInput,
	) (atc.Plan, error)
}

type BuildScheduler struct {
	Factory Factory
	Planner Planner

	DefaultCheckInterval        time.Duration
	DefaultWebhookCheckInterval time.Duration
}

func (scheduler BuildScheduler) Run(context.Context) error {
	resources, err := scheduler.Factory.ResourcesToSchedule(
		scheduler.DefaultCheckInterval,
		scheduler.DefaultWebhookCheckInterval,
	)
	if err != nil {
		return fmt.Errorf("get all resources: %w", err)
	}

	// keep track of scope IDs which have already been checked
	alreadyChecked := make(map[int]bool)

	for _, resource := range resources {
		// evaluate creds to come up with a resource config and scope
		//
		// if resource config scope is different from current one, AND has been
		// checked, update resource_config_id and resource_config_scope_id
		//
		// if value changes, request scheduling for downstream jobs
		scope, err := resource.RefreshResourceConfig()
		if err != nil {
			// XXX: expect to get here if an ancestor type has no version?
			return fmt.Errorf("update resource config: %w", err)
		}

		// make sure we haven't already queued a check for the same scope
		if alreadyChecked[scope.ID()] {
			continue
		} else {
			alreadyChecked[scope.ID()] = true
		}

		// create a build plan with unevaluated creds
		//
		// check step will evaluate creds to come up with its own resource config
		// and scope
		//
		// check step will save versions to scope at runtime
		//
		// check step will update associated resource or resource type's
		// resource_config_id and resource_config_scope_id
		plan, err := scheduler.Planner.Create(
			resource.StepConfig(),
			resource.Resources(),     // NOTE: unevaluated creds
			resource.ResourceTypes(), // NOTE: unevaluated creds
			[]db.BuildInput{},
		)
		if err != nil {
			// XXX: probably shouldn't bail on the whole thing
			return fmt.Errorf("create build plan: %w", err)
		}
	}

	return nil
}
