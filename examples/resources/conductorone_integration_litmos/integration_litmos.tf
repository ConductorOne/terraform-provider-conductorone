resource "conductorone_integration_litmos" "litmos" {
  app_id = conductorone_app.litmos.id
  user_ids = [
    conductorone_user.admin.id
  ]
  litmos_source     = "..."
  litmos_api_key    = "..."
  litmos_course_ids = ["..."]
}
