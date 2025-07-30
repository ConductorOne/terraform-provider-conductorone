resource "conductorone_integration_notion" "notion" {
  app_id = conductorone_app.notion.id
  user_ids = [
    conductorone_user.admin.id
  ]
  scim_token = "..."
}
