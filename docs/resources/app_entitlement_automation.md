---
page_title: "conductorone_app_entitlement_automation Resource - terraform-provider-conductorone"
subcategory: ""
description: |-
  AppEntitlementAutomation Resource
---

# conductorone_app_entitlement_automation (Resource)

AppEntitlementAutomation Resource

This resource allows you to configure an App Entitlement Automation instance in ConductorOne.
It enables automation rules for entitlements within the ConductorOne application, allowing for granular control of access.

When creating an entitlement automation resource, you must specify the associated `app_entitlement_id` and `app_id`.
You can define multiple rule types, including basic expressions, CEL expressions, and entitlement references.

NOTE: This resource can only be created for an entitlement on the ConductorOne application.

## Example Usage

```terraform
resource "conductorone_app_entitlement_automation" "my_app_entitlement_automation" {
  app_entitlement_automation_last_run_status = {
    # ...
  }
  app_entitlement_automation_rule_basic = {
    expression = "...my_expression..."
  }
  app_entitlement_automation_rule_cel = {
    expression = "...my_expression..."
  }
  app_entitlement_automation_rule_entitlement = {
    entitlement_refs = [
      {
        app_id = "...my_app_id..."
        id     = "...my_id..."
      }
    ]
  }
  app_entitlement_automation_rule_none = {
    # ...
  }
  app_entitlement_id = "...my_app_entitlement_id..."
  app_id             = "...my_app_id..."
  description        = "...my_description..."
  display_name       = "...my_display_name..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_entitlement_id` (String)
- `app_id` (String)

### Optional

- `app_entitlement_automation_last_run_status` (Attributes) The AppEntitlementAutomationLastRunStatus message. Requires replacement if changed. (see [below for nested schema](#nestedatt--app_entitlement_automation_last_run_status))
- `app_entitlement_automation_rule_basic` (Attributes) The AppEntitlementAutomationRuleBasic message. (see [below for nested schema](#nestedatt--app_entitlement_automation_rule_basic))
- `app_entitlement_automation_rule_cel` (Attributes) The AppEntitlementAutomationRuleCEL message. (see [below for nested schema](#nestedatt--app_entitlement_automation_rule_cel))
- `app_entitlement_automation_rule_entitlement` (Attributes) The AppEntitlementAutomationRuleEntitlement message. (see [below for nested schema](#nestedatt--app_entitlement_automation_rule_entitlement))
- `app_entitlement_automation_rule_none` (Attributes) The AppEntitlementAutomationRuleNone message. (see [below for nested schema](#nestedatt--app_entitlement_automation_rule_none))
- `description` (String) The description of the app entitlement.
- `display_name` (String) The display name of the app entitlement.

### Read-Only

- `created_at` (String)
- `updated_at` (String)

<a id="nestedatt--app_entitlement_automation_last_run_status"></a>
### Nested Schema for `app_entitlement_automation_last_run_status`

Read-Only:

- `error_message` (String) The errorMessage field.
- `last_completed_at` (String)
- `status` (String) The status field. must be one of ["APP_ENTITLEMENT_AUTOMATION_RUN_STATUS_UNSPECIFIED", "APP_ENTITLEMENT_AUTOMATION_RUN_STATUS_SUCCESS", "APP_ENTITLEMENT_AUTOMATION_RUN_STATUS_FAILED", "APP_ENTITLEMENT_AUTOMATION_RUN_STATUS_IN_PROGRESS"]


<a id="nestedatt--app_entitlement_automation_rule_basic"></a>
### Nested Schema for `app_entitlement_automation_rule_basic`

Optional:

- `expression` (String) The expression field.


<a id="nestedatt--app_entitlement_automation_rule_cel"></a>
### Nested Schema for `app_entitlement_automation_rule_cel`

Optional:

- `expression` (String) The expression field.


<a id="nestedatt--app_entitlement_automation_rule_entitlement"></a>
### Nested Schema for `app_entitlement_automation_rule_entitlement`

Optional:

- `entitlement_refs` (Attributes List) The entitlementRefs field. (see [below for nested schema](#nestedatt--app_entitlement_automation_rule_entitlement--entitlement_refs))

<a id="nestedatt--app_entitlement_automation_rule_entitlement--entitlement_refs"></a>
### Nested Schema for `app_entitlement_automation_rule_entitlement.entitlement_refs`

Optional:

- `app_id` (String) The appId field.
- `id` (String) The id field.



<a id="nestedatt--app_entitlement_automation_rule_none"></a>
### Nested Schema for `app_entitlement_automation_rule_none`