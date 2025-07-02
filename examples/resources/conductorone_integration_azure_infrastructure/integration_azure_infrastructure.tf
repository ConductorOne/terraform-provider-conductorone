resource "conductorone_integration_azure_infrastructure" "azure_infrastructure" {
  app_id = conductorone_app.azure_infrastructure.id
  user_ids = [
    conductorone_user.admin.id
  ]
  azure_client_id              = "..."
  azure_client_secret          = "..."
  azure_tenant_id              = "..."
  mailbox_settings             = false
  skip_ad_groups               = false
  graph_domain                 = "..."
  skip_unused_roles            = false
  skip_sync_storage_containers = false
}
