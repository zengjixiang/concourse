package exec

import (
	"context"
	"errors"
	"time"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/vars/interp"
)

type interpDuration interface {
	Interpolate(interp.Resolver) (atc.Duration, error)
}

// TimeoutStep applies a fixed timeout to a step's Run.
type TimeoutStep struct {
	resolver interp.Resolver
	step     Step
	duration interpDuration
	timedOut bool
}

// Timeout constructs a TimeoutStep factory.
func Timeout(resolver interp.Resolver, step Step, duration interpDuration) *TimeoutStep {
	return &TimeoutStep{
		resolver: resolver,
		step:     step,
		duration: duration,
		timedOut: false,
	}
}

// Run invokes the nested step with the specified duration.
//
// If the nested step takes longer than the duration, it is sent the Interrupt
// signal, and the TimeoutStep returns nil once the nested step exits (ignoring
// the nested step's error).
//
// The result of the nested step's Run is returned.
func (ts *TimeoutStep) Run(ctx context.Context, state RunState) error {
	duration, err := ts.duration.Interpolate(ts.resolver)
	if err != nil {
		return err
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, time.Duration(duration))
	defer cancel()

	err = ts.step.Run(timeoutCtx, state)
	if errors.Is(err, context.DeadlineExceeded) {
		ts.timedOut = true
		return nil
	}

	return err
}

// Succeeded is true if the nested step completed successfully
// and did not time out.
func (ts *TimeoutStep) Succeeded() bool {
	return !ts.timedOut && ts.step.Succeeded()
}
