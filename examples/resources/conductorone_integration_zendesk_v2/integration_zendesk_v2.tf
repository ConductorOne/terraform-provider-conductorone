resource "conductorone_integration_zendesk_v2" "zendesk_v2" {
  app_id = conductorone_app.zendesk_v2.id
  user_ids = [
    conductorone_user.admin.id
  ]
  zendesk_v2_subdomain = "..."
  zendesk_v2_email     = "..."
  zendesk_v2_api_token = "..."
  orgs                 = "..."
}
