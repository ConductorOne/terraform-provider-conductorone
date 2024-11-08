resource "conductorone_integration_celigo" "celigo" {
  app_id = conductorone_app.celigo.id
  user_ids = [
    conductorone_user.admin.id
  ]
  celigo_region       = "..."
  celigo_access_token = "..."
}
