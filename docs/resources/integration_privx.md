---
page_title: "conductorone_integration_privx Resource - terraform-provider-conductorone"
subcategory: ""
description: |-
  Privx Integration Resource
---

# conductorone_integration_privx (Resource)

Privx Integration Resource

This resource allows you to configure an instance of the privx integration in ConductorOne.
It is always associated with an application. Optionally you can specify the list of users who are owners of the integration.
If owners are not specified, the integration will be owned by the user who created the resource.

## Example Usage

```terraform
resource "conductorone_integration_privx" "privx" {
  app_id = conductorone_app.privx.id
  user_ids = [
    conductorone_user.admin.id
  ]
  privx_base_url            = "..."
  privx_client_id           = "..."
  privx_client_secret       = "..."
  privx_oauth_client_id     = "..."
  privx_oauth_client_secret = "..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_id` (String) The ID for the Application that this integration should connected to.

### Optional

- `privx_base_url` (String) Base URL
- `privx_client_id` (String) Client ID
- `privx_client_secret` (String, Sensitive) Client secret
- `privx_oauth_client_id` (String) OAuth client ID
- `privx_oauth_client_secret` (String, Sensitive) OAuth client secret
- `user_ids` (List of String) A list of user IDs of who owns this integration. It defaults to the user who created the integration.

### Read-Only

- `created_at` (String) The time this integration was created.
- `deleted_at` (String) The time this integration was deleted.
- `id` (String) The ID of this integration.
- `updated_at` (String) The time this integration was last updated.
