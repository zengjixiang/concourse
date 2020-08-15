package atc

type ResourceConfig struct {
	Name         string  `json:"name"`
	OldName      string  `json:"old_name,omitempty"`
	Public       bool    `json:"public,omitempty"`
	WebhookToken string  `json:"webhook_token,omitempty"`
	Type         string  `json:"type"`
	Source       Source  `json:"source"`
	CheckEvery   string  `json:"check_every,omitempty"`
	CheckTimeout string  `json:"check_timeout,omitempty"`
	Tags         Tags    `json:"tags,omitempty"`
	Version      Version `json:"version,omitempty"`
	Icon         string  `json:"icon,omitempty"`
}

func (config ResourceConfig) StepConfig() StepConfig {
	return &CheckStep{
		Name:     config.Name,
		Resource: config.Name,
		Tags:     config.Tags,
	}
}
