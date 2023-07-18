package provider

import (
	"conductorone/internal/sdk"
	"conductorone/internal/sdk/pkg/models/shared"
)

const (
	envConfigType = "type.googleapis.com/c1.api.app.v1.EnvConfig"
)

func makeConnectorConfig(config map[string]interface{}) *shared.ConnectorConfig {
	out := shared.ConnectorConfig{
		AtType: sdk.String(envConfigType),
		AdditionalProperties: map[string]interface{}{
			"configuration": config,
		},
	}
	return &out
}
