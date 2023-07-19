---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "conductorone_integration_netsuite Resource - conductorone"
subcategory: ""
description: |-
  Netsuite Integration Resource
---

# conductorone_integration_netsuite (Resource)

Netsuite Integration Resource



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_id` (String) The appId field.

### Optional

- `netsuite_account_id` (String) Netsuite Account ID
- `netsuite_consumer_key` (String) Netsuite Consumer Key
- `netsuite_consumer_secret` (String, Sensitive) Netsuite Consumer Secret
- `netsuite_token_key` (String) Netsuite Token Key
- `netsuite_token_secret` (String, Sensitive) Netsuite Token Secret
- `user_ids` (List of String) The userIds field.

### Read-Only

- `created_at` (String)
- `deleted_at` (String)
- `id` (String) The id field.
- `updated_at` (String)