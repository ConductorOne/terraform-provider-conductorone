// Generated by tf-integration-gen. DO NOT EDIT.
package provider

import (
	"fmt"
	"strconv"
	"time"

	"conductorone/internal/sdk"
	"conductorone/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

const oktaCatalogID = "23w9L3qudsiSZQJ8jUP1KYyQqVW"

func (r *IntegrationOktaResourceModel) ToCreateDelegatedSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	catalogID := sdk.String(oktaCatalogID)
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.ConnectorServiceCreateDelegatedRequest{
		DisplayName: sdk.String("Okta"),
		CatalogID:   catalogID,
		UserIds:     userIds,
	}
	return &out
}

func (r *IntegrationOktaResourceModel) ToCreateSDKType() (*shared.ConnectorServiceCreateRequest, error) {
	catalogID := sdk.String(oktaCatalogID)
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

func (r *IntegrationOktaResourceModel) ToUpdateSDKType() (*shared.Connector, bool) {
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
		DisplayName: sdk.String("Okta"),
		AppID:       sdk.String(r.AppID.ValueString()),
		CatalogID:   sdk.String(oktaCatalogID),
		ID:          sdk.String(r.ID.ValueString()),
		UserIds:     userIds,
		Config:      makeConnectorConfig(configOut),
	}

	return &out, configSet
}

func (r *IntegrationOktaResourceModel) populateConfig() map[string]*string {
	oktaDomain := new(string)
	if !r.OktaDomain.IsUnknown() && !r.OktaDomain.IsNull() {
		*oktaDomain = r.OktaDomain.ValueString()
	} else {
		oktaDomain = nil
	}

	oktaApiKey := new(string)
	if !r.OktaApiKey.IsUnknown() && !r.OktaApiKey.IsNull() {
		*oktaApiKey = r.OktaApiKey.ValueString()
	} else {
		oktaApiKey = nil
	}

	oktaDontSyncInactiveApps := new(string)
	if !r.OktaDontSyncInactiveApps.IsUnknown() && !r.OktaDontSyncInactiveApps.IsNull() {
		*oktaDontSyncInactiveApps = strconv.FormatBool(r.OktaDontSyncInactiveApps.ValueBool())
	} else {
		oktaDontSyncInactiveApps = nil
	}

	oktaExtractAwsSamlRoles := new(string)
	if !r.OktaExtractAwsSamlRoles.IsUnknown() && !r.OktaExtractAwsSamlRoles.IsNull() {
		*oktaExtractAwsSamlRoles = strconv.FormatBool(r.OktaExtractAwsSamlRoles.ValueBool())
	} else {
		oktaExtractAwsSamlRoles = nil
	}

	oktaSyncDeprovisionedUsers := new(string)
	if !r.OktaSyncDeprovisionedUsers.IsUnknown() && !r.OktaSyncDeprovisionedUsers.IsNull() {
		*oktaSyncDeprovisionedUsers = strconv.FormatBool(r.OktaSyncDeprovisionedUsers.ValueBool())
	} else {
		oktaSyncDeprovisionedUsers = nil
	}

	configValues := map[string]*string{
		"okta_domain":                   oktaDomain,
		"okta_api_key":                  oktaApiKey,
		"okta_dont_sync_inactive_apps":  oktaDontSyncInactiveApps,
		"okta_extract_aws_saml_roles":   oktaExtractAwsSamlRoles,
		"okta_sync_deprovisioned_users": oktaSyncDeprovisionedUsers,
	}

	return configValues
}

func (r *IntegrationOktaResourceModel) getConfig() (map[string]string, bool) {
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

func (r *IntegrationOktaResourceModel) ToGetSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationOktaResourceModel) ToDeleteSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationOktaResourceModel) RefreshFromGetResponse(resp *shared.Connector) {
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
				if v, ok := values["okta_domain"]; ok {
					r.OktaDomain = types.StringValue(v.(string))
				}

				if localV, ok := configValues["okta_dont_sync_inactive_apps"]; ok {
					if v, ok := values["okta_dont_sync_inactive_apps"]; ok {
						bv, err := strconv.ParseBool(v.(string))
						if err == nil {
							if localV != nil || (localV == nil && !bv) {
								r.OktaDontSyncInactiveApps = types.BoolValue(bv)
							}
						}
					}
				}

				if localV, ok := configValues["okta_extract_aws_saml_roles"]; ok {
					if v, ok := values["okta_extract_aws_saml_roles"]; ok {
						bv, err := strconv.ParseBool(v.(string))
						if err == nil {
							if localV != nil || (localV == nil && !bv) {
								r.OktaExtractAwsSamlRoles = types.BoolValue(bv)
							}
						}
					}
				}

				if localV, ok := configValues["okta_sync_deprovisioned_users"]; ok {
					if v, ok := values["okta_sync_deprovisioned_users"]; ok {
						bv, err := strconv.ParseBool(v.(string))
						if err == nil {
							if localV != nil || (localV == nil && !bv) {
								r.OktaSyncDeprovisionedUsers = types.BoolValue(bv)
							}
						}
					}
				}

			}
		}
	}
}

func (r *IntegrationOktaResourceModel) RefreshFromUpdateResponse(resp *shared.Connector) {
	r.RefreshFromGetResponse(resp)
}

func (r *IntegrationOktaResourceModel) RefreshFromCreateResponse(resp *shared.Connector) {
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
				if v, ok := values["okta_domain"]; ok {
					r.OktaDomain = types.StringValue(v.(string))
				}

				if localV, ok := configValues["okta_dont_sync_inactive_apps"]; ok {
					if v, ok := values["okta_dont_sync_inactive_apps"]; ok {
						bv, err := strconv.ParseBool(v.(string))
						if err == nil {
							if localV != nil || (localV == nil && !bv) {
								r.OktaDontSyncInactiveApps = types.BoolValue(bv)
							}
						}
					}
				}

				if localV, ok := configValues["okta_extract_aws_saml_roles"]; ok {
					if v, ok := values["okta_extract_aws_saml_roles"]; ok {
						bv, err := strconv.ParseBool(v.(string))
						if err == nil {
							if localV != nil || (localV == nil && !bv) {
								r.OktaExtractAwsSamlRoles = types.BoolValue(bv)
							}
						}
					}
				}

				if localV, ok := configValues["okta_sync_deprovisioned_users"]; ok {
					if v, ok := values["okta_sync_deprovisioned_users"]; ok {
						bv, err := strconv.ParseBool(v.(string))
						if err == nil {
							if localV != nil || (localV == nil && !bv) {
								r.OktaSyncDeprovisionedUsers = types.BoolValue(bv)
							}
						}
					}
				}

			}
		}
	}
}
