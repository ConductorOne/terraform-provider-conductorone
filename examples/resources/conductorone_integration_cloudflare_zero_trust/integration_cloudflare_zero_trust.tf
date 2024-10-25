resource "conductorone_integration_cloudflare_zero_trust" "cloudflare_zero_trust" {
  app_id = conductorone_app.cloudflare_zero_trust.id
  user_ids = [
    conductorone_user.admin.id
  ]
  account_id = "..."
  api_key    = "..."
}
