resource "conductorone_integration_redis_enterprise" "redis_enterprise" {
  app_id = conductorone_app.redis_enterprise.id
  user_ids = [
    conductorone_user.admin.id
  ]
  cluster_host = "..."
  api_port     = "..."
  username     = "..."
  password     = "..."
}
