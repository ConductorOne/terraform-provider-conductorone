resource "conductorone_integration_wiz_query" "wiz_query" {
  app_id = conductorone_app.wiz_query.id
  user_ids = [
    conductorone_user.admin.id
  ]
  wiz_client_id                = "..."
  wiz_client_secret            = "..."
  endpoint_url                 = "..."
  auth_url                     = "..."
  audience                     = "..."
  project_id                   = "..."
  resource_ids                 = ["..."]
  resource_tags                = "..."
  resource_types               = ["..."]
  enable_sync_identities       = false
  enable_sync_service_accounts = false
}
