// Generated by tf-integration-gen. DO NOT EDIT.
package provider

import (
	"time"

	"conductorone/internal/sdk"
	"conductorone/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

const datadogCatalogID = "23NeUC7UYRBJYDXscZNn4p8SQP2"

func (r *IntegrationDatadogResourceModel) ToCreateSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	catalogID := sdk.String(datadogCatalogID)
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.ConnectorServiceCreateDelegatedRequest{
		DisplayName: sdk.String("Datadog"),
		CatalogID:   catalogID,
		UserIds:     userIds,
	}
	return &out
}

func (r *IntegrationDatadogResourceModel) ToUpdateSDKType() (*shared.Connector, bool) {
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
		DisplayName: sdk.String("Datadog"),
		AppID:       sdk.String(r.AppID.ValueString()),
		CatalogID:   sdk.String(datadogCatalogID),
		ID:          sdk.String(r.ID.ValueString()),
		UserIds:     userIds,
		Config:      makeConnectorConfig(configOut),
	}

	return &out, configSet
}

func (r *IntegrationDatadogResourceModel) populateConfig() map[string]*string {
	datadogSite := new(string)
	if !r.DatadogSite.IsUnknown() && !r.DatadogSite.IsNull() {
		*datadogSite = r.DatadogSite.ValueString()
	} else {
		datadogSite = nil
	}

	datadogApiKey := new(string)
	if !r.DatadogApiKey.IsUnknown() && !r.DatadogApiKey.IsNull() {
		*datadogApiKey = r.DatadogApiKey.ValueString()
	} else {
		datadogApiKey = nil
	}

	datadogApplicationKey := new(string)
	if !r.DatadogApplicationKey.IsUnknown() && !r.DatadogApplicationKey.IsNull() {
		*datadogApplicationKey = r.DatadogApplicationKey.ValueString()
	} else {
		datadogApplicationKey = nil
	}

	configValues := map[string]*string{
		"datadog_site":            datadogSite,
		"datadog_api_key":         datadogApiKey,
		"datadog_application_key": datadogApplicationKey,
	}

	return configValues
}

func (r *IntegrationDatadogResourceModel) ToGetSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *IntegrationDatadogResourceModel) ToDeleteSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *IntegrationDatadogResourceModel) RefreshFromGetResponse(resp *shared.Connector) {
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
				if v, ok := values["datadog_site"]; ok {
					r.DatadogSite = types.StringValue(v.(string))
				}

			}
		}
	}
}

func (r *IntegrationDatadogResourceModel) RefreshFromUpdateResponse(resp *shared.Connector) {
	r.RefreshFromGetResponse(resp)
}

func (r *IntegrationDatadogResourceModel) RefreshFromCreateResponse(resp *shared.Connector) {
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
				if v, ok := values["datadog_site"]; ok {
					r.DatadogSite = types.StringValue(v.(string))
				}

			}
		}
	}
}
