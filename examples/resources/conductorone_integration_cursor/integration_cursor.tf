resource "conductorone_integration_cursor" "cursor" {
  app_id = conductorone_app.cursor.id
  user_ids = [
    conductorone_user.admin.id
  ]
  cursor_api_key = "..."
}
