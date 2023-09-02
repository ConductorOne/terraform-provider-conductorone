package provider

import (
	"conductorone/internal/sdk"
	"conductorone/internal/sdk/pkg/models/operations"
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppEntitlementOwnerResource{}
var _ resource.ResourceWithImportState = &AppEntitlementOwnerResource{}

func NewAppEntitlementOwnerResource() resource.Resource {
	return &AppEntitlementOwnerResource{}
}

// AppEntitlementOwnerResource defines the resource implementation.
type AppEntitlementOwnerResource struct {
	client *sdk.ConductoroneAPI
}

// AppEntitlementOwnerResourceModel describes the resource data model.
type AppEntitlementOwnerResourceModel struct {
	AppID         types.String   `tfsdk:"app_id"`
	EntitlementID types.String   `tfsdk:"entitlement_id"`
	UserIds       []types.String `tfsdk:"user_ids"`
}

func (r *AppEntitlementOwnerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_app_entitlement_owner"
}

func (r *AppEntitlementOwnerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "AppEntitlementOwner Resource",

		Attributes: map[string]schema.Attribute{
			"app_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Required: true,
			},
			"entitlement_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Required: true,
			},
			"user_ids": schema.ListAttribute{
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
				Description: `The user_ids field for the users to set as an owner of the app entitlement.`,
			},
		},
	}
}

func (r *AppEntitlementOwnerResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *AppEntitlementOwnerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *AppEntitlementOwnerResourceModel
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

	setAppEntitlementOwnersRequest := data.ToCreateSDKType()
	appID := data.AppID.ValueString()
	entitlementID := data.EntitlementID.ValueString()
	request := operations.C1APIAppV1AppEntitlementOwnersSetRequest{
		SetAppEntitlementOwnersRequest: setAppEntitlementOwnersRequest,
		AppID:                          appID,
		EntitlementID:                  entitlementID,
	}
	res, err := r.client.AppEntitlementOwners.C1APIAppV1AppEntitlementOwnersSet(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
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
	if res.SetAppEntitlementOwnersResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.SetAppEntitlementOwnersResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppEntitlementOwnerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *AppEntitlementOwnerResourceModel
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

	listResp, err := r.client.AppEntitlementOwners.C1APIAppV1AppEntitlementOwnersList(ctx, operations.C1APIAppV1AppEntitlementOwnersListRequest{
		AppID:         data.AppID.ValueString(),
		EntitlementID: data.EntitlementID.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke list API", err.Error())
		return
	}
	if listResp == nil {
		resp.Diagnostics.AddError("unexpected response from list API", fmt.Sprintf("%v", listResp))
		return
	}
	if listResp.ListAppEntitlementOwnersResponse == nil || listResp.ListAppEntitlementOwnersResponse.List == nil {
		resp.Diagnostics.AddError("unexpected response from list API", fmt.Sprintf("%v", listResp))
		return
	}

	data.RefreshFromReadResponse(listResp.ListAppEntitlementOwnersResponse.List)
	// Not Implemented

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppEntitlementOwnerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *AppEntitlementOwnerResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; all attributes marked as RequiresReplace

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppEntitlementOwnerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *AppEntitlementOwnerResourceModel
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

	listResp, err := r.client.AppEntitlementOwners.C1APIAppV1AppEntitlementOwnersList(ctx, operations.C1APIAppV1AppEntitlementOwnersListRequest{
		AppID:         data.AppID.ValueString(),
		EntitlementID: data.EntitlementID.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke list API", err.Error())
		return
	}
	if listResp == nil {
		resp.Diagnostics.AddError("unexpected response from list API", fmt.Sprintf("%v", listResp))
		return
	}

	setAppEntitlementOwnersRequest := data.ToDeleteSDKType(listResp.ListAppEntitlementOwnersResponse.List)
	appID := data.AppID.ValueString()
	entitlementID := data.EntitlementID.ValueString()
	request := operations.C1APIAppV1AppEntitlementOwnersSetRequest{
		SetAppEntitlementOwnersRequest: setAppEntitlementOwnersRequest,
		AppID:                          appID,
		EntitlementID:                  entitlementID,
	}
	res, err := r.client.AppEntitlementOwners.C1APIAppV1AppEntitlementOwnersSet(ctx, request)
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

func (r *AppEntitlementOwnerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource app_entitlement_owner.")
}
