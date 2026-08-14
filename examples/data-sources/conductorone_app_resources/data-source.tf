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
  secret_aging_filter = {
    last_used_after       = "2022-05-16T14:37:08.306Z"
    last_used_before      = "2022-02-03T20:33:50.527Z"
    secret_created_after  = "2022-12-06T02:32:52.852Z"
    secret_created_before = "2021-05-15T13:32:27.740Z"
    secret_expires_after  = "2022-10-07T06:31:43.579Z"
    secret_expires_before = "2022-06-04T02:07:04.673Z"
  }
  sort_field         = "APP_RESOURCE_SORT_FIELD_SECRET_CREATED_AT"
  unowned_only       = false
  with_open_findings = true
}