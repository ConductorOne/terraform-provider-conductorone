resource "conductorone_integration_ironclad" "ironclad" {
  app_id = conductorone_app.ironclad.id
  user_ids = [
    conductorone_user.admin.id
  ]
}
