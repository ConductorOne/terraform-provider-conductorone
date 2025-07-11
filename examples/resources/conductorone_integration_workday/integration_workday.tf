resource "conductorone_integration_workday" "workday" {
  app_id = conductorone_app.workday.id
  user_ids = [
    conductorone_user.admin.id
  ]
  workday_client_id            = "..."
  workday_client_secret        = "..."
  refresh_token                = "..."
  workday_url                  = "..."
  tenant_name                  = "..."
  workday_report_url           = "..."
  workday_report_username      = "..."
  workday_report_user_password = "..."
}
