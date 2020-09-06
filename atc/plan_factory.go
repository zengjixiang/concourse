package atc

import (
	"encoding/json"
	"fmt"
	"sync/atomic"
)

type PlanFactory struct {
	currentNum *int64
}

func NewPlanFactory(startingNum int64) PlanFactory {
	return PlanFactory{
		currentNum: &startingNum,
	}
}

type PlanConfig interface {
	Public() *json.RawMessage
}

func (factory PlanFactory) NewPlan(step PlanConfig) InterpPlan {
	num := atomic.AddInt64(factory.currentNum, 1)

	var plan InterpPlan
	switch t := step.(type) {
	case InterpAggregatePlan:
		plan.Aggregate = &t
	case InterpInParallelPlan:
		plan.InParallel = &t
	case InterpAcrossPlan:
		plan.Across = &t
	case InterpDoPlan:
		plan.Do = &t
	case InterpGetPlan:
		plan.Get = &t
	case InterpPutPlan:
		plan.Put = &t
	case InterpTaskPlan:
		plan.Task = &t
	case InterpSetPipelinePlan:
		plan.SetPipeline = &t
	case InterpLoadVarPlan:
		plan.LoadVar = &t
	case InterpCheckPlan:
		plan.Check = &t
	case InterpOnAbortPlan:
		plan.OnAbort = &t
	case InterpOnErrorPlan:
		plan.OnError = &t
	case InterpEnsurePlan:
		plan.Ensure = &t
	case InterpOnSuccessPlan:
		plan.OnSuccess = &t
	case InterpOnFailurePlan:
		plan.OnFailure = &t
	case InterpTryPlan:
		plan.Try = &t
	case InterpTimeoutPlan:
		plan.Timeout = &t
	case InterpRetryPlan:
		plan.Retry = &t
	case ArtifactInputPlan:
		plan.ArtifactInput = &t
	case ArtifactOutputPlan:
		plan.ArtifactOutput = &t
	default:
		panic(fmt.Sprintf("don't know how to construct plan from %T", step))
	}

	plan.ID = PlanID(fmt.Sprintf("%x", num))

	return plan
}
