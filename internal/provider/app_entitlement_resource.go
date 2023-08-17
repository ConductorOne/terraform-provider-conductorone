package provider

import (
	"conductorone/internal/sdk"
	"context"
	"fmt"

	"conductorone/internal/sdk/pkg/models/operations"
	"conductorone/internal/sdk/pkg/models/shared"
	"conductorone/internal/validators"

	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppEntitlementResource{}
var _ resource.ResourceWithImportState = &AppEntitlementResource{}

func NewAppEntitlementResource() resource.Resource {
	return &AppEntitlementResource{}
}

// AppEntitlementResource defines the resource implementation.
type AppEntitlementResource struct {
	client *sdk.ConductoroneAPI
}

// AppEntitlementResourceModel describes the resource data model.
type AppEntitlementResourceModel struct {
	Alias                       types.String      `tfsdk:"alias"`
	AppID                       types.String      `tfsdk:"app_id"`
	AppResourceID               types.String      `tfsdk:"app_resource_id"`
	AppResourceTypeID           types.String      `tfsdk:"app_resource_type_id"`
	CertifyPolicyID             types.String      `tfsdk:"certify_policy_id"`
	ComplianceFrameworkValueIds []types.String    `tfsdk:"compliance_framework_value_ids"`
	CreatedAt                   types.String      `tfsdk:"created_at"`
	DeletedAt                   types.String      `tfsdk:"deleted_at"`
	Description                 types.String      `tfsdk:"description"`
	DisplayName                 types.String      `tfsdk:"display_name"`
	MaxGrantDuration            *MaxGrantDuration `tfsdk:"max_grant_duration"`
	EmergencyGrantEnabled       types.Bool        `tfsdk:"emergency_grant_enabled"`
	EmergencyGrantPolicyID      types.String      `tfsdk:"emergency_grant_policy_id"`
	GrantPolicyID               types.String      `tfsdk:"grant_policy_id"`
	ID                          types.String      `tfsdk:"id"`
	ProvisionPolicy             *ProvisionPolicy  `tfsdk:"provision_policy"`
	RevokePolicyID              types.String      `tfsdk:"revoke_policy_id"`
	RiskLevelValueID            types.String      `tfsdk:"risk_level_value_id"`
	Slug                        types.String      `tfsdk:"slug"`
	UpdatedAt                   types.String      `tfsdk:"updated_at"`
}

func (r *AppEntitlementResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_app_entitlement"
}

func (r *AppEntitlementResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "AppEntitlement Resource",

		Attributes: map[string]schema.Attribute{
			"alias": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The alias field.`,
			},
			"app_id": schema.StringAttribute{
				Required:    true,
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
				Optional:    true,
				Description: `The certifyPolicyId is the ID of the policy that will be used for access review certify tasks.`,
			},
			"compliance_framework_value_ids": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
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
				Optional:    true,
				Description: `The description field.`,
			},
			"display_name": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The displayName field.`,
			},
			"max_grant_duration": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"duration_grant": schema.StringAttribute{
						Optional: true,
						Description: `The DurationGrant field is a string attribute that represents the maximum duration a grant to this entitlement can last. 
						The format of this is <time in seconds>s. i.e. 1h = 3600s.`,
					},
					"duration_unset": schema.SingleNestedAttribute{
						Optional:    true,
						Attributes:  map[string]schema.Attribute{},
						Description: `The DurationUnset field is set if there is no maximum duration a grant to this entitlement can last.`,
						Validators: []validator.Object{
							objectvalidator.ExactlyOneOf(path.Expressions{
								path.MatchRoot("max_grant_duration").AtName("duration_grant"),
							}...),
						},
					},
				},
				MarkdownDescription: `MaxGrantDuration is a one of.` + "\n" +
					`` +
					`This message contains a oneof. Only a single field of the following list may be set at a time:` + "\n" +
					`  - duration_unset` + "\n" +
					`  - duration_grant` + "\n" +
					"\n" +
					``,
			},
			"emergency_grant_enabled": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The emergencyGrantEnabled field determines whether or not this entitlement has an emergency grant policy.`,
			},
			"emergency_grant_policy_id": schema.StringAttribute{
				Computed: true,
				Optional: true,
				Description: `The emergencyGrantPolicyId field is the ID of the grant policy that will be used for emergency grant tasks. 
				To set this field, emergencyGrantEnabled must be set to true.`,
			},
			"grant_policy_id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The grantPolicyId field is the policy that will be used for access request grant tasks.`,
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `The id field.`,
			},
			// TODO: this is a oneof
			"provision_policy": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"connector_provision": schema.SingleNestedAttribute{
						Optional:    true,
						Attributes:  map[string]schema.Attribute{},
						Description: `The ConnectorProvision message.`,
					},
					"delegated_provision": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"app_id": schema.StringAttribute{
								Optional:    true,
								Description: `The appId field.`,
							},
							"entitlement_id": schema.StringAttribute{
								Optional:    true,
								Description: `The entitlementId field.`,
							},
						},
						Description: `The DelegatedProvision message.`,
					},
					"manual_provision": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"instructions": schema.StringAttribute{
								Optional:    true,
								Description: `The instructions field.`,
							},
							"user_ids": schema.ListAttribute{
								Optional:    true,
								ElementType: types.StringType,
								Description: `The userIds field.`,
							},
						},
						Description: `The ManualProvision message.`,
					},
				},
				MarkdownDescription: `The ProvisionPolicy message is the Provision strategy that will be used for granting access for this entitlement.` + "\n" +
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
				Optional:    true,
				Description: `The revokePolicyId field is the ID of the policy that will be used for revoke access tasks.`,
			},
			"risk_level_value_id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The riskLevelValueId field is the ID of the risk level attribute value that will be set on the app entitlement.`,
			},
			"slug": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
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

func (r *AppEntitlementResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *AppEntitlementResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *AppEntitlementResourceModel
	var item types.Object

	var updateAppEntitlementRequest *shared.UpdateAppEntitlementRequest
	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	appEntitlement := data.ToUpdateSDKType()

	updateAppEntitlementRequest = &shared.UpdateAppEntitlementRequest{
		AppEntitlement: appEntitlement,
	}
	appID := data.AppID.ValueString()
	id := data.ID.ValueString()
	request := operations.C1APIAppV1AppEntitlementsUpdateRequest{
		UpdateAppEntitlementRequest: updateAppEntitlementRequest,
		AppID:                       appID,
		ID:                          id,
	}
	res, err := r.client.AppEntitlements.Update(ctx, request)
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
	if res.UpdateAppEntitlementResponse.AppEntitlementView == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromUpdateResponse(res.UpdateAppEntitlementResponse.AppEntitlementView.AppEntitlement)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppEntitlementResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *AppEntitlementResourceModel
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
	request := operations.C1APIAppV1AppEntitlementsGetRequest{
		AppID: appID,
		ID:    id,
	}
	res, err := r.client.AppEntitlements.Get(ctx, request)
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

	if res.GetAppEntitlementResponse.AppEntitlementView == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}

	if res.GetAppEntitlementResponse.AppEntitlementView.AppEntitlement == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return

	}

	if res.GetAppEntitlementResponse.AppEntitlementView.AppEntitlement.DeletedAt != nil {
		resp.State.RemoveResource(ctx)
		return
	}

	data.RefreshFromGetResponse(res.GetAppEntitlementResponse.AppEntitlementView.AppEntitlement)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppEntitlementResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *AppEntitlementResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	var updateAppEntitlementRequest *shared.UpdateAppEntitlementRequest
	appEntitlement := data.ToUpdateSDKType()

	appID := data.AppID.ValueString()
	id := data.ID.ValueString()

	currentAppEntitlement, err := r.readAppEntitlementAndValidate(ctx, appID, id)
	if err != nil {
		resp.Diagnostics.AddError("failure reading app entitlement", err.Error())
		return
	}

	// If no value was specified for the ProvisionPolicy, use the current value.
	if appEntitlement.ProvisionPolicy == nil {
		appEntitlement.ProvisionPolicy = currentAppEntitlement.ProvisionPolicy
	}

	updateAppEntitlementRequest = &shared.UpdateAppEntitlementRequest{
		AppEntitlement: appEntitlement,
	}

	request := operations.C1APIAppV1AppEntitlementsUpdateRequest{
		UpdateAppEntitlementRequest: updateAppEntitlementRequest,
		AppID:                       appID,
		ID:                          id,
	}
	res, err := r.client.AppEntitlements.Update(ctx, request)
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
	if res.UpdateAppEntitlementResponse.AppEntitlementView == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromUpdateResponse(res.UpdateAppEntitlementResponse.AppEntitlementView.AppEntitlement)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppEntitlementResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *AppEntitlementResourceModel
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

	// Not Implemented; entity does not have a configured DELETE operation
}

func (r *AppEntitlementResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource app_entitlement.")
}

func (r *AppEntitlementResource) readAppEntitlementAndValidate(ctx context.Context, appID string, id string) (*shared.AppEntitlement, error) {
	request := operations.C1APIAppV1AppEntitlementsGetRequest{
		AppID: appID,
		ID:    id,
	}
	res, err := r.client.AppEntitlements.Get(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("failure to invoke API")
	}
	if res == nil {
		return nil, fmt.Errorf("unexpected response from API")
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected response from API. Got an unexpected response code %v", res.StatusCode)
	}

	if res.GetAppEntitlementResponse.AppEntitlementView == nil {
		return nil, fmt.Errorf("unexpected response from API. No response body")
	}

	if res.GetAppEntitlementResponse.AppEntitlementView.AppEntitlement == nil {
		return nil, fmt.Errorf("unexpected response from API. No response body")
	}
	return res.GetAppEntitlementResponse.AppEntitlementView.AppEntitlement, nil
}
