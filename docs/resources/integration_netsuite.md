---
page_title: "conductorone_integration_netsuite Resource - terraform-provider-conductorone"
subcategory: ""
description: |-
  Netsuite Integration Resource
---

# conductorone_integration_netsuite (Resource)

Netsuite Integration Resource

This resource allows you to configure an instance of the netsuite integration in ConductorOne.
It is always associated with an application. Optionally you can specify the list of users who are owners of the integration.
If owners are not specified, the integration will be owned by the user who created the resource.

## Example Usage

```terraform
resource "conductorone_integration_netsuite" "netsuite" {
  app_id = conductorone_app.netsuite.id
  user_ids = [
    conductorone_user.admin.id
  ]
  netsuite_account_id      = "..."
  netsuite_consumer_key    = "..."
  netsuite_consumer_secret = "..."
  netsuite_token_key       = "..."
  netsuite_token_secret    = "..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_id` (String) The ID for the Application that this integration should connected to.

### Optional

- `netsuite_account_id` (String) Account ID
- `netsuite_consumer_key` (String) Consumer key
- `netsuite_consumer_secret` (String, Sensitive) Consumer secret
- `netsuite_token_key` (String) Token key
- `netsuite_token_secret` (String, Sensitive) Token secret
- `user_ids` (List of String) A list of user IDs of who owns this integration. It defaults to the user who created the integration.

### Read-Only

- `created_at` (String) The time this integration was created.
- `deleted_at` (String) The time this integration was deleted.
- `id` (String) The ID of this integration.
- `updated_at` (String) The time this integration was last updated.
