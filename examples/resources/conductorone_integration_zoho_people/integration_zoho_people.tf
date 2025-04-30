resource "conductorone_integration_zoho_people" "zoho_people" {
  app_id = conductorone_app.zoho_people.id
  user_ids = [
    conductorone_user.admin.id
  ]
  client_id      = "..."
  client_secret  = "..."
  refresh_token  = "..."
  account_domain = "..."
}
