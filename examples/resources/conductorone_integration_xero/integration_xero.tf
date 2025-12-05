resource "conductorone_integration_xero" "xero" {
  app_id = conductorone_app.xero.id
  user_ids = [
    conductorone_user.admin.id
  ]
  xero_client_id     = "..."
  xero_client_secret = "..."
}
