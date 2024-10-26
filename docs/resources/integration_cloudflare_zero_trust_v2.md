---
page_title: "conductorone_integration_cloudflare_zero_trust_v2 Resource - terraform-provider-conductorone"
subcategory: ""
description: |-
  Cloudflarezerotrust_v2 Integration Resource
---

# conductorone_integration_cloudflare_zero_trust_v2 (Resource)

Cloudflare_zero_trust_v2 Integration Resource

This resource allows you to configure an instance of the cloudflare_zero_trust_v2 integration in ConductorOne.
It is always associated with an application. Optionally you can specify the list of users who are owners of the integration.
If owners are not specified, the integration will be owned by the user who created the resource.

## Example Usage

```terraform
resource "conductorone_integration_cloudflare_zero_trust_v2" "cloudflare_zero_trust_v2" {
  app_id = conductorone_app.cloudflare_zero_trust_v2.id
  user_ids = [
    conductorone_user.admin.id
  ]
  account_id = "..."
  api_token  = "..."
  api_key    = "..."
  email      = "..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_id` (String) The ID for the Application that this integration should connected to.

### Optional

- `account_id` (String) Account ID (required)
- `api_key` (String, Sensitive) API key (required if API token not provided)
- `api_token` (String, Sensitive) API token
- `email` (String) Email (required if API token not provided)
- `user_ids` (List of String) A list of user IDs of who owns this integration. It defaults to the user who created the integration.

### Read-Only

- `created_at` (String) The time this integration was created.
- `deleted_at` (String) The time this integration was deleted.
- `id` (String) The ID of this integration.
- `updated_at` (String) The time this integration was last updated.