resource "conductorone_integration_zendesk" "zendesk" {
  app_id    = conductorone_app.zendesk.id
  email     = "..."
  subdomain = "..."
  api_token = "..."
}
