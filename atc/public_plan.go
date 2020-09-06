package atc

import (
	"encoding/json"

	"github.com/aoldershaw/interpolate"
)

func (plan InterpPlan) Public() *json.RawMessage {
	var public struct {
		ID PlanID `json:"id"`

		Aggregate      *json.RawMessage `json:"aggregate,omitempty"`
		InParallel     *json.RawMessage `json:"in_parallel,omitempty"`
		Across         *json.RawMessage `json:"across,omitempty"`
		Do             *json.RawMessage `json:"do,omitempty"`
		Get            *json.RawMessage `json:"get,omitempty"`
		Put            *json.RawMessage `json:"put,omitempty"`
		Check          *json.RawMessage `json:"check,omitempty"`
		Task           *json.RawMessage `json:"task,omitempty"`
		SetPipeline    *json.RawMessage `json:"set_pipeline,omitempty"`
		LoadVar        *json.RawMessage `json:"load_var,omitempty"`
		OnAbort        *json.RawMessage `json:"on_abort,omitempty"`
		OnError        *json.RawMessage `json:"on_error,omitempty"`
		Ensure         *json.RawMessage `json:"ensure,omitempty"`
		OnSuccess      *json.RawMessage `json:"on_success,omitempty"`
		OnFailure      *json.RawMessage `json:"on_failure,omitempty"`
		Try            *json.RawMessage `json:"try,omitempty"`
		DependentGet   *json.RawMessage `json:"dependent_get,omitempty"`
		Timeout        *json.RawMessage `json:"timeout,omitempty"`
		Retry          *json.RawMessage `json:"retry,omitempty"`
		ArtifactInput  *json.RawMessage `json:"artifact_input,omitempty"`
		ArtifactOutput *json.RawMessage `json:"artifact_output,omitempty"`
	}

	public.ID = plan.ID

	if plan.Aggregate != nil {
		public.Aggregate = plan.Aggregate.Public()
	}

	if plan.InParallel != nil {
		public.InParallel = plan.InParallel.Public()
	}

	if plan.Across != nil {
		public.Across = plan.Across.Public()
	}

	if plan.Do != nil {
		public.Do = plan.Do.Public()
	}

	if plan.Get != nil {
		public.Get = plan.Get.Public()
	}

	if plan.Put != nil {
		public.Put = plan.Put.Public()
	}

	if plan.Check != nil {
		public.Check = plan.Check.Public()
	}

	if plan.Task != nil {
		public.Task = plan.Task.Public()
	}

	if plan.SetPipeline != nil {
		public.SetPipeline = plan.SetPipeline.Public()
	}

	if plan.LoadVar != nil {
		public.LoadVar = plan.LoadVar.Public()
	}

	if plan.OnAbort != nil {
		public.OnAbort = plan.OnAbort.Public()
	}

	if plan.OnError != nil {
		public.OnError = plan.OnError.Public()
	}

	if plan.Ensure != nil {
		public.Ensure = plan.Ensure.Public()
	}

	if plan.OnSuccess != nil {
		public.OnSuccess = plan.OnSuccess.Public()
	}

	if plan.OnFailure != nil {
		public.OnFailure = plan.OnFailure.Public()
	}

	if plan.Try != nil {
		public.Try = plan.Try.Public()
	}

	if plan.Timeout != nil {
		public.Timeout = plan.Timeout.Public()
	}

	if plan.Retry != nil {
		public.Retry = plan.Retry.Public()
	}

	if plan.ArtifactInput != nil {
		public.ArtifactInput = plan.ArtifactInput.Public()
	}

	if plan.ArtifactOutput != nil {
		public.ArtifactOutput = plan.ArtifactOutput.Public()
	}

	if plan.DependentGet != nil {
		public.DependentGet = plan.DependentGet.Public()
	}

	return enc(public)
}

func (plan InterpAggregatePlan) Public() *json.RawMessage {
	public := make([]*json.RawMessage, len(plan))

	for i := 0; i < len(plan); i++ {
		public[i] = plan[i].Public()
	}

	return enc(public)
}

func (plan InterpInParallelPlan) Public() *json.RawMessage {
	steps := make([]*json.RawMessage, len(plan.Steps))

	for i := 0; i < len(plan.Steps); i++ {
		steps[i] = plan.Steps[i].Public()
	}

	return enc(struct {
		Steps    []*json.RawMessage `json:"steps"`
		Limit    interpInt          `json:"limit,omitempty"`
		FailFast interpBool         `json:"fail_fast,omitempty"`
	}{
		Steps:    steps,
		Limit:    plan.Limit,
		FailFast: plan.FailFast,
	})
}

func (plan InterpAcrossPlan) Public() *json.RawMessage {
	type scopedStep struct {
		Step   *json.RawMessage `json:"step"`
		Values []interface{}    `json:"values"`
	}

	steps := []scopedStep{}
	for _, step := range plan.Steps {
		steps = append(steps, scopedStep{
			Step:   step.Step.Public(),
			Values: step.Values,
		})
	}

	return enc(struct {
		Vars     []interpAcrossVar `json:"vars"`
		Steps    []scopedStep      `json:"steps"`
		FailFast interpBool        `json:"fail_fast,omitempty"`
	}{
		Vars:     plan.Vars,
		Steps:    steps,
		FailFast: plan.FailFast,
	})
}

func (plan InterpDoPlan) Public() *json.RawMessage {
	public := make([]*json.RawMessage, len(plan))

	for i := 0; i < len(plan); i++ {
		public[i] = plan[i].Public()
	}

	return enc(public)
}

func (plan InterpEnsurePlan) Public() *json.RawMessage {
	return enc(struct {
		Step *json.RawMessage `json:"step"`
		Next *json.RawMessage `json:"ensure"`
	}{
		Step: plan.Step.Public(),
		Next: plan.Next.Public(),
	})
}

func (plan InterpGetPlan) Public() *json.RawMessage {
	return enc(struct {
		Type     string   `json:"type"`
		Name     string   `json:"name,omitempty"`
		Resource string   `json:"resource"`
		Version  *Version `json:"version,omitempty"`
	}{
		Type:     plan.Type,
		Name:     plan.Name,
		Resource: plan.Resource,
		Version:  plan.Version,
	})
}

func (plan DependentGetPlan) Public() *json.RawMessage {
	return enc(struct {
		Type     string `json:"type"`
		Name     string `json:"name,omitempty"`
		Resource string `json:"resource"`
	}{
		Type:     plan.Type,
		Name:     plan.Name,
		Resource: plan.Resource,
	})
}

func (plan InterpOnAbortPlan) Public() *json.RawMessage {
	return enc(struct {
		Step *json.RawMessage `json:"step"`
		Next *json.RawMessage `json:"on_abort"`
	}{
		Step: plan.Step.Public(),
		Next: plan.Next.Public(),
	})
}

func (plan InterpOnErrorPlan) Public() *json.RawMessage {
	return enc(struct {
		Step *json.RawMessage `json:"step"`
		Next *json.RawMessage `json:"on_error"`
	}{
		Step: plan.Step.Public(),
		Next: plan.Next.Public(),
	})
}

func (plan InterpOnFailurePlan) Public() *json.RawMessage {
	return enc(struct {
		Step *json.RawMessage `json:"step"`
		Next *json.RawMessage `json:"on_failure"`
	}{
		Step: plan.Step.Public(),
		Next: plan.Next.Public(),
	})
}

func (plan InterpOnSuccessPlan) Public() *json.RawMessage {
	return enc(struct {
		Step *json.RawMessage `json:"step"`
		Next *json.RawMessage `json:"on_success"`
	}{
		Step: plan.Step.Public(),
		Next: plan.Next.Public(),
	})
}

func (plan InterpPutPlan) Public() *json.RawMessage {
	return enc(struct {
		Type     string `json:"type"`
		Name     string `json:"name,omitempty"`
		Resource string `json:"resource"`
	}{
		Type:     plan.Type,
		Name:     plan.Name,
		Resource: plan.Resource,
	})
}

func (plan InterpCheckPlan) Public() *json.RawMessage {
	return enc(struct {
		Type string `json:"type"`
		Name string `json:"name,omitempty"`
	}{
		Type: plan.Type,
		Name: plan.Name,
	})
}

func (plan InterpTaskPlan) Public() *json.RawMessage {
	return enc(struct {
		Name       interpolate.String `json:"name"`
		Privileged interpBool         `json:"privileged"`
	}{
		Name:       plan.Name,
		Privileged: plan.Privileged,
	})
}

func (plan InterpSetPipelinePlan) Public() *json.RawMessage {
	return enc(struct {
		Name interpolate.String `json:"name"`
		Team interpolate.String `json:"team"`
	}{
		Name: plan.Name,
		Team: plan.Team,
	})
}

func (plan InterpLoadVarPlan) Public() *json.RawMessage {
	return enc(struct {
		Name interpolate.String `json:"name"`
	}{
		Name: plan.Name,
	})
}

func (plan InterpTimeoutPlan) Public() *json.RawMessage {
	return enc(struct {
		Step     *json.RawMessage   `json:"step"`
		Duration interpolate.String `json:"duration"`
	}{
		Step:     plan.Step.Public(),
		Duration: plan.Duration,
	})
}

func (plan InterpTryPlan) Public() *json.RawMessage {
	return enc(struct {
		Step *json.RawMessage `json:"step"`
	}{
		Step: plan.Step.Public(),
	})
}

func (plan InterpRetryPlan) Public() *json.RawMessage {
	public := make([]*json.RawMessage, len(plan))

	for i := 0; i < len(plan); i++ {
		public[i] = plan[i].Public()
	}

	return enc(public)
}

func (plan ArtifactInputPlan) Public() *json.RawMessage {
	return enc(plan)
}

func (plan ArtifactOutputPlan) Public() *json.RawMessage {
	return enc(plan)
}

func enc(public interface{}) *json.RawMessage {
	enc, _ := json.Marshal(public)
	return (*json.RawMessage)(&enc)
}
