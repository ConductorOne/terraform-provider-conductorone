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

const microsoftEntraCatalogID = "2UUdYir252rR6PVSASeUFNaJOIB"

func (r *IntegrationMicrosoftEntraResourceModel) ToCreateDelegatedSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	catalogID := sdk.String(microsoftEntraCatalogID)
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.ConnectorServiceCreateDelegatedRequest{
		DisplayName: sdk.String("Microsoft Entra"),
		CatalogID:   catalogID,
		UserIds:     userIds,
	}
	return &out
}

func (r *IntegrationMicrosoftEntraResourceModel) ToCreateSDKType() (*shared.ConnectorServiceCreateRequest, error) {
	catalogID := sdk.String(microsoftEntraCatalogID)
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

func (r *IntegrationMicrosoftEntraResourceModel) ToUpdateSDKType() (*shared.Connector, bool) {
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
		DisplayName: sdk.String("Microsoft Entra"),
		AppID:       sdk.String(r.AppID.ValueString()),
		CatalogID:   sdk.String(microsoftEntraCatalogID),
		ID:          sdk.String(r.ID.ValueString()),
		UserIds:     userIds,
		Config:      makeConnectorConfig(configOut),
	}

	return &out, configSet
}

func (r *IntegrationMicrosoftEntraResourceModel) populateConfig() map[string]*string {
	entraTenantId := new(string)
	if !r.EntraTenantId.IsUnknown() && !r.EntraTenantId.IsNull() {
		*entraTenantId = r.EntraTenantId.ValueString()
	} else {
		entraTenantId = nil
	}

	entraClientId := new(string)
	if !r.EntraClientId.IsUnknown() && !r.EntraClientId.IsNull() {
		*entraClientId = r.EntraClientId.ValueString()
	} else {
		entraClientId = nil
	}

	entraClientSecret := new(string)
	if !r.EntraClientSecret.IsUnknown() && !r.EntraClientSecret.IsNull() {
		*entraClientSecret = r.EntraClientSecret.ValueString()
	} else {
		entraClientSecret = nil
	}

	entraSkipAdGroups := new(string)
	if !r.EntraSkipAdGroups.IsUnknown() && !r.EntraSkipAdGroups.IsNull() {
		*entraSkipAdGroups = strconv.FormatBool(r.EntraSkipAdGroups.ValueBool())
	} else {
		entraSkipAdGroups = nil
	}

	configValues := map[string]*string{
		"entra_tenant_id":      entraTenantId,
		"entra_client_id":      entraClientId,
		"entra_client_secret":  entraClientSecret,
		"entra_skip_ad_groups": entraSkipAdGroups,
	}

	return configValues
}

func (r *IntegrationMicrosoftEntraResourceModel) getConfig() (map[string]string, bool) {
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

func (r *IntegrationMicrosoftEntraResourceModel) ToGetSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationMicrosoftEntraResourceModel) ToDeleteSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationMicrosoftEntraResourceModel) RefreshFromGetResponse(resp *shared.Connector) {
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
				if v, ok := values["entra_tenant_id"]; ok {
					r.EntraTenantId = types.StringValue(v.(string))
				}

				if v, ok := values["entra_client_id"]; ok {
					r.EntraClientId = types.StringValue(v.(string))
				}

				if localV, ok := configValues["entra_skip_ad_groups"]; ok {
					if v, ok := values["entra_skip_ad_groups"]; ok {
						bv, err := strconv.ParseBool(v.(string))
						if err == nil {
							if localV != nil || (localV == nil && !bv) {
								r.EntraSkipAdGroups = types.BoolValue(bv)
							}
						}
					}
				}

			}
		}
	}
}

func (r *IntegrationMicrosoftEntraResourceModel) RefreshFromUpdateResponse(resp *shared.Connector) {
	r.RefreshFromGetResponse(resp)
}

func (r *IntegrationMicrosoftEntraResourceModel) RefreshFromCreateResponse(resp *shared.Connector) {
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
				if v, ok := values["entra_tenant_id"]; ok {
					r.EntraTenantId = types.StringValue(v.(string))
				}

				if v, ok := values["entra_client_id"]; ok {
					r.EntraClientId = types.StringValue(v.(string))
				}

				if localV, ok := configValues["entra_skip_ad_groups"]; ok {
					if v, ok := values["entra_skip_ad_groups"]; ok {
						bv, err := strconv.ParseBool(v.(string))
						if err == nil {
							if localV != nil || (localV == nil && !bv) {
								r.EntraSkipAdGroups = types.BoolValue(bv)
							}
						}
					}
				}

			}
		}
	}
}
