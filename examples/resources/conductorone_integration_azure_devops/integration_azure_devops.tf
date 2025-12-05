resource "conductorone_integration_azure_devops" "azure_devops" {
  app_id = conductorone_app.azure_devops.id
  user_ids = [
    conductorone_user.admin.id
  ]
  azure_devops_group_oauth = {
    organization_url                       = "..."
    azure_devops_tenant_id                 = "..."
    oauth2_client_cred_grant_client_id     = "..."
    oauth2_client_cred_grant_client_secret = "..."
    sync_teams                             = false
    sync_organization                      = false
  }
  azure_devops_group_pat = {
    organization_url      = "..."
    personal_access_token = "..."
    sync_teams            = false
    sync_organization     = false
  }
}
