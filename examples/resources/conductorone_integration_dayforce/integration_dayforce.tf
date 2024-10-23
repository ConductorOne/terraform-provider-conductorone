resource "conductorone_integration_dayforce" "dayforce" {
  app_id = conductorone_app.dayforce.id
  user_ids = [
    conductorone_user.admin.id
  ]
  dayforce_username         = "..."
  dayforce_password         = "..."
  dayforce_environment      = "..."
  dayforce_url              = "..."
  dayforce_client_namespace = "..."
}
