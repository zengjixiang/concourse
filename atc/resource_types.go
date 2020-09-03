package atc

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/concourse/concourse/vars/interp"
)

type MetadataField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Source map[string]interface{}

//interpgen:generate InterpSource

type InterpSource map[interp.String]interp.Any

func (src Source) MarshalJSON() ([]byte, error) {
	return marshalStringOnlyKeys(src)
}

func (src InterpSource) MarshalJSON() ([]byte, error) {
	return marshalStringOnlyKeys(src)
}

type Params map[string]interface{}

//interpgen:generate InterpParams

type InterpParams map[interp.String]interp.Any

func (ps Params) MarshalJSON() ([]byte, error) {
	return marshalStringOnlyKeys(ps)
}

func (ps InterpParams) MarshalJSON() ([]byte, error) {
	return marshalStringOnlyKeys(ps)
}

func marshalStringOnlyKeys(v interface{}) ([]byte, error) {
	if v == nil {
		return json.Marshal(nil)
	}

	strKeys, err := stringifyKeys(v)
	if err != nil {
		return nil, err
	}

	return json.Marshal(strKeys)
}

//interpgen:generate Version

type Version map[string]string

func stringifyKeys(root interface{}) (interface{}, error) {
	val := reflect.ValueOf(root)

	switch val.Kind() {
	case reflect.Map:
		sanitized := map[string]interface{}{}

		iter := val.MapRange()
		for iter.Next() {
			k := iter.Key()
			v := iter.Value()

			if k.Kind() != reflect.String {
				return nil, fmt.Errorf("non-string key: '%s'", k.Interface())
			}
			str := k.String()

			sub, err := stringifyKeys(v.Interface())
			if err != nil {
				return nil, err
			}

			sanitized[str] = sub
		}

		return sanitized, nil

	case reflect.Slice:
		sanitized := make([]interface{}, val.Len())

		for i := range sanitized {
			v := val.Index(i)

			sub, err := stringifyKeys(v.Interface())
			if err != nil {
				return nil, err
			}

			sanitized[i] = sub
		}

		return sanitized, nil

	default:
		return root, nil
	}
}
