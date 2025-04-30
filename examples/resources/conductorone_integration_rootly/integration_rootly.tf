resource "conductorone_integration_rootly" "rootly" {
  app_id = conductorone_app.rootly.id
  user_ids = [
    conductorone_user.admin.id
  ]
  rootly_api_key = "..."
}
