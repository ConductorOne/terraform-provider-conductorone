resource "conductorone_integration_temporal_cloud" "temporal_cloud" {
  app_id = conductorone_app.temporal_cloud.id
  user_ids = [
    conductorone_user.admin.id
  ]
  temporal_cloud_api_key = "..."
}
