---
page_title: "conductorone_integration_jira_cloud Resource - terraform-provider-conductorone"
subcategory: ""
description: |-
  Jira_cloud Integration Resource
---

# conductorone_integration_jira_cloud (Resource)

Jira_cloud Integration Resource

This resource allows you to configure an instance of the jira_cloud integration in ConductorOne.
It is always associated with an application. Optionally you can specify the list of users who are owners of the integration.
If owners are not specified, the integration will be owned by the user who created the resource.

## Example Usage

```terraform
resource "conductorone_integration_jira_cloud" "jira_cloud" {
  app_id = conductorone_app.jira_cloud.id
  user_ids = [
    conductorone_user.admin.id
  ]
  jiracloud_domain   = "..."
  jiracloud_username = "..."
  jiracloud_apikey   = "..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_id` (String) The ID for the Application that this integration should connected to.

### Optional

- `jiracloud_apikey` (String, Sensitive) Jira API Key
- `jiracloud_domain` (String) Jira Site Domain
- `jiracloud_username` (String) Jira Username
- `user_ids` (List of String) A list of user IDs of who owns this integration. It defaults to the user who created the integration.

### Read-Only

- `created_at` (String) The time this integration was created.
- `deleted_at` (String) The time this integration was deleted.
- `id` (String) The ID of this integration.
- `updated_at` (String) The time this integration was last updated.