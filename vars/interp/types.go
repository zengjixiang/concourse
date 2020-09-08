package interp

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/concourse/concourse/vars"
	"github.com/hashicorp/go-multierror"
)

var (
	pathRegex                  = regexp.MustCompile(`("[^"]*"|[^\.]+)+`)
	interpolationRegex         = regexp.MustCompile(`\(\((([-/\.\w\pL]+\:)?[-/\.:"\w\pL]+)\)\)`)
	interpolationAnchoredRegex = regexp.MustCompile("\\A" + interpolationRegex.String() + "\\z")
)

type String string

func (s String) Interpolate(resolver Resolver) (string, error) {
	if interpolationAnchoredRegex.MatchString(string(s)) {
		var dst string
		if err := parseVarName(string(s)).InterpolateInto(resolver, &dst); err != nil {
			return "", err
		}
		return dst, nil
	}
	var merr error
	interpolated := interpolationRegex.ReplaceAllStringFunc(string(s), func(name string) string {
		var val interface{}
		if err := parseVarName(name).InterpolateInto(resolver, &val); err != nil {
			merr = multierror.Append(merr, err)
			return name
		}

		switch val := val.(type) {
		// TODO: don't want float64s, actually
		case string, float64, bool:
			return fmt.Sprint(val)
		default:
			merr = multierror.Append(merr, errInvalidTypeForInterpolation(reflect.TypeOf(val)))
			return name
		}
	})

	return interpolated, merr
}

func (s String) IsStatic() bool {
	return interpolationRegex.MatchString(string(s))
}

func parseVarName(name string) Var {
	name = strings.TrimPrefix(name, "((")
	name = strings.TrimSuffix(name, "))")
	var pathPieces []string

	varRef := vars.VariableReference{Name: name}

	if strings.Index(name, ":") > 0 {
		parts := strings.SplitN(name, ":", 2)
		varRef.Source = parts[0]

		pathPieces = pathRegex.FindAllString(parts[1], -1)

	} else {
		pathPieces = pathRegex.FindAllString(name, -1)
	}

	for i := range pathPieces {
		pathPieces[i] = strings.Trim(pathPieces[i], `"`)
	}
	varRef.Path = pathPieces[0]
	if len(pathPieces) >= 2 {
		varRef.Fields = pathPieces[1:]
	}

	return Var(varRef)
}

type Var vars.VariableReference

func (v *Var) UnmarshalJSON(data []byte) error {
	var dst string
	if err := json.Unmarshal(data, &dst); err != nil {
		return err
	}
	if !interpolationAnchoredRegex.MatchString(dst) {
		return errAssignedNotAVar
	}
	*v = Var(parseVarName(dst))
	return nil
}

func (v Var) MarshalJSON() ([]byte, error) {
	return json.Marshal("((" + v.Name + "))")
}

func (v Var) InterpolateInto(resolver Resolver, dst interface{}) error {
	val, err := resolver.Resolve(vars.VariableReference(v))
	if err != nil {
		return err
	}
	payload, err := json.Marshal(val)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(payload, &dst); err != nil {
		return err
	}
	return nil
}

type Any struct {
	Value interface{}
}

func (a *Any) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &a.Value)
}

func (a Any) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.Value)
}

func (a Any) Interpolate(resolver Resolver) (interface{}, error) {
	return interpolate(a.Value, resolver)
}

func interpolate(node interface{}, resolver Resolver) (interface{}, error) {
	switch typedNode := node.(type) {
	case map[interface{}]interface{}:
		for k, v := range typedNode {
			evaluatedValue, err := interpolate(v, resolver)
			if err != nil {
				return nil, err
			}

			evaluatedKey, err := interpolate(k, resolver)
			if err != nil {
				return nil, err
			}

			delete(typedNode, k) // delete in case key has changed
			typedNode[evaluatedKey] = evaluatedValue
		}

	case []interface{}:
		for idx, x := range typedNode {
			var err error
			typedNode[idx], err = interpolate(x, resolver)
			if err != nil {
				return nil, err
			}
		}

	case string:
		if interpolationAnchoredRegex.MatchString(typedNode) {
			var dst interface{}
			err := parseVarName(typedNode).InterpolateInto(resolver, &dst)
			if err != nil {
				return nil, err
			}
			return dst, nil
		}
		return String(typedNode).Interpolate(resolver)
	}

	return node, nil
}

type Int struct {
	I interface {
		Interpolate(Resolver) (int, error)
	}
}

func StaticInt(i int) Int {
	return Int{I: intVal(i)}
}

type intVal int
type intVar Var

func (v intVal) Interpolate(Resolver) (int, error) {
	return int(v), nil
}

func (v intVar) Interpolate(r Resolver) (int, error) {
	var dst int
	err := Var(v).InterpolateInto(r, &dst)
	return dst, err
}

func (i Int) MarshalJSON() ([]byte, error) { return json.Marshal(i.I) }

func (i *Int) UnmarshalJSON(data []byte) error {
	var v Var
	if err := json.Unmarshal(data, &v); err == nil {
		i.I = intVar(v)
		return nil
	}
	var raw int
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	i.I = intVal(raw)
	return nil
}

type Bool struct {
	I interface {
		Interpolate(Resolver) (bool, error)
	}
}

func StaticBool(b bool) Bool {
	return Bool{I: boolVal(b)}
}

type boolVal bool
type boolVar Var

func (v boolVal) Interpolate(Resolver) (bool, error) {
	return bool(v), nil
}

func (v boolVar) Interpolate(r Resolver) (bool, error) {
	var dst bool
	err := Var(v).InterpolateInto(r, &dst)
	return dst, err
}

func (i Bool) MarshalJSON() ([]byte, error) { return json.Marshal(i.I) }

func (i *Bool) UnmarshalJSON(data []byte) error {
	var v Var
	if err := json.Unmarshal(data, &v); err == nil {
		i.I = boolVar(v)
		return nil
	}
	var raw bool
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	i.I = boolVal(raw)
	return nil
}
