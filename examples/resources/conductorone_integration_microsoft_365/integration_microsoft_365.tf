resource "conductorone_integration_microsoft_365" "microsoft_365" {
  app_id = conductorone_app.microsoft_365.id
  user_ids = [
    conductorone_user.admin.id
  ]
  ms365_tenant_id     = "..."
  ms365_client_id     = "..."
  ms365_client_secret = "..."
}
