resource "conductorone_integration_asana" "asana" {
  app_id = conductorone_app.asana.id
  user_ids = [
    conductorone_user.admin.id
  ]
  asana_api_key            = "..."
  asana_is_service_account = false
  asana_default_workspace  = "..."
  asana_use_scim_api       = false
}
