resource "conductorone_integration_cloudflare" "cloudflare" {
  app_id     = conductorone_app.cloudflare.id
  account_id = "..."
  api_key    = "..."
}
