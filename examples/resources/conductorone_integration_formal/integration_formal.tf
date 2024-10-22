resource "conductorone_integration_formal" "formal" {
  app_id = conductorone_app.formal.id
  user_ids = [
    conductorone_user.admin.id
  ]
  formal_api_key = "..."
}
