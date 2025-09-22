resource "conductorone_app_entitlement_automation" "my_app_entitlement_automation" {
  app_entitlement_id = "...my_app_entitlement_id..."
  app_id             = "...my_app_id..."
  basic = {
    expression = "...my_expression..."
  }
  cel = {
    expression = "...my_expression..."
  }
  created_at   = "2022-11-26T09:39:40.100Z"
  deleted_at   = "2022-05-02T02:32:37.165Z"
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
    last_completed_at = "2022-09-24T04:12:21.151Z"
  }
  none = {
    # ...
  }
  updated_at = "2020-03-04T17:09:55.729Z"
}