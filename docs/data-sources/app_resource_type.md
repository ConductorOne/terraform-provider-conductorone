---
page_title: "conductorone_app_resource_type Data Source - terraform-provider-conductorone"
subcategory: ""
description: |-
  AppResourceType DataSource
---

# conductorone_app_resource_type (Data Source)

AppResourceType DataSource

This data source enables you to retrieve ConductorOne app resource types using the following search criteria:

* `app_ids` - List of app IDs to filter by
* `app_user_ids` - List of app user IDs to filter by
* `display_name` - Filter by display name
* `query` - Search query string
* `resource_type_ids` - List of resource type IDs to include
* `resource_type_trait_ids` - List of resource type trait IDs to include
* `exclude_resource_type_ids` - List of resource type IDs to exclude
* `exclude_resource_type_trait_ids` - List of resource type trait IDs to exclude

## Example Usage

```terraform
data "conductorone_app_resource_type" "my_app_resource_type" {
  app_ids = [
    "..."
  ]
  app_user_ids = [
    "..."
  ]
  display_name = "...my_display_name..."
  exclude_resource_type_ids = [
    "..."
  ]
  exclude_resource_type_trait_ids = [
    "..."
  ]
  page_size  = 4
  page_token = "...my_page_token..."
  query      = "...my_query..."
  resource_type_ids = [
    "..."
  ]
  resource_type_trait_ids = [
    "..."
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `app_ids` (List of String) A list of app IDs to restrict the search by.
- `app_user_ids` (List of String) A list of app user IDs to restrict the search by.
- `display_name` (String) Exact match on display name
- `exclude_resource_type_ids` (List of String) A list of resource type IDs to exclude from the search.
- `exclude_resource_type_trait_ids` (List of String) A list of resource type trait IDs to exclude from the search.
- `page_size` (Number) The pageSize where 10 <= pageSize <= 100, default 25.
- `page_token` (String) The pageToken field.
- `query` (String) Fuzzy search the display name of resource types.
- `resource_type_ids` (List of String) A list of resource type IDs to restrict the search by.
- `resource_type_trait_ids` (List of String) A list of resource type trait IDs to restrict the search by.

### Read-Only

- `app_id` (String) The ID of the app that is associated with the app resource type
- `created_at` (String)
- `deleted_at` (String)
- `id` (String) The unique ID for the app resource type.
- `next_page_token` (String) The nextPageToken is shown for the next page if the number of results is larger than the max page size.
 The server returns one page of results and the nextPageToken until all results are retreived.
 To retrieve the next page, use the same request and append a pageToken field with the value of nextPageToken shown on the previous page.
- `trait_ids` (List of String) Associated trait ids
- `updated_at` (String)