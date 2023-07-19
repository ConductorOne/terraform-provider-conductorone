package provider

import (
	"conductorone/internal/sdk"
	"conductorone/internal/sdk/pkg/models/shared"
)

const (
	envConfigType = "type.googleapis.com/c1.api.app.v1.EnvConfig"
)

func makeConnectorConfig(config map[string]string) *shared.ConnectorConfig {
	if config == nil {
		config = make(map[string]string)
	}

	return &shared.ConnectorConfig{
		AtType: sdk.String(envConfigType),
		AdditionalProperties: map[string]interface{}{
			"configuration": config,
		},
	}
}
