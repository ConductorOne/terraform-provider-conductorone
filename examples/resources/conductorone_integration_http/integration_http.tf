resource "conductorone_integration_http" "http" {
  app_id = conductorone_app.http.id
  user_ids = [
    conductorone_user.admin.id
  ]
  http_connector_map_values  = "..."
  http_connector_config_file = "..."
}
