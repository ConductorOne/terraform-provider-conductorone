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

const gcpGwsCatalogID = "2Te9ag5brAPS25E3hsloNmKRQSM"

func (r *IntegrationGcpGwsResourceModel) ToCreateDelegatedSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	catalogID := sdk.String(gcpGwsCatalogID)
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.ConnectorServiceCreateDelegatedRequest{
		DisplayName: sdk.String("Google Cloud Platform (With Google Workspace)"),
		CatalogID:   catalogID,
		UserIds:     userIds,
	}
	return &out
}

func (r *IntegrationGcpGwsResourceModel) ToCreateSDKType() (*shared.ConnectorServiceCreateRequest, error) {
	catalogID := sdk.String(gcpGwsCatalogID)
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

func (r *IntegrationGcpGwsResourceModel) ToUpdateSDKType() (*shared.Connector, bool) {
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
		DisplayName: sdk.String("Google Cloud Platform (With Google Workspace)"),
		AppID:       sdk.String(r.AppID.ValueString()),
		CatalogID:   sdk.String(gcpGwsCatalogID),
		ID:          sdk.String(r.ID.ValueString()),
		UserIds:     userIds,
		Config:      makeConnectorConfig(configOut),
	}

	return &out, configSet
}

func (r *IntegrationGcpGwsResourceModel) populateConfig() map[string]interface{} {
	configValues := make(map[string]interface{})

	customerId := new(string)
	if !r.CustomerId.IsUnknown() && !r.CustomerId.IsNull() {
		*customerId = r.CustomerId.ValueString()
		configValues["customer_id"] = customerId
	}

	domain := new(string)
	if !r.Domain.IsUnknown() && !r.Domain.IsNull() {
		*domain = r.Domain.ValueString()
		configValues["domain"] = domain
	}

	administratorEmail := new(string)
	if !r.AdministratorEmail.IsUnknown() && !r.AdministratorEmail.IsNull() {
		*administratorEmail = r.AdministratorEmail.ValueString()
		configValues["administrator_email"] = administratorEmail
	}

	credentialsJson := new(string)
	if !r.CredentialsJson.IsUnknown() && !r.CredentialsJson.IsNull() {
		*credentialsJson = r.CredentialsJson.ValueString()
		configValues["credentials_json"] = credentialsJson
	}

	skipSystemAccounts := new(string)
	if !r.SkipSystemAccounts.IsUnknown() && !r.SkipSystemAccounts.IsNull() {
		*skipSystemAccounts = strconv.FormatBool(r.SkipSystemAccounts.ValueBool())
		configValues["skip_system_accounts"] = skipSystemAccounts
	}

	skipDefaultProjects := new(string)
	if !r.SkipDefaultProjects.IsUnknown() && !r.SkipDefaultProjects.IsNull() {
		*skipDefaultProjects = strconv.FormatBool(r.SkipDefaultProjects.ValueBool())
		configValues["skip_default_projects"] = skipDefaultProjects
	}

	return configValues
}

func (r *IntegrationGcpGwsResourceModel) getConfig() (map[string]interface{}, bool) {
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

func (r *IntegrationGcpGwsResourceModel) ToGetSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationGcpGwsResourceModel) ToDeleteSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationGcpGwsResourceModel) RefreshFromGetResponse(resp *shared.Connector) {
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
				if v, ok := values["customer_id"]; ok {
					if val, ok := v.(string); ok {
						r.CustomerId = types.StringValue(val)
					}
				}

				if v, ok := values["domain"]; ok {
					if val, ok := v.(string); ok {
						r.Domain = types.StringValue(val)
					}
				}

				if v, ok := values["administrator_email"]; ok {
					if val, ok := v.(string); ok {
						r.AdministratorEmail = types.StringValue(val)
					}
				}

				if localV, ok := configValues["skip_system_accounts"]; ok {
					if v, ok := values["skip_system_accounts"]; ok {
						if val, ok := v.(string); ok {
							bv, err := strconv.ParseBool(val)
							if err == nil {
								if localV != nil || (localV == nil && !bv) {
									r.SkipSystemAccounts = types.BoolValue(bv)
								}
							}
						}
					}
				}

				if localV, ok := configValues["skip_default_projects"]; ok {
					if v, ok := values["skip_default_projects"]; ok {
						if val, ok := v.(string); ok {
							bv, err := strconv.ParseBool(val)
							if err == nil {
								if localV != nil || (localV == nil && !bv) {
									r.SkipDefaultProjects = types.BoolValue(bv)
								}
							}
						}
					}
				}

			}
		}
	}
}

func (r *IntegrationGcpGwsResourceModel) RefreshFromUpdateResponse(resp *shared.Connector) {
	r.RefreshFromGetResponse(resp)
}

func (r *IntegrationGcpGwsResourceModel) RefreshFromCreateResponse(resp *shared.Connector) {
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
				if v, ok := values["customer_id"]; ok {
					if val, ok := v.(string); ok {
						r.CustomerId = types.StringValue(val)
					}
				}

				if v, ok := values["domain"]; ok {
					if val, ok := v.(string); ok {
						r.Domain = types.StringValue(val)
					}
				}

				if v, ok := values["administrator_email"]; ok {
					if val, ok := v.(string); ok {
						r.AdministratorEmail = types.StringValue(val)
					}
				}

				if localV, ok := configValues["skip_system_accounts"]; ok {
					if v, ok := values["skip_system_accounts"]; ok {
						if val, ok := v.(string); ok {
							bv, err := strconv.ParseBool(val)
							if err == nil {
								if localV != nil || (localV == nil && !bv) {
									r.SkipSystemAccounts = types.BoolValue(bv)
								}
							}
						}
					}
				}

				if localV, ok := configValues["skip_default_projects"]; ok {
					if v, ok := values["skip_default_projects"]; ok {
						if val, ok := v.(string); ok {
							bv, err := strconv.ParseBool(val)
							if err == nil {
								if localV != nil || (localV == nil && !bv) {
									r.SkipDefaultProjects = types.BoolValue(bv)
								}
							}
						}
					}
				}

			}
		}
	}
}