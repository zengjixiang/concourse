package interp

import "github.com/concourse/concourse/vars"

type Resolver interface {
	Resolve(vars.VariableReference) (interface{}, error)
}

type VarsResolver struct {
	vars.Variables
}

func (vr VarsResolver) Resolve(ref vars.VariableReference) (interface{}, error) {
	val, found, err := vr.Get(vars.VariableDefinition{Ref: ref})
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errNotFound(ref)
	}
	return val, nil
}
