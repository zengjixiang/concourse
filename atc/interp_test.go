package atc_test

import (
	"encoding/json"
	"fmt"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/vars"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type errorResolver struct {
}

func (r errorResolver) Resolve(varName string) (interface{}, error) {
	return nil, fmt.Errorf("shouldn't need to resolve anything")
}

var _ = FDescribe("Interpolated Types", func() {
	Describe("InterpPlan", func() {
		It("works with a completely static plan", func() {
			cpu := atc.CPULimit(1)
			mem := atc.MemoryLimit(2)
			plan := atc.Plan{
				ID: "1",
				Task: &atc.TaskPlan{
					Name:       "some-task",
					Privileged: true,
					Tags:       []string{"hello", "world"},
					Config: &atc.TaskConfig{
						Platform: "linux",
						ImageResource: &atc.ImageResource{
							Type: "registry-image",
							Source: atc.Source{
								"repository": "ubuntu",
								"tag":        "latest",
							},
						},
						Limits: &atc.ContainerLimits{
							CPU:    &cpu,
							Memory: &mem,
						},
						Params: atc.TaskEnv{
							"hello": "world",
						},
						Run: atc.TaskRunConfig{
							Path: "echo",
							Args: []string{"hi"},
						},
						Inputs:  []atc.TaskInputConfig{{Name: "input"}},
						Outputs: []atc.TaskOutputConfig{{Name: "output"}},
					},
				},
			}

			data, err := json.Marshal(plan)
			Expect(err).ToNot(HaveOccurred())

			var interpPlan atc.InterpPlan
			err = json.Unmarshal(data, &interpPlan)
			Expect(err).ToNot(HaveOccurred())

			newPlan, err := interpPlan.Interpolate(errorResolver{})
			Expect(err).ToNot(HaveOccurred())

			Expect(newPlan).To(Equal(plan))
		})

		Context("when the plan has interpolation for simple/composite values", func() {
			data := `{
  "id": "1",
  "task": {
    "name": "((name))",
    "privileged": "((privileged))",
    "tags": "((tags))",
    "config": {
      "platform": "((platform))",
      "image_resource": {
        "type": "registry-image",
        "source": {
          "repository": "ubuntu",
          "tag": "((image-tag))"
        }
      },
      "container_limits": "((container-limits))",
      "params": "((params))",
      "run": {
        "path": "ls",
        "args": "((directories))"
      },
      "inputs": [
        {
          "name": "((input-name))"
        }
      ],
      "outputs": [
        "((output))"
      ]
    }
  }
}`
			It("interpolation works", func() {
				var interpPlan atc.InterpPlan
				err := json.Unmarshal([]byte(data), &interpPlan)
				Expect(err).ToNot(HaveOccurred())

				variables := vars.StaticVariables{
					"name":             "task-name",
					"privileged":       true,
					"tags":             []interface{}{"hello", "world"},
					"platform":         "linux",
					"image-tag":        "latest",
					"container-limits": map[string]interface{}{"cpu": 1, "memory": "1GB"},
					"params":           map[string]interface{}{"hello": 1, "world": []interface{}{"v1", "v2"}, "version": "1.2.3"},
					"directories":      []string{"dir1", "dir2"},
					"input-name":       "my-input",
					"output":           map[string]string{"name": "my-output"},
				}

				plan, err := interpPlan.Interpolate(vars.Resolver{Variables: variables})
				Expect(err).ToNot(HaveOccurred())

				cpu := atc.CPULimit(1)
				mem := atc.MemoryLimit(1 << 30) // 1GB
				Expect(plan).To(Equal(atc.Plan{
					ID: "1",
					Task: &atc.TaskPlan{
						Name:       "task-name",
						Privileged: true,
						Tags:       []string{"hello", "world"},
						Config: &atc.TaskConfig{
							Platform: "linux",
							ImageResource: &atc.ImageResource{
								Type: "registry-image",
								Source: atc.Source{
									"repository": "ubuntu",
									"tag":        "latest",
								},
							},
							Limits: &atc.ContainerLimits{
								CPU:    &cpu,
								Memory: &mem,
							},
							Params: atc.TaskEnv{
								"hello":   "1",
								"world":   `["v1","v2"]`,
								"version": "1.2.3",
							},
							Run: atc.TaskRunConfig{
								Path: "ls",
								Args: []string{"dir1", "dir2"},
							},
							Inputs:  []atc.TaskInputConfig{{Name: "my-input"}},
							Outputs: []atc.TaskOutputConfig{{Name: "my-output"}},
						},
					},
				}))
			})

			It("marshalling works", func() {
				var interpPlan atc.InterpPlan
				err := json.Unmarshal([]byte(data), &interpPlan)
				Expect(err).ToNot(HaveOccurred())

				newData, err := json.Marshal(interpPlan)
				Expect(err).ToNot(HaveOccurred())

				Expect(newData).To(MatchJSON(data))
			})
		})
	})
})
