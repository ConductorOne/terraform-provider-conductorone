resource "conductorone_integration_cloudflare_zero_trust_v2" "cloudflare_zero_trust_v2" {
  app_id = conductorone_app.cloudflare_zero_trust_v2.id
  user_ids = [
    conductorone_user.admin.id
  ]
  account_id = "..."
  api_token  = "..."
  api_key    = "..."
  email      = "..."
}
