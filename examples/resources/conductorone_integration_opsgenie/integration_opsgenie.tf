resource "conductorone_integration_opsgenie" "opsgenie" {
  app_id = conductorone_app.opsgenie.id
  user_ids = [
    conductorone_user.admin.id
  ]
  opsgenie_apikey = "..."
}
