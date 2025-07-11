package provider

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
)

const (
	envConfigType = "type.googleapis.com/c1.api.app.v1.EnvConfig"
)

type StringValue struct {
	StringValue interface{} `json:"stringValue"`
}

func makeConnectorConfig(config map[string]interface{}) *shared.Config {
	if config == nil {
		config = make(map[string]interface{})
	}

	return &shared.Config{
		AtType: sdk.String(envConfigType),
		AdditionalProperties: map[string]interface{}{
			"configuration": config,
		},
	}
}

func makeStringValue(value interface{}) *StringValue {
	return &StringValue{
		StringValue: value,
	}
}
