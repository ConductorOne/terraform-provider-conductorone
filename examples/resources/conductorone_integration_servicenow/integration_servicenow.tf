resource "conductorone_integration_servicenow" "servicenow" {
  app_id = conductorone_app.servicenow.id
  user_ids = [
    conductorone_user.admin.id
  ]
  password   = "..."
  username   = "..."
  deployment = "..."
}
