---
page_title: "conductorone_policy Resource - terraform-provider-conductorone"
subcategory: ""
description: |-
  Policy Resource
---

# conductorone_policy (Resource)

Policy Resource

This resource allows you to create and configure a policy in ConductorOne.
When creating a policy you must provide a display name. Optionally you can define the steps of the policy, as well as configuring other options on the policy.

## Example Usage

```terraform
resource "conductorone_policy" "my_policy" {
  annotations = {
    key = "value"
  }
  baseline_policy_id = "...my_baseline_policy_id..."
  description        = "...my_description..."
  display_name       = "...my_display_name..."
  policy_steps = {
    key = {
      steps = [
        {
          accept = {
            accept_message = "...my_accept_message..."
          }
          action = {
            automation = {
              automation_template_id = "...my_automation_template_id..."
            }
            baton_resource_action = {
              baton_resource_action_id = "...my_baton_resource_action_id..."
            }
            client_id_approval = {
              # ...
            }
          }
          approval = {
            agent = {
              agent_failure_action = "APPROVAL_AGENT_FAILURE_ACTION_REASSIGN_TO_USERS"
              agent_mode           = "APPROVAL_AGENT_MODE_COMMENT_ONLY"
              agent_user_id        = "...my_agent_user_id..."
              instructions         = "...my_instructions..."
              policy_ids = [
                "..."
              ]
              reassign_to_user_ids = [
                "..."
              ]
            }
            allow_delegation   = true
            allow_reassignment = false
            allowed_reassignees = [
              "..."
            ]
            app_owners = {
              allow_self_approval        = false
              require_distinct_approvers = false
            }
            entitlement_owners = {
              allow_self_approval = true
              fallback            = false
              fallback_group_ids = [
                {
                  app_entitlement_id = "...my_app_entitlement_id..."
                  app_id             = "...my_app_id..."
                }
              ]
              fallback_user_ids = [
                "..."
              ]
              is_group_fallback_enabled  = false
              require_distinct_approvers = false
            }
            escalation = {
              cancel_ticket = {
                # ...
              }
              escalation_comment = "...my_escalation_comment..."
              expiration         = "...my_expiration..."
              reassign_to_approvers = {
                approver_ids = [
                  "..."
                ]
              }
              replace_policy = {
                policy_id = "...my_policy_id..."
              }
              skip_step = {
                # ...
              }
            }
            escalation_enabled = false
            expression = {
              allow_self_approval = false
              expressions = [
                "..."
              ]
              fallback = true
              fallback_group_ids = [
                {
                  app_entitlement_id = "...my_app_entitlement_id..."
                  app_id             = "...my_app_id..."
                }
              ]
              fallback_user_ids = [
                "..."
              ]
              is_group_fallback_enabled  = true
              require_distinct_approvers = false
            }
            group = {
              allow_self_approval = true
              app_group_id        = "...my_app_group_id..."
              app_id              = "...my_app_id..."
              fallback            = false
              fallback_group_ids = [
                {
                  app_entitlement_id = "...my_app_entitlement_id..."
                  app_id             = "...my_app_id..."
                }
              ]
              fallback_user_ids = [
                "..."
              ]
              is_group_fallback_enabled  = true
              require_distinct_approvers = false
            }
            manager = {
              allow_self_approval = false
              fallback            = true
              fallback_group_ids = [
                {
                  app_entitlement_id = "...my_app_entitlement_id..."
                  app_id             = "...my_app_id..."
                }
              ]
              fallback_user_ids = [
                "..."
              ]
              is_group_fallback_enabled  = true
              require_distinct_approvers = true
            }
            require_approval_reason      = true
            require_denial_reason        = true
            require_reassignment_reason  = true
            requires_step_up_provider_id = "...my_requires_step_up_provider_id..."
            resource_owners = {
              allow_self_approval = true
              fallback            = false
              fallback_group_ids = [
                {
                  app_entitlement_id = "...my_app_entitlement_id..."
                  app_id             = "...my_app_id..."
                }
              ]
              fallback_user_ids = [
                "..."
              ]
              is_group_fallback_enabled  = true
              require_distinct_approvers = true
            }
            self = {
              fallback = false
              fallback_group_ids = [
                {
                  app_entitlement_id = "...my_app_entitlement_id..."
                  app_id             = "...my_app_id..."
                }
              ]
              fallback_user_ids = [
                "..."
              ]
              is_group_fallback_enabled = true
            }
            users = {
              allow_self_approval        = false
              require_distinct_approvers = false
              user_ids = [
                "..."
              ]
            }
            webhook = {
              webhook_id = "...my_webhook_id..."
            }
          }
          form = "{ \"see\": \"documentation\" }"
          provision = {
            assigned = true
            provision_policy = {
              action = {
                action_name  = "...my_action_name..."
                app_id       = "...my_app_id..."
                connector_id = "...my_connector_id..."
                display_name = "...my_display_name..."
              }
              connector = {
                account = {
                  config       = "{ \"see\": \"documentation\" }"
                  connector_id = "...my_connector_id..."
                  do_not_save = {
                    # ...
                  }
                  save_to_vault = {
                    vault_ids = [
                      "..."
                    ]
                  }
                  schema_id = "...my_schema_id..."
                }
                default_behavior = {
                  connector_id = "...my_connector_id..."
                }
                delete_account = {
                  connector_id = "...my_connector_id..."
                }
              }
              delegated = {
                app_id         = "...my_app_id..."
                entitlement_id = "...my_entitlement_id..."
              }
              device_placement = {
                vault_boundary_id = "...my_vault_boundary_id..."
              }
              external_ticket = {
                app_id                                = "...my_app_id..."
                connector_id                          = "...my_connector_id..."
                external_ticket_provisioner_config_id = "...my_external_ticket_provisioner_config_id..."
                instructions                          = "...my_instructions..."
              }
              manual = {
                assignee = {
                  app_owners = {
                    allow_reassignment = true
                    fallback_user_ids = [
                      "..."
                    ]
                  }
                  entitlement_owners = {
                    allow_reassignment = true
                    fallback_user_ids = [
                      "..."
                    ]
                  }
                  expression = {
                    allow_reassignment = false
                    expressions = [
                      "..."
                    ]
                    fallback_user_ids = [
                      "..."
                    ]
                  }
                  group = {
                    allow_reassignment = true
                    app_group_id       = "...my_app_group_id..."
                    app_id             = "...my_app_id..."
                    fallback_user_ids = [
                      "..."
                    ]
                  }
                  manager = {
                    allow_reassignment = true
                    fallback_user_ids = [
                      "..."
                    ]
                  }
                  users = {
                    allow_reassignment = false
                    user_ids = [
                      "..."
                    ]
                  }
                }
                instructions = "...my_instructions..."
                user_ids = [
                  "..."
                ]
              }
              multi_step = "{ \"see\": \"documentation\" }"
              unconfigured = {
                # ...
              }
              webhook = {
                webhook_id = "...my_webhook_id..."
              }
            }
            provision_target = {
              app_entitlement_id = "...my_app_entitlement_id..."
              app_id             = "...my_app_id..."
              app_user_id        = "...my_app_user_id..."
              grant_duration     = "...my_grant_duration..."
            }
          }
          reject = {
            reject_message = "...my_reject_message..."
          }
          wait = {
            comment_on_first_wait = "...my_comment_on_first_wait..."
            comment_on_timeout    = "...my_comment_on_timeout..."
            condition = {
              condition = "...my_condition..."
            }
            duration = {
              duration = "...my_duration..."
            }
            name             = "...my_name..."
            timeout_duration = "...my_timeout_duration..."
            until_time = {
              hours    = 8
              minutes  = 9
              timezone = "...my_timezone..."
            }
          }
        }
      ]
    }
  }
  policy_type = "POLICY_TYPE_UNSPECIFIED"
  post_actions = [
    {
      certify_remediate_immediately = false
    }
  ]
  reassign_tasks_to_delegates = false
  rules = [
    {
      condition  = "...my_condition..."
      policy_id  = "...my_policy_id..."
      policy_key = "...my_policy_key..."
      step_key   = "...my_step_key..."
    }
  ]
  scope = {
    app_entitlement_id = "...my_app_entitlement_id..."
    app_id             = "...my_app_id..."
    slot               = "POLICY_SCOPE_SLOT_EMERGENCY"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `display_name` (String) The display name of the new policy.

### Optional

- `annotations` (Map of String) Bounded key/value metadata bag for IaC marking and customer tags.
 See .rfcs/object-annotations.md §2. Limits: ≤16 entries; keys 1–128
 chars matching ^[A-Za-z][A-Za-z0-9._/-]{0,127}$; values 0–256 chars
 matching URL-safe ASCII; total serialized ≤4096 bytes. Keys starting
 with `c1/` are reserved for server-managed use and rejected on write.

 Well-known keys: `managed_by`, `iac_workspace`,
 `iac_resource_address`, `iac_tool_version`.
- `baseline_policy_id` (String) When set, the new policy's baseline defers to another policy of the same
 type when no rule matches, instead of an inline baseline step list.
 Mutually exclusive with the baseline entry in policy_steps. Requires the
 POLICY_REFERENCES_POLICY feature; obeys the same depth/cycle/self rules as
 Rule.policy_id.
- `description` (String) The description of the new policy.
- `policy_steps` (Attributes Map) Step sequences for this policy. The map must include a baseline entry keyed
 by the lowercased policy type (e.g., "grant"). Additional entries with
 opaque keys can be added for conditional routing via the rules array. (see [below for nested schema](#nestedatt--policy_steps))
- `policy_type` (String) The type of policy to create (grant, revoke, or certify). possible known values include one of ["POLICY_TYPE_UNSPECIFIED", "POLICY_TYPE_GRANT", "POLICY_TYPE_REVOKE", "POLICY_TYPE_CERTIFY", "POLICY_TYPE_ACCESS_REQUEST", "POLICY_TYPE_PROVISION"]
- `post_actions` (Attributes List) Ordered actions to execute after the policy completes processing. (see [below for nested schema](#nestedatt--post_actions))
- `reassign_tasks_to_delegates` (Boolean, Deprecated) This field is no longer used. Configure delegate reassignment in the policy step instead.
- `rules` (Attributes List) Conditional routing rules. See the Policy message for details on evaluation order. (see [below for nested schema](#nestedatt--rules))
- `scope` (Attributes) Scopes a policy to an app or to a single entitlement within an app. (see [below for nested schema](#nestedatt--scope))

### Read-Only

- `created_at` (String)
- `id` (String) The ID of the Policy.
- `system_builtin` (Boolean) Whether this policy is a builtin system policy. Builtin system policies cannot be edited.
- `updated_at` (String)

<a id="nestedatt--policy_steps"></a>
### Nested Schema for `policy_steps`

Optional:

- `steps` (Attributes List) Ordered array of steps. Each step is a oneof -- exactly one step type is
 set per entry. Steps execute sequentially. (see [below for nested schema](#nestedatt--policy_steps--steps))

<a id="nestedatt--policy_steps--steps"></a>
### Nested Schema for `policy_steps.steps`

Optional:

- `accept` (Attributes) This policy step indicates that a ticket should have an approved outcome. This is a terminal approval state and is used to explicitly define the end of approval steps. (see [below for nested schema](#nestedatt--policy_steps--steps--accept))
- `action` (Attributes) The Action message.

This message contains a oneof named target. Only a single field of the following list may be set at a time:
  - automation
  - batonResourceAction
  - clientIdApproval (see [below for nested schema](#nestedatt--policy_steps--steps--action))
- `approval` (Attributes) The Approval message.

This message contains a oneof named typ. Only a single field of the following list may be set at a time:
  - users
  - manager
  - appOwners
  - group
  - self
  - entitlementOwners
  - expression
  - webhook
  - resourceOwners
  - agent (see [below for nested schema](#nestedatt--policy_steps--steps--approval))
- `form` (String) Parsed as JSON.
- `provision` (Attributes) The provision step references a provision policy for this step. (see [below for nested schema](#nestedatt--policy_steps--steps--provision))
- `reject` (Attributes) This policy step indicates that a ticket should have a denied outcome. This is a terminal approval state and is used to explicitly define the end of approval steps. (see [below for nested schema](#nestedatt--policy_steps--steps--reject))
- `wait` (Attributes) Define a Wait step for a policy to wait on a condition to be met.

This message contains a oneof named until. Only a single field of the following list may be set at a time:
  - condition
  - duration
  - untilTime (see [below for nested schema](#nestedatt--policy_steps--steps--wait))

<a id="nestedatt--policy_steps--steps--accept"></a>
### Nested Schema for `policy_steps.steps.accept`

Optional:

- `accept_message` (String) An optional message to include in the comments when a task is automatically accepted.


<a id="nestedatt--policy_steps--steps--action"></a>
### Nested Schema for `policy_steps.steps.action`

Optional:

- `automation` (Attributes) ActionTargetAutomation targets automation templates for policy actions. (see [below for nested schema](#nestedatt--policy_steps--steps--action--automation))
- `baton_resource_action` (Attributes) ActionTargetResource targets resource actions for policy actions. (see [below for nested schema](#nestedatt--policy_steps--steps--action--baton_resource_action))
- `client_id_approval` (Attributes) ActionTargetClientIdApproval targets administrator review of an external
 OAuth client registration (CIMD or DCR) for policy actions. (see [below for nested schema](#nestedatt--policy_steps--steps--action--client_id_approval))

<a id="nestedatt--policy_steps--steps--action--automation"></a>
### Nested Schema for `policy_steps.steps.action.automation`

Optional:

- `automation_template_id` (String) The automationTemplateId field.


<a id="nestedatt--policy_steps--steps--action--baton_resource_action"></a>
### Nested Schema for `policy_steps.steps.action.baton_resource_action`

Optional:

- `baton_resource_action_id` (String) The batonResourceActionId field.


<a id="nestedatt--policy_steps--steps--action--client_id_approval"></a>
### Nested Schema for `policy_steps.steps.action.client_id_approval`



<a id="nestedatt--policy_steps--steps--approval"></a>
### Nested Schema for `policy_steps.steps.approval`

Optional:

- `agent` (Attributes) The agent to assign the task to. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--agent))
- `allow_delegation` (Boolean) Whether ticket delegation is allowed for this step.
- `allow_reassignment` (Boolean) Configuration to allow reassignment by reviewers during this step.
- `allowed_reassignees` (List of String) List of users for whom this step can be reassigned.
- `app_owners` (Attributes) App owner approval provides the configuration for an approval step when the app owner is the target. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--app_owners))
- `entitlement_owners` (Attributes) The entitlement owner approval allows configuration of the approval step when the target approvers are the entitlement owners. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--entitlement_owners))
- `escalation` (Attributes) The Escalation message.

This message contains a oneof named escalation_policy. Only a single field of the following list may be set at a time:
  - replacePolicy
  - reassignToApprovers
  - cancelTicket
  - skipStep (see [below for nested schema](#nestedatt--policy_steps--steps--approval--escalation))
- `escalation_enabled` (Boolean) Whether escalation is enabled for this step.
- `expression` (Attributes) The ExpressionApproval message. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--expression))
- `group` (Attributes) The AppGroupApproval object provides the configuration for setting a group as the approvers of an approval policy step. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--group))
- `manager` (Attributes) The manager approval object provides configuration options for approval when the target of the approval is the manager of the user in the task. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--manager))
- `require_approval_reason` (Boolean) Configuration to require a reason when approving this step.
- `require_denial_reason` (Boolean) Configuration to require a reason when denying this step.
- `require_reassignment_reason` (Boolean) Configuration to require a reason when reassigning this step.
- `requires_step_up_provider_id` (String) The ID of a step-up authentication provider that will be required for approvals on this step.
 If set, approvers must complete the step-up authentication flow before they can approve.
- `resource_owners` (Attributes) The resource owner approval allows configuration of the approval step when the target approvers are the resource owners. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--resource_owners))
- `self` (Attributes) The self approval object describes the configuration of a policy step that needs to be approved by the target of the request. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--self))
- `users` (Attributes) The user approval object describes the approval configuration of a policy step that needs to be approved by a specific list of users. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--users))
- `webhook` (Attributes) The WebhookApproval message. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--webhook))

Read-Only:

- `assigned` (Boolean) A field indicating whether this step is assigned.

<a id="nestedatt--policy_steps--steps--approval--agent"></a>
### Nested Schema for `policy_steps.steps.approval.agent`

Optional:

- `agent_failure_action` (String) The action to take if the agent fails to approve, deny, or reassign the task. possible known values include one of ["APPROVAL_AGENT_FAILURE_ACTION_UNSPECIFIED", "APPROVAL_AGENT_FAILURE_ACTION_REASSIGN_TO_USERS", "APPROVAL_AGENT_FAILURE_ACTION_REASSIGN_TO_SUPER_ADMINS", "APPROVAL_AGENT_FAILURE_ACTION_SKIP_POLICY_STEP"]
- `agent_mode` (String) The mode of the agent, full control, change policy only, or comment only. possible known values include one of ["APPROVAL_AGENT_MODE_UNSPECIFIED", "APPROVAL_AGENT_MODE_FULL_CONTROL", "APPROVAL_AGENT_MODE_CHANGE_POLICY_ONLY", "APPROVAL_AGENT_MODE_COMMENT_ONLY"]
- `agent_user_id` (String, Deprecated) Deprecated: agent steps are evaluated by the system; no agent user is
 selected. Retained so pre-migration policies still validate.
- `instructions` (String) Instructions for the agent.
- `policy_ids` (List of String) The allow list of policy IDs to re-route the task to.
- `reassign_to_user_ids` (List of String) The users to reassign the task to if the agent failure action is reassign to users.


<a id="nestedatt--policy_steps--steps--approval--app_owners"></a>
### Nested Schema for `policy_steps.steps.approval.app_owners`

Optional:

- `allow_self_approval` (Boolean) Configuration that allows a user to self approve if they are an app owner during this approval step.
- `require_distinct_approvers` (Boolean) Configuration to require distinct approvers across approval steps of a rule.


<a id="nestedatt--policy_steps--steps--approval--entitlement_owners"></a>
### Nested Schema for `policy_steps.steps.approval.entitlement_owners`

Optional:

- `allow_self_approval` (Boolean) Configuration to allow self approval if the target user is an entitlement owner during this step.
- `fallback` (Boolean) Configuration to allow a fallback if the entitlement owner cannot be identified.
- `fallback_group_ids` (Attributes List) Configuration to specify which groups to fallback to if fallback is enabled and the entitlement owner cannot be identified. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--entitlement_owners--fallback_group_ids))
- `fallback_user_ids` (List of String) Configuration to specific which users to fallback to if fallback is enabled and the entitlement owner cannot be identified.
- `is_group_fallback_enabled` (Boolean) Configuration to enable fallback for group fallback.
- `require_distinct_approvers` (Boolean) Configuration to require distinct approvers across approval steps of a rule.

<a id="nestedatt--policy_steps--steps--approval--entitlement_owners--fallback_group_ids"></a>
### Nested Schema for `policy_steps.steps.approval.entitlement_owners.fallback_group_ids`

Optional:

- `app_entitlement_id` (String) The ID of the Entitlement.
- `app_id` (String) The ID of the App this entitlement belongs to.



<a id="nestedatt--policy_steps--steps--approval--escalation"></a>
### Nested Schema for `policy_steps.steps.approval.escalation`

Optional:

- `cancel_ticket` (Attributes) The CancelTicket message. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--escalation--cancel_ticket))
- `escalation_comment` (String) The escalationComment field.
- `expiration` (String) The expiration field.
- `reassign_to_approvers` (Attributes) The ReassignToApprovers message. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--escalation--reassign_to_approvers))
- `replace_policy` (Attributes) The ReplacePolicy message. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--escalation--replace_policy))
- `skip_step` (Attributes) The SkipStep message. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--escalation--skip_step))

<a id="nestedatt--policy_steps--steps--approval--escalation--cancel_ticket"></a>
### Nested Schema for `policy_steps.steps.approval.escalation.cancel_ticket`


<a id="nestedatt--policy_steps--steps--approval--escalation--reassign_to_approvers"></a>
### Nested Schema for `policy_steps.steps.approval.escalation.reassign_to_approvers`

Optional:

- `approver_ids` (List of String) The approverIds field.


<a id="nestedatt--policy_steps--steps--approval--escalation--replace_policy"></a>
### Nested Schema for `policy_steps.steps.approval.escalation.replace_policy`

Optional:

- `policy_id` (String) The policyId field.


<a id="nestedatt--policy_steps--steps--approval--escalation--skip_step"></a>
### Nested Schema for `policy_steps.steps.approval.escalation.skip_step`



<a id="nestedatt--policy_steps--steps--approval--expression"></a>
### Nested Schema for `policy_steps.steps.approval.expression`

Optional:

- `allow_self_approval` (Boolean) Configuration to allow self approval of if the user is specified and also the target of the ticket.
- `expressions` (List of String) Array of dynamic expressions to determine the approvers.  The first expression to return a non-empty list of users will be used.
- `fallback` (Boolean) Configuration to allow a fallback if the expression does not return a valid list of users.
- `fallback_group_ids` (Attributes List) Configuration to specify which groups to fallback to if fallback is enabled and the expression does not return a valid list of users. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--expression--fallback_group_ids))
- `fallback_user_ids` (List of String) Configuration to specific which users to fallback to if and the expression does not return a valid list of users.
- `is_group_fallback_enabled` (Boolean) Configuration to enable fallback for group fallback.
- `require_distinct_approvers` (Boolean) Configuration to require distinct approvers across approval steps of a rule.

Read-Only:

- `assigned_user_ids` (List of String) The assignedUserIds field.

<a id="nestedatt--policy_steps--steps--approval--expression--fallback_group_ids"></a>
### Nested Schema for `policy_steps.steps.approval.expression.fallback_group_ids`

Optional:

- `app_entitlement_id` (String) The ID of the Entitlement.
- `app_id` (String) The ID of the App this entitlement belongs to.



<a id="nestedatt--policy_steps--steps--approval--group"></a>
### Nested Schema for `policy_steps.steps.approval.group`

Optional:

- `allow_self_approval` (Boolean) Configuration to allow self approval if the target user is a member of the group during this step.
- `app_group_id` (String) The app entitlement ID of the group specified for approval (not the group resource ID). Use the conductorone_app_entitlement data source to look up the correct entitlement ID.
- `app_id` (String) The ID of the app that contains the group specified for approval.
- `fallback` (Boolean) Configuration to allow a fallback if the group is empty.
- `fallback_group_ids` (Attributes List) Configuration to specify which groups to fallback to if fallback is enabled and the group is empty. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--group--fallback_group_ids))
- `fallback_user_ids` (List of String) Configuration to specific which users to fallback to if fallback is enabled and the group is empty.
- `is_group_fallback_enabled` (Boolean) Configuration to enable fallback for group fallback.
- `require_distinct_approvers` (Boolean) Configuration to require distinct approvers across approval steps of a rule.

<a id="nestedatt--policy_steps--steps--approval--group--fallback_group_ids"></a>
### Nested Schema for `policy_steps.steps.approval.group.fallback_group_ids`

Optional:

- `app_entitlement_id` (String) The ID of the Entitlement.
- `app_id` (String) The ID of the App this entitlement belongs to.



<a id="nestedatt--policy_steps--steps--approval--manager"></a>
### Nested Schema for `policy_steps.steps.approval.manager`

Optional:

- `allow_self_approval` (Boolean) Configuration to allow self approval if the target user is their own manager. This may occur if a service account has an identity user and manager specified as the same person.
- `fallback` (Boolean) Configuration to allow a fallback if no manager is found.
- `fallback_group_ids` (Attributes List) Configuration to specify which groups to fallback to if fallback is enabled and no manager is found. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--manager--fallback_group_ids))
- `fallback_user_ids` (List of String) Configuration to specific which users to fallback to if fallback is enabled and no manager is found.
- `is_group_fallback_enabled` (Boolean) Configuration to enable fallback for group fallback.
- `require_distinct_approvers` (Boolean) Configuration to require distinct approvers across approval steps of a rule.

Read-Only:

- `assigned_user_ids` (List of String) The array of users determined to be the manager during processing time.

<a id="nestedatt--policy_steps--steps--approval--manager--fallback_group_ids"></a>
### Nested Schema for `policy_steps.steps.approval.manager.fallback_group_ids`

Optional:

- `app_entitlement_id` (String) The ID of the Entitlement.
- `app_id` (String) The ID of the App this entitlement belongs to.



<a id="nestedatt--policy_steps--steps--approval--resource_owners"></a>
### Nested Schema for `policy_steps.steps.approval.resource_owners`

Optional:

- `allow_self_approval` (Boolean) Configuration to allow self approval if the target user is an resource owner during this step.
- `fallback` (Boolean) Configuration to allow a fallback if the resource owner cannot be identified.
- `fallback_group_ids` (Attributes List) Configuration to specify which groups to fallback to if fallback is enabled and the resource owner cannot be identified. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--resource_owners--fallback_group_ids))
- `fallback_user_ids` (List of String) Configuration to specific which users to fallback to if fallback is enabled and the resource owner cannot be identified.
- `is_group_fallback_enabled` (Boolean) Configuration to enable fallback for group fallback.
- `require_distinct_approvers` (Boolean) Configuration to require distinct approvers across approval steps of a rule.

<a id="nestedatt--policy_steps--steps--approval--resource_owners--fallback_group_ids"></a>
### Nested Schema for `policy_steps.steps.approval.resource_owners.fallback_group_ids`

Optional:

- `app_entitlement_id` (String) The ID of the Entitlement.
- `app_id` (String) The ID of the App this entitlement belongs to.



<a id="nestedatt--policy_steps--steps--approval--self"></a>
### Nested Schema for `policy_steps.steps.approval.self`

Optional:

- `fallback` (Boolean) Configuration to allow a fallback if the identity user of the target app user cannot be determined.
- `fallback_group_ids` (Attributes List) Configuration to specify which groups to fallback to if fallback is enabled and the identity user of the target app user cannot be determined. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--self--fallback_group_ids))
- `fallback_user_ids` (List of String) Configuration to specific which users to fallback to if fallback is enabled and the identity user of the target app user cannot be determined.
- `is_group_fallback_enabled` (Boolean) Configuration to enable fallback for group fallback.

Read-Only:

- `assigned_user_ids` (List of String) The array of users determined to be themselves during approval. This should only ever be one person, but is saved because it may change if the owner of an app user changes while the ticket is open.

<a id="nestedatt--policy_steps--steps--approval--self--fallback_group_ids"></a>
### Nested Schema for `policy_steps.steps.approval.self.fallback_group_ids`

Optional:

- `app_entitlement_id` (String) The ID of the Entitlement.
- `app_id` (String) The ID of the App this entitlement belongs to.



<a id="nestedatt--policy_steps--steps--approval--users"></a>
### Nested Schema for `policy_steps.steps.approval.users`

Optional:

- `allow_self_approval` (Boolean) Configuration to allow self approval of if the user is specified and also the target of the ticket.
- `require_distinct_approvers` (Boolean) Configuration to require distinct approvers across approval steps of a rule.
- `user_ids` (List of String) Array of users configured for approval.


<a id="nestedatt--policy_steps--steps--approval--webhook"></a>
### Nested Schema for `policy_steps.steps.approval.webhook`

Optional:

- `webhook_id` (String) The ID of the webhook to call for approval.



<a id="nestedatt--policy_steps--steps--provision"></a>
### Nested Schema for `policy_steps.steps.provision`

Optional:

- `assigned` (Boolean) A field indicating whether this step is assigned.
- `provision_policy` (Attributes) ProvisionPolicy is a oneOf that indicates how a provision step should be processed.

This message contains a oneof named typ. Only a single field of the following list may be set at a time:
  - connector
  - manual
  - delegated
  - webhook
  - multiStep
  - externalTicket
  - unconfigured
  - action
  - devicePlacement (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy))
- `provision_target` (Attributes) ProvisionTarget indicates the specific app, app entitlement, and if known, the app user and grant duration of this provision step (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_target))

<a id="nestedatt--policy_steps--steps--provision--provision_policy"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy`

Optional:

- `action` (Attributes) This provision step indicates that account lifecycle action should be called to provision this entitlement. (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--action))
- `connector` (Attributes) Indicates that a connector should perform the provisioning. This object has no fields.

This message contains a oneof named provision_type. Only a single field of the following list may be set at a time:
  - defaultBehavior
  - account
  - deleteAccount (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--connector))
- `delegated` (Attributes) This provision step indicates that we should delegate provisioning to the configuration of another app entitlement. This app entitlement does not have to be one from the same app, but MUST be configured as a proxy binding leading into this entitlement. (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--delegated))
- `device_placement` (Attributes) This provision step is fulfilled by a Latchkey member device producing an MLS Welcome for the recipient. It has no assignee and no instructions because the step is not human-actionable. (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--device_placement))
- `external_ticket` (Attributes) This provision step indicates that we should check an external ticket to provision this entitlement (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--external_ticket))
- `manual` (Attributes) Manual provisioning indicates that a human must intervene for the provisioning of this step. (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--manual))
- `multi_step` (String) Parsed as JSON.
- `unconfigured` (Attributes) The UnconfiguredProvision message. (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--unconfigured))
- `webhook` (Attributes) This provision step indicates that a webhook should be called to provision this entitlement. (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--webhook))

<a id="nestedatt--policy_steps--steps--provision--provision_policy--action"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.action`

Optional:

- `action_name` (String) The actionName field.
- `app_id` (String) The appId field.
- `connector_id` (String) The connectorId field.
- `display_name` (String) The displayName field.


<a id="nestedatt--policy_steps--steps--provision--provision_policy--connector"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.connector`

Optional:

- `account` (Attributes) The AccountProvision message.

This message contains a oneof named storage_type. Only a single field of the following list may be set at a time:
  - saveToVault
  - doNotSave (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--connector--account))
- `default_behavior` (Attributes) The DefaultBehavior message. (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--connector--default_behavior))
- `delete_account` (Attributes) The DeleteAccount message. (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--connector--delete_account))

<a id="nestedatt--policy_steps--steps--provision--provision_policy--connector--account"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.connector.account`

Optional:

- `config` (String) Parsed as JSON.
- `connector_id` (String) The connectorId field.
- `do_not_save` (Attributes) The DoNotSave message. (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--connector--account--do_not_save))
- `save_to_vault` (Attributes) The SaveToVault message. (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--connector--account--save_to_vault))
- `schema_id` (String) The schemaId field.

<a id="nestedatt--policy_steps--steps--provision--provision_policy--connector--account--do_not_save"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.connector.account.do_not_save`


<a id="nestedatt--policy_steps--steps--provision--provision_policy--connector--account--save_to_vault"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.connector.account.save_to_vault`

Optional:

- `vault_ids` (List of String) The vaultIds field.



<a id="nestedatt--policy_steps--steps--provision--provision_policy--connector--default_behavior"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.connector.default_behavior`

Optional:

- `connector_id` (String) this checks if the entitlement is enabled by provisioning in a specific connector
 this can happen automatically and doesn't need any extra info


<a id="nestedatt--policy_steps--steps--provision--provision_policy--connector--delete_account"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.connector.delete_account`

Optional:

- `connector_id` (String) The connectorId field.



<a id="nestedatt--policy_steps--steps--provision--provision_policy--delegated"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.delegated`

Optional:

- `app_id` (String) The AppID of the entitlement to delegate provisioning to.
- `entitlement_id` (String) The ID of the entitlement we are delegating provisioning to.


<a id="nestedatt--policy_steps--steps--provision--provision_policy--device_placement"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.device_placement`

Optional:

- `vault_boundary_id` (String) The vaultBoundaryId field.


<a id="nestedatt--policy_steps--steps--provision--provision_policy--external_ticket"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.external_ticket`

Optional:

- `app_id` (String) The appId field.
- `connector_id` (String) The connectorId field.
- `external_ticket_provisioner_config_id` (String) The externalTicketProvisionerConfigId field.
- `instructions` (String) This field indicates a text body of instructions for the provisioner to indicate.


<a id="nestedatt--policy_steps--steps--provision--provision_policy--manual"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.manual`

Optional:

- `assignee` (Attributes) ProvisionerAssignment defines how a provisioner is dynamically assigned.

This message contains a oneof named typ. Only a single field of the following list may be set at a time:
  - users
  - appOwners
  - group
  - manager
  - expression
  - entitlementOwners (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--manual--assignee))
- `instructions` (String) This field indicates a text body of instructions for the provisioner to indicate.
- `user_ids` (List of String) An array of users that are required to provision during this step.
 Deprecated: Use assignee field instead for dynamic provisioner assignment.

<a id="nestedatt--policy_steps--steps--provision--provision_policy--manual--assignee"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.manual.assignee`

Optional:

- `app_owners` (Attributes) AppOwnerProvisioner resolves to app owners. (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--manual--assignee--app_owners))
- `entitlement_owners` (Attributes) EntitlementOwnerProvisioner resolves to entitlement owners. (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--manual--assignee--entitlement_owners))
- `expression` (Attributes) ExpressionProvisioner evaluates CEL expressions to determine provisioners. (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--manual--assignee--expression))
- `group` (Attributes) GroupProvisioner resolves to members of a specific group. (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--manual--assignee--group))
- `manager` (Attributes) ManagerProvisioner resolves to the user's manager. (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--manual--assignee--manager))
- `users` (Attributes) UserProvisioner assigns specific users as provisioners. (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--manual--assignee--users))

<a id="nestedatt--policy_steps--steps--provision--provision_policy--manual--assignee--app_owners"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.manual.assignee.app_owners`

Optional:

- `allow_reassignment` (Boolean) Whether the provisioner can reassign the task.
- `fallback_user_ids` (List of String) Fallback user IDs if no app owners are found.


<a id="nestedatt--policy_steps--steps--provision--provision_policy--manual--assignee--entitlement_owners"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.manual.assignee.entitlement_owners`

Optional:

- `allow_reassignment` (Boolean) Whether the provisioner can reassign the task.
- `fallback_user_ids` (List of String) Fallback user IDs if no entitlement owners are found.


<a id="nestedatt--policy_steps--steps--provision--provision_policy--manual--assignee--expression"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.manual.assignee.expression`

Optional:

- `allow_reassignment` (Boolean) Whether the provisioner can reassign the task.
- `expressions` (List of String) The CEL expressions to evaluate.
- `fallback_user_ids` (List of String) Fallback user IDs if expression evaluation yields no users.


<a id="nestedatt--policy_steps--steps--provision--provision_policy--manual--assignee--group"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.manual.assignee.group`

Optional:

- `allow_reassignment` (Boolean) Whether the provisioner can reassign the task.
- `app_group_id` (String) The app group ID (entitlement ID).
- `app_id` (String) The app ID containing the group.
- `fallback_user_ids` (List of String) Fallback user IDs if no group members are found.


<a id="nestedatt--policy_steps--steps--provision--provision_policy--manual--assignee--manager"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.manual.assignee.manager`

Optional:

- `allow_reassignment` (Boolean) Whether the provisioner can reassign the task.
- `fallback_user_ids` (List of String) Fallback user IDs if no manager is found.


<a id="nestedatt--policy_steps--steps--provision--provision_policy--manual--assignee--users"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.manual.assignee.users`

Optional:

- `allow_reassignment` (Boolean) Whether the provisioner can reassign the task.
- `user_ids` (List of String) The user IDs to assign as provisioners.




<a id="nestedatt--policy_steps--steps--provision--provision_policy--unconfigured"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.unconfigured`


<a id="nestedatt--policy_steps--steps--provision--provision_policy--webhook"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.webhook`

Optional:

- `webhook_id` (String) The ID of the webhook to call for provisioning.



<a id="nestedatt--policy_steps--steps--provision--provision_target"></a>
### Nested Schema for `policy_steps.steps.provision.provision_target`

Optional:

- `app_entitlement_id` (String) The app entitlement that should be provisioned.
- `app_id` (String) The app in which the entitlement should be provisioned
- `app_user_id` (String) The app user that should be provisioned. May be unset if the app user is unknown
- `grant_duration` (String)



<a id="nestedatt--policy_steps--steps--reject"></a>
### Nested Schema for `policy_steps.steps.reject`

Optional:

- `reject_message` (String) An optional message to include in the comments when a task is automatically rejected.


<a id="nestedatt--policy_steps--steps--wait"></a>
### Nested Schema for `policy_steps.steps.wait`

Optional:

- `comment_on_first_wait` (String) The comment to post on first failed check.
- `comment_on_timeout` (String) The comment to post if we timeout.
- `condition` (Attributes) The WaitCondition message. (see [below for nested schema](#nestedatt--policy_steps--steps--wait--condition))
- `duration` (Attributes) The WaitDuration message. (see [below for nested schema](#nestedatt--policy_steps--steps--wait--duration))
- `name` (String) The name of our condition to show on the task details page
- `timeout_duration` (String)
- `until_time` (Attributes) Waits until a specific time of the day (UTC) (see [below for nested schema](#nestedatt--policy_steps--steps--wait--until_time))

<a id="nestedatt--policy_steps--steps--wait--condition"></a>
### Nested Schema for `policy_steps.steps.wait.condition`

Optional:

- `condition` (String) The condition that has to be true for this wait condition to continue.


<a id="nestedatt--policy_steps--steps--wait--duration"></a>
### Nested Schema for `policy_steps.steps.wait.duration`

Optional:

- `duration` (String)


<a id="nestedatt--policy_steps--steps--wait--until_time"></a>
### Nested Schema for `policy_steps.steps.wait.until_time`

Optional:

- `hours` (Number) The hours field.
- `minutes` (Number) The minutes field.
- `timezone` (String) The timezone field.





<a id="nestedatt--post_actions"></a>
### Nested Schema for `post_actions`

Optional:

- `certify_remediate_immediately` (Boolean) Only valid on certify policies. When true, any revocations resulting from
 the certification are applied immediately when the campaign task closes.
This field is part of the `action` oneof.
See the documentation for `c1.api.policy.v1.PolicyPostActions` for more details.


<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Optional:

- `condition` (String) A CEL expression that is evaluated against the request context. If it
 returns true, the step sequence identified by the outcome is used.
- `policy_id` (String) The ID of another Policy that is evaluated recursively when this
 rule matches. The referenced policy must share this policy's
 policy_type, must not introduce a cycle, and must not push any
 reachable chain over depth 5. Gated by the
 POLICY_REFERENCES_POLICY feature flag.
This field is part of the `outcome` oneof.
See the documentation for `c1.api.policy.v1.Rule` for more details.
- `policy_key` (String, Deprecated) Deprecated: prefer outcome.step_key. Still read by the request path
 for backward compatibility with rules persisted before the outcome
 oneof existed.
- `step_key` (String) A key into the policy's policy_steps map identifying which step
 sequence to execute when this rule's condition matches.
This field is part of the `outcome` oneof.
See the documentation for `c1.api.policy.v1.Rule` for more details.


<a id="nestedatt--scope"></a>
### Nested Schema for `scope`

Optional:

- `app_entitlement_id` (String) Optional. When set, the policy is scoped to this entitlement of app_id
 rather than to the whole app.
- `app_id` (String) The ID of the app this policy is scoped to.
- `slot` (String) Which of the object's local-policy slots this policy occupies. Part of the
 scope, and immutable with it.
possible known values include one of ["POLICY_SCOPE_SLOT_UNSPECIFIED", "POLICY_SCOPE_SLOT_EMERGENCY"]
