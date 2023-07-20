resource "conductorone_integration_github" "github" {
  app_id = conductorone_app.github.id
  user_ids = [
    conductorone_user.admin.id
  ]
  github_org          = "..."
  github_access_token = "..."
}
