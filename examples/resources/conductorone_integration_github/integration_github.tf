resource "conductorone_integration_github" "github" {
  app_id              = conductorone_app.github.id
  github_org          = "..."
  github_access_token = "..."
}
