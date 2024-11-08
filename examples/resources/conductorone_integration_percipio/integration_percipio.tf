resource "conductorone_integration_percipio" "percipio" {
  app_id = conductorone_app.percipio.id
  user_ids = [
    conductorone_user.admin.id
  ]
  percipio_organization_id = "..."
  percipio_api_token       = "..."
  percipio_course_ids      = ["..."]
}
