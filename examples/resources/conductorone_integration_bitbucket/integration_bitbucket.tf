resource "conductorone_integration_bitbucket" "bitbucket" {
  app_id = conductorone_app.bitbucket.id
  user_ids = [
    conductorone_user.admin.id
  ]
  bitbucket_username       = "..."
  bitbucket_app_password   = "..."
  bitbucket_workspace_list = ["..."]
}
