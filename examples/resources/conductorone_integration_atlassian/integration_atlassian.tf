resource "conductorone_integration_atlassian" "atlassian" {
  app_id = conductorone_app.atlassian.id
  user_ids = [
    conductorone_user.admin.id
  ]
  atlassian_user_email      = "..."
  atlassian_api_token       = "..."
  atlassian_organization_id = "..."
  atlassian_site_id         = "..."
}
