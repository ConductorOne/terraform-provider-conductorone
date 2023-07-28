---
page_title: "conductorone_catalog_visibility_bindings Resource - terraform-provider-conductorone"
subcategory: ""
description: |-
  CatalogVisibilityBindings Resource
---

# conductorone_catalog_visibility_bindings (Resource)

CatalogVisibilityBindings Resource

This resource allows you to manage visibility bindings for a specific request catalog in ConductorOne.
When creating a `catalog_visbility_bindings` resource you must provide a `catalog_id`, and specify the `app_id` and `id` of at least one app entitlement.
The entitlements you define here will determine which users have access to the specified catalog.

## Example Usage

```terraform
resource "conductorone_catalog" "test_catalog" {
  display_name        = "terraform created catalog"
  description         = "terraform test"
  visible_to_everyone = "false"
  published           = "true"
}
```

```terraform
resource "conductorone_catalog_visibility_bindings" "test_bindings" {
  catalog_id = conductorone_catalog.test_catalog.id
  access_entitlements = [
    {
      app_id = data.conductorone_app_entitlement.okta_everyone.app_id
      id     = data.conductorone_app_entitlement.okta_everyone.id
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `access_entitlements` (Attributes List) The accessEntitlements field. (see [below for nested schema](#nestedatt--access_entitlements))
- `catalog_id` (String)

<a id="nestedatt--access_entitlements"></a>
### Nested Schema for `access_entitlements`

Required:

- `app_id` (String) The appId field.
- `id` (String) The id field.