resource "conductorone_integration_auth_0" "auth_0" {
  app_id = conductorone_app.auth_0.id
  user_ids = [
    conductorone_user.admin.id
  ]
  auth0_base_url      = "..."
  auth0_client_id     = "..."
  auth0_client_secret = "..."
}
