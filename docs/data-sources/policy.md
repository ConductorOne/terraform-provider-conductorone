---
page_title: "conductorone_policy Data Source - terraform-provider-conductorone"
subcategory: ""
description: |-
  Policy DataSource
---

# conductorone_policy (Data Source)

Policy DataSource

The Policy datasource allows you to retrieve a Policy instance by `display_name` (case insensitive), or `id` in ConductorOne.

## Example Usage

```terraform
data "conductorone_policy" "default_request_policy" {
  display_name = "App Owner Request Policy"
}

data "conductorone_policy" "default_revoke_policy" {
  display_name = "Default Revoke Policy"
}

data "conductorone_policy" "default_review_policy" {
  display_name = "App Owner Review Policy"
}

data "conductorone_policy" "custom_review_policy" {
  id = "<policy_id>"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `display_name` (String) The displayName field.
- `id` (String) The id field.

### Read-Only

- `created_at` (String)
- `deleted_at` (String)
- `description` (String) The description field.
- `policy_steps` (Attributes Map) The policySteps field. (see [below for nested schema](#nestedatt--policy_steps))
- `policy_type` (String) must be one of [POLICY_TYPE_UNSPECIFIED, POLICY_TYPE_GRANT, POLICY_TYPE_REVOKE, POLICY_TYPE_CERTIFY, POLICY_TYPE_ACCESS_REQUEST, POLICY_TYPE_PROVISION]
The policyType field.
- `post_actions` (Attributes List) The postActions field. (see [below for nested schema](#nestedatt--post_actions))
- `reassign_tasks_to_delegates` (Boolean) A policy configuration option that allows for reassinging tasks to delgated users. This level of delegation refers to the individual delegates users set on their account.
- `rules` (Attributes List) The rules field. (see [below for nested schema](#nestedatt--rules))
- `system_builtin` (Boolean) The systemBuiltin field.
- `updated_at` (String)

<a id="nestedatt--policy_steps"></a>
### Nested Schema for `policy_steps`

Read-Only:

- `steps` (Attributes List) The steps field. (see [below for nested schema](#nestedatt--policy_steps--steps))

<a id="nestedatt--policy_steps--steps"></a>
### Nested Schema for `policy_steps.steps`

Read-Only:

- `accept` (Attributes) This policy step indicates that a ticket should have an approved outcome. This is a terminal approval state and is used to explicitly define the end of approval steps. (see [below for nested schema](#nestedatt--policy_steps--steps--accept))
- `approval` (Attributes) The Approval field is used to define who should perform the review.
This message contains a oneof. Only a single field of the following list may be set at a time:
  - users
  - manager
  - appOwners
  - group
  - self
  - entitlementOwners (see [below for nested schema](#nestedatt--policy_steps--steps--approval))
- `provision` (Attributes) The Provision message. (see [below for nested schema](#nestedatt--policy_steps--steps--provision))
- `reject` (Attributes) This policy step indicates that a ticket should have a denied outcome. This is a terminal approval state and is used to explicitly define the end of approval steps. (see [below for nested schema](#nestedatt--policy_steps--steps--reject))

<a id="nestedatt--policy_steps--steps--accept"></a>
### Nested Schema for `policy_steps.steps.accept`


<a id="nestedatt--policy_steps--steps--approval"></a>
### Nested Schema for `policy_steps.steps.approval`

Read-Only:

- `allow_reassignment` (Boolean) The allowReassignment field.
- `app_group_approval` (Attributes) The AppGroupApproval message. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--app_group_approval))
- `app_owner_approval` (Attributes) The AppOwnerApproval message. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--app_owner_approval))
- `assigned` (Boolean) The assigned field.
- `entitlement_owner_approval` (Attributes) The EntitlementOwnerApproval message. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--entitlement_owner_approval))
- `expression_approval` (Attributes) The ExpressionApproval message. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--expression_approval))
- `manager_approval` (Attributes) The ManagerApproval message. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--manager_approval))
- `require_approval_reason` (Boolean) The requireApprovalReason field.
- `require_reassignment_reason` (Boolean) The requireReassignmentReason field.
- `self_approval` (Attributes) The SelfApproval message. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--self_approval))
- `user_approval` (Attributes) The UserApproval message. (see [below for nested schema](#nestedatt--policy_steps--steps--approval--user_approval))

<a id="nestedatt--policy_steps--steps--approval--app_group_approval"></a>
### Nested Schema for `policy_steps.steps.approval.user_approval`

Read-Only:

- `allow_self_approval` (Boolean) The allowSelfApproval field.
- `app_group_id` (String) The appGroupId field.
- `app_id` (String) The appId field.
- `fallback` (Boolean) The fallback field.
- `fallback_user_ids` (List of String) The fallbackUserIds field.


<a id="nestedatt--policy_steps--steps--approval--app_owner_approval"></a>
### Nested Schema for `policy_steps.steps.approval.user_approval`

Read-Only:

- `allow_self_approval` (Boolean) App owner is based on the app id and doesn't need to have self-contained data


<a id="nestedatt--policy_steps--steps--approval--entitlement_owner_approval"></a>
### Nested Schema for `policy_steps.steps.approval.user_approval`

Read-Only:

- `allow_self_approval` (Boolean) Entitlement owner is based on the current entitlement's id and doesn't need to have self-contained data
- `fallback` (Boolean) The fallback field.
- `fallback_user_ids` (List of String) The fallbackUserIds field.


<a id="nestedatt--policy_steps--steps--approval--expression_approval"></a>
### Nested Schema for `policy_steps.steps.approval.user_approval`

Read-Only:

- `allow_self_approval` (Boolean) Configuration to allow self approval of if the user is specified and also the target of the ticket.
- `assigned_user_ids` (List of String) The assignedUserIds field.
- `expressions` (List of String) Array of dynamic expressions to determine the approvers.  The first expression to return a non-empty list of users will be used.
- `fallback` (Boolean) Configuration to allow a fallback if the expression does not return a valid list of users.
- `fallback_user_ids` (List of String) Configuration to specific which users to fallback to if and the expression does not return a valid list of users.


<a id="nestedatt--policy_steps--steps--approval--manager_approval"></a>
### Nested Schema for `policy_steps.steps.approval.user_approval`

Read-Only:

- `allow_self_approval` (Boolean) The allowSelfApproval field.
- `assigned_user_ids` (List of String) The assignedUserIds field.
- `fallback` (Boolean) The fallback field.
- `fallback_user_ids` (List of String) The fallbackUserIds field.


<a id="nestedatt--policy_steps--steps--approval--self_approval"></a>
### Nested Schema for `policy_steps.steps.approval.user_approval`

Read-Only:

- `assigned_user_ids` (List of String) The assignedUserIds field.
- `fallback` (Boolean) The fallback field.
- `fallback_user_ids` (List of String) Self approval is the target of the ticket


<a id="nestedatt--policy_steps--steps--approval--user_approval"></a>
### Nested Schema for `policy_steps.steps.approval.user_approval`

Read-Only:

- `allow_self_approval` (Boolean) The allowSelfApproval field.
- `user_ids` (List of String) The userIds field.



<a id="nestedatt--policy_steps--steps--provision"></a>
### Nested Schema for `policy_steps.steps.provision`


<a id="nestedatt--policy_steps--steps--reject"></a>
### Nested Schema for `policy_steps.steps.reject`




<a id="nestedatt--post_actions"></a>
### Nested Schema for `post_actions`

Read-Only:

- `certify_remediate_immediately` (Boolean) ONLY valid when used in a CERTIFY Ticket Type:
 Causes any deprovision or change in a grant to be applied when Certify Ticket is closed.
See the documentation for `c1.api.policy.v1.PolicyPostActions` for more details.


<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Read-Only:

- `condition` (String) The condition field.
- `policy_key` (String) This is a reference to a list of policy steps from `policy_steps`