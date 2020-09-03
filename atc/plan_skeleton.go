package atc

import "github.com/concourse/concourse/vars/interp"

type PlanSkeleton struct {
	ID       PlanID `json:"id"`
	Attempts []int  `json:"attempts,omitempty"`

	Get         *GetPlanSkeleton         `json:"get,omitempty"`
	Put         *PutPlanSkeleton         `json:"put,omitempty"`
	Check       *CheckPlanSkeleton       `json:"check,omitempty"`
	Task        *TaskPlanSkeleton        `json:"task,omitempty"`
	SetPipeline *SetPipelinePlanSkeleton `json:"set_pipeline,omitempty"`
	LoadVar     *LoadVarPlanSkeleton     `json:"load_var,omitempty"`

	Do         *DoPlanSkeleton         `json:"do,omitempty"`
	InParallel *InParallelPlanSkeleton `json:"in_parallel,omitempty"`
	Aggregate  *AggregatePlanSkeleton  `json:"aggregate,omitempty"`
	Across     *AcrossPlanSkeleton     `json:"across,omitempty"`

	OnSuccess *OnSuccessPlanSkeleton `json:"on_success,omitempty"`
	OnFailure *OnFailurePlanSkeleton `json:"on_failure,omitempty"`
	OnAbort   *OnAbortPlanSkeleton   `json:"on_abort,omitempty"`
	OnError   *OnErrorPlanSkeleton   `json:"on_error,omitempty"`
	Ensure    *EnsurePlanSkeleton    `json:"ensure,omitempty"`

	Try     *TryPlanSkeleton     `json:"try,omitempty"`
	Timeout *TimeoutPlanSkeleton `json:"timeout,omitempty"`
	Retry   *RetryPlanSkeleton   `json:"retry,omitempty"`

	// used for 'fly execute'
	ArtifactInput  *ArtifactInputPlanSkeleton  `json:"artifact_input,omitempty"`
	ArtifactOutput *ArtifactOutputPlanSkeleton `json:"artifact_output,omitempty"`

	// deprecated, kept for backwards compatibility to be able to show old builds
	DependentGet *DependentGetPlanSkeleton `json:"dependent_get,omitempty"`
}

func (plan *PlanSkeleton) Each(f func(*PlanSkeleton)) {
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

type ArtifactInputPlanSkeleton struct {
	ArtifactID int    `json:"artifact_id"`
	Name       string `json:"name"`
}

type ArtifactOutputPlanSkeleton struct {
	Name string `json:"name"`
}

type OnAbortPlanSkeleton struct {
	Step PlanSkeleton `json:"step"`
	Next PlanSkeleton `json:"on_abort"`
}

type OnErrorPlanSkeleton struct {
	Step PlanSkeleton `json:"step"`
	Next PlanSkeleton `json:"on_error"`
}

type OnFailurePlanSkeleton struct {
	Step PlanSkeleton `json:"step"`
	Next PlanSkeleton `json:"on_failure"`
}

type EnsurePlanSkeleton struct {
	Step PlanSkeleton `json:"step"`
	Next PlanSkeleton `json:"ensure"`
}

type OnSuccessPlanSkeleton struct {
	Step PlanSkeleton `json:"step"`
	Next PlanSkeleton `json:"on_success"`
}

type TimeoutPlanSkeleton struct {
	Step     PlanSkeleton   `json:"step"`
	Duration interpDuration `json:"duration"`
}

type TryPlanSkeleton struct {
	Step PlanSkeleton `json:"step"`
}

type AggregatePlanSkeleton []PlanSkeleton

type InParallelPlanSkeleton struct {
	Steps    []PlanSkeleton `json:"steps"`
	Limit    interp.Int     `json:"limit,omitempty"`
	FailFast interp.Bool    `json:"fail_fast,omitempty"`
}

type AcrossVarSkeleton struct {
	Var    string        `json:"name"`
	Values []interface{} `json:"values"`
	// TODO: make a MaxInFlightConfig
	MaxInFlight interp.Int `json:"max_in_flight"`
}

type AcrossPlanSkeleton struct {
	Vars     []AcrossPlanSkeleton    `json:"vars"`
	Steps    []VarScopedPlanSkeleton `json:"steps"`
	FailFast interp.Bool             `json:"fail_fast,omitempty"`
}

type VarScopedPlanSkeleton struct {
	Step   PlanSkeleton  `json:"step"`
	Values []interface{} `json:"values"`
}

type DoPlanSkeleton []PlanSkeleton

type GetPlanSkeleton struct {
	Name string `json:"name,omitempty"`

	Type        string             `json:"type"`
	Resource    string             `json:"resource"`
	Source      Source             `json:"source"`
	Params      interpInterpParams `json:"params,omitempty"`
	Version     *Version           `json:"version,omitempty"`
	VersionFrom *PlanID            `json:"version_from,omitempty"`
	Tags        interpInterpTags   `json:"tags,omitempty"`

	VersionedResourceTypes VersionedResourceTypes `json:"resource_types,omitempty"`
}

type PutPlanSkeleton struct {
	Type     string              `json:"type"`
	Name     string              `json:"name,omitempty"`
	Resource string              `json:"resource"`
	Source   Source              `json:"source"`
	Params   interpInterpParams  `json:"params,omitempty"`
	Tags     interpInterpTags    `json:"tags,omitempty"`
	Inputs   *interpInputsConfig `json:"inputs,omitempty"`

	VersionedResourceTypes VersionedResourceTypes `json:"resource_types,omitempty"`
}

type CheckPlanSkeleton struct {
	Type        string           `json:"type"`
	Name        string           `json:"name,omitempty"`
	Source      Source           `json:"source"`
	Tags        interpInterpTags `json:"tags,omitempty"`
	Timeout     string           `json:"timeout,omitempty"`
	FromVersion Version          `json:"from_version,omitempty"`

	VersionedResourceTypes VersionedResourceTypes `json:"resource_types,omitempty"`
}

type TaskPlanSkeleton struct {
	Name string `json:"name,omitempty"`

	Privileged interp.Bool      `json:"privileged"`
	Tags       interpInterpTags `json:"tags,omitempty"`

	ConfigPath interp.String          `json:"config_path,omitempty"`
	Config     interpInterpTaskConfig `json:"config,omitempty"`
	Vars       interpInterpParams     `json:"vars,omitempty"`

	Params            interpInterpParams          `json:"params,omitempty"`
	InputMapping      interpInterpArtifactMapping `json:"input_mapping,omitempty"`
	OutputMapping     interpInterpArtifactMapping `json:"output_mapping,omitempty"`
	ImageArtifactName interp.String               `json:"image,omitempty"`

	VersionedResourceTypes VersionedResourceTypes `json:"resource_types,omitempty"`
}

type SetPipelinePlanSkeleton struct {
	Name     interp.String          `json:"name"`
	File     interp.String          `json:"file"`
	Team     interp.String          `json:"team,omitempty"`
	Vars     interpInterpParams     `json:"vars,omitempty"`
	VarFiles interpInterpStringList `json:"var_files,omitempty"`
}

type LoadVarPlanSkeleton struct {
	Name   interp.String `json:"name"`
	File   interp.String `json:"file"`
	Format interp.String `json:"format,omitempty"`
	Reveal interp.Bool   `json:"reveal,omitempty"`
}

type RetryPlanSkeleton []PlanSkeleton

type DependentGetPlanSkeleton struct {
	Type     interp.String `json:"type"`
	Name     interp.String `json:"name,omitempty"`
	Resource interp.String `json:"resource"`
}
