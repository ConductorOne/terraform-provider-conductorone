resource "conductorone_integration_verkada" "verkada" {
  app_id = conductorone_app.verkada.id
  user_ids = [
    conductorone_user.admin.id
  ]
  verkada_api_key = "..."
  verkada_region  = "..."
}
