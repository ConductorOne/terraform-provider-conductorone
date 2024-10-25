// Generated by tf-integration-gen. DO NOT EDIT.
package provider

import (
	"fmt"

	"time"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

const ukgCatalogID = "2QFfoa6GzhmBbncvQYjFlCIfzVn"

func (r *IntegrationUkgResourceModel) ToCreateDelegatedSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	catalogID := sdk.String(ukgCatalogID)
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.ConnectorServiceCreateDelegatedRequest{
		DisplayName: sdk.String("UKG"),
		CatalogID:   catalogID,
		UserIds:     userIds,
	}
	return &out
}

func (r *IntegrationUkgResourceModel) ToCreateSDKType() (*shared.ConnectorServiceCreateRequest, error) {
	catalogID := sdk.String(ukgCatalogID)
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

func (r *IntegrationUkgResourceModel) ToUpdateSDKType() (*shared.Connector, bool) {
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}

	configValues := r.populateConfig()

	configOut := make(map[string]interface{})
	configSet := false
	for key, configValue := range configValues {
		configOut[key] = ""
		if configValue != nil {
			configOut[key] = configValue
			configSet = true
		}
	}
	if !configSet {
		configOut = nil
	}

	out := shared.Connector{
		DisplayName: sdk.String("UKG"),
		AppID:       sdk.String(r.AppID.ValueString()),
		CatalogID:   sdk.String(ukgCatalogID),
		ID:          sdk.String(r.ID.ValueString()),
		UserIds:     userIds,
		Config:      makeConnectorConfig(configOut),
	}

	return &out, configSet
}

func (r *IntegrationUkgResourceModel) populateConfig() map[string]interface{} {
	ukgCustomerApiKey := new(string)
	if !r.UkgCustomerApiKey.IsUnknown() && !r.UkgCustomerApiKey.IsNull() {
		*ukgCustomerApiKey = r.UkgCustomerApiKey.ValueString()
	} else {
		ukgCustomerApiKey = nil
	}

	ukgUsername := new(string)
	if !r.UkgUsername.IsUnknown() && !r.UkgUsername.IsNull() {
		*ukgUsername = r.UkgUsername.ValueString()
	} else {
		ukgUsername = nil
	}

	ukgPassword := new(string)
	if !r.UkgPassword.IsUnknown() && !r.UkgPassword.IsNull() {
		*ukgPassword = r.UkgPassword.ValueString()
	} else {
		ukgPassword = nil
	}

	ukgServiceEndpoint := new(string)
	if !r.UkgServiceEndpoint.IsUnknown() && !r.UkgServiceEndpoint.IsNull() {
		*ukgServiceEndpoint = r.UkgServiceEndpoint.ValueString()
	} else {
		ukgServiceEndpoint = nil
	}

	configValues := map[string]interface{}{
		"ukg_customer_api_key": ukgCustomerApiKey,
		"ukg_username":         ukgUsername,
		"ukg_password":         ukgPassword,
		"ukg_service_endpoint": ukgServiceEndpoint,
	}

	return configValues
}

func (r *IntegrationUkgResourceModel) getConfig() (map[string]interface{}, bool) {
	configValues := r.populateConfig()
	configOut := make(map[string]interface{})
	configSet := false
	for key, configValue := range configValues {
		configOut[key] = ""
		if configValue != nil {
			configOut[key] = configValue
			configSet = true
		}
	}
	if !configSet {
		configOut = nil
	}
	return configOut, configSet
}

func (r *IntegrationUkgResourceModel) ToGetSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationUkgResourceModel) ToDeleteSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationUkgResourceModel) RefreshFromGetResponse(resp *shared.Connector) {
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
				if v, ok := values["ukg_customer_api_key"]; ok {
					if val, ok := v.(string); ok {
						r.UkgCustomerApiKey = types.StringValue(val)
					}
				}

				if v, ok := values["ukg_username"]; ok {
					if val, ok := v.(string); ok {
						r.UkgUsername = types.StringValue(val)
					}
				}

				if v, ok := values["ukg_service_endpoint"]; ok {
					if val, ok := v.(string); ok {
						r.UkgServiceEndpoint = types.StringValue(val)
					}
				}

			}
		}
	}
}

func (r *IntegrationUkgResourceModel) RefreshFromUpdateResponse(resp *shared.Connector) {
	r.RefreshFromGetResponse(resp)
}

func (r *IntegrationUkgResourceModel) RefreshFromCreateResponse(resp *shared.Connector) {
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
				if v, ok := values["ukg_customer_api_key"]; ok {
					if val, ok := v.(string); ok {
						r.UkgCustomerApiKey = types.StringValue(val)
					}
				}

				if v, ok := values["ukg_username"]; ok {
					if val, ok := v.(string); ok {
						r.UkgUsername = types.StringValue(val)
					}
				}

				if v, ok := values["ukg_service_endpoint"]; ok {
					if val, ok := v.(string); ok {
						r.UkgServiceEndpoint = types.StringValue(val)
					}
				}

			}
		}
	}
}
