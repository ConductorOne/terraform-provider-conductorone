resource "conductorone_integration_github_enterprise" "github_enterprise" {
  app_id = conductorone_app.github_enterprise.id
  user_ids = [
    conductorone_user.admin.id
  ]
  github_personal_access_token_group = {
    github_instance_url = "..."
    github_access_token = "..."
    github_org_list     = ["..."]
  }
  github_app_group = {
    github_instance_url    = "..."
    github_app_id          = "..."
    github_app_private_key = "..."
    github_app_org         = "..."
  }
}
