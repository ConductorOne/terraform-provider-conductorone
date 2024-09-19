resource "conductorone_app_resource_type" "test_new_app_resource_type" {
  display_name  = "Group test"
  app_id        = data.conductorone_app.test_okta.id
  resource_type = "GROUP"
}
