// Generated by tf-integration-gen. DO NOT EDIT.
package provider

import (
	"fmt"

	"time"

	"conductorone/internal/sdk"
	"conductorone/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

const dayforceCatalogID = "2nDXbxwZ6JBhdZyrTsCZobnfiJl"

func (r *IntegrationDayforceResourceModel) ToCreateDelegatedSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	catalogID := sdk.String(dayforceCatalogID)
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.ConnectorServiceCreateDelegatedRequest{
		DisplayName: sdk.String("Dayforce"),
		CatalogID:   catalogID,
		UserIds:     userIds,
	}
	return &out
}

func (r *IntegrationDayforceResourceModel) ToCreateSDKType() (*shared.ConnectorServiceCreateRequest, error) {
	catalogID := sdk.String(dayforceCatalogID)
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}

	configOut, configSet := r.getConfig()
	if !configSet {
		return nil, fmt.Errorf("config must be set for create request")
	}

	out := shared.ConnectorServiceCreateRequest{
		CatalogID: catalogID,
		UserIds:   userIds,
		Config: &shared.ConnectorServiceCreateRequestConfig{
			AtType: sdk.String(envConfigType),
			AdditionalProperties: map[string]interface{}{
				"configuration": configOut,
			},
		},
	}
	return &out, nil
}

func (r *IntegrationDayforceResourceModel) ToUpdateSDKType() (*shared.Connector, bool) {
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}

	configValues := r.populateConfig()

	configOut := make(map[string]string)
	configSet := false
	for key, configValue := range configValues {
		configOut[key] = ""
		if configValue != nil {
			configOut[key] = *configValue
			configSet = true
		}
	}
	if !configSet {
		configOut = nil
	}

	out := shared.Connector{
		DisplayName: sdk.String("Dayforce"),
		AppID:       sdk.String(r.AppID.ValueString()),
		CatalogID:   sdk.String(dayforceCatalogID),
		ID:          sdk.String(r.ID.ValueString()),
		UserIds:     userIds,
		Config:      makeConnectorConfig(configOut),
	}

	return &out, configSet
}

func (r *IntegrationDayforceResourceModel) populateConfig() map[string]*string {
	dayforceUsername := new(string)
	if !r.DayforceUsername.IsUnknown() && !r.DayforceUsername.IsNull() {
		*dayforceUsername = r.DayforceUsername.ValueString()
	} else {
		dayforceUsername = nil
	}

	dayforcePassword := new(string)
	if !r.DayforcePassword.IsUnknown() && !r.DayforcePassword.IsNull() {
		*dayforcePassword = r.DayforcePassword.ValueString()
	} else {
		dayforcePassword = nil
	}

	dayforceEnvironment := new(string)
	if !r.DayforceEnvironment.IsUnknown() && !r.DayforceEnvironment.IsNull() {
		*dayforceEnvironment = r.DayforceEnvironment.ValueString()
	} else {
		dayforceEnvironment = nil
	}

	dayforceUrl := new(string)
	if !r.DayforceUrl.IsUnknown() && !r.DayforceUrl.IsNull() {
		*dayforceUrl = r.DayforceUrl.ValueString()
	} else {
		dayforceUrl = nil
	}

	dayforceClientNamespace := new(string)
	if !r.DayforceClientNamespace.IsUnknown() && !r.DayforceClientNamespace.IsNull() {
		*dayforceClientNamespace = r.DayforceClientNamespace.ValueString()
	} else {
		dayforceClientNamespace = nil
	}

	configValues := map[string]*string{
		"dayforce_username":         dayforceUsername,
		"dayforce_password":         dayforcePassword,
		"dayforce_environment":      dayforceEnvironment,
		"dayforce_url":              dayforceUrl,
		"dayforce_client_namespace": dayforceClientNamespace,
	}

	return configValues
}

func (r *IntegrationDayforceResourceModel) getConfig() (map[string]string, bool) {
	configValues := r.populateConfig()
	configOut := make(map[string]string)
	configSet := false
	for key, configValue := range configValues {
		configOut[key] = ""
		if configValue != nil {
			configOut[key] = *configValue
			configSet = true
		}
	}
	if !configSet {
		configOut = nil
	}
	return configOut, configSet
}

func (r *IntegrationDayforceResourceModel) ToGetSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationDayforceResourceModel) ToDeleteSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationDayforceResourceModel) RefreshFromGetResponse(resp *shared.Connector) {
	if resp == nil {
		return
	}
	if resp.AppID != nil {
		r.AppID = types.StringValue(*resp.AppID)
	} else {
		r.AppID = types.StringNull()
	}

	if resp.CreatedAt != nil {
		r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339))
	} else {
		r.CreatedAt = types.StringNull()
	}
	if resp.DeletedAt != nil {
		r.DeletedAt = types.StringValue(resp.DeletedAt.Format(time.RFC3339))
	} else {
		r.DeletedAt = types.StringNull()
	}
	if resp.ID != nil {
		r.ID = types.StringValue(*resp.ID)
	} else {
		r.ID = types.StringNull()
	}
	if resp.UpdatedAt != nil {
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339))
	} else {
		r.UpdatedAt = types.StringNull()
	}
	r.UserIds = nil
	for _, v := range resp.UserIds {
		r.UserIds = append(r.UserIds, types.StringValue(v))
	}

	if resp.Config != nil && *resp.Config.AtType == envConfigType {
		if config, ok := resp.Config.AdditionalProperties.(map[string]interface{}); ok {
			if values, ok := config["configuration"].(map[string]interface{}); ok {
				if v, ok := values["dayforce_username"]; ok {
					r.DayforceUsername = types.StringValue(v.(string))
				}

				if v, ok := values["dayforce_environment"]; ok {
					r.DayforceEnvironment = types.StringValue(v.(string))
				}

				if v, ok := values["dayforce_url"]; ok {
					r.DayforceUrl = types.StringValue(v.(string))
				}

				if v, ok := values["dayforce_client_namespace"]; ok {
					r.DayforceClientNamespace = types.StringValue(v.(string))
				}

			}
		}
	}
}

func (r *IntegrationDayforceResourceModel) RefreshFromUpdateResponse(resp *shared.Connector) {
	r.RefreshFromGetResponse(resp)
}

func (r *IntegrationDayforceResourceModel) RefreshFromCreateResponse(resp *shared.Connector) {
	if resp.AppID != nil {
		r.AppID = types.StringValue(*resp.AppID)
	} else {
		r.AppID = types.StringNull()
	}
	if resp.CreatedAt != nil {
		r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339))
	} else {
		r.CreatedAt = types.StringNull()
	}
	if resp.DeletedAt != nil {
		r.DeletedAt = types.StringValue(resp.DeletedAt.Format(time.RFC3339))
	} else {
		r.DeletedAt = types.StringNull()
	}
	if resp.ID != nil {
		r.ID = types.StringValue(*resp.ID)
	} else {
		r.ID = types.StringNull()
	}
	if resp.UpdatedAt != nil {
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339))
	} else {
		r.UpdatedAt = types.StringNull()
	}
	r.UserIds = nil
	for _, v := range resp.UserIds {
		r.UserIds = append(r.UserIds, types.StringValue(v))
	}

	if resp.Config != nil && *resp.Config.AtType == envConfigType {
		if config, ok := resp.Config.AdditionalProperties.(map[string]interface{}); ok {
			if values, ok := config["configuration"].(map[string]interface{}); ok {
				if v, ok := values["dayforce_username"]; ok {
					r.DayforceUsername = types.StringValue(v.(string))
				}

				if v, ok := values["dayforce_environment"]; ok {
					r.DayforceEnvironment = types.StringValue(v.(string))
				}

				if v, ok := values["dayforce_url"]; ok {
					r.DayforceUrl = types.StringValue(v.(string))
				}

				if v, ok := values["dayforce_client_namespace"]; ok {
					r.DayforceClientNamespace = types.StringValue(v.(string))
				}

			}
		}
	}
}
