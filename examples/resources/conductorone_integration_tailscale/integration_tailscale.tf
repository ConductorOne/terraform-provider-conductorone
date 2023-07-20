resource "conductorone_integration_tailscale" "tailscale" {
  app_id            = conductorone_app.tailscale.id
  tailscale_api_key = "..."
  tailnet           = "..."
}
