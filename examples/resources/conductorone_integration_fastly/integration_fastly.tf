resource "conductorone_integration_fastly" "fastly" {
  app_id = conductorone_app.fastly.id
  user_ids = [
    conductorone_user.admin.id
  ]
  fastly_access_token = "..."
}
