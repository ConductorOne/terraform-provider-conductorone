resource "conductorone_integration_freshdesk" "freshdesk" {
  app_id = conductorone_app.freshdesk.id
  user_ids = [
    conductorone_user.admin.id
  ]
  domain  = "..."
  api_key = "..."
}
