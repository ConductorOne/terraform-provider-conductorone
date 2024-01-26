resource "conductorone_integration_snipe_it" "snipe_it" {
  app_id = conductorone_app.snipe_it.id
  user_ids = [
    conductorone_user.admin.id
  ]
  snipeit_base_url     = "..."
  snipeit_access_token = "..."
}
