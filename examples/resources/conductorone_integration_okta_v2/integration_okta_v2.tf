resource "conductorone_integration_okta_v2" "okta_v2" {
  app_id = conductorone_app.okta_v2.id
  user_ids = [
    conductorone_user.admin.id
  ]
  okta_v2_domain         = "..."
  okta_v2_api_token      = "..."
  okta_sync_custom_roles = false
}
