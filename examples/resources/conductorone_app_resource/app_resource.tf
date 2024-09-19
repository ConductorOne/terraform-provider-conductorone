resource "conductorone_app_resource" "test_new_app_resource" {
  display_name         = "Custom app resource"
  app_resource_type_id = "<app_resource_type_id>"
  description          = "This is a resource created via Terraform"
  app_id               = data.conductorone_app.test_okta.id
}
