resource "conductorone_integration_oracle_idcs" "oracle_idcs" {
  app_id = conductorone_app.oracle_idcs.id
  user_ids = [
    conductorone_user.admin.id
  ]
  api_base_domain   = "..."
  api_access_id     = "..."
  api_access_secret = "..."
}
