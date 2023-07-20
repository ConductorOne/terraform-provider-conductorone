resource "conductorone_integration_tailscale" "tailscale" {
  app_id = conductorone_app.tailscale.id
  user_ids = [
    conductorone_user.admin.id
  ]
  tailscale_api_key = "..."
  tailnet           = "..."
}
