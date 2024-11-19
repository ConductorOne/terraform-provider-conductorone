---
page_title: "conductorone_integration_percipio Resource - terraform-provider-conductorone"
subcategory: ""
description: |-
  Percipio Integration Resource
---

# conductorone_integration_percipio (Resource)

Percipio Integration Resource

This resource allows you to configure an instance of the percipio integration in ConductorOne.
It is always associated with an application. Optionally you can specify the list of users who are owners of the integration.
If owners are not specified, the integration will be owned by the user who created the resource.

## Example Usage

```terraform
resource "conductorone_integration_percipio" "percipio" {
  app_id = conductorone_app.percipio.id
  user_ids = [
    conductorone_user.admin.id
  ]
  percipio_organization_id = "..."
  percipio_api_token       = "..."
  percipio_course_ids      = ["..."]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_id` (String) The ID for the Application that this integration should connected to.

### Optional

- `percipio_api_token` (String, Sensitive) API Token
- `percipio_course_ids` (List of String) Course IDs
- `percipio_organization_id` (String) Organization ID
- `user_ids` (List of String) A list of user IDs of who owns this integration. It defaults to the user who created the integration.

### Read-Only

- `created_at` (String) The time this integration was created.
- `deleted_at` (String) The time this integration was deleted.
- `id` (String) The ID of this integration.
- `updated_at` (String) The time this integration was last updated.