resource "conductorone_integration_concur" "concur" {
  app_id = conductorone_app.concur.id
  user_ids = [
    conductorone_user.admin.id
  ]
  concur_base_url      = "..."
  concur_client_id     = "..."
  concur_client_secret = "..."
  concur_refresh_token = "..."
}
