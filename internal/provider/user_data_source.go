package provider

import (
	"context"
	"fmt"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/conductorone/terraform-provider-conductorone/internal/validators"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/operations"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func NewUserDataSource() datasource.DataSource {
	return &UserDataSource{}
}

// UserDataSource defines the data source implementation.
type UserDataSource struct {
	client *sdk.ConductoroneAPI
}

// UserDataSourceModel describes the data source data model.
type UserDataSourceModel struct {
	CreatedAt               types.String                 `tfsdk:"created_at"`
	DelegatedUserID         types.String                 `tfsdk:"delegated_user_id"`
	DeletedAt               types.String                 `tfsdk:"deleted_at"`
	Department              types.String                 `tfsdk:"department"`
	DepartmentSources       []UserAttributeMappingSource `tfsdk:"department_sources"`
	DirectoryIds            []types.String               `tfsdk:"directory_ids"`
	DirectoryStatus         types.String                 `tfsdk:"directory_status"`
	DirectoryStatusSources  []UserAttributeMappingSource `tfsdk:"directory_status_sources"`
	DisplayName             types.String                 `tfsdk:"display_name"`
	Email                   types.String                 `tfsdk:"email"`
	EmploymentStatus        types.String                 `tfsdk:"employment_status"`
	EmploymentStatusSources []UserAttributeMappingSource `tfsdk:"employment_status_sources"`
	EmploymentType          types.String                 `tfsdk:"employment_type"`
	EmploymentTypeSources   []UserAttributeMappingSource `tfsdk:"employment_type_sources"`
	ID                      types.String                 `tfsdk:"id"`
	JobTitle                types.String                 `tfsdk:"job_title"`
	JobTitleSources         []UserAttributeMappingSource `tfsdk:"job_title_sources"`
	ManagerIds              []types.String               `tfsdk:"manager_ids"`
	ManagerSources          []UserAttributeMappingSource `tfsdk:"manager_sources"`
	RoleIds                 []types.String               `tfsdk:"role_ids"`
	Status                  types.String                 `tfsdk:"status"`
	UpdatedAt               types.String                 `tfsdk:"updated_at"`
}

func (r *UserDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

func (r *UserDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "App DataSource",
		Attributes: map[string]schema.Attribute{
			"created_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"delegated_user_id": schema.StringAttribute{
				Computed:    true,
				Description: `The delegatedUserId field.`,
			},
			"deleted_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"department": schema.StringAttribute{
				Computed:    true,
				Description: `The department field.`,
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
				Description: `The departmentSources field.`,
			},
			"directory_ids": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `The directoryIds field.`,
			},
			"directory_status": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"UNKNOWN",
						"ENABLED",
						"DISABLED",
						"DELETED",
					),
				},
				MarkdownDescription: `must be one of [UNKNOWN, ENABLED, DISABLED, DELETED]` + "\n" +
					`The directoryStatus field.`,
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
				Description: `The directoryStatusSources field.`,
			},
			"display_name": schema.StringAttribute{
				Computed:    true,
				Description: `The displayName field.`,
			},
			"email": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The email field.`,
			},
			"employment_status": schema.StringAttribute{
				Computed:    true,
				Description: `The employmentStatus field.`,
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
				Description: `The employmentStatusSources field.`,
			},
			"employment_type": schema.StringAttribute{
				Computed:    true,
				Description: `The employmentType field.`,
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
				Description: `The employmentTypeSources field.`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The id field.`,
			},
			"job_title": schema.StringAttribute{
				Computed:    true,
				Description: `The jobTitle field.`,
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
				Description: `The jobTitleSources field.`,
			},
			"manager_ids": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `The managerIds field.`,
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
				Description: `The managerSources field.`,
			},
			"role_ids": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `The roleIds field.`,
			},
			"status": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"UNKNOWN",
						"ENABLED",
						"DISABLED",
						"DELETED",
					),
				},
				MarkdownDescription: `must be one of [UNKNOWN, ENABLED, DISABLED, DELETED]` + "\n" +
					`The status field.`,
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

func (r *UserDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *UserDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *UserDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
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
	email := data.Email.ValueStringPointer()

	if id == "" && (email == nil || *email == "") {
		resp.Diagnostics.AddError("either id or email must be set", "")
		return
	}

	// If we have an ID, we can use the Get API to fetch the latest data
	if id != "" {
		req := operations.C1APIUserV1UserServiceGetRequest{
			ID: id,
		}
		res, err := r.client.User.Get(ctx, req)
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
		if res.UserServiceGetResponse == nil {
			resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
			return
		}
		if res.UserServiceGetResponse.UserView == nil || res.UserServiceGetResponse.UserView.User == nil {
			resp.Diagnostics.AddError("unexpected response from API. Returned user is nil", debugResponse(res.RawResponse))
			return
		}
		data.RefreshFromGetResponse(res.UserServiceGetResponse.UserView.User)
		// Save updated data into Terraform state
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		return
	}

	// If we don't have an ID, we can use the Search API to fetch the user by email
	request := shared.SearchUsersRequest{
		Email: email,
	}
	res, err := r.client.UserSearch.Search(ctx, request)
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
	if res.SearchUsersResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}

	if len(res.SearchUsersResponse.List) != 1 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Expected 1 user, got %d", len(res.SearchUsersResponse.List)), debugResponse(res.RawResponse))
		return
	}

	if len(res.SearchUsersResponse.List) == 0 {
		resp.Diagnostics.AddError("unexpected response from API. User was not found", debugResponse(res.RawResponse))
		return
	}

	if len(res.SearchUsersResponse.List) > 1 {
		resp.Diagnostics.AddError("unexpected response from API. More than 1 User was found", debugResponse(res.RawResponse))
		return
	}

	data.RefreshFromGetResponse(res.SearchUsersResponse.List[0].User)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
