// Generated by tf-integration-gen. DO NOT EDIT.
package provider

import (
	"context"
	"fmt"

	"conductorone/internal/sdk"
	"conductorone/internal/sdk/pkg/models/operations"

	"conductorone/internal/sdk/pkg/models/shared"
	"conductorone/internal/validators"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &IntegrationDatabricksResource{}
var _ resource.ResourceWithImportState = &IntegrationDatabricksResource{}

func NewIntegrationDatabricksResource() resource.Resource {
	return &IntegrationDatabricksResource{}
}

// IntegrationDatabricksResource defines the resource implementation.
type IntegrationDatabricksResource struct {
	client *sdk.ConductoroneAPI
}

// IntegrationDatabricksResourceModel describes the resource data model.
type IntegrationDatabricksResourceModel struct {
	AppID                  types.String   `tfsdk:"app_id"`
	CreatedAt              types.String   `tfsdk:"created_at"`
	DeletedAt              types.String   `tfsdk:"deleted_at"`
	ID                     types.String   `tfsdk:"id"`
	UpdatedAt              types.String   `tfsdk:"updated_at"`
	UserIds                []types.String `tfsdk:"user_ids"`
	DatabricksAccountId    types.String   `tfsdk:"databricks_account_id"`
	DatabricksClientId     types.String   `tfsdk:"databricks_client_id"`
	DatabricksClientSecret types.String   `tfsdk:"databricks_client_secret"`
	DatabricksAccessToken  types.String   `tfsdk:"databricks_access_token"`
	DatabricksWorkspace    types.String   `tfsdk:"databricks_workspace"`
	DatabricksUsername     types.String   `tfsdk:"databricks_username"`
	DatabricksPassword     types.String   `tfsdk:"databricks_password"`
}

func (r *IntegrationDatabricksResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_integration_databricks"
}

func (r *IntegrationDatabricksResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Databricks Integration Resource",

		Attributes: map[string]schema.Attribute{
			"app_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Required:    true,
				Description: `The ID for the Application that this integration should connected to.`,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
				Description: `The time this integration was created.`,
			},
			"deleted_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
				Description: `The time this integration was deleted.`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `The ID of this integration.`,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
				Description: `The time this integration was last updated.`,
			},
			"user_ids": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `A list of user IDs of who owns this integration. It defaults to the user who created the integration.`,
			},
			"databricks_account_id": &schema.StringAttribute{
				Description: `Account ID`,
			},
			"databricks_client_id": &schema.StringAttribute{
				Description: `OAuth2 Client ID`,
			},
			"databricks_client_secret": &schema.StringAttribute{
				Sensitive:   true,
				Description: `OAuth2 Client Secret`,
			},
			"databricks_access_token": &schema.StringAttribute{
				Sensitive:   true,
				Description: `Personal Access Token`,
			},
			"databricks_workspace": &schema.StringAttribute{
				Description: `Workspace ID`,
			},
			"databricks_username": &schema.StringAttribute{
				Description: `Username`,
			},
			"databricks_password": &schema.StringAttribute{
				Sensitive:   true,
				Description: `Password`,
			},
		},
	}
}

func (r *IntegrationDatabricksResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.ConductoroneAPI)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *IntegrationDatabricksResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *IntegrationDatabricksResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, configSet := data.getConfig()

	var configResp *shared.Connector

	if configSet {
		connectorServiceCreateRequest, err := data.ToCreateSDKType()
		if err != nil {
			resp.Diagnostics.AddError("failure to make create request", err.Error())
			return
		}
		appID := data.AppID.ValueString()
		request := operations.C1APIAppV1ConnectorServiceCreateRequest{
			ConnectorServiceCreateRequest: connectorServiceCreateRequest,
			AppID:                         appID,
		}
		res, err := r.client.Connector.Create(ctx, request)
		if err != nil {
			resp.Diagnostics.AddError("failure to invoke API", err.Error())
			return
		}
		if res == nil {
			resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
			return
		}
		if res.StatusCode != 200 {
			resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
			return
		}
		if res.ConnectorServiceCreateResponse.ConnectorView == nil {
			resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
			return
		}
		configResp = res.ConnectorServiceCreateResponse.ConnectorView.Connector
	} else {
		connectorServiceCreateDelegatedRequest := data.ToCreateDelegatedSDKType()
		appID := data.AppID.ValueString()
		request := operations.C1APIAppV1ConnectorServiceCreateDelegatedRequest{
			ConnectorServiceCreateDelegatedRequest: connectorServiceCreateDelegatedRequest,
			AppID:                                  appID,
		}
		res, err := r.client.Connector.CreateDelegated(ctx, request)
		if err != nil {
			resp.Diagnostics.AddError("failure to invoke API", err.Error())
			return
		}
		if res == nil {
			resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
			return
		}
		if res.StatusCode != 200 {
			resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
			return
		}
		if res.ConnectorServiceCreateResponse.ConnectorView == nil {
			resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
			return
		}
		configResp = res.ConnectorServiceCreateResponse.ConnectorView.Connector
	}
	data.RefreshFromCreateResponse(configResp)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IntegrationDatabricksResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *IntegrationDatabricksResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.get(ctx, data.AppID.ValueString(), data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}

	if res.ConnectorView.Connector.DeletedAt != nil {
		resp.State.RemoveResource(ctx)
		return
	}

	data.RefreshFromGetResponse(res.ConnectorView.Connector)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IntegrationDatabricksResource) get(ctx context.Context, appID string, id string) (*shared.ConnectorServiceGetResponse, error) {
	request := operations.C1APIAppV1ConnectorServiceGetRequest{
		AppID: appID,
		ID:    id,
	}
	res, err := r.client.Connector.Get(ctx, request)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("unexpected response from API")
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected response from API. Got an unexpected response code %v", res.StatusCode)
	}
	if res.ConnectorServiceGetResponse == nil {
		return nil, fmt.Errorf("unexpected response from API. No response body")
	}
	return res.ConnectorServiceGetResponse, nil
}

func (r *IntegrationDatabricksResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *IntegrationDatabricksResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	appID := data.AppID.ValueString()

	updateCon, configSet := data.ToUpdateSDKType()
	if configSet {
		configReq := operations.C1APIAppV1ConnectorServiceUpdateRequest{
			ConnectorServiceUpdateRequest: &shared.ConnectorServiceUpdateRequest{
				Connector:  updateCon,
				UpdateMask: "config",
			},
			AppID: appID,
			ID:    data.ID.ValueString(),
		}
		updateRes, err := r.client.Connector.Update(ctx, configReq)
		if err != nil {
			resp.Diagnostics.AddError("failure to invoke API", err.Error())
			return
		}
		if updateRes == nil {
			resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", updateRes))
			return
		}
		if updateRes.StatusCode != 200 {
			resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", updateRes.StatusCode), debugResponse(updateRes.RawResponse))
			return
		}
		data.RefreshFromUpdateResponse(updateRes.ConnectorServiceUpdateResponse.ConnectorView.Connector)
	} else {
		configReq := operations.C1APIAppV1ConnectorServiceUpdateDelegatedRequest{
			ConnectorServiceUpdateDelegatedRequest: &shared.ConnectorServiceUpdateDelegatedRequest{
				Connector:  updateCon,
				UpdateMask: "displayName,userIds",
			},
			ConnectorAppID: appID,
			ConnectorID:    data.ID.ValueString(),
		}
		updateRes, err := r.client.Connector.UpdateDelegated(ctx, configReq)
		if err != nil {
			resp.Diagnostics.AddError("failure to invoke API", err.Error())
			return
		}
		if updateRes == nil {
			resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", updateRes))
			return
		}
		if updateRes.StatusCode != 200 {
			resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", updateRes.StatusCode), debugResponse(updateRes.RawResponse))
			return
		}
		data.RefreshFromUpdateResponse(updateRes.ConnectorServiceUpdateResponse.ConnectorView.Connector)
	}
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IntegrationDatabricksResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *IntegrationDatabricksResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	connectorServiceDeleteRequest := &shared.ConnectorServiceDeleteRequest{}
	appID := data.AppID.ValueString()
	id := data.ID.ValueString()
	request := operations.C1APIAppV1ConnectorServiceDeleteRequest{
		ConnectorServiceDeleteRequest: connectorServiceDeleteRequest,
		AppID:                         appID,
		ID:                            id,
	}
	res, err := r.client.Connector.Delete(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
}

func (r *IntegrationDatabricksResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource connector.")
}
