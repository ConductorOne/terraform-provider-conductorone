// Generated by tf-integration-gen. DO NOT EDIT.
package provider

import (
    "fmt"
	
	"time"

	"conductorone/internal/sdk"
	"conductorone/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

const pantherCatalogID = "2QtUSl96IHLuymkQgMWpdxyvcou"

func (r *IntegrationPantherResourceModel) ToCreateDelegatedSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	catalogID := sdk.String(pantherCatalogID)
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.ConnectorServiceCreateDelegatedRequest{
		DisplayName: sdk.String("Panther"),
		CatalogID: catalogID,
		UserIds:   userIds,
	}
	return &out
}

func (r *IntegrationPantherResourceModel) ToCreateSDKType() (*shared.ConnectorServiceCreateRequest, error) {
	catalogID := sdk.String(pantherCatalogID)
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

func (r *IntegrationPantherResourceModel) ToUpdateSDKType() (*shared.Connector, bool) {
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
	    DisplayName: sdk.String("Panther"),
		AppID:     sdk.String(r.AppID.ValueString()),
		CatalogID: sdk.String(pantherCatalogID),
		ID:        sdk.String(r.ID.ValueString()),
		UserIds:   userIds,
		Config: makeConnectorConfig(configOut),
	}

	return &out, configSet
}

func (r *IntegrationPantherResourceModel) populateConfig() map[string]*string {
     pantherApiKey := new(string)
if !r.PantherApiKey.IsUnknown() && !r.PantherApiKey.IsNull() {
*pantherApiKey = r.PantherApiKey.ValueString()
} else {
pantherApiKey = nil
}

        pantherUrl := new(string)
if !r.PantherUrl.IsUnknown() && !r.PantherUrl.IsNull() {
*pantherUrl = r.PantherUrl.ValueString()
} else {
pantherUrl = nil
}

        

    	configValues := map[string]*string{
    	"panther_api_key": pantherApiKey,
"panther_url": pantherUrl,

    	}

    	return configValues
}

func (r *IntegrationPantherResourceModel) getConfig() (map[string]string, bool) {
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

func (r *IntegrationPantherResourceModel) ToGetSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationPantherResourceModel) ToDeleteSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationPantherResourceModel) RefreshFromGetResponse(resp *shared.Connector) {
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

    
}

func (r *IntegrationPantherResourceModel) RefreshFromUpdateResponse(resp *shared.Connector) {
	r.RefreshFromGetResponse(resp)
}

func (r *IntegrationPantherResourceModel) RefreshFromCreateResponse(resp *shared.Connector) {
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

   
}
