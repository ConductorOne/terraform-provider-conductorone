resource "conductorone_integration_microsoft_entra" "microsoft_entra" {
  app_id = conductorone_app.microsoft_entra.id
  user_ids = [
    conductorone_user.admin.id
  ]
  entra_tenant_id      = "..."
  entra_client_id      = "..."
  entra_client_secret  = "..."
  entra_skip_ad_groups = false
}
