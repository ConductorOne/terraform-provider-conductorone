resource "conductorone_integration_atlassian" "atlassian" {
  app_id = conductorone_app.atlassian.id
  user_ids = [
    conductorone_user.admin.id
  ]
  access_token    = "..."
  organization_id = "..."
}
