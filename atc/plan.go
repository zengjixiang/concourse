package atc

//go:generate interpolate . interp.go

//interpolate:generate Plan

type Plan struct {
	ID       PlanID `json:"id"`
	Attempts []int  `json:"attempts,omitempty"`

	Get         *GetPlan         `json:"get,omitempty" interpolate:"subtypes,export"`
	Put         *PutPlan         `json:"put,omitempty" interpolate:"subtypes,export"`
	Check       *CheckPlan       `json:"check,omitempty" interpolate:"subtypes,export"`
	Task        *TaskPlan        `json:"task,omitempty" interpolate:"subtypes,export"`
	SetPipeline *SetPipelinePlan `json:"set_pipeline,omitempty" interpolate:"subtypes,export"`
	LoadVar     *LoadVarPlan     `json:"load_var,omitempty" interpolate:"subtypes,export"`

	Do         *DoPlan         `json:"do,omitempty" interpolate:"subtypes,values=subtypes,export"`
	InParallel *InParallelPlan `json:"in_parallel,omitempty" interpolate:"subtypes,export"`
	Aggregate  *AggregatePlan  `json:"aggregate,omitempty" interpolate:"subtypes,values=subtypes,export"`
	Across     *AcrossPlan     `json:"across,omitempty" interpolate:"subtypes,export"`

	OnSuccess *OnSuccessPlan `json:"on_success,omitempty" interpolate:"subtypes,export"`
	OnFailure *OnFailurePlan `json:"on_failure,omitempty" interpolate:"subtypes,export"`
	OnAbort   *OnAbortPlan   `json:"on_abort,omitempty" interpolate:"subtypes,export"`
	OnError   *OnErrorPlan   `json:"on_error,omitempty" interpolate:"subtypes,export"`
	Ensure    *EnsurePlan    `json:"ensure,omitempty" interpolate:"subtypes,export"`

	Try     *TryPlan     `json:"try,omitempty" interpolate:"subtypes,export"`
	Timeout *TimeoutPlan `json:"timeout,omitempty" interpolate:"subtypes,export"`
	Retry   *RetryPlan   `json:"retry,omitempty" interpolate:"subtypes,values=subtypes,export"`

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
	Step Plan `json:"step" interpolate:"subtypes"`
	Next Plan `json:"on_abort" interpolate:"subtypes"`
}

type OnErrorPlan struct {
	Step Plan `json:"step" interpolate:"subtypes"`
	Next Plan `json:"on_error" interpolate:"subtypes"`
}

type OnFailurePlan struct {
	Step Plan `json:"step" interpolate:"subtypes"`
	Next Plan `json:"on_failure" interpolate:"subtypes"`
}

type EnsurePlan struct {
	Step Plan `json:"step" interpolate:"subtypes"`
	Next Plan `json:"ensure" interpolate:"subtypes"`
}

type OnSuccessPlan struct {
	Step Plan `json:"step" interpolate:"subtypes"`
	Next Plan `json:"on_success" interpolate:"subtypes"`
}

type TimeoutPlan struct {
	Step     Plan   `json:"step" interpolate:"subtypes"`
	Duration string `json:"duration" interpolate:"root"`
}

type TryPlan struct {
	Step Plan `json:"step" interpolate:"subtypes"`
}

type AggregatePlan []Plan

type InParallelPlan struct {
	Steps    []Plan `json:"steps" interpolate:"subtypes,values=subtypes"`
	Limit    int    `json:"limit,omitempty" interpolate:"root"`
	FailFast bool   `json:"fail_fast,omitempty" interpolate:"root"`
}

type AcrossPlan struct {
	Vars     []AcrossVar     `json:"vars" interpolate:"subtypes,values=subtypes"`
	Steps    []VarScopedPlan `json:"steps" interpolate:"subtypes,values=subtypes"`
	FailFast bool            `json:"fail_fast,omitempty" interpolate:"root"`
}

type AcrossVar struct {
	Var         string        `json:"name"`
	Values      []interface{} `json:"values"`
	MaxInFlight int           `json:"max_in_flight" interpolate:"root"`
}

type VarScopedPlan struct {
	Step   Plan          `json:"step" interpolate:"subtypes"`
	Values []interface{} `json:"values"`
}

type DoPlan []Plan

type GetPlan struct {
	Name string `json:"name,omitempty"`

	Type        string   `json:"type"`
	Resource    string   `json:"resource"`
	Source      Source   `json:"source"`
	Params      Params   `json:"params,omitempty" interpolate:"root,keys=root,values=root"`
	Version     *Version `json:"version,omitempty"`
	VersionFrom *PlanID  `json:"version_from,omitempty"`
	Tags        Tags     `json:"tags,omitempty" interpolate:"root,values=root"`

	VersionedResourceTypes VersionedResourceTypes `json:"resource_types,omitempty"`
}

type PutPlan struct {
	Type     string        `json:"type"`
	Name     string        `json:"name,omitempty"`
	Resource string        `json:"resource"`
	Source   Source        `json:"source"`
	Params   Params        `json:"params,omitempty" interpolate:"root,keys=root,values=root"`
	Tags     Tags          `json:"tags,omitempty" interpolate:"root,values=root"`
	Inputs   *InputsConfig `json:"inputs,omitempty" interpolate:"root"`

	VersionedResourceTypes VersionedResourceTypes `json:"resource_types,omitempty"`
}

type CheckPlan struct {
	Type        string  `json:"type"`
	Name        string  `json:"name,omitempty"`
	Source      Source  `json:"source"`
	Tags        Tags    `json:"tags,omitempty"`
	Timeout     string  `json:"timeout,omitempty"`
	FromVersion Version `json:"from_version,omitempty"`

	VersionedResourceTypes VersionedResourceTypes `json:"resource_types,omitempty"`
}

type TaskPlan struct {
	Name string `json:"name,omitempty" interpolate:"root"`

	Privileged bool `json:"privileged" interpolate:"root"`
	Tags       Tags `json:"tags,omitempty" interpolate:"root,values=root"`

	ConfigPath string      `json:"config_path,omitempty" interpolate:"root"`
	Config     *TaskConfig `json:"config,omitempty" interpolate:"root"`
	Vars       Params      `json:"vars,omitempty" interpolate:"root,keys=root,values=root"`

	Params            Params            `json:"params,omitempty" interpolate:"root,keys=root,values=root"`
	InputMapping      map[string]string `json:"input_mapping,omitempty" interpolate:"root,keys=root,values=root"`
	OutputMapping     map[string]string `json:"output_mapping,omitempty" interpolate:"root,keys=root,values=root"`
	ImageArtifactName string            `json:"image,omitempty" interpolate:"root"`

	VersionedResourceTypes VersionedResourceTypes `json:"resource_types,omitempty"`
}

type SetPipelinePlan struct {
	Name     string                 `json:"name" interpolate:"root"`
	File     string                 `json:"file" interpolate:"root"`
	Team     string                 `json:"team,omitempty" interpolate:"root"`
	Vars     map[string]interface{} `json:"vars,omitempty" interpolate:"root,keys=root,values=root"`
	VarFiles []string               `json:"var_files,omitempty" interpolate:"root,values=root"`
}

type LoadVarPlan struct {
	Name   string `json:"name" interpolate:"root"`
	File   string `json:"file" interpolate:"root"`
	Format string `json:"format,omitempty" interpolate:"root"`
	Reveal bool   `json:"reveal,omitempty" interpolate:"root"`
}

type RetryPlan []Plan

type DependentGetPlan struct {
	Type     string `json:"type"`
	Name     string `json:"name,omitempty"`
	Resource string `json:"resource"`
}
