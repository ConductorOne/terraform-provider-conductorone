resource "conductorone_integration_github_v2" "github_v2" {
  app_id = conductorone_app.github_v2.id
  user_ids = [
    conductorone_user.admin.id
  ]
  github_access_token = "..."
  github_org_list     = ["..."]
}
