resource "conductorone_integration_crowdstrike" "crowdstrike" {
  app_id = conductorone_app.crowdstrike.id
  user_ids = [
    conductorone_user.admin.id
  ]
  crowdstrike_client_id     = "..."
  crowdstrike_client_secret = "..."
  region                    = "..."
}
