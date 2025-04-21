resource "conductorone_integration_wiz_id" "wiz_id" {
  app_id = conductorone_app.wiz_id.id
  user_ids = [
    conductorone_user.admin.id
  ]
  wiz_client_id     = "..."
  wiz_client_secret = "..."
  endpoint_url      = "..."
  auth_url          = "..."
  audience          = "..."
  project_id        = "..."
  resource_ids      = ["..."]
  resource_tags     = "..."
  resource_types    = ["..."]
}
