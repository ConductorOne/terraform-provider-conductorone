resource "conductorone_integration_workato" "workato" {
  app_id = conductorone_app.workato.id
  user_ids = [
    conductorone_user.admin.id
  ]
  workato_api_key                   = "..."
  workato_data_center               = "..."
  workato_env                       = "..."
  workato_disable_custom_roles_sync = false
}
