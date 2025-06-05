resource "conductorone_integration_teamcity" "teamcity" {
  app_id = conductorone_app.teamcity.id
  user_ids = [
    conductorone_user.admin.id
  ]
  teamcity_access_token       = "..."
  teamcity_instance_url       = "..."
  teamcity_sync_grant_sources = false
}
