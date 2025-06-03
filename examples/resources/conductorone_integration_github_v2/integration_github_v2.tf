resource "conductorone_integration_github_v2" "github_v2" {
  app_id = conductorone_app.github_v2.id
  user_ids = [
    conductorone_user.admin.id
  ]
  github_access_token    = "..."
  github_org_list        = ["..."]
  github_app_org         = "..."
  github_app_id          = "..."
  github_app_private_key = "..."
  github_sync_secrets    = false
}
