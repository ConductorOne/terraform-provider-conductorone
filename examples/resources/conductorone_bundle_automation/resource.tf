resource "conductorone_bundle_automation" "my_bundleautomation" {
  cel = {
    expression = "...my_expression..."
  }
  create_tasks = false
  delete_bundle_automation_request = {
    # ...
  }
  disable_circuit_breaker   = false
  enabled                   = true
  enforce_on_small_profiles = false
  entitlements = {
    entitlement_refs = [
      {
        app_id = "...my_app_id..."
        id     = "...my_id..."
      }
    ]
  }
  removed_members_threshold_percent = "...my_removed_members_threshold_percent..."
  request_catalog_id                = "...my_request_catalog_id..."
}