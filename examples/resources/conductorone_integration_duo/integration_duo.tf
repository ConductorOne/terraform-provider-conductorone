resource "conductorone_integration_duo" "duo" {
  app_id              = conductorone_app.duo.id
  duo_integration_key = "..."
  duo_secret_key      = "..."
  duo_api_hostname    = "..."
}
