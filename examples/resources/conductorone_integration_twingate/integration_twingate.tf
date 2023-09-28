resource "conductorone_integration_twingate" "twingate" {
  app_id = conductorone_app.twingate.id
  user_ids = [
   conductorone_user.admin.id
  ]
  twingate_apikey = "..."
  twingate_domain = "..."
  }
