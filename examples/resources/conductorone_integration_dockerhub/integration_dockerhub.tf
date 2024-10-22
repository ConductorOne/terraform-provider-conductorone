resource "conductorone_integration_dockerhub" "dockerhub" {
  app_id = conductorone_app.dockerhub.id
  user_ids = [
    conductorone_user.admin.id
  ]
  dockerhub_username = "..."
  dockerhub_password = "..."
}
