// Generated by tf-integration-gen. DO NOT EDIT.
package provider

import (
	"time"

	"conductorone/internal/sdk"
	"conductorone/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

const googleWorkspaceCatalogID = "24eqHE201g0mICwdTQoZF3skJyD"

func (r *IntegrationGoogleWorkspaceResourceModel) ToCreateSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	catalogID := sdk.String(googleWorkspaceCatalogID)
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.ConnectorServiceCreateDelegatedRequest{
		DisplayName: sdk.String("Google Workspace"),
		CatalogID:   catalogID,
		UserIds:     userIds,
	}
	return &out
}

func (r *IntegrationGoogleWorkspaceResourceModel) ToUpdateSDKType() (*shared.Connector, bool) {
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}

	customerId := new(string)
	if !r.CustomerId.IsUnknown() && !r.CustomerId.IsNull() {
		*customerId = r.CustomerId.ValueString()
	} else {
		customerId = nil
	}

	domain := new(string)
	if !r.Domain.IsUnknown() && !r.Domain.IsNull() {
		*domain = r.Domain.ValueString()
	} else {
		domain = nil
	}

	administratorEmail := new(string)
	if !r.AdministratorEmail.IsUnknown() && !r.AdministratorEmail.IsNull() {
		*administratorEmail = r.AdministratorEmail.ValueString()
	} else {
		administratorEmail = nil
	}

	credentialsJson := new(string)
	if !r.CredentialsJson.IsUnknown() && !r.CredentialsJson.IsNull() {
		*credentialsJson = r.CredentialsJson.ValueString()
	} else {
		credentialsJson = nil
	}

	configValues := map[string]*string{
		"customer_id":         customerId,
		"domain":              domain,
		"administrator_email": administratorEmail,
		"credentials_json":    credentialsJson,
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
		DisplayName: sdk.String("Google Workspace"),
		AppID:       sdk.String(r.AppID.ValueString()),
		CatalogID:   sdk.String(googleWorkspaceCatalogID),
		ID:          sdk.String(r.ID.ValueString()),
		UserIds:     userIds,
		Config:      makeConnectorConfig(configOut),
	}

	return &out, configSet
}

func (r *IntegrationGoogleWorkspaceResourceModel) ToGetSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *IntegrationGoogleWorkspaceResourceModel) ToDeleteSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *IntegrationGoogleWorkspaceResourceModel) RefreshFromGetResponse(resp *shared.Connector) {
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
				if v, ok := values["customer_id"]; ok {
					r.CustomerId = types.StringValue(v.(string))
				}

				if v, ok := values["domain"]; ok {
					r.Domain = types.StringValue(v.(string))
				}

				if v, ok := values["administrator_email"]; ok {
					r.AdministratorEmail = types.StringValue(v.(string))
				}

				if v, ok := values["credentials_json"]; ok {
					r.CredentialsJson = types.StringValue(v.(string))
				}

			}
		}
	}
}

func (r *IntegrationGoogleWorkspaceResourceModel) RefreshFromUpdateResponse(resp *shared.Connector) {
	r.RefreshFromGetResponse(resp)
}

func (r *IntegrationGoogleWorkspaceResourceModel) RefreshFromCreateResponse(resp *shared.Connector) {
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
				if v, ok := values["customer_id"]; ok {
					r.CustomerId = types.StringValue(v.(string))
				}

				if v, ok := values["domain"]; ok {
					r.Domain = types.StringValue(v.(string))
				}

				if v, ok := values["administrator_email"]; ok {
					r.AdministratorEmail = types.StringValue(v.(string))
				}

				if v, ok := values["credentials_json"]; ok {
					r.CredentialsJson = types.StringValue(v.(string))
				}

			}
		}
	}
}