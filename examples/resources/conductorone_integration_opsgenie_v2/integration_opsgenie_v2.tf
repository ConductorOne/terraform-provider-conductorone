resource "conductorone_integration_opsgenie_v2" "opsgenie_v2" {
  app_id = conductorone_app.opsgenie_v2.id
  user_ids = [
    conductorone_user.admin.id
  ]
  opsgenie_apikey = "..."
}
