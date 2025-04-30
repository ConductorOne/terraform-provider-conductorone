resource "conductorone_integration_gcp_gws" "gcp_gws" {
  app_id = conductorone_app.gcp_gws.id
  user_ids = [
    conductorone_user.admin.id
  ]
  customer_id           = "..."
  domain                = "..."
  administrator_email   = "..."
  credentials_json      = "..."
  skip_system_accounts  = false
  skip_default_projects = false
  sync_secrets          = false
  project_ids           = ["..."]
}
