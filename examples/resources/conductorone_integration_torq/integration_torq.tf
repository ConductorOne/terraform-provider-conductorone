resource "conductorone_integration_torq" "torq" {
  app_id = conductorone_app.torq.id
  user_ids = [
    conductorone_user.admin.id
  ]
  torq_client_id     = "..."
  torq_client_secret = "..."
}
