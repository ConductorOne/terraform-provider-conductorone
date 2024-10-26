---
page_title: "conductorone_integration_databricks Resource - terraform-provider-conductorone"
subcategory: ""
description: |-
  Databricks Integration Resource
---

# conductorone_integration_databricks (Resource)

Databricks Integration Resource

This resource allows you to configure an instance of the databricks integration in ConductorOne.
It is always associated with an application. Optionally you can specify the list of users who are owners of the integration.
If owners are not specified, the integration will be owned by the user who created the resource.

## Example Usage

```terraform
resource "conductorone_integration_databricks" "databricks" {
  app_id = conductorone_app.databricks.id
  user_ids = [
    conductorone_user.admin.id
  ]
  databricks_account_id    = "..."
  databricks_client_id     = "..."
  databricks_client_secret = "..."
  databricks_access_token  = "..."
  databricks_workspace     = "..."
  databricks_username      = "..."
  databricks_password      = "..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_id` (String) The ID for the Application that this integration should connected to.

### Optional

- `databricks_access_token` (String, Sensitive) Personal Access Token
- `databricks_account_id` (String) Account ID
- `databricks_client_id` (String) OAuth2 Client ID
- `databricks_client_secret` (String, Sensitive) OAuth2 Client Secret
- `databricks_password` (String, Sensitive) Password
- `databricks_username` (String) Username
- `databricks_workspace` (String) Workspace ID
- `user_ids` (List of String) A list of user IDs of who owns this integration. It defaults to the user who created the integration.

### Read-Only

- `created_at` (String) The time this integration was created.
- `deleted_at` (String) The time this integration was deleted.
- `id` (String) The ID of this integration.
- `updated_at` (String) The time this integration was last updated.