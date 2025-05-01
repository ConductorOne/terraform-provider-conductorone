resource "conductorone_integration_tray_ai" "tray_ai" {
  app_id = conductorone_app.tray_ai.id
  user_ids = [
    conductorone_user.admin.id
  ]
  trayai_authorization_token = "..."
}
