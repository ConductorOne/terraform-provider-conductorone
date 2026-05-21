resource "conductorone_bundle_automation" "my_bundleautomation" {
  bundle_automation_rule_cel = {
    expression = "...my_expression..."
  }
  bundle_automation_rule_entitlement = {
    entitlement_refs = [
      {
        app_id = "...my_app_id..."
        id     = "...my_id..."
      }
    ]
  }
  create_tasks = false
  delete_bundle_automation_request = {
    # ...
  }
  disable_circuit_breaker           = false
  enabled                           = true
  enforce_on_small_profiles         = false
  removed_members_threshold_percent = "...my_removed_members_threshold_percent..."
  request_catalog_id                = "...my_request_catalog_id..."
}