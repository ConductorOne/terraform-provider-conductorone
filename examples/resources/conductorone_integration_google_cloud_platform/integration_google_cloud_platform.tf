resource "conductorone_integration_google_cloud_platform" "google_cloud_platform" {
  app_id           = conductorone_app.google_cloud_platform.id
  credentials_json = "..."
}
