resource "conductorone_integration_onelogin_v2" "onelogin_v2" {
  app_id = conductorone_app.onelogin_v2.id
  user_ids = [
    conductorone_user.admin.id
  ]
  onelogin_domain                       = "..."
  oauth_client_cred_grant_client_id     = "..."
  oauth_client_cred_grant_client_secret = "..."
}
