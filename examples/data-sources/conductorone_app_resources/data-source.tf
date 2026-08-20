data "conductorone_app_resources" "my_app_resources" {
  agent_statuses = [
    "AGENT_STATUS_DELETED"
  ]
  app_id = "...my_app_id..."
  app_ids = [
    "..."
  ]
  app_user_ids = [
    "..."
  ]
  credential_types = [
    "CREDENTIAL_TYPE_STATIC_SECRET"
  ]
  direction                         = "SORT_DIRECTION_DESC"
  exclude_deleted_apps              = false
  exclude_deleted_resource_bindings = false
  exclude_resource_ids = [
    "..."
  ]
  exclude_resource_type_trait_ids = [
    "..."
  ]
  nhi_types = [
    "NHI_TYPE_ASSUMABLE_ROLE"
  ]
  owner_user_ids = [
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
  secret_aging = {
    last_used_after       = "2021-03-15T09:45:25.932Z"
    last_used_before      = "2022-05-11T06:31:41.789Z"
    secret_created_after  = "2021-03-21T00:47:32.671Z"
    secret_created_before = "2022-04-20T09:16:36.843Z"
    secret_expires_after  = "2022-08-17T03:52:45.203Z"
    secret_expires_before = "2022-10-03T05:33:21.174Z"
  }
  sort_field         = "APP_RESOURCE_SORT_FIELD_SECRET_CREATED_AT"
  unowned_only       = false
  with_open_findings = true
}