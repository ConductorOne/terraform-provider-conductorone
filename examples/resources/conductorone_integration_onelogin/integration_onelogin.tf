resource "conductorone_integration_onelogin" "onelogin" {
  app_id = conductorone_app.onelogin.id
  user_ids = [
   conductorone_user.admin.id
  ]
  onelogin_domain = "..."
  oauth_client_cred_grant_client_id = "..."
  oauth_client_cred_grant_client_secret = "..."
  }
