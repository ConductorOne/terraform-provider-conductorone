resource "conductorone_access_profile_requestable_entry" "my_access_profile_requestable_entry" {
  app_id          = "...my_app_id..."
  catalog_id      = "...my_catalog_id..."
  create_requests = false
  entitlement_id  = "...my_entitlement_id..."
}