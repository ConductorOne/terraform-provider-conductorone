resource "conductorone_integration_tailscale_v2" "tailscale_v2" {
  app_id = conductorone_app.tailscale_v2.id
  user_ids = [
    conductorone_user.admin.id
  ]
  api_key_flow_group = {
    api_key                  = "..."
    tailnet                  = "..."
    ignore_ephemeral_devices = false
  }
  oauth_client_credentials_flow_group = {
    tailscale_client_id      = "..."
    tailscale_client_secret  = "..."
    tailnet                  = "..."
    ignore_ephemeral_devices = false
  }
}
