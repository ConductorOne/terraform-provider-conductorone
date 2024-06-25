// Generated by tf-integration-gen. DO NOT EDIT.
package provider

import (
	"fmt"

	"time"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

const microsoft365CatalogID = "2ZHQTiZa8KcWvZ6m2Z2hVy0aOlS"

func (r *IntegrationMicrosoft365ResourceModel) ToCreateDelegatedSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	catalogID := sdk.String(microsoft365CatalogID)
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.ConnectorServiceCreateDelegatedRequest{
		DisplayName: sdk.String("Microsoft 365"),
		CatalogID:   catalogID,
		UserIds:     userIds,
	}
	return &out
}

func (r *IntegrationMicrosoft365ResourceModel) ToCreateSDKType() (*shared.ConnectorServiceCreateRequest, error) {
	catalogID := sdk.String(microsoft365CatalogID)
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

func (r *IntegrationMicrosoft365ResourceModel) ToUpdateSDKType() (*shared.Connector, bool) {
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
		DisplayName: sdk.String("Microsoft 365"),
		AppID:       sdk.String(r.AppID.ValueString()),
		CatalogID:   sdk.String(microsoft365CatalogID),
		ID:          sdk.String(r.ID.ValueString()),
		UserIds:     userIds,
		Config:      makeConnectorConfig(configOut),
	}

	return &out, configSet
}

func (r *IntegrationMicrosoft365ResourceModel) populateConfig() map[string]*string {
	configValues := map[string]*string{}

	ms365TenantId := new(string)
	if !r.Ms365TenantId.IsUnknown() && !r.Ms365TenantId.IsNull() {
		*ms365TenantId = r.Ms365TenantId.ValueString()
		configValues["ms365_tenant_id"] = ms365TenantId
	}

	ms365ClientId := new(string)
	if !r.Ms365ClientId.IsUnknown() && !r.Ms365ClientId.IsNull() {
		*ms365ClientId = r.Ms365ClientId.ValueString()
		configValues["ms365_client_id"] = ms365ClientId
	}

	ms365ClientSecret := new(string)
	if !r.Ms365ClientSecret.IsUnknown() && !r.Ms365ClientSecret.IsNull() {
		*ms365ClientSecret = r.Ms365ClientSecret.ValueString()
		configValues["ms365_client_secret"] = ms365ClientSecret
	}

	return configValues
}

func (r *IntegrationMicrosoft365ResourceModel) getConfig() (map[string]string, bool) {
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

func (r *IntegrationMicrosoft365ResourceModel) ToGetSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationMicrosoft365ResourceModel) ToDeleteSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationMicrosoft365ResourceModel) RefreshFromGetResponse(resp *shared.Connector) {
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
				if v, ok := values["ms365_tenant_id"]; ok {
					r.Ms365TenantId = types.StringValue(v.(string))
				}

				if v, ok := values["ms365_client_id"]; ok {
					r.Ms365ClientId = types.StringValue(v.(string))
				}

			}
		}
	}
}

func (r *IntegrationMicrosoft365ResourceModel) RefreshFromUpdateResponse(resp *shared.Connector) {
	r.RefreshFromGetResponse(resp)
}

func (r *IntegrationMicrosoft365ResourceModel) RefreshFromCreateResponse(resp *shared.Connector) {
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
				if v, ok := values["ms365_tenant_id"]; ok {
					r.Ms365TenantId = types.StringValue(v.(string))
				}

				if v, ok := values["ms365_client_id"]; ok {
					r.Ms365ClientId = types.StringValue(v.(string))
				}

			}
		}
	}
}
