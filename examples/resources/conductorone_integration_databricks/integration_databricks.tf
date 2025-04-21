resource "conductorone_integration_databricks" "databricks" {
  app_id = conductorone_app.databricks.id
  user_ids = [
    conductorone_user.admin.id
  ]
  databricks_account_hostname = "..."
  databricks_hostname         = "..."
  databricks_account_id       = "..."
  databricks_client_id        = "..."
  databricks_client_secret    = "..."
  databricks_access_token     = "..."
  databricks_workspace        = "..."
  databricks_username         = "..."
  databricks_password         = "..."
}
