resource "conductorone_integration_grafana" "grafana" {
  app_id = "2xVg4U1nyh7SRqjC5FT10fqT5Ju"
  user_ids = [
    "2x9crk1Dj2vo6FdohbBTslEOI4A"
  ]
  grafana_url      = "https://grafana.bsu.d2.ductone.com:2443/"
  grafana_username = "admin"
  grafana_password = "admin"
}

data "conductorone_app_resources" "my_app_resources" {
  app_id = "2xVg4U1nyh7SRqjC5FT10fqT5Ju"
}