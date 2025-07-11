// Generated by tf-integration-gen. DO NOT EDIT.
package provider

import (
	"fmt"
	"strconv"
	"time"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

const linearCatalogID = "2OHqazgOJBJ9IqnQ0TbbNjdZuoH"

func (r *IntegrationLinearResourceModel) ToCreateDelegatedSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	catalogID := sdk.String(linearCatalogID)
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.ConnectorServiceCreateDelegatedRequest{
		DisplayName: sdk.String("Linear"),
		CatalogID:   catalogID,
		UserIds:     userIds,
	}
	return &out
}

func (r *IntegrationLinearResourceModel) ToCreateSDKType() (*shared.ConnectorServiceCreateRequest, error) {
	catalogID := sdk.String(linearCatalogID)
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

func (r *IntegrationLinearResourceModel) ToUpdateSDKType() (*shared.ConnectorInput, bool) {
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
		DisplayName: sdk.String("Linear"),
		AppID:       sdk.String(r.AppID.ValueString()),
		CatalogID:   sdk.String(linearCatalogID),
		ID:          sdk.String(r.ID.ValueString()),
		UserIds:     userIds,
		Config:      makeConnectorConfig(configOut),
	}

	return &out, configSet
}

func (r *IntegrationLinearResourceModel) populateConfig() map[string]interface{} {
	configValues := make(map[string]interface{})

	linearApiKey := new(string)
	if !r.LinearApiKey.IsUnknown() && !r.LinearApiKey.IsNull() {
		*linearApiKey = r.LinearApiKey.ValueString()
		configValues["linear_api_key"] = linearApiKey
	}

	enableExternalTicketProvisioning := new(string)
	if !r.EnableExternalTicketProvisioning.IsUnknown() && !r.EnableExternalTicketProvisioning.IsNull() {
		*enableExternalTicketProvisioning = strconv.FormatBool(r.EnableExternalTicketProvisioning.ValueBool())
		configValues["enable_external_ticket_provisioning"] = enableExternalTicketProvisioning
	}

	return configValues
}

func (r *IntegrationLinearResourceModel) getConfig() (map[string]interface{}, bool) {
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

func (r *IntegrationLinearResourceModel) ToGetSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationLinearResourceModel) ToDeleteSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationLinearResourceModel) RefreshFromGetResponse(resp *shared.Connector) {
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

				if localV, ok := configValues["enable_external_ticket_provisioning"]; ok {
					if v, ok := values["enable_external_ticket_provisioning"]; ok {
						if val, ok := v.(string); ok {
							bv, err := strconv.ParseBool(val)
							if err == nil {
								if localV != nil || (localV == nil && !bv) {
									r.EnableExternalTicketProvisioning = types.BoolValue(bv)
								}
							}
						}
					}
				}

			}
		}
	}
}

func (r *IntegrationLinearResourceModel) RefreshFromUpdateResponse(resp *shared.Connector) {
	r.RefreshFromGetResponse(resp)
}

func (r *IntegrationLinearResourceModel) RefreshFromCreateResponse(resp *shared.Connector) {
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

				if localV, ok := configValues["enable_external_ticket_provisioning"]; ok {
					if v, ok := values["enable_external_ticket_provisioning"]; ok {
						if val, ok := v.(string); ok {
							bv, err := strconv.ParseBool(val)
							if err == nil {
								if localV != nil || (localV == nil && !bv) {
									r.EnableExternalTicketProvisioning = types.BoolValue(bv)
								}
							}
						}
					}
				}

			}
		}
	}
}
