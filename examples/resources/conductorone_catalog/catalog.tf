resource "conductorone_catalog" "test_catalog" {
  display_name = "terraform created catalog"
  description  = "terraform test"
  owner_ids = [
    data.conductorone_user.my_user.id
  ]
  visible_to_everyone = "false"
  published           = "true"
}

resource "conductorone_catalog_visibility_bindings" "test_bindings" {
  catalog_id = conductorone_catalog.test_catalog.id
  access_entitlements = [
    {
      app_id = data.conductorone_app_entitlement.okta_everyone.app_id
      id     = data.conductorone_app_entitlement.okta_everyone.id
    }
  ]
}

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
