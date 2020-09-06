package vars

import (
	"fmt"
)

type Resolver struct {
	Variables
}

func (r Resolver) Resolve(varName string) (interface{}, error) {
	val, found, err := r.Get(VariableDefinition{Ref: parseVarName(varName)})
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, fmt.Errorf("var %q not found (%#v)", varName, parseVarName(varName))
	}
	return val, nil
}
