package provider

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
)

const (
	envConfigType = "type.googleapis.com/c1.api.app.v1.EnvConfig"
)

func makeConnectorConfig(config map[string]interface{}) *shared.ConnectorConfig {
	if config == nil {
		config = make(map[string]interface{})
	}

	return &shared.ConnectorConfig{
		AtType: sdk.String(envConfigType),
		AdditionalProperties: map[string]interface{}{
			"configuration": config,
		},
	}
}
