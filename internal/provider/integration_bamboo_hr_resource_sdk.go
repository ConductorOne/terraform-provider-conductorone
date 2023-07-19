// Generated by tf-integration-gen. DO NOT EDIT.
package provider

import (
	"time"

	"conductorone/internal/sdk"
	"conductorone/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

const bambooHrCatalogID = "2CrErEM4PZ0cFwdLKjHVnGqrBVF"

func (r *IntegrationBambooHrResourceModel) ToCreateSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	catalogID := sdk.String(bambooHrCatalogID)
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.ConnectorServiceCreateDelegatedRequest{
		DisplayName: sdk.String("BambooHR"),
		CatalogID:   catalogID,
		UserIds:     userIds,
	}
	return &out
}

func (r *IntegrationBambooHrResourceModel) ToUpdateSDKType() (*shared.Connector, bool) {
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}

	companyDomain := new(string)
	if !r.CompanyDomain.IsUnknown() && !r.CompanyDomain.IsNull() {
		*companyDomain = r.CompanyDomain.ValueString()
	} else {
		companyDomain = nil
	}

	apiKey := new(string)
	if !r.ApiKey.IsUnknown() && !r.ApiKey.IsNull() {
		*apiKey = r.ApiKey.ValueString()
	} else {
		apiKey = nil
	}

	configValues := map[string]*string{
		"company_domain": companyDomain,
		"api_key":        apiKey,
	}

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
		DisplayName: sdk.String("BambooHR"),
		AppID:       sdk.String(r.AppID.ValueString()),
		CatalogID:   sdk.String(bambooHrCatalogID),
		ID:          sdk.String(r.ID.ValueString()),
		UserIds:     userIds,
		Config:      makeConnectorConfig(configOut),
	}

	return &out, configSet
}

func (r *IntegrationBambooHrResourceModel) ToGetSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *IntegrationBambooHrResourceModel) ToDeleteSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *IntegrationBambooHrResourceModel) RefreshFromGetResponse(resp *shared.Connector) {
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
				if v, ok := values["company_domain"]; ok {
					r.CompanyDomain = types.StringValue(v.(string))
				}

				if v, ok := values["api_key"]; ok {
					r.ApiKey = types.StringValue(v.(string))
				}

			}
		}
	}
}

func (r *IntegrationBambooHrResourceModel) RefreshFromUpdateResponse(resp *shared.Connector) {
	r.RefreshFromGetResponse(resp)
}

func (r *IntegrationBambooHrResourceModel) RefreshFromCreateResponse(resp *shared.Connector) {
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
				if v, ok := values["company_domain"]; ok {
					r.CompanyDomain = types.StringValue(v.(string))
				}

				if v, ok := values["api_key"]; ok {
					r.ApiKey = types.StringValue(v.(string))
				}

			}
		}
	}
}