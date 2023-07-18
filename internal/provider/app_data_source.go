package provider

import (
	"conductorone/internal/sdk"
	"conductorone/internal/validators"
	"context"
	"fmt"

	"conductorone/internal/sdk/pkg/models/operations"
	"conductorone/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func NewAppDataSource() datasource.DataSource {
	return &AppDataSource{}
}

// AppDataSource defines the data source implementation.
type AppDataSource struct {
	client *sdk.ConductoroneAPI
}

// AppDataSourceModel describes the data source data model.
type AppDataSourceModel struct {
	AppAccountID    types.String   `tfsdk:"app_account_id"`
	AppAccountName  types.String   `tfsdk:"app_account_name"`
	CertifyPolicyID types.String   `tfsdk:"certify_policy_id"`
	CreatedAt       types.String   `tfsdk:"created_at"`
	DeletedAt       types.String   `tfsdk:"deleted_at"`
	Description     types.String   `tfsdk:"description"`
	DisplayName     types.String   `tfsdk:"display_name"`
	FieldMask       types.String   `tfsdk:"field_mask"`
	GrantPolicyID   types.String   `tfsdk:"grant_policy_id"`
	IconURL         types.String   `tfsdk:"icon_url"`
	ID              types.String   `tfsdk:"id"`
	LogoURI         types.String   `tfsdk:"logo_uri"`
	MonthlyCostUsd  types.Number   `tfsdk:"monthly_cost_usd"`
	Owners          []types.String `tfsdk:"owners"`
	ParentAppID     types.String   `tfsdk:"parent_app_id"`
	RevokePolicyID  types.String   `tfsdk:"revoke_policy_id"`
	UpdatedAt       types.String   `tfsdk:"updated_at"`
	UserCount       types.String   `tfsdk:"user_count"`
}

func (r *AppDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_app"
}

func (r *AppDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "App DataSource",
		Attributes: map[string]schema.Attribute{
			"app_account_id": schema.StringAttribute{
				Computed:    true,
				Description: `The appAccountId field.`,
			},
			"app_account_name": schema.StringAttribute{
				Computed:    true,
				Description: `The appAccountName field.`,
			},
			"certify_policy_id": schema.StringAttribute{
				Computed:    true,
				Description: `The certifyPolicyId field.`,
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
				Optional:    true,
				Computed:    true,
				Description: `The displayName field.`,
			},
			"field_mask": schema.StringAttribute{
				Computed: true,
			},
			"grant_policy_id": schema.StringAttribute{
				Computed:    true,
				Description: `The grantPolicyId field.`,
			},
			"icon_url": schema.StringAttribute{
				Computed:    true,
				Description: `The iconUrl field.`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The id field.`,
			},
			"logo_uri": schema.StringAttribute{
				Computed:    true,
				Description: `The logoUri field.`,
			},
			"monthly_cost_usd": schema.NumberAttribute{
				Computed:    true,
				Description: `The monthlyCostUsd field.`,
			},
			"owners": schema.ListAttribute{
				ElementType: types.StringType,
				Description: `The owners field.`,
				Computed:    true,
			},
			"parent_app_id": schema.StringAttribute{
				Computed:    true,
				Description: `The parentAppId field.`,
			},
			"revoke_policy_id": schema.StringAttribute{
				Computed:    true,
				Description: `The revokePolicyId field.`,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"user_count": schema.StringAttribute{
				Computed:    true,
				Description: `The userCount field.`,
			},
		},
	}
}

func (r *AppDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *AppDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *AppDataSourceModel
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
	displayName := data.DisplayName.ValueStringPointer()

	if id == "" && (displayName == nil || *displayName == "") {
		resp.Diagnostics.AddError("either id or display_name must be set", "")
		return
	}

	// If the ID is set, we can use the Get API to retrieve the resource.
	if id != "" {
		request := operations.C1APIAppV1AppsGetRequest{
			ID: id,
		}
		res, err := r.client.Apps.Get(ctx, request)
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
		if res.GetAppResponse == nil {
			resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
			return
		}
		if res.GetAppResponse.App == nil {
			resp.Diagnostics.AddError("unexpected response from API. Returned App is nil", debugResponse(res.RawResponse))
			return
		}
		data.RefreshFromGetResponse(res.GetAppResponse.App)

		// Save updated data into Terraform state
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		return
	}

	// If the ID is not set, we can use the Search API to try to retrieve the resource.
	request := shared.SearchAppsRequest{
		DisplayName: displayName,
	}
	res, err := r.client.AppSearch.Search(ctx, request)
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
	if res.SearchAppsResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}

	if len(res.SearchAppsResponse.List) != 1 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Expected 1 app, got %d", len(res.SearchAppsResponse.List)), debugResponse(res.RawResponse))
		return
	}

	data.RefreshFromGetResponse(&res.SearchAppsResponse.List[0])

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
