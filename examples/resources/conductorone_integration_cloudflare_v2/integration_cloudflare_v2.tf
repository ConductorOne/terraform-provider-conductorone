resource "conductorone_integration_cloudflare_v2" "cloudflare_v2" {
  app_id = conductorone_app.cloudflare_v2.id
  user_ids = [
    conductorone_user.admin.id
  ]
  group_account_id = {
    account_id_v2 = "..."
    api_token     = "..."
  }
  group_email = {
    account_id_v2 = "..."
    email         = "..."
    api_key_v2    = "..."
  }
}
