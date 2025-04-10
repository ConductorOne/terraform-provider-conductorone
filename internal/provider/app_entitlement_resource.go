package provider

import (
	"context"
	"fmt"
	"strings"

	tfTypes "github.com/conductorone/terraform-provider-conductorone/internal/provider/types"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/operations"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"github.com/conductorone/terraform-provider-conductorone/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
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
	Alias                          types.String                                      `tfsdk:"alias"`
	AppID                          types.String                                      `tfsdk:"app_id"`
	AppResourceID                  types.String                                      `tfsdk:"app_resource_id"`
	AppResourceTypeID              types.String                                      `tfsdk:"app_resource_type_id"`
	CertifyPolicyID                types.String                                      `tfsdk:"certify_policy_id"`
	ComplianceFrameworkValueIds    []types.String                                    `tfsdk:"compliance_framework_value_ids"`
	CreatedAt                      types.String                                      `tfsdk:"created_at"`
	DeletedAt                      types.String                                      `tfsdk:"deleted_at"`
	Description                    types.String                                      `tfsdk:"description"`
	DisplayName                    types.String                                      `tfsdk:"display_name"`
	DurationGrant                  types.String                                      `tfsdk:"duration_grant"`
	DurationUnset                  *tfTypes.CreateAppEntitlementRequestDurationUnset `tfsdk:"duration_unset"`
	EmergencyGrantEnabled          types.Bool                                        `tfsdk:"emergency_grant_enabled"`
	EmergencyGrantPolicyID         types.String                                      `tfsdk:"emergency_grant_policy_id"`
	GrantPolicyID                  types.String                                      `tfsdk:"grant_policy_id"`
	ID                             types.String                                      `tfsdk:"id"`
	OverrideAccessRequestsDefaults types.Bool                                        `tfsdk:"override_access_requests_defaults"`
	ProvisionPolicy                *tfTypes.ProvisionPolicy                          `tfsdk:"provision_policy" tfPlanOnly:"true"`
	RevokePolicyID                 types.String                                      `tfsdk:"revoke_policy_id"`
	RiskLevelValueID               types.String                                      `tfsdk:"risk_level_value_id"`
	Slug                           types.String                                      `tfsdk:"slug"`
	UpdatedAt                      types.String                                      `tfsdk:"updated_at"`
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
			"duration_grant": schema.StringAttribute{
				Computed: true,
				Optional: true,
				Description: `The DurationGrant field is a string attribute that represents the maximum duration a grant to this entitlement can last. 
				The format of this is <time in seconds>s. i.e. 1h = 3600s.`,
			},
			"duration_unset": schema.SingleNestedAttribute{
				Computed:    true,
				Optional:    true,
				Attributes:  map[string]schema.Attribute{},
				Description: `The DurationUnset field is set if there is no maximum duration a grant to this entitlement can last.`,
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
			"override_access_requests_defaults": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The overrideAccessRequestsDefaults field.`,
			},
			"provision_policy": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"connector_provision": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"account_provision": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"config": schema.SingleNestedAttribute{
										Computed: true,
										Optional: true,
									},
									"connector_id": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Description: `The connectorId field.`,
									},
									"schema_id": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Description: `The schemaId field.`,
									},
								},
								Description: `The AccountProvision message.`,
							},
							"default_behavior": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"connector_id": schema.StringAttribute{
										Computed: true,
										Optional: true,
										MarkdownDescription: `this checks if the entitlement is enabled by provisioning in a specific connector` + "\n" +
											` this can happen automatically and doesn't need any extra info`,
									},
								},
								Description: `The DefaultBehavior message.`,
							},
						},
						MarkdownDescription: `Indicates that a connector should perform the provisioning. This object has no fields.` + "\n" +
							`` + "\n" +
							`This message contains a oneof named provision_type. Only a single field of the following list may be set at a time:` + "\n" +
							`  - defaultBehavior` + "\n" +
							`  - account`,
						Validators: []validator.Object{
							objectvalidator.ConflictsWith(path.Expressions{
								path.MatchRelative().AtParent().AtName("delegated_provision"),
								path.MatchRelative().AtParent().AtName("external_ticket_provision"),
								path.MatchRelative().AtParent().AtName("manual_provision"),
								path.MatchRelative().AtParent().AtName("multi_step"),
								path.MatchRelative().AtParent().AtName("webhook_provision"),
							}...),
						},
					},
					"delegated_provision": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"app_id": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `The AppID of the entitlement to delegate provisioning to.`,
							},
							"entitlement_id": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `The ID of the entitlement we are delegating provisioning to.`,
							},
						},
						Description: `This provision step indicates that we should delegate provisioning to the configuration of another app entitlement. This app entitlement does not have to be one from the same app, but MUST be configured as a proxy binding leading into this entitlement.`,
						Validators: []validator.Object{
							objectvalidator.ConflictsWith(path.Expressions{
								path.MatchRelative().AtParent().AtName("connector_provision"),
								path.MatchRelative().AtParent().AtName("external_ticket_provision"),
								path.MatchRelative().AtParent().AtName("manual_provision"),
								path.MatchRelative().AtParent().AtName("multi_step"),
								path.MatchRelative().AtParent().AtName("webhook_provision"),
							}...),
						},
					},
					"external_ticket_provision": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"app_id": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `The appId field.`,
							},
							"connector_id": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `The connectorId field.`,
							},
							"external_ticket_provisioner_config_id": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `The externalTicketProvisionerConfigId field.`,
							},
							"instructions": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `This field indicates a text body of instructions for the provisioner to indicate.`,
							},
						},
						Description: `This provision step indicates that we should check an external ticket to provision this entitlement`,
						Validators: []validator.Object{
							objectvalidator.ConflictsWith(path.Expressions{
								path.MatchRelative().AtParent().AtName("connector_provision"),
								path.MatchRelative().AtParent().AtName("delegated_provision"),
								path.MatchRelative().AtParent().AtName("manual_provision"),
								path.MatchRelative().AtParent().AtName("multi_step"),
								path.MatchRelative().AtParent().AtName("webhook_provision"),
							}...),
						},
					},
					"manual_provision": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"instructions": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `This field indicates a text body of instructions for the provisioner to indicate.`,
							},
							"user_ids": schema.ListAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
								Description: `An array of users that are required to provision during this step.`,
							},
						},
						Description: `Manual provisioning indicates that a human must intervene for the provisioning of this step.`,
						Validators: []validator.Object{
							objectvalidator.ConflictsWith(path.Expressions{
								path.MatchRelative().AtParent().AtName("connector_provision"),
								path.MatchRelative().AtParent().AtName("delegated_provision"),
								path.MatchRelative().AtParent().AtName("external_ticket_provision"),
								path.MatchRelative().AtParent().AtName("multi_step"),
								path.MatchRelative().AtParent().AtName("webhook_provision"),
							}...),
						},
					},
					"multi_step": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `MultiStep indicates that this provision step has multiple steps to process. Parsed as JSON.`,
						Validators: []validator.String{
							stringvalidator.ConflictsWith(path.Expressions{
								path.MatchRelative().AtParent().AtName("connector_provision"),
								path.MatchRelative().AtParent().AtName("delegated_provision"),
								path.MatchRelative().AtParent().AtName("external_ticket_provision"),
								path.MatchRelative().AtParent().AtName("manual_provision"),
								path.MatchRelative().AtParent().AtName("webhook_provision"),
							}...),
							validators.IsValidJSON(),
						},
					},
					"webhook_provision": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"webhook_id": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `The ID of the webhook to call for provisioning.`,
							},
						},
						Description: `This provision step indicates that a webhook should be called to provision this entitlement.`,
						Validators: []validator.Object{
							objectvalidator.ConflictsWith(path.Expressions{
								path.MatchRelative().AtParent().AtName("connector_provision"),
								path.MatchRelative().AtParent().AtName("delegated_provision"),
								path.MatchRelative().AtParent().AtName("external_ticket_provision"),
								path.MatchRelative().AtParent().AtName("manual_provision"),
								path.MatchRelative().AtParent().AtName("multi_step"),
							}...),
						},
					},
				},
				MarkdownDescription: `ProvisionPolicy is a oneOf that indicates how a provision step should be processed.` + "\n" +
					`` + "\n" +
					`This message contains a oneof named typ. Only a single field of the following list may be set at a time:` + "\n" +
					`  - connector` + "\n" +
					`  - manual` + "\n" +
					`  - delegated` + "\n" +
					`  - webhook` + "\n" +
					`  - multiStep` + "\n" +
					`  - externalTicket`,
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
	if resp.Diagnostics.HasError() {
		return
	}

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

	overrideAccessRequestsDefaults := new(bool)
	if !data.OverrideAccessRequestsDefaults.IsUnknown() && !data.OverrideAccessRequestsDefaults.IsNull() {
		*overrideAccessRequestsDefaults = data.OverrideAccessRequestsDefaults.ValueBool()
	} else {
		overrideAccessRequestsDefaults = nil
	}
	updateAppEntitlementRequest = &shared.UpdateAppEntitlementRequest{
		AppEntitlement:                 appEntitlement,
		OverrideAccessRequestsDefaults: overrideAccessRequestsDefaults,
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
	idParts := strings.Split(req.ID, "_")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: id_appId. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("app_id"), idParts[1])...)
}
