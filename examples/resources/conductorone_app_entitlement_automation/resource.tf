resource "conductorone_app_entitlement_automation" "my_app_entitlement_automation" {
  app_entitlement_id = "...my_app_entitlement_id..."
  app_id             = "...my_app_id..."
  basic = {
    expression = "...my_expression..."
  }
  cel = {
    expression = "...my_expression..."
  }
  description  = "...my_description..."
  display_name = "...my_display_name..."
  entitlements = {
    entitlement_refs = [
      {
        app_id = "...my_app_id..."
        id     = "...my_id..."
      }
    ]
  }
  last_run_status = {
    # ...
  }
  none = {
    # ...
  }
}