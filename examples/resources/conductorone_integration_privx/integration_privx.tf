resource "conductorone_integration_privx" "privx" {
  app_id = conductorone_app.privx.id
  user_ids = [
    conductorone_user.admin.id
  ]
  privx_base_url            = "..."
  privx_client_id           = "..."
  privx_client_secret       = "..."
  privx_oauth_client_id     = "..."
  privx_oauth_client_secret = "..."
}
