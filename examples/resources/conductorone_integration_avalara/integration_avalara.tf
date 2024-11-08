resource "conductorone_integration_avalara" "avalara" {
  app_id = conductorone_app.avalara.id
  user_ids = [
    conductorone_user.admin.id
  ]
  avalara_username    = "..."
  avalara_password    = "..."
  avalara_environment = "..."
}
