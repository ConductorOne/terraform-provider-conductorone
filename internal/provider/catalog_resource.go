// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"conductorone/internal/sdk"
	"context"
	"fmt"

	"conductorone/internal/sdk/pkg/models/operations"
	"conductorone/internal/sdk/pkg/models/shared"
	"conductorone/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CatalogResource{}
var _ resource.ResourceWithImportState = &CatalogResource{}

func NewCatalogResource() resource.Resource {
	return &CatalogResource{}
}

// CatalogResource defines the resource implementation.
type CatalogResource struct {
	client *sdk.ConductoroneAPI
}

// CatalogResourceModel describes the resource data model.
type CatalogResourceModel struct {
	AccessEntitlements       []AppEntitlement                                     `tfsdk:"access_entitlements"`
	AccessEntitlementsPath   types.String                                         `tfsdk:"access_entitlements_path"`
	AppIds                   []types.String                                       `tfsdk:"app_ids"`
	AppPaths                 types.String                                         `tfsdk:"app_paths"`
	CreatedAt                types.String                                         `tfsdk:"created_at"`
	CreatedByUserID          types.String                                         `tfsdk:"created_by_user_id"`
	CreatedByUserPath        types.String                                         `tfsdk:"created_by_user_path"`
	DeletedAt                types.String                                         `tfsdk:"deleted_at"`
	Description              types.String                                         `tfsdk:"description"`
	DisplayName              types.String                                         `tfsdk:"display_name"`
	Expanded                 []RequestCatalogManagementServiceGetResponseExpanded `tfsdk:"expanded"`
	ID                       types.String                                         `tfsdk:"id"`
	Published                types.Bool                                           `tfsdk:"published"`
	RequestCatalogExpandMask *RequestCatalogExpandMask                            `tfsdk:"request_catalog_expand_mask"`
	UpdatedAt                types.String                                         `tfsdk:"updated_at"`
	VisibleToEveryone        types.Bool                                           `tfsdk:"visible_to_everyone"`
}

func (r *CatalogResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_catalog"
}

func (r *CatalogResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Catalog Resource",

		Attributes: map[string]schema.Attribute{
			"access_entitlements": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"alias": schema.StringAttribute{
							Computed:    true,
							Description: `The alias field.`,
						},
						"app_id": schema.StringAttribute{
							Computed:    true,
							Description: `The appId field.`,
						},
						"app_resource_id": schema.StringAttribute{
							Computed:    true,
							Description: `The appResourceId field.`,
						},
						"app_resource_type_id": schema.StringAttribute{
							Computed:    true,
							Description: `The appResourceTypeId field.`,
						},
						"certify_policy_id": schema.StringAttribute{
							Computed:    true,
							Description: `The certifyPolicyId field.`,
						},
						"compliance_framework_value_ids": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
							Description: `The complianceFrameworkValueIds field.`,
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
						"description": schema.StringAttribute{
							Computed:    true,
							Description: `The description field.`,
						},
						"display_name": schema.StringAttribute{
							Computed:    true,
							Description: `The displayName field.`,
						},
						"duration_grant": schema.StringAttribute{
							Computed: true,
						},
						"duration_unset": schema.SingleNestedAttribute{
							Computed:   true,
							Attributes: map[string]schema.Attribute{},
						},
						"emergency_grant_enabled": schema.BoolAttribute{
							Computed:    true,
							Description: `The emergencyGrantEnabled field.`,
						},
						"emergency_grant_policy_id": schema.StringAttribute{
							Computed:    true,
							Description: `The emergencyGrantPolicyId field.`,
						},
						"grant_count": schema.StringAttribute{
							Computed:    true,
							Description: `The grantCount field.`,
						},
						"grant_policy_id": schema.StringAttribute{
							Computed:    true,
							Description: `The grantPolicyId field.`,
						},
						"id": schema.StringAttribute{
							Computed:    true,
							Description: `The id field.`,
						},
						"provision_policy": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"connector_provision": schema.SingleNestedAttribute{
									Computed:    true,
									Attributes:  map[string]schema.Attribute{},
									Description: `The ConnectorProvision message.`,
								},
								"delegated_provision": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"app_id": schema.StringAttribute{
											Computed:    true,
											Description: `The appId field.`,
										},
										"entitlement_id": schema.StringAttribute{
											Computed:    true,
											Description: `The entitlementId field.`,
										},
									},
									Description: `The DelegatedProvision message.`,
								},
								"manual_provision": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"instructions": schema.StringAttribute{
											Computed:    true,
											Description: `The instructions field.`,
										},
										"user_ids": schema.ListAttribute{
											Computed:    true,
											ElementType: types.StringType,
											Description: `The userIds field.`,
										},
									},
									Description: `The ManualProvision message.`,
								},
							},
							MarkdownDescription: `The ProvisionPolicy message.` + "\n" +
								`` +
								`This message contains a oneof. Only a single field of the following list may be set at a time:` + "\n" +
								`  - connector` + "\n" +
								`  - manual` + "\n" +
								`  - delegated` + "\n" +
								"\n" +
								``,
						},
						"revoke_policy_id": schema.StringAttribute{
							Computed:    true,
							Description: `The revokePolicyId field.`,
						},
						"risk_level_value_id": schema.StringAttribute{
							Computed:    true,
							Description: `The riskLevelValueId field.`,
						},
						"slug": schema.StringAttribute{
							Computed:    true,
							Description: `The slug field.`,
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
				},
				Description: `The accessEntitlements field.`,
			},
			"access_entitlements_path": schema.StringAttribute{
				Computed:    true,
				Description: `The accessEntitlementsPath field.`,
			},
			"app_ids": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `The appIds field.`,
			},
			"app_paths": schema.StringAttribute{
				Computed:    true,
				Description: `The appPaths field.`,
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
			"created_by_user_id": schema.StringAttribute{
				Computed:    true,
				Description: `The createdByUserId field.`,
			},
			"created_by_user_path": schema.StringAttribute{
				Computed:    true,
				Description: `The createdByUserPath field.`,
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The description field.`,
			},
			"display_name": schema.StringAttribute{
				Required:    true,
				Description: `The displayName field.`,
			},
			"expanded": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"at_type": schema.StringAttribute{
							Computed:    true,
							Description: `The type of the serialized message.`,
						},
					},
				},
				Description: `The expanded field.`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `The id field.`,
			},
			"published": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The published field is used to determine if the catalog is active and visible to users.`,
			},
			"request_catalog_expand_mask": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"paths": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `The paths field.`,
					},
				},
				Description: `The RequestCatalogExpandMask message.`,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"visible_to_everyone": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The visibleToEveryone field is used to determine if this catalog should be visible to all users.`,
			},
		},
	}
}

func (r *CatalogResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *CatalogResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *CatalogResourceModel
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

	request := *data.ToCreateSDKType()
	res, err := r.client.RequestCatalogManagement.Create(ctx, request)
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
	if res.RequestCatalogManagementServiceGetResponse.RequestCatalogView == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.RequestCatalogManagementServiceGetResponse.RequestCatalogView.RequestCatalog)
	if res.RequestCatalogManagementServiceGetResponse.RequestCatalogView.AccessEntitlementsPath != nil {
		data.AccessEntitlementsPath = types.StringValue(*res.RequestCatalogManagementServiceGetResponse.RequestCatalogView.AccessEntitlementsPath)
	} else {
		data.AccessEntitlementsPath = types.StringNull()
	}
	if res.RequestCatalogManagementServiceGetResponse.RequestCatalogView.AppPaths != nil {
		data.AppPaths = types.StringValue(*res.RequestCatalogManagementServiceGetResponse.RequestCatalogView.AppPaths)
	} else {
		data.AppPaths = types.StringNull()
	}
	if res.RequestCatalogManagementServiceGetResponse.RequestCatalogView.CreatedByUserPath != nil {
		data.CreatedByUserPath = types.StringValue(*res.RequestCatalogManagementServiceGetResponse.RequestCatalogView.CreatedByUserPath)
	} else {
		data.CreatedByUserPath = types.StringNull()
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CatalogResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *CatalogResourceModel
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

	id := data.ID.ValueString()
	request := operations.C1APIRequestcatalogV1RequestCatalogManagementServiceGetRequest{
		ID: id,
	}
	res, err := r.client.RequestCatalogManagement.Get(ctx, request)
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
	if res.RequestCatalogManagementServiceGetResponse.RequestCatalogView == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	if res.RequestCatalogManagementServiceGetResponse.RequestCatalogView.RequestCatalog == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.RequestCatalogManagementServiceGetResponse.RequestCatalogView.RequestCatalog)
	if res.RequestCatalogManagementServiceGetResponse.RequestCatalogView.AccessEntitlementsPath != nil {
		data.AccessEntitlementsPath = types.StringValue(*res.RequestCatalogManagementServiceGetResponse.RequestCatalogView.AccessEntitlementsPath)
	} else {
		data.AccessEntitlementsPath = types.StringNull()
	}
	if res.RequestCatalogManagementServiceGetResponse.RequestCatalogView.AppPaths != nil {
		data.AppPaths = types.StringValue(*res.RequestCatalogManagementServiceGetResponse.RequestCatalogView.AppPaths)
	} else {
		data.AppPaths = types.StringNull()
	}
	if res.RequestCatalogManagementServiceGetResponse.RequestCatalogView.CreatedByUserPath != nil {
		data.CreatedByUserPath = types.StringValue(*res.RequestCatalogManagementServiceGetResponse.RequestCatalogView.CreatedByUserPath)
	} else {
		data.CreatedByUserPath = types.StringNull()
	}

	if res.RequestCatalogManagementServiceGetResponse.RequestCatalogView.RequestCatalog.DeletedAt != nil {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CatalogResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *CatalogResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	var requestCatalogManagementServiceUpdateRequest *shared.RequestCatalogManagementServiceUpdateRequest
	requestCatalog := data.ToUpdateSDKType()

	requestCatalogManagementServiceUpdateRequest = &shared.RequestCatalogManagementServiceUpdateRequest{
		RequestCatalog: requestCatalog,
	}
	id := data.ID.ValueString()
	request := operations.C1APIRequestcatalogV1RequestCatalogManagementServiceUpdateRequest{
		RequestCatalogManagementServiceUpdateRequest: requestCatalogManagementServiceUpdateRequest,
		ID: id,
	}
	res, err := r.client.RequestCatalogManagement.Update(ctx, request)
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
	if res.RequestCatalogManagementServiceGetResponse.RequestCatalogView == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromUpdateResponse(res.RequestCatalogManagementServiceGetResponse.RequestCatalogView.RequestCatalog)
	if res.RequestCatalogManagementServiceGetResponse.RequestCatalogView.AccessEntitlementsPath != nil {
		data.AccessEntitlementsPath = types.StringValue(*res.RequestCatalogManagementServiceGetResponse.RequestCatalogView.AccessEntitlementsPath)
	} else {
		data.AccessEntitlementsPath = types.StringNull()
	}
	if res.RequestCatalogManagementServiceGetResponse.RequestCatalogView.AppPaths != nil {
		data.AppPaths = types.StringValue(*res.RequestCatalogManagementServiceGetResponse.RequestCatalogView.AppPaths)
	} else {
		data.AppPaths = types.StringNull()
	}
	if res.RequestCatalogManagementServiceGetResponse.RequestCatalogView.CreatedByUserPath != nil {
		data.CreatedByUserPath = types.StringValue(*res.RequestCatalogManagementServiceGetResponse.RequestCatalogView.CreatedByUserPath)
	} else {
		data.CreatedByUserPath = types.StringNull()
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CatalogResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *CatalogResourceModel
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

	requestCatalogManagementServiceDeleteRequest := data.ToDeleteSDKType()
	id := data.ID.ValueString()
	request := operations.C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteRequest{
		RequestCatalogManagementServiceDeleteRequest: requestCatalogManagementServiceDeleteRequest,
		ID: id,
	}
	res, err := r.client.RequestCatalogManagement.Delete(ctx, request)
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

func (r *CatalogResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
