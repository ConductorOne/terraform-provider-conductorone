resource "conductorone_integration_google_workspace" "google_workspace" {
  app_id = conductorone_app.google_workspace.id
  user_ids = [
   conductorone_user.admin.id
  ]
  customer_id = "..."
  domain = "..."
  administrator_email = "..."
  credentials_json = "..."
  }
