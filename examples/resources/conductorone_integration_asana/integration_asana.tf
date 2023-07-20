resource "conductorone_integration_asana" "asana" {
  app_id        = conductorone_app.asana.id
  asana_api_key = "..."
}
