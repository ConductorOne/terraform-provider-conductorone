resource "conductorone_integration_confluence" "confluence" {
  app_id              = conductorone_app.confluence.id
  confluence_domain   = "..."
  confluence_username = "..."
  confluence_apikey   = "..."
}
