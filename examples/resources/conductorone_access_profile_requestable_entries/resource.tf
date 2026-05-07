resource "conductorone_access_profile_requestable_entries" "my_access_profile_requestable_entries" {
  app_entitlements = [
    {
      app_id = "...my_app_id..."
      id     = "...my_id..."
    }
  ]
  catalog_id      = "...my_catalog_id..."
  create_requests = true
}