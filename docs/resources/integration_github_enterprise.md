---
page_title: "conductorone_integration_github_enterprise Resource - terraform-provider-conductorone"
subcategory: ""
description: |-
  Github_enterprise Integration Resource
---

# conductorone_integration_github_enterprise (Resource)

Github_enterprise Integration Resource

This resource allows you to configure an instance of the github_enterprise integration in ConductorOne.
It is always associated with an application. Optionally you can specify the list of users who are owners of the integration.
If owners are not specified, the integration will be owned by the user who created the resource.

## Example Usage

```terraform
resource "conductorone_integration_github_enterprise" "github_enterprise" {
  app_id = conductorone_app.github_enterprise.id
  user_ids = [
    conductorone_user.admin.id
  ]
  github_instance_url = "..."
  github_access_token = "..."
  github_org_list     = ["..."]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_id` (String) The ID for the Application that this integration should connected to.

### Optional

- `github_access_token` (String, Sensitive) Personal access token
- `github_instance_url` (String) Instance URL
- `github_org_list` (List of String) Organizations (optional)
- `user_ids` (List of String) A list of user IDs of who owns this integration. It defaults to the user who created the integration.

### Read-Only

- `created_at` (String) The time this integration was created.
- `deleted_at` (String) The time this integration was deleted.
- `id` (String) The ID of this integration.
- `updated_at` (String) The time this integration was last updated.
