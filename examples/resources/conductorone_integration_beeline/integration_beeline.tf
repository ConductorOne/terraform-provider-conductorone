resource "conductorone_integration_beeline" "beeline" {
  app_id = conductorone_app.beeline.id
  user_ids = [
    conductorone_user.admin.id
  ]
  beeline_client_site_id = "..."
  beeline_client_id      = "..."
  beeline_client_secret  = "..."
  base_url               = "..."
  auth_server_url        = "..."
}
