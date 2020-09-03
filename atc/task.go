package atc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aoldershaw/interpolate"
	"sigs.k8s.io/yaml"
)

type TaskConfig struct {
	// The platform the task must run on (e.g. linux, windows).
	Platform string `json:"platform,omitempty" interpolate:"root"`

	// Optional string specifying an image to use for the build. Depending on the
	// platform, this may or may not be required (e.g. Windows/OS X vs. Linux).
	RootfsURI string `json:"rootfs_uri,omitempty" interpolate:"root"`

	ImageResource *ImageResource `json:"image_resource,omitempty" interpolate:"root"`

	// Limits to set on the Task Container
	Limits *ContainerLimits `json:"container_limits,omitempty" interpolate:"root"`

	// Parameters to pass to the task via environment variables.
	Params TaskEnv `json:"params,omitempty" interpolate:"root,keys=root,values=root"`

	// Script to execute.
	Run TaskRunConfig `json:"run,omitempty" interpolate:"root"`

	// The set of (logical, name-only) inputs required by the task.
	Inputs []TaskInputConfig `json:"inputs,omitempty" interpolate:"root,values=root"`

	// The set of (logical, name-only) outputs provided by the task.
	Outputs []TaskOutputConfig `json:"outputs,omitempty" interpolate:"root,values=root"`

	// Path to cached directory that will be shared between builds for the same task.
	Caches []TaskCacheConfig `json:"caches,omitempty" interpolate:"root,values=root"`
}

type ImageResource struct {
	Type   string `json:"type" interpolate:"root"`
	Source Source `json:"source" interpolate:"root,keys=root,values=root"`

	Params  Params  `json:"params,omitempty" interpolate:"root,keys=root,values=root"`
	Version Version `json:"version,omitempty" interpolate:"root,keys=root,values=root"`
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

	if config.Run.Path == "" {
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
	var messages []string

	for i, output := range config.Outputs {
		if output.Name == "" {
			messages = append(messages, fmt.Sprintf("  output in position %d is missing a name", i))
		}
	}

	return messages
}

func (config TaskConfig) validateInputContainsNames() []string {
	messages := []string{}

	for i, input := range config.Inputs {
		if input.Name == "" {
			messages = append(messages, fmt.Sprintf("  input in position %d is missing a name", i))
		}
	}

	return messages
}

type TaskRunConfig struct {
	Path string   `json:"path" interpolate:"root"`
	Args []string `json:"args,omitempty" interpolate:"root,values=root"`
	Dir  string   `json:"dir,omitempty" interpolate:"root"`

	// The user that the task will run as (defaults to whatever the docker image specifies)
	User string `json:"user,omitempty" interpolate:"root"`
}

type TaskInputConfig struct {
	Name     string `json:"name" interpolate:"root"`
	Path     string `json:"path,omitempty" interpolate:"root"`
	Optional bool   `json:"optional,omitempty" interpolate:"root"`
}

type TaskOutputConfig struct {
	Name string `json:"name" interpolate:"root"`
	Path string `json:"path,omitempty" interpolate:"root"`
}

type TaskCacheConfig struct {
	Path string `json:"path,omitempty" interpolate:"root"`
}

type TaskEnv map[string]string

func (t *TaskEnv) UnmarshalJSON(p []byte) error {
	raw := map[string]CoercedString{}
	err := json.Unmarshal(p, &raw)
	if err != nil {
		return err
	}

	m := map[string]string{}
	for k, v := range raw {
		m[k] = string(v)
	}

	*t = m

	return nil
}

//interpolate:skip UnmarshalJSON TaskEnv

func (i *interpTaskEnv) UnmarshalJSON(p []byte) error {
	if err := json.Unmarshal(p, &i.Ref); err == nil {
		return nil
	}
	raw := map[string]CoercedString{}
	err := json.Unmarshal(p, &raw)
	if err != nil {
		return err
	}

	m := map[interpolate.String]interpolate.String{}
	for k, v := range raw {
		m[interpolate.String(k)] = interpolate.String(v)
	}

	i.Val = m

	return nil
}

func (te TaskEnv) Env() []string {
	env := make([]string, 0, len(te))

	for k, v := range te {
		env = append(env, k+"="+v)
	}

	return env
}

type CoercedString string

func (cs *CoercedString) UnmarshalJSON(p []byte) error {
	var raw interface{}
	dec := json.NewDecoder(bytes.NewReader(p))
	dec.UseNumber()
	err := dec.Decode(&raw)
	if err != nil {
		return err
	}

	if raw == nil {
		*cs = CoercedString("")
		return nil
	}
	switch v := raw.(type) {
	case string:
		*cs = CoercedString(v)

	case json.Number:
		*cs = CoercedString(v)

	default:
		j, err := json.Marshal(v)
		if err != nil {
			return err
		}

		*cs = CoercedString(j)
	}

	return nil
}
