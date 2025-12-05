resource "conductorone_integration_zoho_people" "zoho_people" {
  app_id = conductorone_app.zoho_people.id
  user_ids = [
    conductorone_user.admin.id
  ]
  zoho_client_id     = "..."
  zoho_client_secret = "..."
  zoho_refresh_token = "..."
  domain_account     = "..."
}
