resource "conductorone_integration_hubspot" "hubspot" {
  app_id = conductorone_app.hubspot.id
  user_ids = [
    conductorone_user.admin.id
  ]
  hubspot_token = "..."
}
