resource "conductorone_integration_galileo_ft" "galileo_ft" {
  app_id = conductorone_app.galileo_ft.id
  user_ids = [
    conductorone_user.admin.id
  ]
  galileoft_provider_id   = "..."
  galileoft_api_login     = "..."
  galileoft_api_trans_key = "..."
  galileoft_hostname      = "..."
}
