resource "conductorone_integration_box" "box" {
  app_id = conductorone_app.box.id
  user_ids = [
    conductorone_user.admin.id
  ]
  box_client_id     = "..."
  box_client_secret = "..."
  box_enterprise_id = "..."
}
