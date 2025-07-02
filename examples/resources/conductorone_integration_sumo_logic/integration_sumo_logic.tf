resource "conductorone_integration_sumo_logic" "sumo_logic" {
  app_id = conductorone_app.sumo_logic.id
  user_ids = [
    conductorone_user.admin.id
  ]
  api_access_id            = "..."
  api_access_key           = "..."
  api_base_url             = "..."
  include_service_accounts = false
}
