resource "conductorone_integration_vultr" "vultr" {
  app_id = conductorone_app.vultr.id
  user_ids = [
    conductorone_user.admin.id
  ]
  vultr_api_key = "..."
}
