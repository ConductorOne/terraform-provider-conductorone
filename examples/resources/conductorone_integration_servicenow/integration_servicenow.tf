resource "conductorone_integration_servicenow" "servicenow" {
  app_id = conductorone_app.servicenow.id
  user_ids = [
    conductorone_user.admin.id
  ]
  deployment                          = "..."
  username                            = "..."
  password                            = "..."
  enable_external_ticket_provisioning = false
  catalog_id                          = "..."
  category_id                         = "..."
}
