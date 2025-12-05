resource "conductorone_app_entitlement_automation" "my_app_entitlement_automation" {
  app_entitlement_automation_last_run_status = {
    # ...
  }
  app_entitlement_automation_rule_basic = {
    expression = "...my_expression..."
  }
  app_entitlement_automation_rule_cel = {
    expression = "...my_expression..."
  }
  app_entitlement_automation_rule_entitlement = {
    entitlement_refs = [
      {
        app_id = "...my_app_id..."
        id     = "...my_id..."
      }
    ]
  }
  app_entitlement_automation_rule_none = {
    # ...
  }
  app_entitlement_id = "...my_app_entitlement_id..."
  app_id             = "...my_app_id..."
  description        = "...my_description..."
  display_name       = "...my_display_name..."
}