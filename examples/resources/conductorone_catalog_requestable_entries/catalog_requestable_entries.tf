resource "conductorone_catalog_requestable_entries" "test_entries" {
  catalog_id = conductorone_catalog.test_catalog.id
  app_entitlements = [
    {
      app_id = data.conductorone_app_entitlement.okta_test_admin.app_id
      id     = data.conductorone_app_entitlement.okta_test_admin.id
    },
    {
      app_id = data.conductorone_app_entitlement.okta_administrators.app_id
      id     = data.conductorone_app_entitlement.okta_administrators.id
    }
  ]
}
