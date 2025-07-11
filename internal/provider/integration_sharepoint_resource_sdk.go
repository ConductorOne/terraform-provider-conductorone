// Generated by tf-integration-gen. DO NOT EDIT.
package provider

import (
	"fmt"

	"time"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

const sharepointCatalogID = "2xQBFEgKJaw94aPsmYyjyoO1Pl8"

func (r *IntegrationSharepointResourceModel) ToCreateDelegatedSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	catalogID := sdk.String(sharepointCatalogID)
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.ConnectorServiceCreateDelegatedRequest{
		DisplayName: sdk.String("SharePoint"),
		CatalogID:   catalogID,
		UserIds:     userIds,
	}
	return &out
}

func (r *IntegrationSharepointResourceModel) ToCreateSDKType() (*shared.ConnectorServiceCreateRequest, error) {
	catalogID := sdk.String(sharepointCatalogID)
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

func (r *IntegrationSharepointResourceModel) ToUpdateSDKType() (*shared.ConnectorInput, bool) {
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
		DisplayName: sdk.String("SharePoint"),
		AppID:       sdk.String(r.AppID.ValueString()),
		CatalogID:   sdk.String(sharepointCatalogID),
		ID:          sdk.String(r.ID.ValueString()),
		UserIds:     userIds,
		Config:      makeConnectorConfig(configOut),
	}

	return &out, configSet
}

func (r *IntegrationSharepointResourceModel) populateConfig() map[string]interface{} {
	configValues := make(map[string]interface{})

	azureClientId := new(string)
	if !r.AzureClientId.IsUnknown() && !r.AzureClientId.IsNull() {
		*azureClientId = r.AzureClientId.ValueString()
		configValues["azure_client_id"] = azureClientId
	}

	azureClientSecret := new(string)
	if !r.AzureClientSecret.IsUnknown() && !r.AzureClientSecret.IsNull() {
		*azureClientSecret = r.AzureClientSecret.ValueString()
		configValues["azure_client_secret"] = azureClientSecret
	}

	azureTenantId := new(string)
	if !r.AzureTenantId.IsUnknown() && !r.AzureTenantId.IsNull() {
		*azureTenantId = r.AzureTenantId.ValueString()
		configValues["azure_tenant_id"] = azureTenantId
	}

	azureGraphDomain := new(string)
	if !r.AzureGraphDomain.IsUnknown() && !r.AzureGraphDomain.IsNull() {
		*azureGraphDomain = r.AzureGraphDomain.ValueString()
		configValues["azure_graph_domain"] = azureGraphDomain
	}

	pfxCertificate := new(string)
	if !r.PfxCertificate.IsUnknown() && !r.PfxCertificate.IsNull() {
		*pfxCertificate = r.PfxCertificate.ValueString()
		configValues["pfx_certificate"] = pfxCertificate
	}

	pfxCertificatePassword := new(string)
	if !r.PfxCertificatePassword.IsUnknown() && !r.PfxCertificatePassword.IsNull() {
		*pfxCertificatePassword = r.PfxCertificatePassword.ValueString()
		configValues["pfx_certificate_password"] = pfxCertificatePassword
	}

	sharepointDomain := new(string)
	if !r.SharepointDomain.IsUnknown() && !r.SharepointDomain.IsNull() {
		*sharepointDomain = r.SharepointDomain.ValueString()
		configValues["sharepoint_domain"] = sharepointDomain
	}

	return configValues
}

func (r *IntegrationSharepointResourceModel) getConfig() (map[string]interface{}, bool) {
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

func (r *IntegrationSharepointResourceModel) ToGetSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationSharepointResourceModel) ToDeleteSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationSharepointResourceModel) RefreshFromGetResponse(resp *shared.Connector) {
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
				if v, ok := values["azure_client_id"]; ok {
					if val, ok := v.(string); ok {
						r.AzureClientId = types.StringValue(val)
					}
				}

				if v, ok := values["azure_tenant_id"]; ok {
					if val, ok := v.(string); ok {
						r.AzureTenantId = types.StringValue(val)
					}
				}

				if v, ok := values["azure_graph_domain"]; ok {
					if val, ok := v.(string); ok {
						r.AzureGraphDomain = types.StringValue(val)
					}
				}

				if v, ok := values["sharepoint_domain"]; ok {
					if val, ok := v.(string); ok {
						r.SharepointDomain = types.StringValue(val)
					}
				}

			}
		}
	}
}

func (r *IntegrationSharepointResourceModel) RefreshFromUpdateResponse(resp *shared.Connector) {
	r.RefreshFromGetResponse(resp)
}

func (r *IntegrationSharepointResourceModel) RefreshFromCreateResponse(resp *shared.Connector) {
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
				if v, ok := values["azure_client_id"]; ok {
					if val, ok := v.(string); ok {
						r.AzureClientId = types.StringValue(val)
					}
				}

				if v, ok := values["azure_tenant_id"]; ok {
					if val, ok := v.(string); ok {
						r.AzureTenantId = types.StringValue(val)
					}
				}

				if v, ok := values["azure_graph_domain"]; ok {
					if val, ok := v.(string); ok {
						r.AzureGraphDomain = types.StringValue(val)
					}
				}

				if v, ok := values["sharepoint_domain"]; ok {
					if val, ok := v.(string); ok {
						r.SharepointDomain = types.StringValue(val)
					}
				}

			}
		}
	}
}
