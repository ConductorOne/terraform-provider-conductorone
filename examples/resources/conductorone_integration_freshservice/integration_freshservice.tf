resource "conductorone_integration_freshservice" "freshservice" {
  app_id = conductorone_app.freshservice.id
  user_ids = [
    conductorone_user.admin.id
  ]
  domain                              = "..."
  api_key                             = "..."
  enable_external_ticket_provisioning = false
  category_id                         = "..."
}
