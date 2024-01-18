resource "conductorone_integration_segment" "segment" {
  app_id = conductorone_app.segment.id
  user_ids = [
    conductorone_user.admin.id
  ]
  segment_access_token = "..."
}
