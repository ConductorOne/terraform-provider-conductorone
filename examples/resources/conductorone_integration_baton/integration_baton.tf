resource "conductorone_integration_baton" "baton" {
  app_id = conductorone_app.baton.id
  user_ids = [
    conductorone_user.admin.id
  ]
}
