package atc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/concourse/concourse/vars/interp"
	"sigs.k8s.io/yaml"
)

type TaskConfig struct {
	// The platform the task must run on (e.g. linux, windows).
	Platform interp.String `json:"platform,omitempty"`

	// Optional string specifying an image to use for the build. Depending on the
	// platform, this may or may not be required (e.g. Windows/OS X vs. Linux).
	RootfsURI interp.String `json:"rootfs_uri,omitempty"`

	ImageResource *interpImageResource `json:"image_resource,omitempty"`

	// Limits to set on the Task Container
	Limits *interpContainerLimits `json:"container_limits,omitempty"`

	// Parameters to pass to the task via environment variables.
	Params interpTaskEnv `json:"params,omitempty"`

	// Script to execute.
	Run interpTaskRunConfig `json:"run,omitempty"`

	// The set of (logical, name-only) inputs required by the task.
	Inputs interpTaskInputConfigs `json:"inputs,omitempty"`

	// The set of (logical, name-only) outputs provided by the task.
	Outputs interpTaskOutputConfigs `json:"outputs,omitempty"`

	// Path to cached directory that will be shared between builds for the same task.
	Caches interpTaskCacheConfigs `json:"caches,omitempty"`
}

//interpgen:generate ImageResource
//interpgen:generate Source
//interpgen:generate Version

type ImageResource struct {
	Type   interp.String `json:"type"`
	Source interpSource  `json:"source"`

	Params  interpParams  `json:"params,omitempty"`
	Version interpVersion `json:"version,omitempty"`
}

func NewTaskConfig(configBytes []byte) (TaskConfig, error) {
	var config TaskConfig
	err := yaml.UnmarshalStrict(configBytes, &config, yaml.DisallowUnknownFields)
	if err != nil {
		return TaskConfig{}, err
	}

	err = config.Validate()
	if err != nil {
		return TaskConfig{}, err
	}

	return config, nil
}

type TaskValidationError struct {
	Errors []string
}

func (err TaskValidationError) Error() string {
	return fmt.Sprintf("invalid task configuration:\n%s", strings.Join(err.Errors, "\n"))
}

func (config TaskConfig) Validate() error {
	var errors []string

	if config.Platform == "" {
		errors = append(errors, "missing 'platform'")
	}

	if run, ok := config.Run.I.(TaskRunConfig); ok && run.Path == "" {
		errors = append(errors, "missing path to executable to run")
	}

	errors = append(errors, config.validateInputContainsNames()...)
	errors = append(errors, config.validateOutputContainsNames()...)

	if len(errors) > 0 {
		return TaskValidationError{
			Errors: errors,
		}
	}

	return nil
}

func (config TaskConfig) validateOutputContainsNames() []string {
	outputs, ok := config.Outputs.I.(TaskOutputConfigs)
	if !ok {
		return nil
	}

	var messages []string
	for i, output := range outputs {
		if output, ok := output.I.(TaskOutputConfig); ok && output.Name == "" {
			messages = append(messages, fmt.Sprintf("  output in position %d is missing a name", i))
		}
	}

	return messages
}

func (config TaskConfig) validateInputContainsNames() []string {
	inputs, ok := config.Inputs.I.(TaskInputConfigs)
	if !ok {
		return nil
	}

	messages := []string{}
	for i, input := range inputs {
		if input, ok := input.I.(TaskInputConfig); ok && input.Name == "" {
			messages = append(messages, fmt.Sprintf("  input in position %d is missing a name", i))
		}
	}

	return messages
}

//interpgen:generate TaskRunConfig

type TaskRunConfig struct {
	Path interp.String    `json:"path"`
	Args interpStringList `json:"args,omitempty"`
	Dir  interp.String    `json:"dir,omitempty"`

	// The user that the task will run as (defaults to whatever the docker image specifies)
	User interp.String `json:"user,omitempty"`
}

//interpgen:generate TaskInputConfig

type TaskInputConfig struct {
	Name     interp.String `json:"name"`
	Path     interp.String `json:"path,omitempty"`
	Optional interp.Bool   `json:"optional,omitempty"`
}

//interpgen:generate TaskInputConfigs

type TaskInputConfigs []interpTaskInputConfig

//interpgen:generate TaskOutputConfig

type TaskOutputConfig struct {
	Name interp.String `json:"name"`
	Path interp.String `json:"path,omitempty"`
}

//interpgen:generate TaskOutputConfigs

type TaskOutputConfigs []interpTaskOutputConfig

//interpgen:generate TaskCacheConfig

type TaskCacheConfig struct {
	Path interp.String `json:"path,omitempty"`
}

//interpgen:generate TaskCacheConfigs

type TaskCacheConfigs []interpTaskCacheConfig

//interpgen:generate TaskEnv

type TaskEnv map[interp.String]interpCoercedString

func (te TaskEnv) Env(resolver interp.Resolver) ([]string, error) {
	env := make([]string, 0, len(te))

	for k, v := range te {
		ki, err := k.Interpolate(resolver)
		if err != nil {
			return nil, err
		}
		vi, err := v.I.Interpolate(resolver)
		if err != nil {
			return nil, err
		}
		vii, err := vi.InterpolateToString(resolver)
		if err != nil {
			return nil, err
		}
		env = append(env, ki+"="+vii)
	}

	return env, nil
}

//interpgen:generate CoercedString

type CoercedString struct {
	raw interface{}
}

func (cs *CoercedString) UnmarshalJSON(p []byte) error {
	dec := json.NewDecoder(bytes.NewReader(p))
	dec.UseNumber()
	return dec.Decode(&cs.raw)
}

func (cs CoercedString) InterpolateToString(resolver interp.Resolver) (string, error) {
	if cs.raw == nil {
		return "", nil
	}
	switch v := cs.raw.(type) {
	case string:
		return interp.String(v).Interpolate(resolver)

	case json.Number:
		return string(v), nil

	default:
		j, err := json.Marshal(v)
		if err != nil {
			return "", err
		}
		return string(j), nil
	}
}
