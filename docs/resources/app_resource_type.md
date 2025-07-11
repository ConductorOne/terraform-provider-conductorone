---
page_title: "conductorone_app_resource_type Resource - terraform-provider-conductorone"
subcategory: ""
description: |-
  AppResourceType Resource
---

# conductorone_app_resource_type (Resource)

AppResourceType Resource

This resource allows you to create and manage an app resource type.
When creating an app resource type you must provide an app_id, display_name and resource_type.

## Example Usage

```terraform
resource "conductorone_app_resource_type" "my_app_resource_type" {
  app_id        = "...my_app_id..."
  display_name  = "...my_display_name..."
  resource_type = "ROLE"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_id` (String)
- `display_name` (String) The displayName field.
- `resource_type` (String) The resourceType field. must be one of ["ROLE", "GROUP", "LICENSE", "PROJECT", "CATALOG", "CUSTOM", "VAULT"]; Requires replacement if changed.

### Read-Only

- `created_at` (String)
- `expanded` (Attributes List) The expanded field. (see [below for nested schema](#nestedatt--expanded))
- `id` (String) The unique ID for the app resource type.
- `trait_ids` (List of String) Associated trait ids
- `updated_at` (String)

<a id="nestedatt--expanded"></a>
### Nested Schema for `expanded`