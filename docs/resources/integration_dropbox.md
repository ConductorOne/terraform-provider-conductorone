---
page_title: "conductorone_integration_dropbox Resource - terraform-provider-conductorone"
subcategory: ""
description: |-
  Dropbox Integration Resource
---

# conductorone_integration_dropbox (Resource)

Dropbox Integration Resource

This resource allows you to configure an instance of the dropbox integration in ConductorOne.
It is always associated with an application. Optionally you can specify the list of users who are owners of the integration.
If owners are not specified, the integration will be owned by the user who created the resource.

## Example Usage

```terraform
resource "conductorone_integration_dropbox" "dropbox" {
  app_id = conductorone_app.dropbox.id
  user_ids = [
    conductorone_user.admin.id
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_id` (String) The ID for the Application that this integration should connected to.

### Optional

- `user_ids` (List of String) A list of user IDs of who owns this integration. It defaults to the user who created the integration.

### Read-Only

- `created_at` (String) The time this integration was created.
- `deleted_at` (String) The time this integration was deleted.
- `id` (String) The ID of this integration.
- `updated_at` (String) The time this integration was last updated.
