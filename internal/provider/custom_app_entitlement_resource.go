package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/resourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/operations"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"github.com/conductorone/terraform-provider-conductorone/internal/validators"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CustomAppEntitlementResource{}
var _ resource.ResourceWithImportState = &CustomAppEntitlementResource{}

func NewCustomAppEntitlementResource() resource.Resource {
	return &CustomAppEntitlementResource{}
}

// CustomAppEntitlementResource defines the resource implementation.
type CustomAppEntitlementResource struct {
	client *sdk.ConductoroneAPI
}

// CustomAppEntitlementResourceModel describes the resource data model.
type CustomAppEntitlementResourceModel struct {
	Alias                          types.String            `tfsdk:"alias"`
	AppID                          types.String            `tfsdk:"app_id"`
	AppResourceID                  types.String            `tfsdk:"app_resource_id"`
	AppResourceTypeID              types.String            `tfsdk:"app_resource_type_id"`
	CertifyPolicyID                types.String            `tfsdk:"certify_policy_id"`
	ComplianceFrameworkValueIds    []types.String          `tfsdk:"compliance_framework_value_ids"`
	CreatedAt                      types.String            `tfsdk:"created_at"`
	DefaultValuesApplied           types.Bool              `tfsdk:"default_values_applied"`
	DeletedAt                      types.String            `tfsdk:"deleted_at"`
	Description                    types.String            `tfsdk:"description"`
	DisplayName                    types.String            `tfsdk:"display_name"`
	DurationGrant                  types.String            `tfsdk:"duration_grant"`
	DurationUnset                  *DurationUnset          `tfsdk:"duration_unset"`
	EmergencyGrantEnabled          types.Bool              `tfsdk:"emergency_grant_enabled"`
	EmergencyGrantPolicyID         types.String            `tfsdk:"emergency_grant_policy_id"`
	GrantCount                     types.String            `tfsdk:"grant_count"`
	GrantPolicyID                  types.String            `tfsdk:"grant_policy_id"`
	ID                             types.String            `tfsdk:"id"`
	IsAutomationEnabled            types.Bool              `tfsdk:"is_automation_enabled"`
	IsManuallyManaged              types.Bool              `tfsdk:"is_manually_managed"`
	OverrideAccessRequestsDefaults types.Bool              `tfsdk:"override_access_requests_defaults"`
	ProvisionPolicy                *ProvisionPolicy        `tfsdk:"provision_policy"`
	Purpose                        types.String            `tfsdk:"purpose"`
	RevokePolicyID                 types.String            `tfsdk:"revoke_policy_id"`
	RiskLevelValueID               types.String            `tfsdk:"risk_level_value_id"`
	Slug                           types.String            `tfsdk:"slug"`
	SourceConnectorIds             map[string]types.String `tfsdk:"source_connector_ids"`
	SystemBuiltin                  types.Bool              `tfsdk:"system_builtin"`
	UpdatedAt                      types.String            `tfsdk:"updated_at"`
}

func (r *CustomAppEntitlementResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_custom_app_entitlement"
}

func (r *CustomAppEntitlementResource) ConfigValidators(ctx context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{
		resourcevalidator.Conflicting(
			path.MatchRoot("duration_grant"),
			path.MatchRoot("duration_unset"),
		),
	}
}

func (r *CustomAppEntitlementResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "CustomAppEntitlement Resource",
		Attributes: map[string]schema.Attribute{
			"alias": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The alias field.`,
			},
			"app_id": schema.StringAttribute{
				Required: true,
			},
			"app_resource_id": schema.StringAttribute{
				Required:    true,
				Description: `The appResourceId field.`,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
			},
			"app_resource_type_id": schema.StringAttribute{
				Required:    true,
				Description: `The appResourceTypeId field.`,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
			},
			"certify_policy_id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The certifyPolicyId field.`,
			},
			"compliance_framework_value_ids": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `The complianceFrameworkValueIds field.`,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"default_values_applied": schema.BoolAttribute{
				Computed:    true,
				Description: `Flag to indicate if app-level access request defaults have been applied to the entitlement`,
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
				Required:    true,
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
				Description: `The emergencyGrantEnabled field.`,
			},
			"emergency_grant_policy_id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The emergencyGrantPolicyId field.`,
			},
			"grant_count": schema.StringAttribute{
				Computed:    true,
				Description: `The amount of grants open for this entitlement`,
			},
			"grant_policy_id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The grantPolicyId field.`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `The unique ID for the App Entitlement.`,
			},
			"is_automation_enabled": schema.BoolAttribute{
				Computed:    true,
				Description: `Flag to indicate whether automation (for adding users to entitlement based on rules) has been enabled.`,
			},
			"is_manually_managed": schema.BoolAttribute{
				Computed:    true,
				Description: `Flag to indicate if the app entitlement is manually managed.`,
			},
			"override_access_requests_defaults": schema.BoolAttribute{
				Computed:    true,
				Description: `The overrideAccessRequestsDefaults field.`,
			},
			"provision_policy": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"connector_provision": schema.SingleNestedAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Indicates that a connector should perform the provisioning. This object has no fields.`,
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
							"implicit": schema.BoolAttribute{
								Computed:    true,
								Optional:    true,
								Description: `If true, a binding will be automatically created from the entitlement of the parent app.`,
							},
						},
						Description: `This provision step indicates that we should delegate provisioning to the configuration of another app entitlement. This app entitlement does not have to be one from the same app, but MUST be configured as a proxy binding leading into this entitlement.`,
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
					},
					"multi_step": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `MultiStep indicates that this provision step has multiple steps to process. Parsed as JSON.`,
						Validators: []validator.String{
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
			"purpose": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The purpose field. must be one of ["APP_ENTITLEMENT_PURPOSE_VALUE_UNSPECIFIED", "APP_ENTITLEMENT_PURPOSE_VALUE_ASSIGNMENT", "APP_ENTITLEMENT_PURPOSE_VALUE_PERMISSION"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"APP_ENTITLEMENT_PURPOSE_VALUE_UNSPECIFIED",
						"APP_ENTITLEMENT_PURPOSE_VALUE_ASSIGNMENT",
						"APP_ENTITLEMENT_PURPOSE_VALUE_PERMISSION",
					),
				},
			},
			"revoke_policy_id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The revokePolicyId field.`,
			},
			"risk_level_value_id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The riskLevelValueId field.`,
			},
			"slug": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The slug field.`,
			},
			"source_connector_ids": schema.MapAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `Map to tell us which connector the entitlement came from.`,
			},
			"system_builtin": schema.BoolAttribute{
				Computed:    true,
				Description: `This field indicates if this is a system builtin entitlement.`,
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

func (r *CustomAppEntitlementResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *CustomAppEntitlementResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *CustomAppEntitlementResourceModel
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

	appID := data.AppID.ValueString()

	createAppEntitlementRequest := data.ToSharedCreateAppEntitlementRequest()
	request := operations.C1APIAppV1AppEntitlementsCreateRequest{
		AppID:                       appID,
		CreateAppEntitlementRequest: createAppEntitlementRequest,
	}
	res, err := r.client.AppEntitlements.Create(ctx, request)
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
	if !(res.CreateAppEntitlementResponse != nil && res.CreateAppEntitlementResponse.AppEntitlementView != nil && res.CreateAppEntitlementResponse.AppEntitlementView.AppEntitlement != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAppEntitlement(res.CreateAppEntitlementResponse.AppEntitlementView.AppEntitlement)
	appId1 := data.AppID.ValueString()

	id := data.ID.ValueString()

	request1 := operations.C1APIAppV1AppEntitlementsGetRequest{
		AppID: appId1,
		ID:    id,
	}
	res1, err := r.client.AppEntitlements.Get(ctx, request1)
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
	if !(res1.GetAppEntitlementResponse != nil && res1.GetAppEntitlementResponse.AppEntitlementView != nil && res1.GetAppEntitlementResponse.AppEntitlementView.AppEntitlement != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedAppEntitlement(res1.GetAppEntitlementResponse.AppEntitlementView.AppEntitlement)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CustomAppEntitlementResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *CustomAppEntitlementResourceModel
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
	if !(res.GetAppEntitlementResponse != nil && res.GetAppEntitlementResponse.AppEntitlementView != nil && res.GetAppEntitlementResponse.AppEntitlementView.AppEntitlement != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAppEntitlement(res.GetAppEntitlementResponse.AppEntitlementView.AppEntitlement)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CustomAppEntitlementResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *CustomAppEntitlementResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	appID := data.AppID.ValueString()

	id := data.ID.ValueString()

	var updateAppEntitlementRequest *shared.UpdateAppEntitlementRequest
	appEntitlement := data.ToSharedAppEntitlementInput()
	var appEntitlementExpandMask *shared.AppEntitlementExpandMask
	if data != nil {
		appEntitlementExpandMask = &shared.AppEntitlementExpandMask{}
	}
	overrideAccessRequestsDefaults := new(bool)
	*overrideAccessRequestsDefaults = true

	currentAppEntitlement, err := r.readAppEntitlementAndValidate(ctx, appID, id)
	if err != nil {
		resp.Diagnostics.AddError("failure reading app entitlement", err.Error())
		return
	}

	// TODO(mstanbCO): Need a better pattern for stuff like this
	// These two fields are a oneof in the API, so we need to ensure that only one is set before making the request.
	if currentAppEntitlement.DurationGrant != nil && appEntitlement.DurationUnset != nil {
		appEntitlement.DurationGrant = nil
	} else if currentAppEntitlement.DurationUnset != nil && appEntitlement.DurationGrant != nil {
		appEntitlement.DurationUnset = nil
	}

	// TODO(mstanbCO): Need a better pattern for handling this in the merge step, instead of doing this.
	if currentAppEntitlement.ProvisionPolicy != nil {
		isDelegatedSet := appEntitlement.ProvisionPolicy.DelegatedProvision != nil
		isManualSet := appEntitlement.ProvisionPolicy.ManualProvision != nil
		isWebhookSet := appEntitlement.ProvisionPolicy.WebhookProvision != nil
		isConnectorSet := appEntitlement.ProvisionPolicy.ConnectorProvision != nil
		isExternalTicketSet := appEntitlement.ProvisionPolicy.ExternalTicketProvision != nil
		isMultiStepSet := appEntitlement.ProvisionPolicy.MultiStep != nil

		if currentAppEntitlement.ProvisionPolicy.ConnectorProvision != nil {
			if isDelegatedSet || isManualSet || isWebhookSet || isMultiStepSet || isExternalTicketSet {
				appEntitlement.ProvisionPolicy.ConnectorProvision = nil
			}
		} else if currentAppEntitlement.ProvisionPolicy.DelegatedProvision != nil {
			if isConnectorSet || isManualSet || isWebhookSet || isMultiStepSet || isExternalTicketSet {
				appEntitlement.ProvisionPolicy.DelegatedProvision = nil
			}
		} else if currentAppEntitlement.ProvisionPolicy.ManualProvision != nil {
			if isConnectorSet || isDelegatedSet || isWebhookSet || isMultiStepSet || isExternalTicketSet {
				appEntitlement.ProvisionPolicy.ManualProvision = nil
			}
		} else if currentAppEntitlement.ProvisionPolicy.WebhookProvision != nil {
			if isConnectorSet || isDelegatedSet || isManualSet || isMultiStepSet || isExternalTicketSet {
				appEntitlement.ProvisionPolicy.WebhookProvision = nil
			}
		} else if currentAppEntitlement.ProvisionPolicy.ExternalTicketProvision != nil {
			if isConnectorSet || isDelegatedSet || isManualSet || isWebhookSet || isMultiStepSet {
				appEntitlement.ProvisionPolicy.ExternalTicketProvision = nil
			}
		} else if currentAppEntitlement.ProvisionPolicy.MultiStep != nil {
			if isConnectorSet || isDelegatedSet || isManualSet || isWebhookSet || isExternalTicketSet {
				appEntitlement.ProvisionPolicy.MultiStep = nil
			}
		}
	}

	updateAppEntitlementRequest = &shared.UpdateAppEntitlementRequest{
		AppEntitlement:                 appEntitlement,
		AppEntitlementExpandMask:       appEntitlementExpandMask,
		OverrideAccessRequestsDefaults: overrideAccessRequestsDefaults,
	}
	request := operations.C1APIAppV1AppEntitlementsUpdateRequest{
		AppID:                       appID,
		ID:                          id,
		UpdateAppEntitlementRequest: updateAppEntitlementRequest,
	}
	res, err := r.client.AppEntitlements.Update(ctx, request)
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
	if !(res.UpdateAppEntitlementResponse != nil && res.UpdateAppEntitlementResponse.AppEntitlementView != nil && res.UpdateAppEntitlementResponse.AppEntitlementView.AppEntitlement != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAppEntitlement(res.UpdateAppEntitlementResponse.AppEntitlementView.AppEntitlement)

	appId1 := data.AppID.ValueString()

	id1 := data.ID.ValueString()

	request1 := operations.C1APIAppV1AppEntitlementsGetRequest{
		AppID: appId1,
		ID:    id1,
	}
	res1, err := r.client.AppEntitlements.Get(ctx, request1)
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
	if !(res1.GetAppEntitlementResponse != nil && res1.GetAppEntitlementResponse.AppEntitlementView != nil && res1.GetAppEntitlementResponse.AppEntitlementView.AppEntitlement != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedAppEntitlement(res1.GetAppEntitlementResponse.AppEntitlementView.AppEntitlement)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CustomAppEntitlementResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *CustomAppEntitlementResourceModel
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

	deleteAppEntitlementRequest := data.ToSharedDeleteAppEntitlementRequest()
	request := operations.C1APIAppV1AppEntitlementsDeleteRequest{
		AppID:                       appID,
		ID:                          id,
		DeleteAppEntitlementRequest: deleteAppEntitlementRequest,
	}
	res, err := r.client.AppEntitlements.Delete(ctx, request)
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

func (r *CustomAppEntitlementResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	dec := json.NewDecoder(bytes.NewReader([]byte(req.ID)))
	dec.DisallowUnknownFields()
	var data struct {
		AppID string `json:"app_id"`
		ID    string `json:"id"`
	}

	if err := dec.Decode(&data); err != nil {
		resp.Diagnostics.AddError("Invalid ID", `The ID is not valid. It's expected to be a JSON object alike '{ "app_id": "",  "id": ""}': `+err.Error())
		return
	}

	if len(data.AppID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field app_id is required but was not found in the json encoded ID. It's expected to be a value alike '""`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("app_id"), data.AppID)...)
	if len(data.ID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field id is required but was not found in the json encoded ID. It's expected to be a value alike '""`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), data.ID)...)

}

func (r *CustomAppEntitlementResource) readAppEntitlementAndValidate(ctx context.Context, appID string, id string) (*shared.AppEntitlement, error) {
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
