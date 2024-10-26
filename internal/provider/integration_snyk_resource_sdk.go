// Generated by tf-integration-gen. DO NOT EDIT.
package provider

import (
	"fmt"

	"time"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

const snykCatalogID = "2j75wwVsHdBnOduHXlJ0v0TuoKU"

func (r *IntegrationSnykResourceModel) ToCreateDelegatedSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	catalogID := sdk.String(snykCatalogID)
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.ConnectorServiceCreateDelegatedRequest{
		DisplayName: sdk.String("Snyk"),
		CatalogID:   catalogID,
		UserIds:     userIds,
	}
	return &out
}

func (r *IntegrationSnykResourceModel) ToCreateSDKType() (*shared.ConnectorServiceCreateRequest, error) {
	catalogID := sdk.String(snykCatalogID)
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

func (r *IntegrationSnykResourceModel) ToUpdateSDKType() (*shared.Connector, bool) {
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
		DisplayName: sdk.String("Snyk"),
		AppID:       sdk.String(r.AppID.ValueString()),
		CatalogID:   sdk.String(snykCatalogID),
		ID:          sdk.String(r.ID.ValueString()),
		UserIds:     userIds,
		Config:      makeConnectorConfig(configOut),
	}

	return &out, configSet
}

func (r *IntegrationSnykResourceModel) populateConfig() map[string]interface{} {
	configValues := make(map[string]interface{})

	snykApiToken := new(string)
	if !r.SnykApiToken.IsUnknown() && !r.SnykApiToken.IsNull() {
		*snykApiToken = r.SnykApiToken.ValueString()
		configValues["snyk_api_token"] = snykApiToken
	}

	snykGroupId := new(string)
	if !r.SnykGroupId.IsUnknown() && !r.SnykGroupId.IsNull() {
		*snykGroupId = r.SnykGroupId.ValueString()
		configValues["snyk_group_id"] = snykGroupId
	}

	snykOrgIds := make([]string, 0)
	for _, item := range r.SnykOrgIds {
		snykOrgIds = append(snykOrgIds, item.ValueString())
	}
	configValues["snyk_org_ids"] = snykOrgIds

	return configValues
}

func (r *IntegrationSnykResourceModel) getConfig() (map[string]interface{}, bool) {
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

func (r *IntegrationSnykResourceModel) ToGetSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationSnykResourceModel) ToDeleteSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationSnykResourceModel) RefreshFromGetResponse(resp *shared.Connector) {
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

				if v, ok := values["snyk_group_id"]; ok {
					if val, ok := v.(string); ok {
						r.SnykGroupId = types.StringValue(val)
					}
				}

				r.SnykOrgIds = nil
				if v, ok := values["snyk_org_ids"]; ok {
					if val, ok := v.([]string); ok {
						for _, item := range val {
							r.SnykOrgIds = append(r.SnykOrgIds, types.StringValue(item))
						}
					}
				}

			}
		}
	}
}

func (r *IntegrationSnykResourceModel) RefreshFromUpdateResponse(resp *shared.Connector) {
	r.RefreshFromGetResponse(resp)
}

func (r *IntegrationSnykResourceModel) RefreshFromCreateResponse(resp *shared.Connector) {
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

				if v, ok := values["snyk_group_id"]; ok {
					if val, ok := v.(string); ok {
						r.SnykGroupId = types.StringValue(val)
					}
				}

				r.SnykOrgIds = nil
				if v, ok := values["snyk_org_ids"]; ok {
					if val, ok := v.([]string); ok {
						for _, item := range val {
							r.SnykOrgIds = append(r.SnykOrgIds, types.StringValue(item))
						}
					}
				}

			}
		}
	}
}
