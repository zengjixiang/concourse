package atc_test

import (
	"github.com/concourse/concourse/atc"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Plan", func() {
	Describe("Public", func() {
		It("returns a sanitized form of the plan", func() {
			plan := atc.PlanSkeleton{
				ID: "0",
				Aggregate: &atc.AggregatePlanSkeleton{
					atc.PlanSkeleton{
						ID: "1",
						Aggregate: &atc.AggregatePlanSkeleton{
							atc.PlanSkeleton{
								ID: "2",
								Task: &atc.TaskPlanSkeleton{
									Name:       "name",
									ConfigPath: "some/config/path.yml",
									Config: &atc.TaskConfig{
										Params: atc.TaskEnv{"some": "secret"},
									},
								},
							},
						},
					},

					atc.PlanSkeleton{
						ID: "3",
						Get: &atc.GetPlanSkeleton{
							Type:     "type",
							Name:     "name",
							Resource: "resource",
							Source:   atc.Source{"some": "source"},
							Params:   atc.Params{"some": "params"},
							Version:  &atc.Version{"some": "version"},
							Tags:     atc.Tags{"tags"},
						},
					},

					atc.PlanSkeleton{
						ID: "4",
						Put: &atc.PutPlanSkeleton{
							Type:     "type",
							Name:     "name",
							Resource: "resource",
							Source:   atc.Source{"some": "source"},
							Params:   atc.Params{"some": "params"},
							Tags:     atc.Tags{"tags"},
						},
					},

					atc.PlanSkeleton{
						ID: "4.2",
						Check: &atc.CheckPlanSkeleton{
							Type:   "type",
							Name:   "name",
							Source: atc.Source{"some": "source"},
							Tags:   atc.Tags{"tags"},
						},
					},

					atc.PlanSkeleton{
						ID: "5",
						Task: &atc.TaskPlanSkeleton{
							Name:       "name",
							Privileged: true,
							Tags:       atc.Tags{"tags"},
							ConfigPath: "some/config/path.yml",
							Config: &atc.TaskConfig{
								Params: atc.TaskEnv{"some": "secret"},
							},
						},
					},

					atc.PlanSkeleton{
						ID: "6",
						Ensure: &atc.EnsurePlanSkeleton{
							Step: atc.PlanSkeleton{
								ID: "7",
								Task: &atc.TaskPlanSkeleton{
									Name:       "name",
									ConfigPath: "some/config/path.yml",
									Config: &atc.TaskConfig{
										Params: atc.TaskEnv{"some": "secret"},
									},
								},
							},
							Next: atc.PlanSkeleton{
								ID: "8",
								Task: &atc.TaskPlanSkeleton{
									Name:       "name",
									ConfigPath: "some/config/path.yml",
									Config: &atc.TaskConfig{
										Params: atc.TaskEnv{"some": "secret"},
									},
								},
							},
						},
					},

					atc.PlanSkeleton{
						ID: "9",
						OnSuccess: &atc.OnSuccessPlanSkeleton{
							Step: atc.PlanSkeleton{
								ID: "10",
								Task: &atc.TaskPlanSkeleton{
									Name:       "name",
									ConfigPath: "some/config/path.yml",
									Config: &atc.TaskConfig{
										Params: atc.TaskEnv{"some": "secret"},
									},
								},
							},
							Next: atc.PlanSkeleton{
								ID: "11",
								Task: &atc.TaskPlanSkeleton{
									Name:       "name",
									ConfigPath: "some/config/path.yml",
									Config: &atc.TaskConfig{
										Params: atc.TaskEnv{"some": "secret"},
									},
								},
							},
						},
					},

					atc.PlanSkeleton{
						ID: "12",
						OnFailure: &atc.OnFailurePlanSkeleton{
							Step: atc.PlanSkeleton{
								ID: "13",
								Task: &atc.TaskPlanSkeleton{
									Name:       "name",
									ConfigPath: "some/config/path.yml",
									Config: &atc.TaskConfig{
										Params: atc.TaskEnv{"some": "secret"},
									},
								},
							},
							Next: atc.PlanSkeleton{
								ID: "14",
								Task: &atc.TaskPlanSkeleton{
									Name:       "name",
									ConfigPath: "some/config/path.yml",
									Config: &atc.TaskConfig{
										Params: atc.TaskEnv{"some": "secret"},
									},
								},
							},
						},
					},

					atc.PlanSkeleton{
						ID: "15",
						OnAbort: &atc.OnAbortPlanSkeleton{
							Step: atc.PlanSkeleton{
								ID: "16",
								Task: &atc.TaskPlanSkeleton{
									Name:       "name",
									ConfigPath: "some/config/path.yml",
									Config: &atc.TaskConfig{
										Params: atc.TaskEnv{"some": "secret"},
									},
								},
							},
							Next: atc.PlanSkeleton{
								ID: "17",
								Task: &atc.TaskPlanSkeleton{
									Name:       "name",
									ConfigPath: "some/config/path.yml",
									Config: &atc.TaskConfig{
										Params: atc.TaskEnv{"some": "secret"},
									},
								},
							},
						},
					},

					atc.PlanSkeleton{
						ID: "18",
						Try: &atc.TryPlanSkeleton{
							Step: atc.PlanSkeleton{
								ID: "19",
								Task: &atc.TaskPlanSkeleton{
									Name:       "name",
									ConfigPath: "some/config/path.yml",
									Config: &atc.TaskConfig{
										Params: atc.TaskEnv{"some": "secret"},
									},
								},
							},
						},
					},

					atc.PlanSkeleton{
						ID: "20",
						Timeout: &atc.TimeoutPlanSkeleton{
							Step: atc.PlanSkeleton{
								ID: "21",
								Task: &atc.TaskPlanSkeleton{
									Name:       "name",
									ConfigPath: "some/config/path.yml",
									Config: &atc.TaskConfig{
										Params: atc.TaskEnv{"some": "secret"},
									},
								},
							},
							Duration: "lol",
						},
					},

					atc.PlanSkeleton{
						ID: "22",
						Do: &atc.DoPlanSkeleton{
							atc.PlanSkeleton{
								ID: "23",
								Task: &atc.TaskPlanSkeleton{
									Name:       "name",
									ConfigPath: "some/config/path.yml",
									Config: &atc.TaskConfig{
										Params: atc.TaskEnv{"some": "secret"},
									},
								},
							},
						},
					},

					atc.PlanSkeleton{
						ID: "24",
						Retry: &atc.RetryPlanSkeleton{
							atc.PlanSkeleton{
								ID: "25",
								Task: &atc.TaskPlanSkeleton{
									Name:       "name",
									ConfigPath: "some/config/path.yml",
									Config: &atc.TaskConfig{
										Params: atc.TaskEnv{"some": "secret"},
									},
								},
							},
							atc.PlanSkeleton{
								ID: "26",
								Task: &atc.TaskPlanSkeleton{
									Name:       "name",
									ConfigPath: "some/config/path.yml",
									Config: &atc.TaskConfig{
										Params: atc.TaskEnv{"some": "secret"},
									},
								},
							},
							atc.PlanSkeleton{
								ID: "27",
								Task: &atc.TaskPlanSkeleton{
									Name:       "name",
									ConfigPath: "some/config/path.yml",
									Config: &atc.TaskConfig{
										Params: atc.TaskEnv{"some": "secret"},
									},
								},
							},
						},
					},

					atc.PlanSkeleton{
						ID: "28",
						OnAbort: &atc.OnAbortPlanSkeleton{
							Step: atc.PlanSkeleton{
								ID: "29",
								Task: &atc.TaskPlanSkeleton{
									Name:       "name",
									ConfigPath: "some/config/path.yml",
									Config: &atc.TaskConfig{
										Params: atc.TaskEnv{"some": "secret"},
									},
								},
							},
							Next: atc.PlanSkeleton{
								ID: "30",
								Task: &atc.TaskPlanSkeleton{
									Name:       "name",
									ConfigPath: "some/config/path.yml",
									Config: &atc.TaskConfig{
										Params: atc.TaskEnv{"some": "secret"},
									},
								},
							},
						},
					},

					atc.PlanSkeleton{
						ID: "31",
						ArtifactInput: &atc.ArtifactInputPlanSkeleton{
							ArtifactID: 17,
							Name:       "some-name",
						},
					},

					atc.PlanSkeleton{
						ID: "32",
						ArtifactOutput: &atc.ArtifactOutputPlanSkeleton{
							Name: "some-name",
						},
					},

					atc.PlanSkeleton{
						ID: "33",
						OnError: &atc.OnErrorPlanSkeleton{
							Step: atc.PlanSkeleton{
								ID: "34",
								Task: &atc.TaskPlanSkeleton{
									Name:       "name",
									ConfigPath: "some/config/path.yml",
									Config: &atc.TaskConfig{
										Params: atc.TaskEnv{"some": "secret"},
									},
								},
							},
							Next: atc.PlanSkeleton{
								ID: "35",
								Task: &atc.TaskPlanSkeleton{
									Name:       "name",
									ConfigPath: "some/config/path.yml",
									Config: &atc.TaskConfig{
										Params: atc.TaskEnv{"some": "secret"},
									},
								},
							},
						},
					},
					atc.PlanSkeleton{
						ID: "36",
						InParallel: &atc.InParallelPlanSkeleton{
							Limit:    1,
							FailFast: true,
							Steps: []atc.PlanSkeleton{
								{
									ID: "37",
									Task: &atc.TaskPlanSkeleton{
										Name:       "name",
										ConfigPath: "some/config/path.yml",
										Config: &atc.TaskConfig{
											Params: atc.TaskEnv{"some": "secret"},
										},
									},
								},
							},
						},
					},
					atc.PlanSkeleton{
						ID: "38",
						SetPipeline: &atc.SetPipelinePlanSkeleton{
							Name:     "some-pipeline",
							Team:     "some-team",
							File:     "some-file",
							VarFiles: []string{"vf"},
							Vars:     map[string]interface{}{"k1": "v1"},
						},
					},
					atc.PlanSkeleton{
						ID: "39",
						Across: &atc.AcrossPlanSkeleton{
							Vars: []atc.AcrossVar{
								{
									Var:         "v1",
									Values:      []interface{}{"a"},
									MaxInFlight: 1,
								},
								{
									Var:         "v2",
									Values:      []interface{}{"b"},
									MaxInFlight: 2,
								},
							},
							Steps: []atc.VarScopedPlanSkeleton{
								{
									Step: atc.PlanSkeleton{
										ID: "40",
										Task: &atc.TaskPlanSkeleton{
											Name:       "name",
											ConfigPath: "some/config/path.yml",
											Config: &atc.TaskConfig{
												Params: atc.TaskEnv{"some": "secret"},
											},
										},
									},
									Values: []interface{}{"a", "b"},
								},
							},
							FailFast: true,
						},
					},
				},
			}

			json := plan.Public()
			Expect(json).ToNot(BeNil())
			Expect([]byte(*json)).To(MatchJSON(`{
  "id": "0",
  "aggregate": [
    {
      "id": "1",
      "aggregate": [
        {
          "id": "2",
          "task": {
            "name": "name",
            "privileged": false
          }
        }
      ]
    },
    {
      "id": "3",
      "get": {
        "type": "type",
        "name": "name",
        "resource": "resource",
        "version": {
          "some": "version"
        }
      }
    },
    {
      "id": "4",
      "put": {
        "type": "type",
        "name": "name",
        "resource": "resource"
      }
    },
    {
      "id": "4.2",
      "check": {
        "type": "type",
        "name": "name"
      }
    },
    {
      "id": "5",
      "task": {
        "name": "name",
        "privileged": true
      }
    },
    {
      "id": "6",
      "ensure": {
        "step": {
          "id": "7",
          "task": {
            "name": "name",
            "privileged": false
          }
        },
        "ensure": {
          "id": "8",
          "task": {
            "name": "name",
            "privileged": false
          }
        }
      }
    },
    {
      "id": "9",
      "on_success": {
        "step": {
          "id": "10",
          "task": {
            "name": "name",
            "privileged": false
          }
        },
        "on_success": {
          "id": "11",
          "task": {
            "name": "name",
            "privileged": false
          }
        }
      }
    },
    {
      "id": "12",
      "on_failure": {
        "step": {
          "id": "13",
          "task": {
            "name": "name",
            "privileged": false
          }
        },
        "on_failure": {
          "id": "14",
          "task": {
            "name": "name",
            "privileged": false
          }
        }
      }
    },
    {
      "id": "15",
      "on_abort": {
        "step": {
          "id": "16",
          "task": {
            "name": "name",
            "privileged": false
          }
        },
        "on_abort": {
          "id": "17",
          "task": {
            "name": "name",
            "privileged": false
          }
        }
      }
    },
    {
      "id": "18",
      "try": {
        "step": {
          "id": "19",
          "task": {
            "name": "name",
            "privileged": false
          }
        }
      }
    },
    {
      "id": "20",
      "timeout": {
        "step": {
          "id": "21",
          "task": {
            "name": "name",
            "privileged": false
          }
        },
        "duration": "lol"
      }
    },
    {
      "id": "22",
      "do": [
        {
          "id": "23",
          "task": {
            "name": "name",
            "privileged": false
          }
        }
      ]
    },
    {
      "id": "24",
      "retry": [
        {
          "id": "25",
          "task": {
            "name": "name",
            "privileged": false
          }
        },
        {
          "id": "26",
          "task": {
            "name": "name",
            "privileged": false
          }
        },
        {
          "id": "27",
          "task": {
            "name": "name",
            "privileged": false
          }
        }
      ]
    },
    {
      "id": "28",
      "on_abort": {
				"step": {
						"id": "29",
						"task": {
							"name": "name",
							"privileged": false
						}
				},
        "on_abort": {
          "id": "30",
          "task": {
            "name": "name",
            "privileged": false
          }
				}
      }
    },
		{
			"id": "31",
			"artifact_input": {
				"artifact_id": 17,
				"name" : "some-name"
			}
		},
		{
			"id": "32",
			"artifact_output": {
				"name": "some-name"
			}
		},
		{
      "id": "33",
      "on_error": {
        "step": {
          "id": "34",
          "task": {
            "name": "name",
            "privileged": false
          }
        },
        "on_error": {
          "id": "35",
          "task": {
            "name": "name",
            "privileged": false
          }
        }
      }
	},
	{
	"id": "36",
	  "in_parallel": {
		"steps": [
		  {
			"id": "37",
			"task": {
              "name": "name",
              "privileged": false
			}
		  }
		],
		"limit": 1,
		"fail_fast": true
	  }
	},
	{
	  "id": "38",
	  "set_pipeline": {
		"name": "some-pipeline",
		"team": "some-team"
	  }
	},
	{
	  "id": "39",
	  "across": {
	    "vars": [
	      {
	        "name": "v1",
	        "values": ["a"],
	        "max_in_flight": 1
	      },
	      {
	        "name": "v2",
	        "values": ["b"],
	        "max_in_flight": 2
	      }
	    ],
	    "steps": [
	      {
	        "step": {
	          "id": "40",
	          "task": {
	            "name": "name",
	            "privileged": false
	          }
	        },
	        "values": ["a", "b"]
	      }
	    ],
	    "fail_fast": true
	  }
	}
  ]
}
`))
		})
	})
})
