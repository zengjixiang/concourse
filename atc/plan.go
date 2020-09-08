package atc

import (
	"github.com/concourse/concourse/vars/interp"
)

//go:generate go run github.com/concourse/concourse/vars/interp/interpgen . interpolate.go

//interpgen:generate Params
//interpgen:generate Tags
//interpgen:generate TaskConfig export

type Plan struct {
	ID       PlanID `json:"id"`
	Attempts []int  `json:"attempts,omitempty"`

	Get         *GetPlan         `json:"get,omitempty"`
	Put         *PutPlan         `json:"put,omitempty"`
	Check       *CheckPlan       `json:"check,omitempty"`
	Task        *TaskPlan        `json:"task,omitempty"`
	SetPipeline *SetPipelinePlan `json:"set_pipeline,omitempty"`
	LoadVar     *LoadVarPlan     `json:"load_var,omitempty"`

	Do         *DoPlan         `json:"do,omitempty"`
	InParallel *InParallelPlan `json:"in_parallel,omitempty"`
	Aggregate  *AggregatePlan  `json:"aggregate,omitempty"`
	Across     *AcrossPlan     `json:"across,omitempty"`

	OnSuccess *OnSuccessPlan `json:"on_success,omitempty"`
	OnFailure *OnFailurePlan `json:"on_failure,omitempty"`
	OnAbort   *OnAbortPlan   `json:"on_abort,omitempty"`
	OnError   *OnErrorPlan   `json:"on_error,omitempty"`
	Ensure    *EnsurePlan    `json:"ensure,omitempty"`

	Try     *TryPlan     `json:"try,omitempty"`
	Timeout *TimeoutPlan `json:"timeout,omitempty"`
	Retry   *RetryPlan   `json:"retry,omitempty"`

	// used for 'fly execute'
	ArtifactInput  *ArtifactInputPlan  `json:"artifact_input,omitempty"`
	ArtifactOutput *ArtifactOutputPlan `json:"artifact_output,omitempty"`

	// deprecated, kept for backwards compatibility to be able to show old builds
	DependentGet *DependentGetPlan `json:"dependent_get,omitempty"`
}

func (plan *Plan) Each(f func(*Plan)) {
	f(plan)

	if plan.Do != nil {
		for i, p := range *plan.Do {
			p.Each(f)
			(*plan.Do)[i] = p
		}
	}

	if plan.InParallel != nil {
		for i, p := range plan.InParallel.Steps {
			p.Each(f)
			plan.InParallel.Steps[i] = p
		}
	}

	if plan.Aggregate != nil {
		for i, p := range *plan.Aggregate {
			p.Each(f)
			(*plan.Aggregate)[i] = p
		}
	}

	if plan.Across != nil {
		for i, p := range plan.Across.Steps {
			p.Step.Each(f)
			plan.Across.Steps[i] = p
		}
	}

	if plan.OnSuccess != nil {
		plan.OnSuccess.Step.Each(f)
		plan.OnSuccess.Next.Each(f)
	}

	if plan.OnFailure != nil {
		plan.OnFailure.Step.Each(f)
		plan.OnFailure.Next.Each(f)
	}

	if plan.OnAbort != nil {
		plan.OnAbort.Step.Each(f)
		plan.OnAbort.Next.Each(f)
	}

	if plan.OnError != nil {
		plan.OnError.Step.Each(f)
		plan.OnError.Next.Each(f)
	}

	if plan.Ensure != nil {
		plan.Ensure.Step.Each(f)
		plan.Ensure.Next.Each(f)
	}

	if plan.Try != nil {
		plan.Try.Step.Each(f)
	}

	if plan.Timeout != nil {
		plan.Timeout.Step.Each(f)
	}

	if plan.Retry != nil {
		for i, p := range *plan.Retry {
			p.Each(f)
			(*plan.Retry)[i] = p
		}
	}
}

type PlanID string

type ArtifactInputPlan struct {
	ArtifactID int    `json:"artifact_id"`
	Name       string `json:"name"`
}

type ArtifactOutputPlan struct {
	Name string `json:"name"`
}

type OnAbortPlan struct {
	Step Plan `json:"step"`
	Next Plan `json:"on_abort"`
}

type OnErrorPlan struct {
	Step Plan `json:"step"`
	Next Plan `json:"on_error"`
}

type OnFailurePlan struct {
	Step Plan `json:"step"`
	Next Plan `json:"on_failure"`
}

type EnsurePlan struct {
	Step Plan `json:"step"`
	Next Plan `json:"ensure"`
}

type OnSuccessPlan struct {
	Step Plan `json:"step"`
	Next Plan `json:"on_success"`
}

type TimeoutPlan struct {
	Step     Plan           `json:"step"`
	Duration interpDuration `json:"duration"`
}

type TryPlan struct {
	Step Plan `json:"step"`
}

type AggregatePlan []Plan

type InParallelPlan struct {
	Steps    []Plan      `json:"steps"`
	Limit    interp.Int  `json:"limit,omitempty"`
	FailFast interp.Bool `json:"fail_fast,omitempty"`
}

type AcrossPlan struct {
	Vars     []AcrossVar     `json:"vars"`
	Steps    []VarScopedPlan `json:"steps"`
	FailFast interp.Bool     `json:"fail_fast,omitempty"`
}

type AcrossVar struct {
	Var         string                  `json:"name"`
	Values      []interface{}           `json:"values"`
	MaxInFlight interpMaxInFlightConfig `json:"max_in_flight"`
}

type VarScopedPlan struct {
	Step   Plan          `json:"step"`
	Values []interface{} `json:"values"`
}

type DoPlan []Plan

type GetPlan struct {
	Name string `json:"name,omitempty"`

	Type        string       `json:"type"`
	Resource    string       `json:"resource"`
	Source      interpSource `json:"source"`
	Params      interpParams `json:"params,omitempty"`
	Version     *Version     `json:"version,omitempty"`
	VersionFrom *PlanID      `json:"version_from,omitempty"`
	Tags        interpTags   `json:"tags,omitempty"`

	VersionedResourceTypes VersionedResourceTypes `json:"resource_types,omitempty"`
}

type PutPlan struct {
	Type     string              `json:"type"`
	Name     string              `json:"name,omitempty"`
	Resource string              `json:"resource"`
	Source   interpSource        `json:"source"`
	Params   interpParams        `json:"params,omitempty"`
	Tags     interpTags          `json:"tags,omitempty"`
	Inputs   *interpInputsConfig `json:"inputs,omitempty"`

	VersionedResourceTypes VersionedResourceTypes `json:"resource_types,omitempty"`
}

type CheckPlan struct {
	Type        string       `json:"type"`
	Name        string       `json:"name,omitempty"`
	Source      interpSource `json:"source"`
	Tags        interpTags   `json:"tags,omitempty"`
	Timeout     string       `json:"timeout,omitempty"`
	FromVersion Version      `json:"from_version,omitempty"`

	VersionedResourceTypes VersionedResourceTypes `json:"resource_types,omitempty"`
}

type TaskPlan struct {
	Name string `json:"name,omitempty"`

	Privileged interp.Bool `json:"privileged"`
	Tags       interpTags  `json:"tags,omitempty"`

	ConfigPath interp.String     `json:"config_path,omitempty"`
	Config     *InterpTaskConfig `json:"config,omitempty"`
	Vars       interpParams      `json:"vars,omitempty"`

	Params            interpParams      `json:"params,omitempty"`
	InputMapping      interpFileMapping `json:"input_mapping,omitempty"`
	OutputMapping     interpFileMapping `json:"output_mapping,omitempty"`
	ImageArtifactName interp.String     `json:"image,omitempty"`

	VersionedResourceTypes VersionedResourceTypes `json:"resource_types,omitempty"`
}

//interpgen:generate FileMapping

type FileMapping map[interp.String]interp.String

type SetPipelinePlan struct {
	Name     interp.String    `json:"name"`
	File     interp.String    `json:"file"`
	Team     interp.String    `json:"team,omitempty"`
	Vars     interpParams     `json:"vars,omitempty"`
	VarFiles interpStringList `json:"var_files,omitempty"`
}

//interpgen:generate stringList

type stringList []interp.String

type LoadVarPlan struct {
	Name   string        `json:"name"`
	File   interp.String `json:"file"`
	Format interp.String `json:"format,omitempty"`
	Reveal interp.Bool   `json:"reveal,omitempty"`
}

type RetryPlan []Plan

type DependentGetPlan struct {
	Type     string `json:"type"`
	Name     string `json:"name,omitempty"`
	Resource string `json:"resource"`
}
