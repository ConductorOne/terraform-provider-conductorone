package provider

import (
	"context"
	"fmt"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &RequestCatalogResource{}
var _ resource.ResourceWithImportState = &RequestCatalogResource{}

func NewRequestCatalogResource() resource.Resource {
	return &RequestCatalogResource{}
}

// RequestCatalogResource defines the resource implementation.
type RequestCatalogResource struct {
	// Provider configured SDK client.
	client *sdk.ConductoroneAPI
}

// RequestCatalogResourceModel describes the resource data model.
type RequestCatalogResourceModel struct {
	CreatedAt                       types.String `tfsdk:"created_at"`
	CreatedByUserID                 types.String `tfsdk:"created_by_user_id"`
	DeletedAt                       types.String `tfsdk:"-"`
	Description                     types.String `tfsdk:"description"`
	DisplayName                     types.String `tfsdk:"display_name"`
	EnrollmentBehavior              types.String `tfsdk:"enrollment_behavior"`
	ID                              types.String `tfsdk:"id"`
	Published                       types.Bool   `tfsdk:"published"`
	RequestBundle                   types.Bool   `tfsdk:"request_bundle"`
	UnenrollmentBehavior            types.String `tfsdk:"unenrollment_behavior"`
	UnenrollmentEntitlementBehavior types.String `tfsdk:"unenrollment_entitlement_behavior"`
	UpdatedAt                       types.String `tfsdk:"updated_at"`
	VisibleToEveryone               types.Bool   `tfsdk:"visible_to_everyone"`
}

func (r *RequestCatalogResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_request_catalog"
}

func (r *RequestCatalogResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "RequestCatalog Resource",
		Attributes: map[string]schema.Attribute{
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"created_by_user_id": schema.StringAttribute{
				Computed:    true,
				Description: `The id of the user this request catalog was created by.`,
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The description of the request catalog.`,
			},
			"display_name": schema.StringAttribute{
				Required:    true,
				Description: `The display name of the request catalog.`,
			},
			"enrollment_behavior": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Defines how to handle the request policies of the entitlements in the catalog during enrollment.`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `The id of the request catalog.`,
			},
			"published": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Whether or not this catalog is published.`,
			},
			"request_bundle": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Whether all the entitlements in the catalog can be requested at once. Your tenant must have the bundles feature to use this.`,
			},
			"unenrollment_behavior": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Defines how to handle the revocation of the entitlements in the catalog during unenrollment.`,
			},
			"unenrollment_entitlement_behavior": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Defines how to handle the revoke policies of the entitlements in the catalog during unenrollment.`,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
			"visible_to_everyone": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `If this is true, the access entitlement requirement is ignored.`,
			},
		},
	}
}

func (r *RequestCatalogResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.ConductoroneAPI)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.ConductoroneAPI, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *RequestCatalogResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *RequestCatalogResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request, requestDiags := data.ToSharedRequestCatalogManagementServiceCreateRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.RequestCatalogManagement.Create(ctx, request)
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
	if !(res.RequestCatalogManagementServiceGetResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedRequestCatalogManagementServiceGetResponse(ctx, res.RequestCatalogManagementServiceGetResponse)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(refreshPlan(ctx, plan, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RequestCatalogResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *RequestCatalogResourceModel
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

	request, requestDiags := data.ToOperationsC1APIRequestcatalogV1RequestCatalogManagementServiceGetRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.RequestCatalogManagement.Get(ctx, *request)
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
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.RequestCatalogManagementServiceGetResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedRequestCatalogManagementServiceGetResponse(ctx, res.RequestCatalogManagementServiceGetResponse)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if !data.DeletedAt.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RequestCatalogResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *RequestCatalogResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	request, requestDiags := data.ToOperationsC1APIRequestcatalogV1RequestCatalogManagementServiceUpdateRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.RequestCatalogManagement.Update(ctx, *request)
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
	if !(res.RequestCatalogManagementServiceGetResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedRequestCatalogManagementServiceGetResponse(ctx, res.RequestCatalogManagementServiceGetResponse)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(refreshPlan(ctx, plan, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	request1, request1Diags := data.ToOperationsC1APIRequestcatalogV1RequestCatalogManagementServiceGetRequest(ctx)
	resp.Diagnostics.Append(request1Diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res1, err := r.client.RequestCatalogManagement.Get(ctx, *request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if !(res1.RequestCatalogManagementServiceGetResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedRequestCatalogManagementServiceGetResponse(ctx, res1.RequestCatalogManagementServiceGetResponse)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(refreshPlan(ctx, plan, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RequestCatalogResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *RequestCatalogResourceModel
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

	request, requestDiags := data.ToOperationsC1APIRequestcatalogV1RequestCatalogManagementServiceDeleteRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.RequestCatalogManagement.Delete(ctx, *request)
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

}

func (r *RequestCatalogResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
