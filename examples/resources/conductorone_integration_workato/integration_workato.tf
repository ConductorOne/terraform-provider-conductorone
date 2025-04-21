resource "conductorone_integration_workato" "workato" {
  app_id = conductorone_app.workato.id
  user_ids = [
    conductorone_user.admin.id
  ]
  workato_api_key     = "..."
  workato_data_center = "..."
  workato_env         = "..."
}
