data "conductorone_app_resources" "my_app_resources" {
  app_id = "...my_app_id..."
  app_user_ids = [
    "..."
  ]
  exclude_deleted_resource_bindings = false
  exclude_resource_ids = [
    "..."
  ]
  exclude_resource_type_trait_ids = [
    "..."
  ]
  page_size  = 7
  page_token = "...my_page_token..."
  query      = "...my_query..."
  refs = [
    {
      app_id               = "...my_app_id..."
      app_resource_type_id = "...my_app_resource_type_id..."
      id                   = "...my_id..."
    }
  ]
  resource_ids = [
    "..."
  ]
  resource_type_ids = [
    "..."
  ]
  resource_type_trait_ids = [
    "..."
  ]
}