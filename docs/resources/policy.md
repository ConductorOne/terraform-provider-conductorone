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
  description  = "...my_description..."
  display_name = "...my_display_name..."
  policy_steps = {
    key = {
      steps = [
        {
          accept = {
            accept_message = "...my_accept_message..."
          }
          approval = {
            agent_approval = {
              agent_failure_action = "APPROVAL_AGENT_FAILURE_ACTION_REASSIGN_TO_USERS"
              agent_mode           = "APPROVAL_AGENT_MODE_FULL_CONTROL"
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
            app_group_approval = {
              allow_self_approval = true
              app_group_id        = "...my_app_group_id..."
              app_id              = "...my_app_id..."
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
              is_group_fallback_enabled = false
            }
            app_owner_approval = {
              allow_self_approval = true
            }
            entitlement_owner_approval = {
              allow_self_approval = false
              fallback            = true
              fallback_user_ids = [
                "..."
              ]
            }
            escalation = {
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
            }
            escalation_enabled = false
            expression_approval = {
              allow_self_approval = true
              expressions = [
                "..."
              ]
              fallback = false
              fallback_user_ids = [
                "..."
              ]
            }
            manager_approval = {
              allow_self_approval = true
              fallback            = true
              fallback_user_ids = [
                "..."
              ]
            }
            require_approval_reason      = true
            require_denial_reason        = true
            require_reassignment_reason  = true
            requires_step_up_provider_id = "...my_requires_step_up_provider_id..."
            resource_owner_approval = {
              allow_self_approval = true
              fallback            = false
              fallback_user_ids = [
                "..."
              ]
            }
            self_approval = {
              fallback = false
              fallback_user_ids = [
                "..."
              ]
            }
            user_approval = {
              allow_self_approval = false
              user_ids = [
                "..."
              ]
            }
            webhook_approval = {
              webhook_id = "...my_webhook_id..."
            }
          }
          form = "{ \"see\": \"documentation\" }"
          provision = {
            assigned = true
            provision_policy = {
              connector_provision = {
                account_provision = {
                  config = {
                    # ...
                  }
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
              delegated_provision = {
                app_id         = "...my_app_id..."
                entitlement_id = "...my_entitlement_id..."
              }
              external_ticket_provision = {
                app_id                                = "...my_app_id..."
                connector_id                          = "...my_connector_id..."
                external_ticket_provisioner_config_id = "...my_external_ticket_provisioner_config_id..."
                instructions                          = "...my_instructions..."
              }
              manual_provision = {
                instructions = "...my_instructions..."
                user_ids = [
                  "..."
                ]
              }
              multi_step = "{ \"see\": \"documentation\" }"
              unconfigured_provision = {
                # ...
              }
              webhook_provision = {
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
            name                  = "...my_name..."
            timeout_duration      = "...my_timeout_duration..."
            wait_condition = {
              condition = "...my_condition..."
            }
            wait_duration = {
              duration = "...my_duration..."
            }
            wait_until_time = {
              hours    = 5
              minutes  = 8
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
      policy_key = "...my_policy_key..."
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `display_name` (String) The display name of the new policy.

### Optional

- `description` (String) The description of the new policy.
- `policy_steps` (Attributes Map) The map of policy type to policy steps. The key is the stringified version of the enum. See other policies for examples. (see [below for nested schema](#nestedatt--policy_steps))
- `policy_type` (String) The enum of the policy type. must be one of ["POLICY_TYPE_UNSPECIFIED", "POLICY_TYPE_GRANT", "POLICY_TYPE_REVOKE", "POLICY_TYPE_CERTIFY", "POLICY_TYPE_ACCESS_REQUEST", "POLICY_TYPE_PROVISION"]
- `post_actions` (Attributes List) Actions to occur after a policy finishes. As of now this is only valid on a certify policy to remediate a denied certification immediately. (see [below for nested schema](#nestedatt--post_actions))
- `reassign_tasks_to_delegates` (Boolean, Deprecated) Deprecated. Use setting in policy step instead
- `rules` (Attributes List) The rules field. (see [below for nested schema](#nestedatt--rules))

### Read-Only

- `created_at` (String)
- `id` (String) The ID of the Policy.
- `system_builtin` (Boolean) Whether this policy is a builtin system policy. Builtin system policies cannot be edited.
- `updated_at` (String)

<a id="nestedatt--policy_steps"></a>
### Nested Schema for `policy_steps`

Optional:

- `steps` (Attributes List) An array of policy steps indicating the processing flow of a policy. These steps are oneOfs, and only one property may be set for each array index at a time. (see [below for nested schema](#nestedatt--policy_steps--steps))

<a id="nestedatt--policy_steps--steps"></a>
### Nested Schema for `policy_steps.steps`

Optional:

- `accept` (Attributes) This policy step indicates that a ticket should have an approved outcome. This is a terminal approval state and is used to explicitly define the end of approval steps. (see [below for nested schema](#nestedatt--policy_steps--steps--accept))
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
- `form` (String) The Form message. Parsed as JSON.
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


<a id="nestedatt--policy_steps--steps--approval"></a>
### Nested Schema for `policy_steps.steps.approval`

Optional:

- `agent_approval` (Attributes) The agent to assign the task to. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--agent_approval))
- `allow_delegation` (Boolean) Whether ticket delegation is allowed for this step.
- `allow_reassignment` (Boolean) Configuration to allow reassignment by reviewers during this step.
- `allowed_reassignees` (List of String) List of users for whom this step can be reassigned.
- `app_group_approval` (Attributes) The AppGroupApproval object provides the configuration for setting a group as the approvers of an approval policy step. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--app_group_approval))
- `app_owner_approval` (Attributes) App owner approval provides the configuration for an approval step when the app owner is the target. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--app_owner_approval))
- `entitlement_owner_approval` (Attributes) The entitlement owner approval allows configuration of the approval step when the target approvers are the entitlement owners. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--entitlement_owner_approval))
- `escalation` (Attributes) The Escalation message.

This message contains a oneof named escalation_policy. Only a single field of the following list may be set at a time:
  - replacePolicy
  - reassignToApprovers (see [below for nested schema](#nestedatt--policy_steps--steps--approval--escalation))
- `escalation_enabled` (Boolean) Whether escalation is enabled for this step.
- `expression_approval` (Attributes) The ExpressionApproval message. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--expression_approval))
- `manager_approval` (Attributes) The manager approval object provides configuration options for approval when the target of the approval is the manager of the user in the task. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--manager_approval))
- `require_approval_reason` (Boolean) Configuration to require a reason when approving this step.
- `require_denial_reason` (Boolean) Configuration to require a reason when denying this step.
- `require_reassignment_reason` (Boolean) Configuration to require a reason when reassigning this step.
- `requires_step_up_provider_id` (String) The ID of a step-up authentication provider that will be required for approvals on this step.
 If set, approvers must complete the step-up authentication flow before they can approve.
- `resource_owner_approval` (Attributes) The resource owner approval allows configuration of the approval step when the target approvers are the resource owners. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--resource_owner_approval))
- `self_approval` (Attributes) The self approval object describes the configuration of a policy step that needs to be approved by the target of the request. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--self_approval))
- `user_approval` (Attributes) The user approval object describes the approval configuration of a policy step that needs to be approved by a specific list of users. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--user_approval))
- `webhook_approval` (Attributes) The WebhookApproval message. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--webhook_approval))

Read-Only:

- `assigned` (Boolean) A field indicating whether this step is assigned.

<a id="nestedatt--policy_steps--steps--approval--agent_approval"></a>
### Nested Schema for `policy_steps.steps.approval.agent_approval`

Optional:

- `agent_failure_action` (String) The action to take if the agent fails to approve, deny, or reassign the task. must be one of ["APPROVAL_AGENT_FAILURE_ACTION_UNSPECIFIED", "APPROVAL_AGENT_FAILURE_ACTION_REASSIGN_TO_USERS", "APPROVAL_AGENT_FAILURE_ACTION_REASSIGN_TO_SUPER_ADMINS", "APPROVAL_AGENT_FAILURE_ACTION_SKIP_POLICY_STEP"]
- `agent_mode` (String) The mode of the agent, full control, change policy only, or comment only. must be one of ["APPROVAL_AGENT_MODE_UNSPECIFIED", "APPROVAL_AGENT_MODE_FULL_CONTROL", "APPROVAL_AGENT_MODE_CHANGE_POLICY_ONLY", "APPROVAL_AGENT_MODE_COMMENT_ONLY"]
- `agent_user_id` (String) The agent user ID to assign the task to.
- `instructions` (String) Instructions for the agent.
- `policy_ids` (List of String) The allow list of policy IDs to re-route the task to.
- `reassign_to_user_ids` (List of String) The users to reassign the task to if the agent failure action is reassign to users.


<a id="nestedatt--policy_steps--steps--approval--app_group_approval"></a>
### Nested Schema for `policy_steps.steps.approval.app_group_approval`

Optional:

- `allow_self_approval` (Boolean) Configuration to allow self approval if the target user is a member of the group during this step.
- `app_group_id` (String) The ID of the group specified for approval.
- `app_id` (String) The ID of the app that contains the group specified for approval.
- `fallback` (Boolean) Configuration to allow a fallback if the group is empty.
- `fallback_group_ids` (Attributes List) Configuration to specify which groups to fallback to if fallback is enabled and the group is empty. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--app_group_approval--fallback_group_ids))
- `fallback_user_ids` (List of String) Configuration to specific which users to fallback to if fallback is enabled and the group is empty.
- `is_group_fallback_enabled` (Boolean) Configuration to enable fallback for group fallback.

<a id="nestedatt--policy_steps--steps--approval--app_group_approval--fallback_group_ids"></a>
### Nested Schema for `policy_steps.steps.approval.app_group_approval.fallback_group_ids`

Optional:

- `app_entitlement_id` (String) The ID of the Entitlement.
- `app_id` (String) The ID of the App this entitlement belongs to.



<a id="nestedatt--policy_steps--steps--approval--app_owner_approval"></a>
### Nested Schema for `policy_steps.steps.approval.app_owner_approval`

Optional:

- `allow_self_approval` (Boolean) Configuration that allows a user to self approve if they are an app owner during this approval step.


<a id="nestedatt--policy_steps--steps--approval--entitlement_owner_approval"></a>
### Nested Schema for `policy_steps.steps.approval.entitlement_owner_approval`

Optional:

- `allow_self_approval` (Boolean) Configuration to allow self approval if the target user is an entitlement owner during this step.
- `fallback` (Boolean) Configuration to allow a fallback if the entitlement owner cannot be identified.
- `fallback_user_ids` (List of String) Configuration to specific which users to fallback to if fallback is enabled and the entitlement owner cannot be identified.


<a id="nestedatt--policy_steps--steps--approval--escalation"></a>
### Nested Schema for `policy_steps.steps.approval.escalation`

Optional:

- `escalation_comment` (String) The escalationComment field.
- `expiration` (String) The expiration field.
- `reassign_to_approvers` (Attributes) The ReassignToApprovers message. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--escalation--reassign_to_approvers))
- `replace_policy` (Attributes) The ReplacePolicy message. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--escalation--replace_policy))

<a id="nestedatt--policy_steps--steps--approval--escalation--reassign_to_approvers"></a>
### Nested Schema for `policy_steps.steps.approval.escalation.reassign_to_approvers`

Optional:

- `approver_ids` (List of String) The approverIds field.


<a id="nestedatt--policy_steps--steps--approval--escalation--replace_policy"></a>
### Nested Schema for `policy_steps.steps.approval.escalation.replace_policy`

Optional:

- `policy_id` (String) The policyId field.



<a id="nestedatt--policy_steps--steps--approval--expression_approval"></a>
### Nested Schema for `policy_steps.steps.approval.expression_approval`

Optional:

- `allow_self_approval` (Boolean) Configuration to allow self approval of if the user is specified and also the target of the ticket.
- `expressions` (List of String) Array of dynamic expressions to determine the approvers.  The first expression to return a non-empty list of users will be used.
- `fallback` (Boolean) Configuration to allow a fallback if the expression does not return a valid list of users.
- `fallback_user_ids` (List of String) Configuration to specific which users to fallback to if and the expression does not return a valid list of users.

Read-Only:

- `assigned_user_ids` (List of String) The assignedUserIds field.


<a id="nestedatt--policy_steps--steps--approval--manager_approval"></a>
### Nested Schema for `policy_steps.steps.approval.manager_approval`

Optional:

- `allow_self_approval` (Boolean) Configuration to allow self approval if the target user is their own manager. This may occur if a service account has an identity user and manager specified as the same person.
- `fallback` (Boolean) Configuration to allow a fallback if no manager is found.
- `fallback_user_ids` (List of String) Configuration to specific which users to fallback to if fallback is enabled and no manager is found.

Read-Only:

- `assigned_user_ids` (List of String) The array of users determined to be the manager during processing time.


<a id="nestedatt--policy_steps--steps--approval--resource_owner_approval"></a>
### Nested Schema for `policy_steps.steps.approval.resource_owner_approval`

Optional:

- `allow_self_approval` (Boolean) Configuration to allow self approval if the target user is an resource owner during this step.
- `fallback` (Boolean) Configuration to allow a fallback if the resource owner cannot be identified.
- `fallback_user_ids` (List of String) Configuration to specific which users to fallback to if fallback is enabled and the resource owner cannot be identified.


<a id="nestedatt--policy_steps--steps--approval--self_approval"></a>
### Nested Schema for `policy_steps.steps.approval.self_approval`

Optional:

- `fallback` (Boolean) Configuration to allow a fallback if the identity user of the target app user cannot be determined.
- `fallback_user_ids` (List of String) Configuration to specific which users to fallback to if fallback is enabled and the identity user of the target app user cannot be determined.

Read-Only:

- `assigned_user_ids` (List of String) The array of users determined to be themselves during approval. This should only ever be one person, but is saved because it may change if the owner of an app user changes while the ticket is open.


<a id="nestedatt--policy_steps--steps--approval--user_approval"></a>
### Nested Schema for `policy_steps.steps.approval.user_approval`

Optional:

- `allow_self_approval` (Boolean) Configuration to allow self approval of if the user is specified and also the target of the ticket.
- `user_ids` (List of String) Array of users configured for approval.


<a id="nestedatt--policy_steps--steps--approval--webhook_approval"></a>
### Nested Schema for `policy_steps.steps.approval.webhook_approval`

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
  - unconfigured (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy))
- `provision_target` (Attributes) ProvisionTarget indicates the specific app, app entitlement, and if known, the app user and grant duration of this provision step (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_target))

<a id="nestedatt--policy_steps--steps--provision--provision_policy"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy`

Optional:

- `connector_provision` (Attributes) Indicates that a connector should perform the provisioning. This object has no fields.

This message contains a oneof named provision_type. Only a single field of the following list may be set at a time:
  - defaultBehavior
  - account
  - deleteAccount (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--connector_provision))
- `delegated_provision` (Attributes) This provision step indicates that we should delegate provisioning to the configuration of another app entitlement. This app entitlement does not have to be one from the same app, but MUST be configured as a proxy binding leading into this entitlement. (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--delegated_provision))
- `external_ticket_provision` (Attributes) This provision step indicates that we should check an external ticket to provision this entitlement (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--external_ticket_provision))
- `manual_provision` (Attributes) Manual provisioning indicates that a human must intervene for the provisioning of this step. (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--manual_provision))
- `multi_step` (String) MultiStep indicates that this provision step has multiple steps to process. Parsed as JSON.
- `unconfigured_provision` (Attributes) The UnconfiguredProvision message. (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--unconfigured_provision))
- `webhook_provision` (Attributes) This provision step indicates that a webhook should be called to provision this entitlement. (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--webhook_provision))

<a id="nestedatt--policy_steps--steps--provision--provision_policy--connector_provision"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.connector_provision`

Optional:

- `account_provision` (Attributes) The AccountProvision message.

This message contains a oneof named storage_type. Only a single field of the following list may be set at a time:
  - saveToVault
  - doNotSave (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--connector_provision--account_provision))
- `default_behavior` (Attributes) The DefaultBehavior message. (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--connector_provision--default_behavior))
- `delete_account` (Attributes) The DeleteAccount message. (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--connector_provision--delete_account))

<a id="nestedatt--policy_steps--steps--provision--provision_policy--connector_provision--account_provision"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.connector_provision.account_provision`

Optional:

- `config` (Attributes) (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--connector_provision--account_provision--config))
- `connector_id` (String) The connectorId field.
- `do_not_save` (Attributes) The DoNotSave message. (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--connector_provision--account_provision--do_not_save))
- `save_to_vault` (Attributes) The SaveToVault message. (see [below for nested schema](#nestedatt--policy_steps--steps--provision--provision_policy--connector_provision--account_provision--save_to_vault))
- `schema_id` (String) The schemaId field.

<a id="nestedatt--policy_steps--steps--provision--provision_policy--connector_provision--account_provision--config"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.connector_provision.account_provision.config`


<a id="nestedatt--policy_steps--steps--provision--provision_policy--connector_provision--account_provision--do_not_save"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.connector_provision.account_provision.do_not_save`


<a id="nestedatt--policy_steps--steps--provision--provision_policy--connector_provision--account_provision--save_to_vault"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.connector_provision.account_provision.save_to_vault`

Optional:

- `vault_ids` (List of String) The vaultIds field.



<a id="nestedatt--policy_steps--steps--provision--provision_policy--connector_provision--default_behavior"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.connector_provision.default_behavior`

Optional:

- `connector_id` (String) this checks if the entitlement is enabled by provisioning in a specific connector
 this can happen automatically and doesn't need any extra info


<a id="nestedatt--policy_steps--steps--provision--provision_policy--connector_provision--delete_account"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.connector_provision.delete_account`

Optional:

- `connector_id` (String) The connectorId field.



<a id="nestedatt--policy_steps--steps--provision--provision_policy--delegated_provision"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.delegated_provision`

Optional:

- `app_id` (String) The AppID of the entitlement to delegate provisioning to.
- `entitlement_id` (String) The ID of the entitlement we are delegating provisioning to.


<a id="nestedatt--policy_steps--steps--provision--provision_policy--external_ticket_provision"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.external_ticket_provision`

Optional:

- `app_id` (String) The appId field.
- `connector_id` (String) The connectorId field.
- `external_ticket_provisioner_config_id` (String) The externalTicketProvisionerConfigId field.
- `instructions` (String) This field indicates a text body of instructions for the provisioner to indicate.


<a id="nestedatt--policy_steps--steps--provision--provision_policy--manual_provision"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.manual_provision`

Optional:

- `instructions` (String) This field indicates a text body of instructions for the provisioner to indicate.
- `user_ids` (List of String) An array of users that are required to provision during this step.


<a id="nestedatt--policy_steps--steps--provision--provision_policy--unconfigured_provision"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.unconfigured_provision`


<a id="nestedatt--policy_steps--steps--provision--provision_policy--webhook_provision"></a>
### Nested Schema for `policy_steps.steps.provision.provision_policy.webhook_provision`

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
- `name` (String) The name of our condition to show on the task details page
- `timeout_duration` (String)
- `wait_condition` (Attributes) The WaitCondition message. (see [below for nested schema](#nestedatt--policy_steps--steps--wait--wait_condition))
- `wait_duration` (Attributes) The WaitDuration message. (see [below for nested schema](#nestedatt--policy_steps--steps--wait--wait_duration))
- `wait_until_time` (Attributes) Waits until a specific time of the day (UTC) (see [below for nested schema](#nestedatt--policy_steps--steps--wait--wait_until_time))

<a id="nestedatt--policy_steps--steps--wait--wait_condition"></a>
### Nested Schema for `policy_steps.steps.wait.wait_condition`

Optional:

- `condition` (String) The condition that has to be true for this wait condition to continue.


<a id="nestedatt--policy_steps--steps--wait--wait_duration"></a>
### Nested Schema for `policy_steps.steps.wait.wait_duration`

Optional:

- `duration` (String)


<a id="nestedatt--policy_steps--steps--wait--wait_until_time"></a>
### Nested Schema for `policy_steps.steps.wait.wait_until_time`

Optional:

- `hours` (Number) The hours field.
- `minutes` (Number) The minutes field.
- `timezone` (String) The timezone field.





<a id="nestedatt--post_actions"></a>
### Nested Schema for `post_actions`

Optional:

- `certify_remediate_immediately` (Boolean) ONLY valid when used in a CERTIFY Ticket Type:
 Causes any deprovision or change in a grant to be applied when Certify Ticket is closed.
This field is part of the `action` oneof.
See the documentation for `c1.api.policy.v1.PolicyPostActions` for more details.


<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Optional:

- `condition` (String) The condition field.
- `policy_key` (String) This is a reference to a list of policy steps from `policy_steps`
