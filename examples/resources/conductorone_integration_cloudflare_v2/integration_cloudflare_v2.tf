resource "conductorone_integration_cloudflare_v2" "cloudflare_v2" {
  app_id = conductorone_app.cloudflare_v2.id
  user_ids = [
    conductorone_user.admin.id
  ]
  account_id_v2 = "..."
  api_token     = "..."
  email         = "..."
  api_key_v2    = "..."
}
