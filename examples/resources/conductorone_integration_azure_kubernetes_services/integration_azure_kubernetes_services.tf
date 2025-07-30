resource "conductorone_integration_azure_kubernetes_services" "azure_kubernetes_services" {
  app_id = conductorone_app.azure_kubernetes_services.id
  user_ids = [
    conductorone_user.admin.id
  ]
  subscription_id     = "..."
  resource_group_name = "..."
  cluster_name        = "..."
  tenant_id           = "..."
  sp_client_id        = "..."
  sp_client_secret    = "..."
}
