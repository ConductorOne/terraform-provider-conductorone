resource "conductorone_integration_google_workspace_v2" "google_workspace_v2" {
  app_id = conductorone_app.google_workspace_v2.id
  user_ids = [
    conductorone_user.admin.id
  ]
  customer_id         = "..."
  domain              = "..."
  administrator_email = "..."
  credentials_json    = "..."
}
