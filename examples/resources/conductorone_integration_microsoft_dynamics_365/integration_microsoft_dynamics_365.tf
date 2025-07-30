resource "conductorone_integration_microsoft_dynamics_365" "microsoft_dynamics_365" {
  app_id = conductorone_app.microsoft_dynamics_365.id
  user_ids = [
    conductorone_user.admin.id
  ]
  environment_url        = "..."
  dynamics_client_id     = "..."
  dynamics_client_secret = "..."
}
