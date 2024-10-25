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

const salesforceV2CatalogID = "2k4W58EZuLALKVuoMIaFKRHvx9p"

func (r *IntegrationSalesforceV2ResourceModel) ToCreateDelegatedSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	catalogID := sdk.String(salesforceV2CatalogID)
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.ConnectorServiceCreateDelegatedRequest{
		DisplayName: sdk.String("Salesforce v2"),
		CatalogID:   catalogID,
		UserIds:     userIds,
	}
	return &out
}

func (r *IntegrationSalesforceV2ResourceModel) ToCreateSDKType() (*shared.ConnectorServiceCreateRequest, error) {
	catalogID := sdk.String(salesforceV2CatalogID)
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

func (r *IntegrationSalesforceV2ResourceModel) ToUpdateSDKType() (*shared.Connector, bool) {
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
		DisplayName: sdk.String("Salesforce v2"),
		AppID:       sdk.String(r.AppID.ValueString()),
		CatalogID:   sdk.String(salesforceV2CatalogID),
		ID:          sdk.String(r.ID.ValueString()),
		UserIds:     userIds,
		Config:      makeConnectorConfig(configOut),
	}

	return &out, configSet
}

func (r *IntegrationSalesforceV2ResourceModel) populateConfig() map[string]interface{} {
	salesforceInstanceUrl := new(string)
	if !r.SalesforceInstanceUrl.IsUnknown() && !r.SalesforceInstanceUrl.IsNull() {
		*salesforceInstanceUrl = r.SalesforceInstanceUrl.ValueString()
	} else {
		salesforceInstanceUrl = nil
	}

	salesforceUsernameForEmail := new(string)
	if !r.SalesforceUsernameForEmail.IsUnknown() && !r.SalesforceUsernameForEmail.IsNull() {
		*salesforceUsernameForEmail = strconv.FormatBool(r.SalesforceUsernameForEmail.ValueBool())
	} else {
		salesforceUsernameForEmail = nil
	}

	salesforceUsername := new(string)
	if !r.SalesforceUsername.IsUnknown() && !r.SalesforceUsername.IsNull() {
		*salesforceUsername = r.SalesforceUsername.ValueString()
	} else {
		salesforceUsername = nil
	}

	salesforcePassword := new(string)
	if !r.SalesforcePassword.IsUnknown() && !r.SalesforcePassword.IsNull() {
		*salesforcePassword = r.SalesforcePassword.ValueString()
	} else {
		salesforcePassword = nil
	}

	salesforceSecurityToken := new(string)
	if !r.SalesforceSecurityToken.IsUnknown() && !r.SalesforceSecurityToken.IsNull() {
		*salesforceSecurityToken = r.SalesforceSecurityToken.ValueString()
	} else {
		salesforceSecurityToken = nil
	}

	configValues := map[string]interface{}{
		"salesforce_instance_url":       salesforceInstanceUrl,
		"salesforce_username_for_email": salesforceUsernameForEmail,
		"salesforce_username":           salesforceUsername,
		"salesforce_password":           salesforcePassword,
		"salesforce_security_token":     salesforceSecurityToken,
	}

	return configValues
}

func (r *IntegrationSalesforceV2ResourceModel) getConfig() (map[string]interface{}, bool) {
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

func (r *IntegrationSalesforceV2ResourceModel) ToGetSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationSalesforceV2ResourceModel) ToDeleteSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationSalesforceV2ResourceModel) RefreshFromGetResponse(resp *shared.Connector) {
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
				if v, ok := values["salesforce_instance_url"]; ok {
					if val, ok := v.(string); ok {
						r.SalesforceInstanceUrl = types.StringValue(val)
					}
				}

				if localV, ok := configValues["salesforce_username_for_email"]; ok {
					if v, ok := values["salesforce_username_for_email"]; ok {
						if val, ok := v.(string); ok {
							bv, err := strconv.ParseBool(val)
							if err == nil {
								if localV != nil || (localV == nil && !bv) {
									r.SalesforceUsernameForEmail = types.BoolValue(bv)
								}
							}
						}
					}
				}

				if v, ok := values["salesforce_username"]; ok {
					if val, ok := v.(string); ok {
						r.SalesforceUsername = types.StringValue(val)
					}
				}

			}
		}
	}
}

func (r *IntegrationSalesforceV2ResourceModel) RefreshFromUpdateResponse(resp *shared.Connector) {
	r.RefreshFromGetResponse(resp)
}

func (r *IntegrationSalesforceV2ResourceModel) RefreshFromCreateResponse(resp *shared.Connector) {
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
				if v, ok := values["salesforce_instance_url"]; ok {
					if val, ok := v.(string); ok {
						r.SalesforceInstanceUrl = types.StringValue(val)
					}
				}

				if localV, ok := configValues["salesforce_username_for_email"]; ok {
					if v, ok := values["salesforce_username_for_email"]; ok {
						if val, ok := v.(string); ok {
							bv, err := strconv.ParseBool(val)
							if err == nil {
								if localV != nil || (localV == nil && !bv) {
									r.SalesforceUsernameForEmail = types.BoolValue(bv)
								}
							}
						}
					}
				}

				if v, ok := values["salesforce_username"]; ok {
					if val, ok := v.(string); ok {
						r.SalesforceUsername = types.StringValue(val)
					}
				}

			}
		}
	}
}
