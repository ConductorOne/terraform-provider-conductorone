resource "conductorone_catalog_visibility_bindings" "my_catalog_visibility_bindings" {
  access_entitlements = [
    {
      app_id = "...my_app_id..."
      id     = "4358dcf8-4177-493f-8ad8-4b71e7e0e2b9"
    },
  ]
  catalog_id = "...my_catalog_id..."
}