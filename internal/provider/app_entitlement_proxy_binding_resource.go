package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	speakeasy_stringplanmodifier "github.com/conductorone/terraform-provider-conductorone/internal/planmodifiers/stringplanmodifier"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/operations"
	"github.com/conductorone/terraform-provider-conductorone/internal/validators"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppEntitlementProxyBindingResource{}
var _ resource.ResourceWithImportState = &AppEntitlementProxyBindingResource{}

func NewAppEntitlementProxyBindingResource() resource.Resource {
	return &AppEntitlementProxyBindingResource{}
}

// AppEntitlementProxyBindingResource defines the resource implementation.
type AppEntitlementProxyBindingResource struct {
	client *sdk.ConductoroneAPI
}

// AppEntitlementProxyBindingResourceModel describes the resource data model.
type AppEntitlementProxyBindingResourceModel struct {
	CreatedAt           types.String `tfsdk:"created_at"`
	DeletedAt           types.String `tfsdk:"deleted_at"`
	DstAppEntitlementID types.String `tfsdk:"dst_app_entitlement_id"`
	DstAppID            types.String `tfsdk:"dst_app_id"`
	Implicit            types.Bool   `tfsdk:"implicit"`
	SrcAppEntitlementID types.String `tfsdk:"src_app_entitlement_id"`
	SrcAppID            types.String `tfsdk:"src_app_id"`
	SystemBuiltin       types.Bool   `tfsdk:"system_builtin"`
	UpdatedAt           types.String `tfsdk:"updated_at"`
}

func (r *AppEntitlementProxyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_app_entitlement_proxy_binding"
}

func (r *AppEntitlementProxyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "AppEntitlementProxyBinding Resource",
		Attributes: map[string]schema.Attribute{
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
			"dst_app_entitlement_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(),
				},
				Description: `The dstAppEntitlementId field. Requires replacement if changed.`,
			},
			"dst_app_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(),
				},
				Description: `The dstAppId field. Requires replacement if changed.`,
			},
			"implicit": schema.BoolAttribute{
				Computed: true,
				MarkdownDescription: `If true, the binding doesn't not exist yet and is from the list of the entitlements from the parent app.` + "\n" +
					` typically the IdP that handles provisioning for the app instead of C1s connector.` + "\n" +
					`This field is part of the ` + "`" + `_implicit` + "`" + ` oneof.` + "\n" +
					`See the documentation for ` + "`" + `c1.api.app.v1.AppEntitlementProxy` + "`" + ` for more details.`,
			},
			"src_app_entitlement_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(),
				},
				Description: `Requires replacement if changed.`,
			},
			"src_app_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(),
				},
				Description: `Requires replacement if changed.`,
			},
			"system_builtin": schema.BoolAttribute{
				Computed:    true,
				Description: `The systemBuiltin field.`,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
		},
	}
}

func (r *AppEntitlementProxyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.ConductoroneAPI)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.Conductorone, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *AppEntitlementProxyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *AppEntitlementProxyBindingResourceModel
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

	srcAppID := data.SrcAppID.ValueString()

	srcAppEntitlementID := data.SrcAppEntitlementID.ValueString()

	dstAppEntitlementID := data.DstAppEntitlementID.ValueString()

	dstAppID := data.DstAppID.ValueString()

	createAppEntitlementProxyRequest := data.ToSharedCreateAppEntitlementProxyRequest()
	request := operations.C1APIAppV1AppEntitlementsProxyCreateRequest{
		SrcAppID:            srcAppID,
		SrcAppEntitlementID: srcAppEntitlementID,
		DstAppID:            dstAppID,
		DstAppEntitlementID: dstAppEntitlementID,

		CreateAppEntitlementProxyRequest: createAppEntitlementProxyRequest,
	}
	res, err := r.client.AppEntitlementsProxy.Create(ctx, request)
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
	if !(res.CreateAppEntitlementProxyResponse != nil && res.CreateAppEntitlementProxyResponse.AppEntitlementProxyView != nil && res.CreateAppEntitlementProxyResponse.AppEntitlementProxyView.AppEntitlementProxy != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAppEntitlementProxy(res.CreateAppEntitlementProxyResponse.AppEntitlementProxyView.AppEntitlementProxy)

	srcAppId1 := data.SrcAppID.ValueString()
	srcAppEntitlementId1 := data.SrcAppEntitlementID.ValueString()
	dstAppId1 := data.DstAppID.ValueString()
	dstAppEntitlementId1 := data.DstAppEntitlementID.ValueString()

	request1 := operations.C1APIAppV1AppEntitlementsProxyGetRequest{
		SrcAppID:            srcAppId1,
		SrcAppEntitlementID: srcAppEntitlementId1,
		DstAppID:            dstAppId1,
		DstAppEntitlementID: dstAppEntitlementId1,
	}
	res1, err := r.client.AppEntitlementsProxy.Get(ctx, request1)
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
	if !(res1.GetAppEntitlementProxyResponse != nil && res1.GetAppEntitlementProxyResponse.AppEntitlementProxyView != nil && res1.GetAppEntitlementProxyResponse.AppEntitlementProxyView.AppEntitlementProxy != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedAppEntitlementProxy(res1.GetAppEntitlementProxyResponse.AppEntitlementProxyView.AppEntitlementProxy)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppEntitlementProxyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *AppEntitlementProxyBindingResourceModel
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

	srcAppID := data.SrcAppID.ValueString()
	srcAppEntitlementID := data.SrcAppEntitlementID.ValueString()
	dstAppEntitlementID := data.DstAppEntitlementID.ValueString()
	dstAppID := data.DstAppID.ValueString()

	request := operations.C1APIAppV1AppEntitlementsProxyGetRequest{
		SrcAppID:            srcAppID,
		SrcAppEntitlementID: srcAppEntitlementID,
		DstAppEntitlementID: dstAppEntitlementID,
		DstAppID:            dstAppID,
	}
	res, err := r.client.AppEntitlementsProxy.Get(ctx, request)
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
	if !(res.GetAppEntitlementProxyResponse != nil && res.GetAppEntitlementProxyResponse.AppEntitlementProxyView != nil && res.GetAppEntitlementProxyResponse.AppEntitlementProxyView.AppEntitlementProxy != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAppEntitlementProxy(res.GetAppEntitlementProxyResponse.AppEntitlementProxyView.AppEntitlementProxy)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppEntitlementProxyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *AppEntitlementProxyBindingResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; all attributes marked as RequiresReplace

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppEntitlementProxyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *AppEntitlementProxyBindingResourceModel
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

	srcAppID := data.SrcAppID.ValueString()
	srcAppEntitlementID := data.SrcAppEntitlementID.ValueString()
	dstAppEntitlementID := data.DstAppEntitlementID.ValueString()
	dstAppID := data.DstAppID.ValueString()

	deleteAppEntitlementProxyRequest := data.ToSharedDeleteAppEntitlementProxyRequest()
	request := operations.C1APIAppV1AppEntitlementsProxyDeleteRequest{
		SrcAppID:                         srcAppID,
		SrcAppEntitlementID:              srcAppEntitlementID,
		DstAppID:                         dstAppID,
		DstAppEntitlementID:              dstAppEntitlementID,
		DeleteAppEntitlementProxyRequest: deleteAppEntitlementProxyRequest,
	}
	res, err := r.client.AppEntitlementsProxy.Delete(ctx, request)
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

func (r *AppEntitlementProxyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	dec := json.NewDecoder(bytes.NewReader([]byte(req.ID)))
	dec.DisallowUnknownFields()
	var data struct {
		SrcAppEntitlementID string `json:"src_app_entitlement_id"`
		SrcAppID            string `json:"src_app_id"`
		DstAppEntitlementID string `json:"dst_app_entitlement_id"`
		DstAppID            string `json:"dst_app_id"`
	}

	if err := dec.Decode(&data); err != nil {
		resp.Diagnostics.AddError("Invalid ID", `The ID is not valid. It's expected to be a JSON object alike '{ "src_app_entitlement_id": "",  "src_app_id": ""}': `+err.Error())
		return
	}

	if len(data.SrcAppEntitlementID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field src_app_entitlement_id is required but was not found in the json encoded ID. It's expected to be a value alike '""`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("src_app_entitlement_id"), data.SrcAppEntitlementID)...)
	if len(data.SrcAppID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field src_app_id is required but was not found in the json encoded ID. It's expected to be a value alike '""`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("src_app_id"), data.SrcAppID)...)
	if len(data.DstAppEntitlementID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field dst_app_entitlement_id is required but was not found in the json encoded ID. It's expected to be a value alike '""`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("dst_app_entitlement_id"), data.DstAppEntitlementID)...)
	if len(data.DstAppID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field dst_app_id is required but was not found in the json encoded ID. It's expected to be a value alike '""`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("dst_app_id"), data.DstAppID)...)

}
