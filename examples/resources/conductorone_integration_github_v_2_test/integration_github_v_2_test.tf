resource "conductorone_integration_github_v_2_test" "github_v_2_test" {
  app_id = conductorone_app.github_v_2_test.id
  user_ids = [
    conductorone_user.admin.id
  ]
  github_access_token = "..."
  github_org_list     = ["..."]
}
