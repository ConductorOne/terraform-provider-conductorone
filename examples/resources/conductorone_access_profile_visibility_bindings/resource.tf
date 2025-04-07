resource "conductorone_access_profile_visibility_bindings" "my_access_profile_visibility_bindings" {
  access_entitlements = [
    {
      app_id = "...my_app_id..."
      id     = "...my_id..."
    }
  ]
  catalog_id = "...my_catalog_id..."
}