resource "conductorone_integration_jumpcloud" "jumpcloud" {
  app_id = conductorone_app.jumpcloud.id
  user_ids = [
    conductorone_user.admin.id
  ]
  jumpcloud_api_key = "..."
}
