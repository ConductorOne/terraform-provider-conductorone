resource "conductorone_integration_grafana" "grafana" {
  app_id = conductorone_app.grafana.id
  user_ids = [
    conductorone_user.admin.id
  ]
  grafana_url      = "..."
  grafana_username = "..."
  grafana_password = "..."
}
