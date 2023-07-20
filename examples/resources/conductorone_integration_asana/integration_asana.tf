resource "conductorone_integration_asana" "asana" {
  app_id = conductorone_app.asana.id
  user_ids = [
    conductorone_user.admin.id
  ]
  asana_api_key = "..."
}
