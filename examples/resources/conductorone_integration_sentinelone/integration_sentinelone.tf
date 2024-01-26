resource "conductorone_integration_sentinelone" "sentinelone" {
  app_id = conductorone_app.sentinelone.id
  user_ids = [
    conductorone_user.admin.id
  ]
  sentinelone_base_url = "..."
  sentinelone_token    = "..."
}
