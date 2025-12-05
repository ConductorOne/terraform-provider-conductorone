resource "conductorone_integration_dropbox" "dropbox" {
  app_id = conductorone_app.dropbox.id
  user_ids = [
    conductorone_user.admin.id
  ]
}
