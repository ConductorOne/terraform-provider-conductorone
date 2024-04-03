// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/speakeasy/terraform-provider-terraform/internal/provider/types"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/models/operations"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/models/shared"
	"github.com/speakeasy/terraform-provider-terraform/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppResource{}
var _ resource.ResourceWithImportState = &AppResource{}

func NewAppResource() resource.Resource {
	return &AppResource{}
}

// AppResource defines the resource implementation.
type AppResource struct {
	client *sdk.SDK
}

// AppResourceModel describes the resource data model.
type AppResourceModel struct {
	AppAccountID     types.String   `tfsdk:"app_account_id"`
	AppAccountName   types.String   `tfsdk:"app_account_name"`
	AppOwners        []tfTypes.User `tfsdk:"app_owners"`
	CertifyPolicyID  types.String   `tfsdk:"certify_policy_id"`
	CreatedAt        types.String   `tfsdk:"created_at"`
	DeleteAppRequest *tfTypes.Three `tfsdk:"delete_app_request"`
	DeletedAt        types.String   `tfsdk:"deleted_at"`
	Description      types.String   `tfsdk:"description"`
	DisplayName      types.String   `tfsdk:"display_name"`
	FieldMask        types.String   `tfsdk:"field_mask"`
	GrantPolicyID    types.String   `tfsdk:"grant_policy_id"`
	IconURL          types.String   `tfsdk:"icon_url"`
	ID               types.String   `tfsdk:"id"`
	IdentityMatching types.String   `tfsdk:"identity_matching"`
	IsDirectory      types.Bool     `tfsdk:"is_directory"`
	LogoURI          types.String   `tfsdk:"logo_uri"`
	MonthlyCostUsd   types.Int64    `tfsdk:"monthly_cost_usd"`
	Owners           []types.String `tfsdk:"owners"`
	ParentAppID      types.String   `tfsdk:"parent_app_id"`
	RevokePolicyID   types.String   `tfsdk:"revoke_policy_id"`
	UpdatedAt        types.String   `tfsdk:"updated_at"`
	UserCount        types.String   `tfsdk:"user_count"`
}

func (r *AppResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_app"
}

func (r *AppResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "App Resource",

		Attributes: map[string]schema.Attribute{
			"app_account_id": schema.StringAttribute{
				Computed:    true,
				Description: `The ID of the Account named by AccountName.`,
			},
			"app_account_name": schema.StringAttribute{
				Computed:    true,
				Description: `The AccountName of the app. For example, AWS is AccountID, Github is Org Name, and Okta is Okta Subdomain.`,
			},
			"app_owners": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"created_at": schema.StringAttribute{
							Computed: true,
							Validators: []validator.String{
								validators.IsRFC3339(),
							},
						},
						"delegated_user_id": schema.StringAttribute{
							Computed:    true,
							Description: `The id of the user to whom tasks will be automatically reassigned to.`,
						},
						"deleted_at": schema.StringAttribute{
							Computed: true,
							Validators: []validator.String{
								validators.IsRFC3339(),
							},
						},
						"department": schema.StringAttribute{
							Computed:    true,
							Description: `The department which the user belongs to in the organization.`,
						},
						"department_sources": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"app_id": schema.StringAttribute{
										Computed:    true,
										Description: `The appId field.`,
									},
									"app_user_id": schema.StringAttribute{
										Computed:    true,
										Description: `The appUserId field.`,
									},
									"app_user_profile_attribute_key": schema.StringAttribute{
										Computed:    true,
										Description: `The appUserProfileAttributeKey field.`,
									},
									"user_attribute_mapping_id": schema.StringAttribute{
										Computed:    true,
										Description: `The userAttributeMappingId field.`,
									},
									"value": schema.StringAttribute{
										Computed:    true,
										Description: `The value field.`,
									},
								},
							},
							Description: `A list of objects mapped based on department attribute mappings configured in the system.`,
						},
						"directory_ids": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
							Description: `A list of unique ids that represent different directories.`,
						},
						"directory_status": schema.StringAttribute{
							Computed:    true,
							Description: `The status of the user in the directory. must be one of ["UNKNOWN", "ENABLED", "DISABLED", "DELETED"]`,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"UNKNOWN",
									"ENABLED",
									"DISABLED",
									"DELETED",
								),
							},
						},
						"directory_status_sources": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"app_id": schema.StringAttribute{
										Computed:    true,
										Description: `The appId field.`,
									},
									"app_user_id": schema.StringAttribute{
										Computed:    true,
										Description: `The appUserId field.`,
									},
									"app_user_profile_attribute_key": schema.StringAttribute{
										Computed:    true,
										Description: `The appUserProfileAttributeKey field.`,
									},
									"user_attribute_mapping_id": schema.StringAttribute{
										Computed:    true,
										Description: `The userAttributeMappingId field.`,
									},
									"value": schema.StringAttribute{
										Computed:    true,
										Description: `The value field.`,
									},
								},
							},
							Description: `A list of objects mapped based on directoryStatus attribute mappings configured in the system.`,
						},
						"display_name": schema.StringAttribute{
							Computed:    true,
							Description: `The display name of the user.`,
						},
						"email": schema.StringAttribute{
							Computed:    true,
							Description: `This is the user's email.`,
						},
						"emails": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
							Description: `This is a list of all of the user's emails from app users.`,
						},
						"employment_status": schema.StringAttribute{
							Computed:    true,
							Description: `The users employment status.`,
						},
						"employment_status_sources": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"app_id": schema.StringAttribute{
										Computed:    true,
										Description: `The appId field.`,
									},
									"app_user_id": schema.StringAttribute{
										Computed:    true,
										Description: `The appUserId field.`,
									},
									"app_user_profile_attribute_key": schema.StringAttribute{
										Computed:    true,
										Description: `The appUserProfileAttributeKey field.`,
									},
									"user_attribute_mapping_id": schema.StringAttribute{
										Computed:    true,
										Description: `The userAttributeMappingId field.`,
									},
									"value": schema.StringAttribute{
										Computed:    true,
										Description: `The value field.`,
									},
								},
							},
							Description: `A list of objects mapped based on employmentStatus attribute mappings configured in the system.`,
						},
						"employment_type": schema.StringAttribute{
							Computed:    true,
							Description: `The employment type of the user.`,
						},
						"employment_type_sources": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"app_id": schema.StringAttribute{
										Computed:    true,
										Description: `The appId field.`,
									},
									"app_user_id": schema.StringAttribute{
										Computed:    true,
										Description: `The appUserId field.`,
									},
									"app_user_profile_attribute_key": schema.StringAttribute{
										Computed:    true,
										Description: `The appUserProfileAttributeKey field.`,
									},
									"user_attribute_mapping_id": schema.StringAttribute{
										Computed:    true,
										Description: `The userAttributeMappingId field.`,
									},
									"value": schema.StringAttribute{
										Computed:    true,
										Description: `The value field.`,
									},
								},
							},
							Description: `A list of objects mapped based on employmentType attribute mappings configured in the system.`,
						},
						"id": schema.StringAttribute{
							Computed:    true,
							Description: `A unique identifier of the user.`,
						},
						"job_title": schema.StringAttribute{
							Computed:    true,
							Description: `The job title of the user.`,
						},
						"job_title_sources": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"app_id": schema.StringAttribute{
										Computed:    true,
										Description: `The appId field.`,
									},
									"app_user_id": schema.StringAttribute{
										Computed:    true,
										Description: `The appUserId field.`,
									},
									"app_user_profile_attribute_key": schema.StringAttribute{
										Computed:    true,
										Description: `The appUserProfileAttributeKey field.`,
									},
									"user_attribute_mapping_id": schema.StringAttribute{
										Computed:    true,
										Description: `The userAttributeMappingId field.`,
									},
									"value": schema.StringAttribute{
										Computed:    true,
										Description: `The value field.`,
									},
								},
							},
							Description: `A list of objects mapped based on jobTitle attribute mappings configured in the system.`,
						},
						"manager_ids": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
							Description: `A list of ids of the user's managers.`,
						},
						"manager_sources": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"app_id": schema.StringAttribute{
										Computed:    true,
										Description: `The appId field.`,
									},
									"app_user_id": schema.StringAttribute{
										Computed:    true,
										Description: `The appUserId field.`,
									},
									"app_user_profile_attribute_key": schema.StringAttribute{
										Computed:    true,
										Description: `The appUserProfileAttributeKey field.`,
									},
									"user_attribute_mapping_id": schema.StringAttribute{
										Computed:    true,
										Description: `The userAttributeMappingId field.`,
									},
									"value": schema.StringAttribute{
										Computed:    true,
										Description: `The value field.`,
									},
								},
							},
							Description: `A list of objects mapped based on managerId attribute mappings configured in the system.`,
						},
						"profile": schema.MapNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"str": schema.StringAttribute{
										Computed: true,
									},
									"number": schema.NumberAttribute{
										Computed: true,
									},
									"array_of_any": schema.ListAttribute{
										Computed:    true,
										ElementType: types.StringType,
										Validators: []validator.List{
											listvalidator.ValueStringsAre(validators.IsValidJSON()),
										},
									},
									"boolean": schema.BoolAttribute{
										Computed: true,
									},
									"three": schema.SingleNestedAttribute{
										Computed:   true,
										Attributes: map[string]schema.Attribute{},
									},
								},
								Validators: []validator.Object{
									validators.ExactlyOneChild(),
								},
							},
						},
						"role_ids": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
							Description: `A list of unique identifiers that maps to ConductorOne’s user roles let you assign users permissions tailored to the work they do in the software.`,
						},
						"status": schema.StringAttribute{
							Computed:    true,
							Description: `The status of the user in the system. must be one of ["UNKNOWN", "ENABLED", "DISABLED", "DELETED"]`,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"UNKNOWN",
									"ENABLED",
									"DISABLED",
									"DELETED",
								),
							},
						},
						"updated_at": schema.StringAttribute{
							Computed: true,
							Validators: []validator.String{
								validators.IsRFC3339(),
							},
						},
						"username": schema.StringAttribute{
							Computed:    true,
							Description: `This is the user's primary username. Typically sourced from the primary directory.`,
						},
						"usernames": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
							Description: `This is a list of all of the user's usernames from app users.`,
						},
					},
				},
				Description: `The owners of the app.`,
			},
			"certify_policy_id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Creates the app with this certify policy.`,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"delete_app_request": schema.SingleNestedAttribute{
				Optional:    true,
				Attributes:  map[string]schema.Attribute{},
				Description: `Empty request body`,
			},
			"deleted_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Creates the app with this description.`,
			},
			"display_name": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Creates the app with this display name.`,
			},
			"field_mask": schema.StringAttribute{
				Computed: true,
			},
			"grant_policy_id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Creates the app with this grant policy.`,
			},
			"icon_url": schema.StringAttribute{
				Computed:    true,
				Description: `The URL of an icon to display for the app.`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `The ID of the app.`,
			},
			"identity_matching": schema.StringAttribute{
				Computed:    true,
				Description: `The identityMatching field. must be one of ["APP_USER_IDENTITY_MATCHING_UNSPECIFIED", "APP_USER_IDENTITY_MATCHING_STRICT", "APP_USER_IDENTITY_MATCHING_DISPLAY_NAME"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"APP_USER_IDENTITY_MATCHING_UNSPECIFIED",
						"APP_USER_IDENTITY_MATCHING_STRICT",
						"APP_USER_IDENTITY_MATCHING_DISPLAY_NAME",
					),
				},
			},
			"is_directory": schema.BoolAttribute{
				Computed:    true,
				Description: `Specifies if the app is a directory.`,
			},
			"logo_uri": schema.StringAttribute{
				Computed:    true,
				Description: `The URL of a logo to display for the app.`,
			},
			"monthly_cost_usd": schema.Int64Attribute{
				Computed:    true,
				Optional:    true,
				Description: `Creates the app with this monthly cost per seat.`,
			},
			"owners": schema.ListAttribute{
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				ElementType: types.StringType,
				Description: `Creates the app with this array of owners. Requires replacement if changed. `,
			},
			"parent_app_id": schema.StringAttribute{
				Computed:    true,
				Description: `The ID of the app that created this app, if any.`,
			},
			"revoke_policy_id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Creates the app with this revoke policy.`,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"user_count": schema.StringAttribute{
				Computed:    true,
				Description: `The number of users with grants to this app.`,
			},
		},
	}
}

func (r *AppResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *AppResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *AppResourceModel
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

	request := data.ToSharedCreateAppRequest()
	res, err := r.client.Apps.Create(ctx, request)
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
	if res.CreateAppResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedApp(res.CreateAppResponse.App)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	id := data.ID.ValueString()
	request1 := operations.C1APIAppV1AppsGetRequest{
		ID: id,
	}
	res1, err := r.client.Apps.Get(ctx, request1)
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
	if res1.GetAppResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedApp(res1.GetAppResponse.App)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *AppResourceModel
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
	request := operations.C1APIAppV1AppsGetRequest{
		ID: id,
	}
	res, err := r.client.Apps.Get(ctx, request)
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
	if res.GetAppResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedApp(res.GetAppResponse.App)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *AppResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	id := data.ID.ValueString()
	var updateAppRequest *shared.UpdateAppRequest
	app := data.ToSharedAppInput()
	updateAppRequest = &shared.UpdateAppRequest{
		App: app,
	}
	request := operations.C1APIAppV1AppsUpdateRequest{
		ID:               id,
		UpdateAppRequest: updateAppRequest,
	}
	res, err := r.client.Apps.Update(ctx, request)
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
	if res.UpdateAppResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedApp(res.UpdateAppResponse.App)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	id1 := data.ID.ValueString()
	request1 := operations.C1APIAppV1AppsGetRequest{
		ID: id1,
	}
	res1, err := r.client.Apps.Get(ctx, request1)
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
	if res1.GetAppResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedApp(res1.GetAppResponse.App)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *AppResourceModel
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
	deleteAppRequest := data.ToSharedDeleteAppRequest()
	request := operations.C1APIAppV1AppsDeleteRequest{
		ID:               id,
		DeleteAppRequest: deleteAppRequest,
	}
	res, err := r.client.Apps.Delete(ctx, request)
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

func (r *AppResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
