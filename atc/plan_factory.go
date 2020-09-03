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

func (factory PlanFactory) NewPlan(step PlanConfig) PlanSkeleton {
	num := atomic.AddInt64(factory.currentNum, 1)

	var plan PlanSkeleton
	switch t := step.(type) {
	case AggregatePlanSkeleton:
		plan.Aggregate = &t
	case InParallelPlanSkeleton:
		plan.InParallel = &t
	case AcrossPlanSkeleton:
		plan.Across = &t
	case DoPlanSkeleton:
		plan.Do = &t
	case GetPlanSkeleton:
		plan.Get = &t
	case PutPlanSkeleton:
		plan.Put = &t
	case TaskPlanSkeleton:
		plan.Task = &t
	case SetPipelinePlanSkeleton:
		plan.SetPipeline = &t
	case LoadVarPlanSkeleton:
		plan.LoadVar = &t
	case CheckPlanSkeleton:
		plan.Check = &t
	case OnAbortPlanSkeleton:
		plan.OnAbort = &t
	case OnErrorPlanSkeleton:
		plan.OnError = &t
	case EnsurePlanSkeleton:
		plan.Ensure = &t
	case OnSuccessPlanSkeleton:
		plan.OnSuccess = &t
	case OnFailurePlanSkeleton:
		plan.OnFailure = &t
	case TryPlanSkeleton:
		plan.Try = &t
	case TimeoutPlanSkeleton:
		plan.Timeout = &t
	case RetryPlanSkeleton:
		plan.Retry = &t
	case ArtifactInputPlanSkeleton:
		plan.ArtifactInput = &t
	case ArtifactOutputPlanSkeleton:
		plan.ArtifactOutput = &t
	default:
		panic(fmt.Sprintf("don't know how to construct plan from %T", step))
	}

	plan.ID = PlanID(fmt.Sprintf("%x", num))

	return plan
}
