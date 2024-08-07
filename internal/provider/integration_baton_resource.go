// Not generated. Written by jirwin.

package provider

import (
	"context"
	"fmt"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/operations"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"github.com/conductorone/terraform-provider-conductorone/internal/validators"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &IntegrationBatonResource{}
var _ resource.ResourceWithImportState = &IntegrationBatonResource{}

func NewIntegrationBatonResource() resource.Resource {
	return &IntegrationBatonResource{}
}

// IntegrationBatonResource defines the resource implementation.
type IntegrationBatonResource struct {
	client *sdk.ConductoroneAPI
}

// IntegrationBatonResourceModel describes the resource data model.
type IntegrationBatonResourceModel struct {
	AppID     types.String   `tfsdk:"app_id"`
	CreatedAt types.String   `tfsdk:"created_at"`
	DeletedAt types.String   `tfsdk:"deleted_at"`
	ID        types.String   `tfsdk:"id"`
	UpdatedAt types.String   `tfsdk:"updated_at"`
	UserIds   []types.String `tfsdk:"user_ids"`
}

func (r *IntegrationBatonResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_integration_baton"
}

func (r *IntegrationBatonResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Connector Resource",

		Attributes: map[string]schema.Attribute{
			"app_id": schema.StringAttribute{
				Required:    true,
				Description: `The appId field.`,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"deleted_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `The id field.`,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"user_ids": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `The userIds field.`,
			},
		},
	}
}

func (r *IntegrationBatonResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *IntegrationBatonResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *IntegrationBatonResourceModel
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

	connectorServiceCreateDelegatedRequest := data.ToCreateSDKType()
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
	data.RefreshFromCreateResponse(res.ConnectorServiceCreateResponse.ConnectorView.Connector)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IntegrationBatonResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *IntegrationBatonResourceModel
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

	appID := data.AppID.ValueString()
	id := data.ID.ValueString()
	request := operations.C1APIAppV1ConnectorServiceGetRequest{
		AppID: appID,
		ID:    id,
	}
	res, err := r.client.Connector.Get(ctx, request)
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
	if res.ConnectorServiceGetResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}

	if res.ConnectorServiceGetResponse.ConnectorView.Connector.DeletedAt != nil {
		resp.State.RemoveResource(ctx)
		return
	}

	data.RefreshFromGetResponse(res.ConnectorServiceGetResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IntegrationBatonResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *IntegrationBatonResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}
	appID := data.AppID.ValueString()

	updateCon, _ := data.ToUpdateSDKType()

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
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IntegrationBatonResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *IntegrationBatonResourceModel
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

func (r *IntegrationBatonResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource connector.")
}
