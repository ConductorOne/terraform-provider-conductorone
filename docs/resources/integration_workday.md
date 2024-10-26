---
page_title: "conductorone_integration_workday Resource - terraform-provider-conductorone"
subcategory: ""
description: |-
  Workday Integration Resource
---

# conductorone_integration_workday (Resource)

Workday Integration Resource

This resource allows you to configure an instance of the workday integration in ConductorOne.
It is always associated with an application. Optionally you can specify the list of users who are owners of the integration.
If owners are not specified, the integration will be owned by the user who created the resource.

## Example Usage

```terraform
resource "conductorone_integration_workday" "workday" {
  app_id = conductorone_app.workday.id
  user_ids = [
    conductorone_user.admin.id
  ]
  workday_client_id     = "..."
  workday_client_secret = "..."
  refresh_token         = "..."
  workday_url           = "..."
  tenant_name           = "..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_id` (String) The ID for the Application that this integration should connected to.

### Optional

- `refresh_token` (String, Sensitive) Refresh Token
- `tenant_name` (String) Tenant Name
- `user_ids` (List of String) A list of user IDs of who owns this integration. It defaults to the user who created the integration.
- `workday_client_id` (String) Client ID
- `workday_client_secret` (String, Sensitive) Client Secret
- `workday_url` (String) Workday URL

### Read-Only

- `created_at` (String) The time this integration was created.
- `deleted_at` (String) The time this integration was deleted.
- `id` (String) The ID of this integration.
- `updated_at` (String) The time this integration was last updated.