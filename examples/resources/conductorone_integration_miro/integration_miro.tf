resource "conductorone_integration_miro" "miro" {
  app_id = conductorone_app.miro.id
  user_ids = [
    conductorone_user.admin.id
  ]
  miro_access_token = "..."
}
