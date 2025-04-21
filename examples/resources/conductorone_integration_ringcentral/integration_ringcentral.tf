resource "conductorone_integration_ringcentral" "ringcentral" {
  app_id = conductorone_app.ringcentral.id
  user_ids = [
    conductorone_user.admin.id
  ]
  ringcentral_client_id     = "..."
  ringcentral_client_secret = "..."
  ringcentral_jwt           = "..."
}
