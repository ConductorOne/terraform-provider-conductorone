---
page_title: "conductorone_integration_galileo_ft Resource - terraform-provider-conductorone"
subcategory: ""
description: |-
  Galileo_ft Integration Resource
---

# conductorone_integration_galileo_ft (Resource)

Galileo_ft Integration Resource

This resource allows you to configure an instance of the galileo_ft integration in ConductorOne.
It is always associated with an application. Optionally you can specify the list of users who are owners of the integration.
If owners are not specified, the integration will be owned by the user who created the resource.

## Example Usage

```terraform
resource "conductorone_integration_galileo_ft" "galileo_ft" {
  app_id = conductorone_app.galileo_ft.id
  user_ids = [
    conductorone_user.admin.id
  ]
  galileoft_provider_id   = "..."
  galileoft_api_login     = "..."
  galileoft_api_trans_key = "..."
  galileoft_hostname      = "..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_id` (String) The ID for the Application that this integration should connected to.

### Optional

- `galileoft_api_login` (String) Username
- `galileoft_api_trans_key` (String, Sensitive) Password
- `galileoft_hostname` (String) Hostname
- `galileoft_provider_id` (String) Organization ID
- `user_ids` (List of String) A list of user IDs of who owns this integration. It defaults to the user who created the integration.

### Read-Only

- `created_at` (String) The time this integration was created.
- `deleted_at` (String) The time this integration was deleted.
- `id` (String) The ID of this integration.
- `updated_at` (String) The time this integration was last updated.