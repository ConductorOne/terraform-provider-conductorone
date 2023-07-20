resource "conductorone_integration_ramp" "ramp" {
  app_id = conductorone_app.ramp.id
  user_ids = [
    conductorone_user.admin.id
  ]
}
