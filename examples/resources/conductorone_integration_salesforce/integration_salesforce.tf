resource "conductorone_integration_salesforce" "salesforce" {
  app_id                        = conductorone_app.salesforce.id
  salesforce_instance_url       = "..."
  salesforce_username_for_email = false
}
