// Generated by tf-integration-gen. DO NOT EDIT.
package provider

import (
	"fmt"
	"strconv"
	"time"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

const oktaV2CatalogID = "2kG3aTftCcAvWtHZt2Ywry0dPn7"

func (r *IntegrationOktaV2ResourceModel) ToCreateDelegatedSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	catalogID := sdk.String(oktaV2CatalogID)
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.ConnectorServiceCreateDelegatedRequest{
		DisplayName: sdk.String("Okta v2"),
		CatalogID:   catalogID,
		UserIds:     userIds,
	}
	return &out
}

func (r *IntegrationOktaV2ResourceModel) ToCreateSDKType() (*shared.ConnectorServiceCreateRequest, error) {
	catalogID := sdk.String(oktaV2CatalogID)
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

func (r *IntegrationOktaV2ResourceModel) ToUpdateSDKType() (*shared.Connector, bool) {
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
		DisplayName: sdk.String("Okta v2"),
		AppID:       sdk.String(r.AppID.ValueString()),
		CatalogID:   sdk.String(oktaV2CatalogID),
		ID:          sdk.String(r.ID.ValueString()),
		UserIds:     userIds,
		Config:      makeConnectorConfig(configOut),
	}

	return &out, configSet
}

func (r *IntegrationOktaV2ResourceModel) populateConfig() map[string]interface{} {
	configValues := make(map[string]interface{})

	oktaV2Domain := new(string)
	if !r.OktaV2Domain.IsUnknown() && !r.OktaV2Domain.IsNull() {
		*oktaV2Domain = r.OktaV2Domain.ValueString()
		configValues["okta_v2_domain"] = oktaV2Domain
	}

	oktaV2ApiToken := new(string)
	if !r.OktaV2ApiToken.IsUnknown() && !r.OktaV2ApiToken.IsNull() {
		*oktaV2ApiToken = r.OktaV2ApiToken.ValueString()
		configValues["okta_v2_api_token"] = oktaV2ApiToken
	}

	oktaSyncCustomRoles := new(string)
	if !r.OktaSyncCustomRoles.IsUnknown() && !r.OktaSyncCustomRoles.IsNull() {
		*oktaSyncCustomRoles = strconv.FormatBool(r.OktaSyncCustomRoles.ValueBool())
		configValues["okta_sync_custom_roles"] = oktaSyncCustomRoles
	}

	return configValues
}

func (r *IntegrationOktaV2ResourceModel) getConfig() (map[string]interface{}, bool) {
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

func (r *IntegrationOktaV2ResourceModel) ToGetSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationOktaV2ResourceModel) ToDeleteSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationOktaV2ResourceModel) RefreshFromGetResponse(resp *shared.Connector) {
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

	configValues := r.populateConfig()
	if resp.Config != nil && *resp.Config.AtType == envConfigType {
		if config, ok := resp.Config.AdditionalProperties.(map[string]interface{}); ok {
			if values, ok := config["configuration"].(map[string]interface{}); ok {
				if v, ok := values["okta_v2_domain"]; ok {
					if val, ok := v.(string); ok {
						r.OktaV2Domain = types.StringValue(val)
					}
				}

				if localV, ok := configValues["okta_sync_custom_roles"]; ok {
					if v, ok := values["okta_sync_custom_roles"]; ok {
						if val, ok := v.(string); ok {
							bv, err := strconv.ParseBool(val)
							if err == nil {
								if localV != nil || (localV == nil && !bv) {
									r.OktaSyncCustomRoles = types.BoolValue(bv)
								}
							}
						}
					}
				}

			}
		}
	}
}

func (r *IntegrationOktaV2ResourceModel) RefreshFromUpdateResponse(resp *shared.Connector) {
	r.RefreshFromGetResponse(resp)
}

func (r *IntegrationOktaV2ResourceModel) RefreshFromCreateResponse(resp *shared.Connector) {
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

	configValues := r.populateConfig()
	if resp.Config != nil && *resp.Config.AtType == envConfigType {
		if config, ok := resp.Config.AdditionalProperties.(map[string]interface{}); ok {
			if values, ok := config["configuration"].(map[string]interface{}); ok {
				if v, ok := values["okta_v2_domain"]; ok {
					if val, ok := v.(string); ok {
						r.OktaV2Domain = types.StringValue(val)
					}
				}

				if localV, ok := configValues["okta_sync_custom_roles"]; ok {
					if v, ok := values["okta_sync_custom_roles"]; ok {
						if val, ok := v.(string); ok {
							bv, err := strconv.ParseBool(val)
							if err == nil {
								if localV != nil || (localV == nil && !bv) {
									r.OktaSyncCustomRoles = types.BoolValue(bv)
								}
							}
						}
					}
				}

			}
		}
	}
}