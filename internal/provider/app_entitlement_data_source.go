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

// Ensure provider defined types fully satisfy framework interfaces.
// var _ datasource.Resource = &CatalogResource{}
// var _ datasource.ResourceWithImportState = &CatalogResource{}

func NewAppEntitlementDataSource() datasource.DataSource {
	return &AppEntitlementDataSource{}
}

// CatalogResource defines the resource implementation.
type AppEntitlementDataSource struct {
	client *sdk.ConductoroneAPI
}

// CatalogResourceModel describes the resource data model.
type AppEntitlementDataSourceModel struct {
	Alias                       types.String                 `tfsdk:"alias"`
	AppID                       types.String                 `tfsdk:"app_id"`
	AppResourceID               types.String                 `tfsdk:"app_resource_id"`
	AppResourceTypeID           types.String                 `tfsdk:"app_resource_type_id"`
	CertifyPolicyID             types.String                 `tfsdk:"certify_policy_id"`
	ComplianceFrameworkValueIds []types.String               `tfsdk:"compliance_framework_value_ids"`
	CreatedAt                   types.String                 `tfsdk:"created_at"`
	DeletedAt                   types.String                 `tfsdk:"deleted_at"`
	Description                 types.String                 `tfsdk:"description"`
	DisplayName                 types.String                 `tfsdk:"display_name"`
	DurationGrant               types.String                 `tfsdk:"duration_grant"`
	DurationUnset               *AppEntitlementDurationUnset `tfsdk:"duration_unset"`
	EmergencyGrantEnabled       types.Bool                   `tfsdk:"emergency_grant_enabled"`
	EmergencyGrantPolicyID      types.String                 `tfsdk:"emergency_grant_policy_id"`
	GrantPolicyID               types.String                 `tfsdk:"grant_policy_id"`
	ID                          types.String                 `tfsdk:"id"`
	ProvisionPolicy             *ProvisionPolicy             `tfsdk:"provision_policy"`
	RevokePolicyID              types.String                 `tfsdk:"revoke_policy_id"`
	RiskLevelValueID            types.String                 `tfsdk:"risk_level_value_id"`
	Slug                        types.String                 `tfsdk:"slug"`
	UpdatedAt                   types.String                 `tfsdk:"updated_at"`
}

func (r *AppEntitlementDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_app_entitlement"
}

func (r *AppEntitlementDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "AppEntitlement Datasource",

		Attributes: map[string]schema.Attribute{
			"alias": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: `The alias field.`,
			},
			"app_id": schema.StringAttribute{
				Optional:    true,
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
				Description: `The certifyPolicyId is the ID of the policy that will be used for access review certify tasks.`,
			},
			"compliance_framework_value_ids": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `ComplianceFrameworkValueIds is a list of ComplianceFramework attributes that are set on the app entitlement.`,
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
				Description: `The emergencyGrantEnabled field determines whether or not this entitlement has an emergency grant policy.`,
			},
			"emergency_grant_policy_id": schema.StringAttribute{
				Computed: true,
				Description: `The emergencyGrantPolicyId field is the ID of the grant policy that will be used for emergency grant tasks. 
				To set this field, emergencyGrantEnabled must be set to true.`,
			},
			"grant_policy_id": schema.StringAttribute{
				Computed:    true,
				Description: `The grantPolicyId field is the policy that will be used for access request grant tasks.`,
			},
			"id": schema.StringAttribute{
				Optional:    true,
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
				MarkdownDescription: `The ProvisionPolicy message is the Provision strategy that will be used for granting access for this entitlement.` + "\n" +
					`` + "\n" +
					`This message contains a oneof. Only a single field of the following list may be set at a time:` + "\n" +
					`  - connector` + "\n" +
					`  - manual` + "\n" +
					`  - delegated` + "\n" +
					``,
			},
			"revoke_policy_id": schema.StringAttribute{
				Computed:    true,
				Description: `The revokePolicyId field is the ID of the policy that will be used for revoke access tasks.`,
			},
			"risk_level_value_id": schema.StringAttribute{
				Computed:    true,
				Description: `The riskLevelValueId field is the ID of the risk level attribute value that will be set on the app entitlement.`,
			},
			"slug": schema.StringAttribute{
				Computed:    true,
				Description: `The slug field.`,
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

func (r *AppEntitlementDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *AppEntitlementDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *AppEntitlementDataSourceModel
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

	appId := data.AppID.ValueString()
	entitlementId := data.ID.ValueString()
	alias := data.Alias.ValueStringPointer()
	if (appId == "" || entitlementId == "") && (alias == nil || *alias == "") {
		resp.Diagnostics.AddError("either app_id and id, or alias must be set", "")
		return
	}

	// If the AppID and EntitlementID are set, use the Get API
	if appId != "" && entitlementId != "" {
		res, err := r.client.AppEntitlements.Get(ctx, operations.C1APIAppV1AppEntitlementsGetRequest{AppID: appId, ID: entitlementId})
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
		if res.GetAppEntitlementResponse == nil {
			resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
			return
		}
		if res.GetAppEntitlementResponse.AppEntitlementView == nil || res.GetAppEntitlementResponse.AppEntitlementView.AppEntitlement == nil {
			resp.Diagnostics.AddError("unexpected response from API. app entitlement is nil", debugResponse(res.RawResponse))
			return
		}
		data.RefreshFromGetResponse(res.GetAppEntitlementResponse.AppEntitlementView.AppEntitlement)

		// Save updated data into Terraform state
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		return
	}

	// If the Alias is set, use the Search API
	request := shared.AppEntitlementSearchServiceSearchRequest{}
	if alias != nil && *alias != "" {
		request.Alias = alias
	}
	res, err := r.client.AppEntitlementSearch.Search(ctx, request)
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
	if res.AppEntitlementSearchServiceSearchResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}

	if len(res.AppEntitlementSearchServiceSearchResponse.List) == 0 {
		resp.Diagnostics.AddError("unexpected response from API. App entitlement was not found", debugResponse(res.RawResponse))
		return
	}

	if len(res.AppEntitlementSearchServiceSearchResponse.List) > 2 {
		resp.Diagnostics.AddError("unexpected response from API. More than 1 app entitlement was found", debugResponse(res.RawResponse))
		return
	}

	if res.AppEntitlementSearchServiceSearchResponse.List[0].AppEntitlement == nil {
		resp.Diagnostics.AddError("unexpected response from API. app entitlement is nil", debugResponse(res.RawResponse))
		return
	}

	data.RefreshFromGetResponse(res.AppEntitlementSearchServiceSearchResponse.List[0].AppEntitlement)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
