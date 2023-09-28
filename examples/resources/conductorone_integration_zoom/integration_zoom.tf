resource "conductorone_integration_zoom" "zoom" {
  app_id = conductorone_app.zoom.id
  user_ids = [
   conductorone_user.admin.id
  ]
  zoom_account_id = "..."
  zoom_client_id = "..."
  zoom_client_secret = "..."
  }
