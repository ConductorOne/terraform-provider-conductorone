resource "conductorone_integration_azure_devops" "azure_devops" {
  app_id = conductorone_app.azure_devops.id
  user_ids = [
    conductorone_user.admin.id
  ]
  organization_url      = "..."
  personal_access_token = "..."
  sync_grant_sources    = false
}
