resource "conductorone_integration_okta" "okta" {
  app_id = conductorone_app.okta.id
  user_ids = [
    conductorone_user.admin.id
  ]
  okta_domain                   = "..."
  okta_api_key                  = "..."
  okta_dont_sync_inactive_apps  = false
  okta_extract_aws_saml_roles   = false
  okta_sync_deprovisioned_users = false
}
