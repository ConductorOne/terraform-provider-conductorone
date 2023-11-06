// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"conductorone/internal/sdk"
	"conductorone/internal/sdk/pkg/models/operations"
	"context"
	"fmt"

	"conductorone/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &PolicyDataSource{}
var _ datasource.DataSourceWithConfigure = &PolicyDataSource{}

func NewPolicyDataSource() datasource.DataSource {
	return &PolicyDataSource{}
}

// PolicyDataSource is the data source implementation.
type PolicyDataSource struct {
	client *sdk.ConductoroneSDKTerraform
}

// PolicyDataSourceModel describes the data model.
type PolicyDataSourceModel struct {
	CreatedAt                types.String           `tfsdk:"created_at"`
	DeletedAt                types.String           `tfsdk:"deleted_at"`
	Description              types.String           `tfsdk:"description"`
	DisplayName              types.String           `tfsdk:"display_name"`
	ID                       types.String           `tfsdk:"id"`
	PolicySteps              map[string]PolicySteps `tfsdk:"policy_steps"`
	PolicyType               types.String           `tfsdk:"policy_type"`
	PostActions              []PolicyPostActions    `tfsdk:"post_actions"`
	ReassignTasksToDelegates types.Bool             `tfsdk:"reassign_tasks_to_delegates"`
	Rules                    []Rule                 `tfsdk:"rules"`
	SystemBuiltin            types.Bool             `tfsdk:"system_builtin"`
	UpdatedAt                types.String           `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *PolicyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policy"
}

// Schema defines the schema for the data source.
func (r *PolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Policy DataSource",

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
			"description": schema.StringAttribute{
				Computed:    true,
				Description: `The description of the new policy.`,
			},
			"display_name": schema.StringAttribute{
				Computed:    true,
				Description: `The display name of the new policy.`,
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `The ID of the Policy.`,
			},
			"policy_steps": schema.MapNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"steps": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"accept": schema.SingleNestedAttribute{
										Computed:    true,
										Attributes:  map[string]schema.Attribute{},
										Description: `This policy step indicates that a ticket should have an approved outcome. This is a terminal approval state and is used to explicitly define the end of approval steps.`,
									},
									"approval": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"allow_reassignment": schema.BoolAttribute{
												Computed:    true,
												Description: `Configuration to allow reassignment by reviewers during this step.`,
											},
											"app_group_approval": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"allow_self_approval": schema.BoolAttribute{
														Computed:    true,
														Description: `Configuration to allow self approval if the target user is a member of the group during this step.`,
													},
													"app_group_id": schema.StringAttribute{
														Computed:    true,
														Description: `The ID of the group specified for approval.`,
													},
													"app_id": schema.StringAttribute{
														Computed:    true,
														Description: `The ID of the app that conatins the group specified for approval.`,
													},
													"fallback": schema.BoolAttribute{
														Computed:    true,
														Description: `Configuration to allow a fallback if the group is empty.`,
													},
													"fallback_user_ids": schema.ListAttribute{
														Computed:    true,
														ElementType: types.StringType,
														Description: `Configuration to specific which users to fallback to if fallback is enabled and the group is empty.`,
													},
												},
												Description: `The AppGroupApproval object provides the configuration for setting a group as the approvers of an approval policy step.`,
											},
											"app_owner_approval": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"allow_self_approval": schema.BoolAttribute{
														Computed:    true,
														Description: `Configuration that allows a user to self approve if they are an app owner during this approval step.`,
													},
												},
												Description: `App owner approval provides the configuration for an approval step when the app owner is the target.`,
											},
											"assigned": schema.BoolAttribute{
												Computed:    true,
												Description: `A field indicating whether this step is assigned.`,
											},
											"entitlement_owner_approval": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"allow_self_approval": schema.BoolAttribute{
														Computed:    true,
														Description: `Configuration to allow self approval if the target user is an entitlement owner during this step.`,
													},
													"fallback": schema.BoolAttribute{
														Computed:    true,
														Description: `Configuration to allow a fallback if the entitlement owner cannot be identified.`,
													},
													"fallback_user_ids": schema.ListAttribute{
														Computed:    true,
														ElementType: types.StringType,
														Description: `Configuration to specific which users to fallback to if fallback is enabled and the entitlement owner cannot be identified.`,
													},
												},
												Description: `The entitlement owner approval allows configuration of the approval step when the target approvers are the entitlement owners.`,
											},
											"expression_approval": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"allow_self_approval": schema.BoolAttribute{
														Computed:    true,
														Description: `Configuration to allow self approval of if the user is specified and also the target of the ticket.`,
													},
													"assigned_user_ids": schema.ListAttribute{
														Computed:    true,
														ElementType: types.StringType,
														Description: `The assignedUserIds field.`,
													},
													"expressions": schema.ListAttribute{
														Computed:    true,
														ElementType: types.StringType,
														Description: `Array of dynamic expressions to determine the approvers.  The first expression to return a non-empty list of users will be used.`,
													},
													"fallback": schema.BoolAttribute{
														Computed:    true,
														Description: `Configuration to allow a fallback if the expression does not return a valid list of users.`,
													},
													"fallback_user_ids": schema.ListAttribute{
														Computed:    true,
														ElementType: types.StringType,
														Description: `Configuration to specific which users to fallback to if and the expression does not return a valid list of users.`,
													},
												},
												Description: `The ExpressionApproval message.`,
											},
											"manager_approval": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"allow_self_approval": schema.BoolAttribute{
														Computed:    true,
														Description: `Configuration to allow self approval if the target user is their own manager. This may occur if a service account has an identity user and manager specified as the same person.`,
													},
													"assigned_user_ids": schema.ListAttribute{
														Computed:    true,
														ElementType: types.StringType,
														Description: `The array of users determined to be the manager during processing time.`,
													},
													"fallback": schema.BoolAttribute{
														Computed:    true,
														Description: `Configuration to allow a fallback if no manager is found.`,
													},
													"fallback_user_ids": schema.ListAttribute{
														Computed:    true,
														ElementType: types.StringType,
														Description: `Configuration to specific which users to fallback to if fallback is enabled and no manager is found.`,
													},
												},
												Description: `The manager approval object provides configuration options for approval when the target of the approval is the manager of the user in the task.`,
											},
											"require_approval_reason": schema.BoolAttribute{
												Computed:    true,
												Description: `Configuration to require a reason when approving this step.`,
											},
											"require_reassignment_reason": schema.BoolAttribute{
												Computed:    true,
												Description: `Configuration to require a reason when reassigning this step.`,
											},
											"self_approval": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"assigned_user_ids": schema.ListAttribute{
														Computed:    true,
														ElementType: types.StringType,
														Description: `The array of users determined to be themselves during approval. This should only ever be one person, but is saved because it may change if the owner of an app user changes while the ticket is open.`,
													},
													"fallback": schema.BoolAttribute{
														Computed:    true,
														Description: `Configuration to allow a fallback if the identity user of the target app user cannot be determined.`,
													},
													"fallback_user_ids": schema.ListAttribute{
														Computed:    true,
														ElementType: types.StringType,
														Description: `Configuration to specific which users to fallback to if fallback is enabled and the identity user of the target app user cannot be determined.`,
													},
												},
												Description: `The self approval object describes the configuration of a policy step that needs to be approved by the target of the request.`,
											},
											"user_approval": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"allow_self_approval": schema.BoolAttribute{
														Computed:    true,
														Description: `Configuration to allow self approval of if the user is specified and also the target of the ticket.`,
													},
													"user_ids": schema.ListAttribute{
														Computed:    true,
														ElementType: types.StringType,
														Description: `Array of users configured for approval.`,
													},
												},
												Description: `The user approval object describes the approval configuration of a policy step that needs to be approved by a specific list of users.`,
											},
										},
										MarkdownDescription: `The Approval message.` + "\n" +
											`` + "\n" +
											`This message contains a oneof named typ. Only a single field of the following list may be set at a time:` + "\n" +
											`  - users` + "\n" +
											`  - manager` + "\n" +
											`  - appOwners` + "\n" +
											`  - group` + "\n" +
											`  - self` + "\n" +
											`  - entitlementOwners` + "\n" +
											`  - expression` + "\n" +
											``,
									},
									"provision": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"assigned": schema.BoolAttribute{
												Computed:    true,
												Description: `A field indicating whether this step is assigned.`,
											},
											"provision_policy": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"connector_provision": schema.SingleNestedAttribute{
														Computed:    true,
														Attributes:  map[string]schema.Attribute{},
														Description: `Indicates that a connector should perform the provisioning. This object has no fields.`,
													},
													"delegated_provision": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"app_id": schema.StringAttribute{
																Computed:    true,
																Description: `The AppID of the entitlement to delegate provisioning to.`,
															},
															"entitlement_id": schema.StringAttribute{
																Computed:    true,
																Description: `The ID of the entitlement we are delegating provisioning to.`,
															},
															"implicit": schema.BoolAttribute{
																Computed:    true,
																Description: `If true, a binding will be automatically created from the entitlement of the parent app.`,
															},
														},
														Description: `This provision step indicates that we should delegate provisioning to the configuration of another app entitlement. This app entitlement does not have to be one from the same app, but MUST be configured as a proxy binding leading into this entitlement.`,
													},
													"manual_provision": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"instructions": schema.StringAttribute{
																Computed:    true,
																Description: `This field indicates a text body of instructions for the provisioner to indicate.`,
															},
															"user_ids": schema.ListAttribute{
																Computed:    true,
																ElementType: types.StringType,
																Description: `An array of users that are required to provision during this step.`,
															},
														},
														Description: `Manual provisioning indicates that a human must intervene for the provisioning of this step.`,
													},
												},
												MarkdownDescription: `ProvisionPolicy is a oneOf that indicates how a provision step should be processed.` + "\n" +
													`` + "\n" +
													`This message contains a oneof named typ. Only a single field of the following list may be set at a time:` + "\n" +
													`  - connector` + "\n" +
													`  - manual` + "\n" +
													`  - delegated` + "\n" +
													``,
											},
											"provision_target": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"app_entitlement_id": schema.StringAttribute{
														Computed:    true,
														Description: `The app entitlement that should be provisioned.`,
													},
													"app_id": schema.StringAttribute{
														Computed:    true,
														Description: `The app in which the entitlement should be provisioned`,
													},
													"app_user_id": schema.StringAttribute{
														Computed:    true,
														Description: `The app user that should be provisioned. May be unset if the app user is unknown`,
													},
													"grant_duration": schema.StringAttribute{
														Computed: true,
													},
												},
												Description: `ProvisionTarget indicates the specific app, app entitlement, and if known, the app user and grant duration of this provision step`,
											},
										},
										Description: `The provision step references a provision policy for this step.`,
									},
									"reject": schema.SingleNestedAttribute{
										Computed:    true,
										Attributes:  map[string]schema.Attribute{},
										Description: `This policy step indicates that a ticket should have a denied outcome. This is a terminal approval state and is used to explicitly define the end of approval steps.`,
									},
								},
							},
							Description: `An array of policy steps indicating the processing flow of a policy. These steps are oneOfs, and only one property may be set for each array index at a time.`,
						},
					},
				},
				Description: `The map of policy type to policy steps. The key is the stringified version of the enum. See other policies for examples.`,
			},
			"policy_type": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"POLICY_TYPE_UNSPECIFIED",
						"POLICY_TYPE_GRANT",
						"POLICY_TYPE_REVOKE",
						"POLICY_TYPE_CERTIFY",
						"POLICY_TYPE_ACCESS_REQUEST",
						"POLICY_TYPE_PROVISION",
					),
				},
				MarkdownDescription: `must be one of ["POLICY_TYPE_UNSPECIFIED", "POLICY_TYPE_GRANT", "POLICY_TYPE_REVOKE", "POLICY_TYPE_CERTIFY", "POLICY_TYPE_ACCESS_REQUEST", "POLICY_TYPE_PROVISION"]` + "\n" +
					`The enum of the policy type.`,
			},
			"post_actions": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"certify_remediate_immediately": schema.BoolAttribute{
							Computed: true,
							MarkdownDescription: `ONLY valid when used in a CERTIFY Ticket Type:` + "\n" +
								` Causes any deprovision or change in a grant to be applied when Certify Ticket is closed.` + "\n" +
								`This field is part of the ` + "`" + `action` + "`" + ` oneof.` + "\n" +
								`See the documentation for ` + "`" + `c1.api.policy.v1.PolicyPostActions` + "`" + ` for more details.`,
						},
					},
				},
				Description: `Actions to occur after a policy finishes. As of now this is only valid on a certify policy to remediate a denied certification immediately.`,
			},
			"reassign_tasks_to_delegates": schema.BoolAttribute{
				Computed:    true,
				Description: `Allows reassigning tasks to delegates.`,
			},
			"rules": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"condition": schema.StringAttribute{
							Computed:    true,
							Description: `The condition field.`,
						},
						"policy_key": schema.StringAttribute{
							Computed:    true,
							Description: `This is a reference to a list of policy steps from ` + "`" + `policy_steps` + "`" + ``,
						},
					},
				},
				Description: `The rules field.`,
			},
			"system_builtin": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether this policy is a builtin system policy. Builtin system policies cannot be edited.`,
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

func (r *PolicyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.ConductoroneSDKTerraform)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.ConductoroneSDKTerraform, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *PolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *PolicyDataSourceModel
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
	request := operations.C1APIPolicyV1PoliciesGetRequest{
		ID: id,
	}
	res, err := r.client.Policies.Get(ctx, request)
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
	if res.GetPolicyResponse == nil || res.GetPolicyResponse.Policy == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.GetPolicyResponse.Policy)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
