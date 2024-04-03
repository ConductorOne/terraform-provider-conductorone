resource "terraform_catalog_requestable_entries" "my_catalog_requestable_entries" {
  app_entitlements = [
    {
      app_id = "...my_app_id..."
      id     = "7f1b39cf-70f0-45c8-9f5e-bc1cc76007fe"
    },
  ]
  catalog_id = "...my_catalog_id..."
}