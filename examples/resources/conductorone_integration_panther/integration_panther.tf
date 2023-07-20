resource "conductorone_integration_panther" "panther" {
  app_id          = conductorone_app.panther.id
  panther_api_key = "..."
  panther_url     = "..."
}
