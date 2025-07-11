---
page_title: "conductorone_integration_okta_ciam Resource - terraform-provider-conductorone"
subcategory: ""
description: |-
  Okta_ciam Integration Resource
---

# conductorone_integration_okta_ciam (Resource)

Okta_ciam Integration Resource

This resource allows you to configure an instance of the okta_ciam integration in ConductorOne.
It is always associated with an application. Optionally you can specify the list of users who are owners of the integration.
If owners are not specified, the integration will be owned by the user who created the resource.

## Example Usage

```terraform
resource "conductorone_integration_okta_ciam" "okta_ciam" {
  app_id = conductorone_app.okta_ciam.id
  user_ids = [
    conductorone_user.admin.id
  ]
  okta_ciam_domain        = "..."
  okta_ciam_api_token     = "..."
  okta_ciam_email_domains = "..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_id` (String) The ID for the Application that this integration should connected to.

### Optional

- `okta_ciam_api_token` (String, Sensitive) API token
- `okta_ciam_domain` (String) Okta domain
- `okta_ciam_email_domains` (String) Okta email domains (optional)
- `user_ids` (List of String) A list of user IDs of who owns this integration. It defaults to the user who created the integration.

### Read-Only

- `created_at` (String) The time this integration was created.
- `deleted_at` (String) The time this integration was deleted.
- `id` (String) The ID of this integration.
- `updated_at` (String) The time this integration was last updated.
