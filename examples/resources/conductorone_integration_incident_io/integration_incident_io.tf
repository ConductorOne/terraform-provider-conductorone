resource "conductorone_integration_incident_io" "incident_io" {
  app_id = conductorone_app.incident_io.id
  user_ids = [
    conductorone_user.admin.id
  ]
  incident_io_token = "..."
}
