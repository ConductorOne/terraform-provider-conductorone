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
var _ resource.Resource = &IntegrationSalesforceResource{}
var _ resource.ResourceWithImportState = &IntegrationSalesforceResource{}

func NewIntegrationSalesforceResource() resource.Resource {
	return &IntegrationSalesforceResource{}
}

// IntegrationSalesforceResource defines the resource implementation.
type IntegrationSalesforceResource struct {
	client *sdk.ConductoroneAPI
}

// IntegrationSalesforceResourceModel describes the resource data model.
type IntegrationSalesforceResourceModel struct {
	AppID                      types.String   `tfsdk:"app_id"`
	CreatedAt                  types.String   `tfsdk:"created_at"`
	DeletedAt                  types.String   `tfsdk:"deleted_at"`
	ID                         types.String   `tfsdk:"id"`
	UpdatedAt                  types.String   `tfsdk:"updated_at"`
	UserIds                    []types.String `tfsdk:"user_ids"`
	SalesforceInstanceUrl      types.String   `tfsdk:"salesforce_instance_url"`
	SalesforceUsernameForEmail types.Bool     `tfsdk:"salesforce_username_for_email"`
}

func (r *IntegrationSalesforceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_integration_salesforce"
}

func (r *IntegrationSalesforceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Salesforce Integration Resource",

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
			"salesforce_instance_url": &schema.StringAttribute{
				Description: `Domain`,
			},
			"salesforce_username_for_email": &schema.BoolAttribute{
				Optional:    true,
				Description: `Use Salesforce usernames for email`,
			},
		},
	}
}

func (r *IntegrationSalesforceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *IntegrationSalesforceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *IntegrationSalesforceResourceModel
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

func (r *IntegrationSalesforceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *IntegrationSalesforceResourceModel
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

func (r *IntegrationSalesforceResource) get(ctx context.Context, appID string, id string) (*shared.ConnectorServiceGetResponse, error) {
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

func (r *IntegrationSalesforceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *IntegrationSalesforceResourceModel
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

func (r *IntegrationSalesforceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *IntegrationSalesforceResourceModel
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

func (r *IntegrationSalesforceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource connector.")
}
