resource "conductorone_catalog_requestable_entries" "my_catalog_requestable_entries" {
  app_entitlements = [
    {
      app_id = "...my_app_id..."
      id     = "fa63f71c-8a50-4788-b9b0-353705617c66"
    },
  ]
  catalog_id = "...my_catalog_id..."
}