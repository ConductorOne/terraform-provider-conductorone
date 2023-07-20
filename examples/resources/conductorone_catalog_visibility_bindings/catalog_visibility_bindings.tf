resource "conductorone_catalog_visibility_bindings" "test_bindings" {
  catalog_id = conductorone_catalog.test_catalog.id
  access_entitlements = [
    {
      app_id = data.conductorone_app_entitlement.okta_everyone.app_id
      id     = data.conductorone_app_entitlement.okta_everyone.id
    }
  ]
}
