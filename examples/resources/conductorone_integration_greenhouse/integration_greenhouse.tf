resource "conductorone_integration_greenhouse" "greenhouse" {
  app_id = conductorone_app.greenhouse.id
  user_ids = [
    conductorone_user.admin.id
  ]
  greenhouse_username     = "..."
  greenhouse_on_behalf_of = "..."
}
