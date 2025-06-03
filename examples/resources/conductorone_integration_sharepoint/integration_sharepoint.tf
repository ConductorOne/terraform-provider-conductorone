resource "conductorone_integration_sharepoint" "sharepoint" {
  app_id = conductorone_app.sharepoint.id
  user_ids = [
    conductorone_user.admin.id
  ]
  azure_client_id          = "..."
  azure_client_secret      = "..."
  azure_tenant_id          = "..."
  azure_graph_domain       = "..."
  pfx_certificate          = "..."
  pfx_certificate_password = "..."
  sharepoint_domain        = "..."
}
