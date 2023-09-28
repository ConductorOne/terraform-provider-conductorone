// Generated by tf-integration-gen. DO NOT EDIT.
package provider

import (
    "fmt"
	
	"time"

	"conductorone/internal/sdk"
	"conductorone/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

const tailscaleCatalogID = "26EAV0cMlU4hRYPBfKg42KhgJZI"

func (r *IntegrationTailscaleResourceModel) ToCreateDelegatedSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	catalogID := sdk.String(tailscaleCatalogID)
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.ConnectorServiceCreateDelegatedRequest{
		DisplayName: sdk.String("Tailscale"),
		CatalogID: catalogID,
		UserIds:   userIds,
	}
	return &out
}

func (r *IntegrationTailscaleResourceModel) ToCreateSDKType() (*shared.ConnectorServiceCreateRequest, error) {
	catalogID := sdk.String(tailscaleCatalogID)
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

func (r *IntegrationTailscaleResourceModel) ToUpdateSDKType() (*shared.Connector, bool) {
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
	    DisplayName: sdk.String("Tailscale"),
		AppID:     sdk.String(r.AppID.ValueString()),
		CatalogID: sdk.String(tailscaleCatalogID),
		ID:        sdk.String(r.ID.ValueString()),
		UserIds:   userIds,
		Config: makeConnectorConfig(configOut),
	}

	return &out, configSet
}

func (r *IntegrationTailscaleResourceModel) populateConfig() map[string]*string {
     tailscaleApiKey := new(string)
if !r.TailscaleApiKey.IsUnknown() && !r.TailscaleApiKey.IsNull() {
*tailscaleApiKey = r.TailscaleApiKey.ValueString()
} else {
tailscaleApiKey = nil
}

        tailnet := new(string)
if !r.Tailnet.IsUnknown() && !r.Tailnet.IsNull() {
*tailnet = r.Tailnet.ValueString()
} else {
tailnet = nil
}

        

    	configValues := map[string]*string{
    	"tailscale_api_key": tailscaleApiKey,
"tailnet": tailnet,

    	}

    	return configValues
}

func (r *IntegrationTailscaleResourceModel) getConfig() (map[string]string, bool) {
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

func (r *IntegrationTailscaleResourceModel) ToGetSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationTailscaleResourceModel) ToDeleteSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationTailscaleResourceModel) RefreshFromGetResponse(resp *shared.Connector) {
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
               
               if v, ok := values["tailnet"]; ok {
r.Tailnet = types.StringValue(v.(string))
}

               
           }
       }
    }
}

func (r *IntegrationTailscaleResourceModel) RefreshFromUpdateResponse(resp *shared.Connector) {
	r.RefreshFromGetResponse(resp)
}

func (r *IntegrationTailscaleResourceModel) RefreshFromCreateResponse(resp *shared.Connector) {
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
                  
                  if v, ok := values["tailnet"]; ok {
r.Tailnet = types.StringValue(v.(string))
}

                  
              }
          }
       }
}
