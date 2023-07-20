resource "conductorone_integration_buildkite" "buildkite" {
  app_id       = conductorone_app.buildkite.id
  api_token    = "..."
  organization = "..."
}
