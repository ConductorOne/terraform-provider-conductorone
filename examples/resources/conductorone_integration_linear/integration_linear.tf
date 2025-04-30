resource "conductorone_integration_linear" "linear" {
  app_id = conductorone_app.linear.id
  user_ids = [
    conductorone_user.admin.id
  ]
  linear_api_key                      = "..."
  enable_external_ticket_provisioning = false
}
