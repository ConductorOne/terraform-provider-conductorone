resource "conductorone_integration_calendly" "calendly" {
  app_id = conductorone_app.calendly.id
  user_ids = [
    conductorone_user.admin.id
  ]
  calendly_token = "..."
}
