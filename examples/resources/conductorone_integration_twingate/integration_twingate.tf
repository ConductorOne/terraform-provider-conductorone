resource "conductorone_integration_twingate" "twingate" {
  app_id          = conductorone_app.twingate.id
  twingate_apikey = "..."
  twingate_domain = "..."
}
