resource "conductorone_integration_ping_federate" "ping_federate" {
  app_id = conductorone_app.ping_federate.id
  user_ids = [
    conductorone_user.admin.id
  ]
  pingfed_instance_url = "..."
  pingfed_username     = "..."
  pingfed_password     = "..."
}
