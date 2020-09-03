package interp

import (
	"fmt"
	"reflect"

	"github.com/concourse/concourse/vars"
)

var (
	errAssignedNotAVar             = fmt.Errorf("assigned value is not a var reference")
	errInvalidTypeForInterpolation = func(t reflect.Type) error {
		return fmt.Errorf("cannot interpolate %s into a string (only strings, numbers, and bools are supported)", t.Kind().String())
	}
	errNotFound = func(ref vars.VariableReference) error {
		return fmt.Errorf("requested var %q was not found", ref.Name)
	}
)
