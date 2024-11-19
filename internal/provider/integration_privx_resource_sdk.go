// Generated by tf-integration-gen. DO NOT EDIT.
package provider

import (
	"fmt"

	"time"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

const privxCatalogID = "2jcBmRnl40QiOGhV3srq5rAYpjE"

func (r *IntegrationPrivxResourceModel) ToCreateDelegatedSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	catalogID := sdk.String(privxCatalogID)
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.ConnectorServiceCreateDelegatedRequest{
		DisplayName: sdk.String("PrivX"),
		CatalogID:   catalogID,
		UserIds:     userIds,
	}
	return &out
}

func (r *IntegrationPrivxResourceModel) ToCreateSDKType() (*shared.ConnectorServiceCreateRequest, error) {
	catalogID := sdk.String(privxCatalogID)
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

func (r *IntegrationPrivxResourceModel) ToUpdateSDKType() (*shared.Connector, bool) {
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
		DisplayName: sdk.String("PrivX"),
		AppID:       sdk.String(r.AppID.ValueString()),
		CatalogID:   sdk.String(privxCatalogID),
		ID:          sdk.String(r.ID.ValueString()),
		UserIds:     userIds,
		Config:      makeConnectorConfig(configOut),
	}

	return &out, configSet
}

func (r *IntegrationPrivxResourceModel) populateConfig() map[string]interface{} {
	configValues := make(map[string]interface{})

	privxBaseUrl := new(string)
	if !r.PrivxBaseUrl.IsUnknown() && !r.PrivxBaseUrl.IsNull() {
		*privxBaseUrl = r.PrivxBaseUrl.ValueString()
		configValues["privx_base_url"] = privxBaseUrl
	}

	privxClientId := new(string)
	if !r.PrivxClientId.IsUnknown() && !r.PrivxClientId.IsNull() {
		*privxClientId = r.PrivxClientId.ValueString()
		configValues["privx_client_id"] = privxClientId
	}

	privxClientSecret := new(string)
	if !r.PrivxClientSecret.IsUnknown() && !r.PrivxClientSecret.IsNull() {
		*privxClientSecret = r.PrivxClientSecret.ValueString()
		configValues["privx_client_secret"] = privxClientSecret
	}

	privxOauthClientId := new(string)
	if !r.PrivxOauthClientId.IsUnknown() && !r.PrivxOauthClientId.IsNull() {
		*privxOauthClientId = r.PrivxOauthClientId.ValueString()
		configValues["privx_oauth_client_id"] = privxOauthClientId
	}

	privxOauthClientSecret := new(string)
	if !r.PrivxOauthClientSecret.IsUnknown() && !r.PrivxOauthClientSecret.IsNull() {
		*privxOauthClientSecret = r.PrivxOauthClientSecret.ValueString()
		configValues["privx_oauth_client_secret"] = privxOauthClientSecret
	}

	return configValues
}

func (r *IntegrationPrivxResourceModel) getConfig() (map[string]interface{}, bool) {
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

func (r *IntegrationPrivxResourceModel) ToGetSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationPrivxResourceModel) ToDeleteSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationPrivxResourceModel) RefreshFromGetResponse(resp *shared.Connector) {
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
				if v, ok := values["privx_base_url"]; ok {
					if val, ok := v.(string); ok {
						r.PrivxBaseUrl = types.StringValue(val)
					}
				}

				if v, ok := values["privx_client_id"]; ok {
					if val, ok := v.(string); ok {
						r.PrivxClientId = types.StringValue(val)
					}
				}

				if v, ok := values["privx_oauth_client_id"]; ok {
					if val, ok := v.(string); ok {
						r.PrivxOauthClientId = types.StringValue(val)
					}
				}

			}
		}
	}
}

func (r *IntegrationPrivxResourceModel) RefreshFromUpdateResponse(resp *shared.Connector) {
	r.RefreshFromGetResponse(resp)
}

func (r *IntegrationPrivxResourceModel) RefreshFromCreateResponse(resp *shared.Connector) {
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
				if v, ok := values["privx_base_url"]; ok {
					if val, ok := v.(string); ok {
						r.PrivxBaseUrl = types.StringValue(val)
					}
				}

				if v, ok := values["privx_client_id"]; ok {
					if val, ok := v.(string); ok {
						r.PrivxClientId = types.StringValue(val)
					}
				}

				if v, ok := values["privx_oauth_client_id"]; ok {
					if val, ok := v.(string); ok {
						r.PrivxOauthClientId = types.StringValue(val)
					}
				}

			}
		}
	}
}