---
page_title: "conductorone_app_owner Resource - terraform-provider-conductorone"
subcategory: ""
description: |-
  AppOwner Resource
---

# conductorone_app_owner (Resource)

AppOwner Resource

This resource allows you to set App owners.
In order to set owners, you must provide the list of `user_ids` for the users you wish to set as owners.
You must also provide the `app_id` of the app you wish to add the owners to.
Note: The owners you set here will replace any existing owners for the app.

## Example Usage

```terraform
resource "conductorone_app_owner" "my_app_owner" {
  app_id = "...my_app_id..."
  user_ids = [
    "..."
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_id` (String) Requires replacement if changed.

### Optional

- `user_ids` (List of String) The user_ids field for the users to set as an owner of the app. Requires replacement if changed.