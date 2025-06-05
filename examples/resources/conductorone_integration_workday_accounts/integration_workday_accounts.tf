resource "conductorone_integration_workday_accounts" "workday_accounts" {
  app_id = conductorone_app.workday_accounts.id
  user_ids = [
    conductorone_user.admin.id
  ]
  workday_rest_api_endpoint               = "..."
  workday_client_id                       = "..."
  workday_client_secret                   = "..."
  workday_refresh_token                   = "..."
  workday_sync_service_centers            = false
  workday_sync_user_based_security_groups = false
  workday_user_based_security_groups      = ["..."]
  workday_security_group_types            = ["..."]
  workday_security_groups                 = ["..."]
}
