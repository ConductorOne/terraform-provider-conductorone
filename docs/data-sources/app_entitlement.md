---
page_title: "conductorone_app_entitlement Data Source - terraform-provider-conductorone"
subcategory: ""
description: |-
  AppEntitlement DataSource
---

# conductorone_app_entitlement (Data Source)

AppEntitlement DataSource

This data source enables you to retrieve ConductorOne app entitlements using the following search criteria:

* `access_review_id` - Filter by access review ID
* `alias` - Filter by alias (case sensitive)
* `app_ids` - List of app IDs to filter by
* `app_user_ids` - List of app user IDs to filter by
* `compliance_framework_ids` - List of compliance framework IDs to filter by
* `exclude_app_ids` - List of app IDs to exclude
* `exclude_app_user_ids` - List of app user IDs to exclude
* `exclude_resource_type_ids` - List of resource type IDs to exclude
* `include_deleted` - Include deleted entitlements
* `is_automated` - Filter by automation status
* `membership_type` - Filter by membership type
* `only_get_expiring` - Only return expiring entitlements
* `query` - Search query string
* `resource_ids` - List of resource IDs to filter by
* `resource_trait_ids` - List of resource trait IDs to filter by
* `resource_type_ids` - List of resource type IDs to filter by
* `risk_level_ids` - List of risk level IDs to filter by
* `source_connector_id` - Filter by source connector ID

## Example Usage

```terraform
data "conductorone_app_entitlement" "my_app_entitlement" {
  access_review_id = "...my_access_review_id..."
  alias            = "...my_alias..."
  app_ids = [
    "..."
  ]
  app_user_ids = [
    "..."
  ]
  compliance_framework_ids = [
    "..."
  ]
  display_name = "...my_display_name..."
  exclude_app_ids = [
    "..."
  ]
  exclude_app_user_ids = [
    "..."
  ]
  exclude_resource_type_ids = [
    "..."
  ]
  include_deleted = true
  is_automated    = false
  membership_type = [
    "APP_ENTITLEMENT_MEMBERSHIP_TYPE_OWNER"
  ]
  only_get_expiring = true
  page_size         = 5
  page_token        = "...my_page_token..."
  query             = "...my_query..."
  refs = [
    {
      app_id = "...my_app_id..."
      id     = "...my_id..."
    }
  ]
  resource_ids = [
    "..."
  ]
  resource_trait_ids = [
    "..."
  ]
  resource_type_ids = [
    "..."
  ]
  risk_level_ids = [
    "..."
  ]
  source_connector_id = "...my_source_connector_id..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `access_review_id` (String) Search for app entitlements that are being reviewed as part of this access review campaign.
- `alias` (String) Search for app entitlements that have this alias (exact match).
- `app_ids` (List of String) Search for app entitlements contained in any of these apps.
- `app_user_ids` (List of String) Search for app entitlements that are granted to any of these app user ids.
- `compliance_framework_ids` (List of String) Search for app entitlements that are part of these compliace frameworks.
- `display_name` (String) The displayName field.
- `exclude_app_ids` (List of String) Exclude app entitlements from the results that are in these app IDs.
- `exclude_app_user_ids` (List of String) Exclude app entitlements from the results that these app users have granted.
- `exclude_resource_type_ids` (List of String) The excludeResourceTypeIds field.
- `include_deleted` (Boolean) Include deleted app entitlements, this includes app entitlements that have a deleted parent object (app, app resource, app resource type)
- `is_automated` (Boolean) The isAutomated field.
- `membership_type` (List of String) The membershipType field.
- `only_get_expiring` (Boolean) Restrict results to only those who have expiring app entitlement user bindings.
- `page_size` (Number) The pageSize where 0 <= pageSize <= 100. Values < 10 will be set to 10. A value of 0 returns the default page size (currently 25)
- `page_token` (String) The pageToken field.
- `query` (String) Query the app entitlements with a fuzzy search on display name and description.
- `refs` (Attributes List) The refs field. (see [below for nested schema](#nestedatt--refs))
- `resource_ids` (List of String) Search for app entitlements that belongs to these resources.
- `resource_trait_ids` (List of String) The resourceTraitIds field.
- `resource_type_ids` (List of String) Search for app entitlements that are for items with resources types that have matching names. Example names are "group", "role", and "app".
- `risk_level_ids` (List of String) Search for app entitlements with these risk levels.
- `source_connector_id` (String) The sourceConnectorId field.

### Read-Only

- `app_id` (String) The ID of the app that is associated with the app entitlement.
- `app_resource_id` (String) The ID of the app resource that is associated with the app entitlement
- `app_resource_type_id` (String) The ID of the app resource type that is associated with the app entitlement
- `certify_policy_id` (String) The ID of the policy that will be used for certify tickets related to the app entitlement.
- `compliance_framework_value_ids` (List of String) The IDs of different compliance frameworks associated with this app entitlement ex (SOX, HIPAA, PCI, etc.)
- `created_at` (String)
- `default_values_applied` (Boolean) Flag to indicate if app-level access request defaults have been applied to the entitlement
- `deleted_at` (String)
- `deprovisioner_policy` (Attributes) ProvisionPolicy is a oneOf that indicates how a provision step should be processed.

This message contains a oneof named typ. Only a single field of the following list may be set at a time:
  - connector
  - manual
  - delegated
  - webhook
  - multiStep
  - externalTicket
  - unconfigured (see [below for nested schema](#nestedatt--deprovisioner_policy))
- `description` (String) The description of the app entitlement.
- `duration_grant` (String)
- `duration_unset` (Attributes) (see [below for nested schema](#nestedatt--duration_unset))
- `emergency_grant_enabled` (Boolean) This enables tasks to be created in an emergency and use a selected emergency access policy.
- `emergency_grant_policy_id` (String) The ID of the policy that will be used for emergency access grant tasks.
- `grant_count` (String) The amount of grants open for this entitlement
- `grant_policy_id` (String) The ID of the policy that will be used for grant tickets related to the app entitlement.
- `id` (String) The unique ID for the App Entitlement.
- `is_automation_enabled` (Boolean) Flag to indicate whether automation (for adding users to entitlement based on rules) has been enabled.
- `is_manually_managed` (Boolean) Flag to indicate if the app entitlement is manually managed.
- `match_baton_id` (String) The matchBatonId field.
- `next_page_token` (String) The nextPageToken is shown for the next page if the number of results is larger than the max page size. The server returns one page of results and the nextPageToken until all results are retreived. To retrieve the next page, use the same request and append a pageToken field with the value of nextPageToken shown on the previous page.
- `override_access_requests_defaults` (Boolean) Flag to indicate if the app-level access request settings have been overridden for the entitlement
- `provision_policy` (Attributes) ProvisionPolicy is a oneOf that indicates how a provision step should be processed.

This message contains a oneof named typ. Only a single field of the following list may be set at a time:
  - connector
  - manual
  - delegated
  - webhook
  - multiStep
  - externalTicket
  - unconfigured (see [below for nested schema](#nestedatt--provision_policy))
- `purpose` (String) The purpose field.
- `revoke_policy_id` (String) The ID of the policy that will be used for revoke tickets related to the app entitlement
- `risk_level_value_id` (String) The riskLevelValueId field.
- `slug` (String) The slug is displayed as an oval next to the name in the frontend of C1, it tells you what permission the entitlement grants. See https://www.conductorone.com/docs/product/admin/entitlements/
- `source_connector_ids` (Map of String) Map to tell us which connector the entitlement came from.
- `system_builtin` (Boolean) This field indicates if this is a system builtin entitlement.
- `updated_at` (String)

<a id="nestedatt--refs"></a>
### Nested Schema for `refs`

Optional:

- `app_id` (String) The appId field.
- `id` (String) The id field.


<a id="nestedatt--deprovisioner_policy"></a>
### Nested Schema for `deprovisioner_policy`

Read-Only:

- `connector_provision` (Attributes) Indicates that a connector should perform the provisioning. This object has no fields.

This message contains a oneof named provision_type. Only a single field of the following list may be set at a time:
  - defaultBehavior
  - account
  - deleteAccount (see [below for nested schema](#nestedatt--deprovisioner_policy--connector_provision))
- `delegated_provision` (Attributes) This provision step indicates that we should delegate provisioning to the configuration of another app entitlement. This app entitlement does not have to be one from the same app, but MUST be configured as a proxy binding leading into this entitlement. (see [below for nested schema](#nestedatt--deprovisioner_policy--delegated_provision))
- `external_ticket_provision` (Attributes) This provision step indicates that we should check an external ticket to provision this entitlement (see [below for nested schema](#nestedatt--deprovisioner_policy--external_ticket_provision))
- `manual_provision` (Attributes) Manual provisioning indicates that a human must intervene for the provisioning of this step. (see [below for nested schema](#nestedatt--deprovisioner_policy--manual_provision))
- `multi_step` (String) MultiStep indicates that this provision step has multiple steps to process. Parsed as JSON.
- `unconfigured_provision` (Attributes) The UnconfiguredProvision message. (see [below for nested schema](#nestedatt--deprovisioner_policy--unconfigured_provision))
- `webhook_provision` (Attributes) This provision step indicates that a webhook should be called to provision this entitlement. (see [below for nested schema](#nestedatt--deprovisioner_policy--webhook_provision))

<a id="nestedatt--deprovisioner_policy--connector_provision"></a>
### Nested Schema for `deprovisioner_policy.connector_provision`

Read-Only:

- `account_provision` (Attributes) The AccountProvision message.

This message contains a oneof named storage_type. Only a single field of the following list may be set at a time:
  - saveToVault
  - doNotSave (see [below for nested schema](#nestedatt--deprovisioner_policy--connector_provision--account_provision))
- `default_behavior` (Attributes) The DefaultBehavior message. (see [below for nested schema](#nestedatt--deprovisioner_policy--connector_provision--default_behavior))
- `delete_account` (Attributes) The DeleteAccount message. (see [below for nested schema](#nestedatt--deprovisioner_policy--connector_provision--delete_account))

<a id="nestedatt--deprovisioner_policy--connector_provision--account_provision"></a>
### Nested Schema for `deprovisioner_policy.connector_provision.account_provision`

Read-Only:

- `config` (Attributes) (see [below for nested schema](#nestedatt--deprovisioner_policy--connector_provision--account_provision--config))
- `connector_id` (String) The connectorId field.
- `do_not_save` (Attributes) The DoNotSave message. (see [below for nested schema](#nestedatt--deprovisioner_policy--connector_provision--account_provision--do_not_save))
- `save_to_vault` (Attributes) The SaveToVault message. (see [below for nested schema](#nestedatt--deprovisioner_policy--connector_provision--account_provision--save_to_vault))
- `schema_id` (String) The schemaId field.

<a id="nestedatt--deprovisioner_policy--connector_provision--account_provision--config"></a>
### Nested Schema for `deprovisioner_policy.connector_provision.account_provision.config`


<a id="nestedatt--deprovisioner_policy--connector_provision--account_provision--do_not_save"></a>
### Nested Schema for `deprovisioner_policy.connector_provision.account_provision.do_not_save`


<a id="nestedatt--deprovisioner_policy--connector_provision--account_provision--save_to_vault"></a>
### Nested Schema for `deprovisioner_policy.connector_provision.account_provision.save_to_vault`

Read-Only:

- `vault_ids` (List of String) The vaultIds field.



<a id="nestedatt--deprovisioner_policy--connector_provision--default_behavior"></a>
### Nested Schema for `deprovisioner_policy.connector_provision.default_behavior`

Read-Only:

- `connector_id` (String) this checks if the entitlement is enabled by provisioning in a specific connector
 this can happen automatically and doesn't need any extra info


<a id="nestedatt--deprovisioner_policy--connector_provision--delete_account"></a>
### Nested Schema for `deprovisioner_policy.connector_provision.delete_account`

Read-Only:

- `connector_id` (String) The connectorId field.



<a id="nestedatt--deprovisioner_policy--delegated_provision"></a>
### Nested Schema for `deprovisioner_policy.delegated_provision`

Read-Only:

- `app_id` (String) The AppID of the entitlement to delegate provisioning to.
- `entitlement_id` (String) The ID of the entitlement we are delegating provisioning to.


<a id="nestedatt--deprovisioner_policy--external_ticket_provision"></a>
### Nested Schema for `deprovisioner_policy.external_ticket_provision`

Read-Only:

- `app_id` (String) The appId field.
- `connector_id` (String) The connectorId field.
- `external_ticket_provisioner_config_id` (String) The externalTicketProvisionerConfigId field.
- `instructions` (String) This field indicates a text body of instructions for the provisioner to indicate.


<a id="nestedatt--deprovisioner_policy--manual_provision"></a>
### Nested Schema for `deprovisioner_policy.manual_provision`

Read-Only:

- `instructions` (String) This field indicates a text body of instructions for the provisioner to indicate.
- `user_ids` (List of String) An array of users that are required to provision during this step.


<a id="nestedatt--deprovisioner_policy--unconfigured_provision"></a>
### Nested Schema for `deprovisioner_policy.unconfigured_provision`


<a id="nestedatt--deprovisioner_policy--webhook_provision"></a>
### Nested Schema for `deprovisioner_policy.webhook_provision`

Read-Only:

- `webhook_id` (String) The ID of the webhook to call for provisioning.



<a id="nestedatt--duration_unset"></a>
### Nested Schema for `duration_unset`


<a id="nestedatt--provision_policy"></a>
### Nested Schema for `provision_policy`

Read-Only:

- `connector_provision` (Attributes) Indicates that a connector should perform the provisioning. This object has no fields.

This message contains a oneof named provision_type. Only a single field of the following list may be set at a time:
  - defaultBehavior
  - account
  - deleteAccount (see [below for nested schema](#nestedatt--provision_policy--connector_provision))
- `delegated_provision` (Attributes) This provision step indicates that we should delegate provisioning to the configuration of another app entitlement. This app entitlement does not have to be one from the same app, but MUST be configured as a proxy binding leading into this entitlement. (see [below for nested schema](#nestedatt--provision_policy--delegated_provision))
- `external_ticket_provision` (Attributes) This provision step indicates that we should check an external ticket to provision this entitlement (see [below for nested schema](#nestedatt--provision_policy--external_ticket_provision))
- `manual_provision` (Attributes) Manual provisioning indicates that a human must intervene for the provisioning of this step. (see [below for nested schema](#nestedatt--provision_policy--manual_provision))
- `multi_step` (String) MultiStep indicates that this provision step has multiple steps to process. Parsed as JSON.
- `unconfigured_provision` (Attributes) The UnconfiguredProvision message. (see [below for nested schema](#nestedatt--provision_policy--unconfigured_provision))
- `webhook_provision` (Attributes) This provision step indicates that a webhook should be called to provision this entitlement. (see [below for nested schema](#nestedatt--provision_policy--webhook_provision))

<a id="nestedatt--provision_policy--connector_provision"></a>
### Nested Schema for `provision_policy.connector_provision`

Read-Only:

- `account_provision` (Attributes) The AccountProvision message.

This message contains a oneof named storage_type. Only a single field of the following list may be set at a time:
  - saveToVault
  - doNotSave (see [below for nested schema](#nestedatt--provision_policy--connector_provision--account_provision))
- `default_behavior` (Attributes) The DefaultBehavior message. (see [below for nested schema](#nestedatt--provision_policy--connector_provision--default_behavior))
- `delete_account` (Attributes) The DeleteAccount message. (see [below for nested schema](#nestedatt--provision_policy--connector_provision--delete_account))

<a id="nestedatt--provision_policy--connector_provision--account_provision"></a>
### Nested Schema for `provision_policy.connector_provision.account_provision`

Read-Only:

- `config` (Attributes) (see [below for nested schema](#nestedatt--provision_policy--connector_provision--account_provision--config))
- `connector_id` (String) The connectorId field.
- `do_not_save` (Attributes) The DoNotSave message. (see [below for nested schema](#nestedatt--provision_policy--connector_provision--account_provision--do_not_save))
- `save_to_vault` (Attributes) The SaveToVault message. (see [below for nested schema](#nestedatt--provision_policy--connector_provision--account_provision--save_to_vault))
- `schema_id` (String) The schemaId field.

<a id="nestedatt--provision_policy--connector_provision--account_provision--config"></a>
### Nested Schema for `provision_policy.connector_provision.account_provision.config`


<a id="nestedatt--provision_policy--connector_provision--account_provision--do_not_save"></a>
### Nested Schema for `provision_policy.connector_provision.account_provision.do_not_save`


<a id="nestedatt--provision_policy--connector_provision--account_provision--save_to_vault"></a>
### Nested Schema for `provision_policy.connector_provision.account_provision.save_to_vault`

Read-Only:

- `vault_ids` (List of String) The vaultIds field.



<a id="nestedatt--provision_policy--connector_provision--default_behavior"></a>
### Nested Schema for `provision_policy.connector_provision.default_behavior`

Read-Only:

- `connector_id` (String) this checks if the entitlement is enabled by provisioning in a specific connector
 this can happen automatically and doesn't need any extra info


<a id="nestedatt--provision_policy--connector_provision--delete_account"></a>
### Nested Schema for `provision_policy.connector_provision.delete_account`

Read-Only:

- `connector_id` (String) The connectorId field.



<a id="nestedatt--provision_policy--delegated_provision"></a>
### Nested Schema for `provision_policy.delegated_provision`

Read-Only:

- `app_id` (String) The AppID of the entitlement to delegate provisioning to.
- `entitlement_id` (String) The ID of the entitlement we are delegating provisioning to.


<a id="nestedatt--provision_policy--external_ticket_provision"></a>
### Nested Schema for `provision_policy.external_ticket_provision`

Read-Only:

- `app_id` (String) The appId field.
- `connector_id` (String) The connectorId field.
- `external_ticket_provisioner_config_id` (String) The externalTicketProvisionerConfigId field.
- `instructions` (String) This field indicates a text body of instructions for the provisioner to indicate.


<a id="nestedatt--provision_policy--manual_provision"></a>
### Nested Schema for `provision_policy.manual_provision`

Read-Only:

- `instructions` (String) This field indicates a text body of instructions for the provisioner to indicate.
- `user_ids` (List of String) An array of users that are required to provision during this step.


<a id="nestedatt--provision_policy--unconfigured_provision"></a>
### Nested Schema for `provision_policy.unconfigured_provision`


<a id="nestedatt--provision_policy--webhook_provision"></a>
### Nested Schema for `provision_policy.webhook_provision`

Read-Only:

- `webhook_id` (String) The ID of the webhook to call for provisioning.