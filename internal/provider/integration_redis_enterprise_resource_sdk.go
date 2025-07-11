// Generated by tf-integration-gen. DO NOT EDIT.
package provider

import (
	"fmt"

	"time"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

const redisEnterpriseCatalogID = "2ueHfYDCTi4fKYcF64p23BWBcHb"

func (r *IntegrationRedisEnterpriseResourceModel) ToCreateDelegatedSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	catalogID := sdk.String(redisEnterpriseCatalogID)
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.ConnectorServiceCreateDelegatedRequest{
		DisplayName: sdk.String("Redis Enterprise"),
		CatalogID:   catalogID,
		UserIds:     userIds,
	}
	return &out
}

func (r *IntegrationRedisEnterpriseResourceModel) ToCreateSDKType() (*shared.ConnectorServiceCreateRequest, error) {
	catalogID := sdk.String(redisEnterpriseCatalogID)
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

func (r *IntegrationRedisEnterpriseResourceModel) ToUpdateSDKType() (*shared.ConnectorInput, bool) {
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
			configOut[key] = makeStringValue(configValue)
			configSet = true
		}
	}
	if !configSet {
		configOut = nil
	}

	out := shared.ConnectorInput{
		DisplayName: sdk.String("Redis Enterprise"),
		AppID:       sdk.String(r.AppID.ValueString()),
		CatalogID:   sdk.String(redisEnterpriseCatalogID),
		ID:          sdk.String(r.ID.ValueString()),
		UserIds:     userIds,
		Config:      makeConnectorConfig(configOut),
	}

	return &out, configSet
}

func (r *IntegrationRedisEnterpriseResourceModel) populateConfig() map[string]interface{} {
	configValues := make(map[string]interface{})

	clusterHost := new(string)
	if !r.ClusterHost.IsUnknown() && !r.ClusterHost.IsNull() {
		*clusterHost = r.ClusterHost.ValueString()
		configValues["cluster-host"] = clusterHost
	}

	apiPort := new(string)
	if !r.ApiPort.IsUnknown() && !r.ApiPort.IsNull() {
		*apiPort = r.ApiPort.ValueString()
		configValues["api-port"] = apiPort
	}

	username := new(string)
	if !r.Username.IsUnknown() && !r.Username.IsNull() {
		*username = r.Username.ValueString()
		configValues["username"] = username
	}

	password := new(string)
	if !r.Password.IsUnknown() && !r.Password.IsNull() {
		*password = r.Password.ValueString()
		configValues["password"] = password
	}

	return configValues
}

func (r *IntegrationRedisEnterpriseResourceModel) getConfig() (map[string]interface{}, bool) {
	configValues := r.populateConfig()
	configOut := make(map[string]interface{})
	configSet := false
	for key, configValue := range configValues {
		configOut[key] = ""
		if configValue != nil {
			configOut[key] = makeStringValue(configValue)
			configSet = true
		}
	}
	if !configSet {
		configOut = nil
	}
	return configOut, configSet
}

func (r *IntegrationRedisEnterpriseResourceModel) ToGetSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationRedisEnterpriseResourceModel) ToDeleteSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationRedisEnterpriseResourceModel) RefreshFromGetResponse(resp *shared.Connector) {
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
				if v, ok := values["cluster-host"]; ok {
					if val, ok := v.(string); ok {
						r.ClusterHost = types.StringValue(val)
					}
				}

				if v, ok := values["api-port"]; ok {
					if val, ok := v.(string); ok {
						r.ApiPort = types.StringValue(val)
					}
				}

				if v, ok := values["username"]; ok {
					if val, ok := v.(string); ok {
						r.Username = types.StringValue(val)
					}
				}

			}
		}
	}
}

func (r *IntegrationRedisEnterpriseResourceModel) RefreshFromUpdateResponse(resp *shared.Connector) {
	r.RefreshFromGetResponse(resp)
}

func (r *IntegrationRedisEnterpriseResourceModel) RefreshFromCreateResponse(resp *shared.Connector) {
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
				if v, ok := values["cluster-host"]; ok {
					if val, ok := v.(string); ok {
						r.ClusterHost = types.StringValue(val)
					}
				}

				if v, ok := values["api-port"]; ok {
					if val, ok := v.(string); ok {
						r.ApiPort = types.StringValue(val)
					}
				}

				if v, ok := values["username"]; ok {
					if val, ok := v.(string); ok {
						r.Username = types.StringValue(val)
					}
				}

			}
		}
	}
}
