resource "conductorone_integration_fullstory" "fullstory" {
  app_id = conductorone_app.fullstory.id
  user_ids = [
    conductorone_user.admin.id
  ]
  fullstory_api_key = "..."
}
