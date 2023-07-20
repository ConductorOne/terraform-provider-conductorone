resource "conductorone_integration_jumpcloud" "jumpcloud" {
  app_id            = conductorone_app.jumpcloud.id
  jumpcloud_api_key = "..."
}
