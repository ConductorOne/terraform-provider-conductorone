resource "terraform_catalog_visibility_bindings" "my_catalog_visibility_bindings" {
  access_entitlements = [
    {
      app_id = "...my_app_id..."
      id     = "ebaa7211-6952-4127-aebe-9e4e40c31d53"
    },
  ]
  catalog_id = "...my_catalog_id..."
}