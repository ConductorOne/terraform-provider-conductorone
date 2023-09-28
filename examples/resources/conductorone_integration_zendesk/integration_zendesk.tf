resource "conductorone_integration_zendesk" "zendesk" {
  app_id = conductorone_app.zendesk.id
  user_ids = [
    conductorone_user.admin.id
  ]
  email     = "..."
  subdomain = "..."
  api_token = "..."
}
