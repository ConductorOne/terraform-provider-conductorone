resource "conductorone_integration_opsgenie" "opsgenie" {
  app_id          = conductorone_app.opsgenie.id
  opsgenie_apikey = "..."
}
