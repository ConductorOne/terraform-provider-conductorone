resource "terraform_catalog_requestable_entries" "my_catalog_requestable_entries" {
  app_entitlements = [
    {
      app_id = "...my_app_id..."
      id     = "b7f1b39c-f70f-405c-89f5-ebc1cc76007f"
    },
  ]
  catalog_id = "...my_catalog_id..."
}