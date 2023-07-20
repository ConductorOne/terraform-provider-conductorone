resource "conductorone_integration_azure_ad" "azure_ad" {
  app_id = conductorone_app.azure_ad.id
}
