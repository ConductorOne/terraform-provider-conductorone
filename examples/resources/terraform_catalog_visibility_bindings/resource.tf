resource "terraform_catalog_visibility_bindings" "my_catalog_visibility_bindings" {
  access_entitlements = [
    {
      app_id = "...my_app_id..."
      id     = "baa72116-9521-427e-abe9-e4e40c31d533"
    },
  ]
  catalog_id = "...my_catalog_id..."
}