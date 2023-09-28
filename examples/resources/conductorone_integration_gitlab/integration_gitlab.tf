resource "conductorone_integration_gitlab" "gitlab" {
  app_id = conductorone_app.gitlab.id
  user_ids = [
   conductorone_user.admin.id
  ]
  gitlab_group = "..."
  gitlab_access_token = "..."
  }
