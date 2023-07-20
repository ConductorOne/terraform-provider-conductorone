resource "conductorone_integration_google_identity_platform" "google_identity_platform" {
  app_id           = conductorone_app.google_identity_platform.id
  project_id       = "..."
  tenant_id        = "..."
  credentials_json = "..."
}
