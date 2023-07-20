resource "conductorone_integration_github_enterprise" "github_enterprise" {
  app_id              = conductorone_app.github_enterprise.id
  github_instance_url = "..."
  github_access_token = "..."
}
