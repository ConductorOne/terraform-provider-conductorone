resource "conductorone_bundle_automation" "my_bundleautomation" {
  create_tasks = false
  delete_bundle_automation_request = {
    # ...
  }
  disable_circuit_breaker = false
  enabled                 = true
  entitlements = {
    entitlement_refs = [
      {
        app_id = "...my_app_id..."
        id     = "...my_id..."
      }
    ]
  }
  request_catalog_id = "...my_request_catalog_id..."
}