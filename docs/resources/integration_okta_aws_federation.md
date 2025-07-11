---
page_title: "conductorone_integration_okta_aws_federation Resource - terraform-provider-conductorone"
subcategory: ""
description: |-
  Okta_aws_federation Integration Resource
---

# conductorone_integration_okta_aws_federation (Resource)

Okta_aws_federation Integration Resource

This resource allows you to configure an instance of the okta_aws_federation integration in ConductorOne.
It is always associated with an application. Optionally you can specify the list of users who are owners of the integration.
If owners are not specified, the integration will be owned by the user who created the resource.

## Example Usage

```terraform
resource "conductorone_integration_okta_aws_federation" "okta_aws_federation" {
  app_id = conductorone_app.okta_aws_federation.id
  user_ids = [
    conductorone_user.admin.id
  ]
  okta_aws_federation_domain                                            = "..."
  okta_aws_federation_api_token                                         = "..."
  okta_aws_federation_aws_okta_app_id                                   = "..."
  okta_aws_federation_allow_group_to_direct_conversion_for_provisioning = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_id` (String) The ID for the Application that this integration should connected to.

### Optional

- `okta_aws_federation_allow_group_to_direct_conversion_for_provisioning` (Boolean) Allow group to direct assignment conversion for provisioning
- `okta_aws_federation_api_token` (String, Sensitive) API token
- `okta_aws_federation_aws_okta_app_id` (String) AWS Okta App ID
- `okta_aws_federation_domain` (String) Okta domain
- `user_ids` (List of String) A list of user IDs of who owns this integration. It defaults to the user who created the integration.

### Read-Only

- `created_at` (String) The time this integration was created.
- `deleted_at` (String) The time this integration was deleted.
- `id` (String) The ID of this integration.
- `updated_at` (String) The time this integration was last updated.
