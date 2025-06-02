resource "conductorone_catalog" "test_catalog" {
  display_name        = "terraform created catalog"
  description         = "terraform test"
  visible_to_everyone = "true"
  published           = "true"
}     

resource "conductorone_catalog_requestable_entries" "test_entries" {
  catalog_id = conductorone_catalog.test_catalog.id
  app_entitlements = [
    {
      app_id = data.conductorone_app_entitlement.okta_test_integration.app_id
      id     = "287oY0rG4UirjDNFEYguMBvxyim"
    },
    {
      app_id = data.conductorone_app_entitlement.okta_test_integration.app_id
      id     = "2xpUq7cnIRgFyoV752cxj5jBwId"
    },
    {
      app_id = data.conductorone_app_entitlement.okta_test_integration.app_id
      id     = "2xpUqDccUIl7hoci6IEeCT2Aw8O"
    },
  ]
}

data "conductorone_app_entitlement" "okta_test_integration" {
  app_id = "2xVg4U1nyh7SRqjC5FT10fqT5Ju"
  id = "287oY0rG4UirjDNFEYguMBvxyim"
}