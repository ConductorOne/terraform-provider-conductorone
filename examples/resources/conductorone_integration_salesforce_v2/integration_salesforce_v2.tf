resource "conductorone_integration_salesforce_v2" "salesforce_v2" {
  app_id = conductorone_app.salesforce_v2.id
  user_ids = [
    conductorone_user.admin.id
  ]
  salesforce_group_oauth = {
    salesforce_instance_url        = "..."
    salesforce_username_for_email  = false
    salesforce_sync_connected_apps = false
  }
  salesforce_group_access_token = {
    salesforce_username            = "..."
    salesforce_password            = "..."
    salesforce_security_token      = "..."
    salesforce_instance_url        = "..."
    salesforce_username_for_email  = false
    salesforce_sync_connected_apps = false
  }
}
