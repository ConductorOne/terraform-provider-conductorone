---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "conductorone_bundle_automation Resource - terraform-provider-conductorone"
subcategory: ""
description: |-
  BundleAutomation Resource
---

# conductorone_bundle_automation (Resource)

BundleAutomation Resource

## Example Usage

```terraform
resource "conductorone_bundle_automation" "my_bundleautomation" {
  bundle_automation_rule_entitlement = {
    entitlement_refs = [
      {
        app_id = "...my_app_id..."
        id     = "...my_id..."
      }
    ]
  }
  create_tasks = false
  delete_bundle_automation_request = {
    # ...
  }
  enabled            = true
  request_catalog_id = "...my_request_catalog_id..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `request_catalog_id` (String)

### Optional

- `bundle_automation_rule_entitlement` (Attributes) The BundleAutomationRuleEntitlement message. (see [below for nested schema](#nestedatt--bundle_automation_rule_entitlement))
- `create_tasks` (Boolean) The createTasks field.
- `delete_bundle_automation_request` (Attributes) The DeleteBundleAutomationRequest message. (see [below for nested schema](#nestedatt--delete_bundle_automation_request))
- `enabled` (Boolean) The enabled field.

### Read-Only

- `bundle_automation_last_run_state` (Attributes) The BundleAutomationLastRunState message. (see [below for nested schema](#nestedatt--bundle_automation_last_run_state))
- `created_at` (String)
- `tenant_id` (String) The tenantId field.
- `updated_at` (String)

<a id="nestedatt--bundle_automation_rule_entitlement"></a>
### Nested Schema for `bundle_automation_rule_entitlement`

Optional:

- `entitlement_refs` (Attributes List) The entitlementRefs field. (see [below for nested schema](#nestedatt--bundle_automation_rule_entitlement--entitlement_refs))

<a id="nestedatt--bundle_automation_rule_entitlement--entitlement_refs"></a>
### Nested Schema for `bundle_automation_rule_entitlement.entitlement_refs`

Optional:

- `app_id` (String) The appId field.
- `id` (String) The id field.



<a id="nestedatt--delete_bundle_automation_request"></a>
### Nested Schema for `delete_bundle_automation_request`


<a id="nestedatt--bundle_automation_last_run_state"></a>
### Nested Schema for `bundle_automation_last_run_state`

Read-Only:

- `error_message` (String) The errorMessage field.
- `last_run_at` (String)
- `status` (String) The status field. must be one of ["BUNDLE_AUTOMATION_RUN_STATUS_UNSPECIFIED", "BUNDLE_AUTOMATION_RUN_STATUS_SUCCESS", "BUNDLE_AUTOMATION_RUN_STATUS_FAILURE", "BUNDLE_AUTOMATION_RUN_STATUS_IN_PROGRESS"]

## Import

Import is supported using the following syntax:

```shell
terraform import conductorone_bundle_automation.my_conductorone_bundle_automation ""
```
