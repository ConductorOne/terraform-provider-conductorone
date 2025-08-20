data "conductorone_app_entitlement_monitor_binding" "my_app_entitlement_monitor_binding" {
  app_entitlement_id = "...my_app_entitlement_id..."
  app_id             = "...my_app_id..."
  entitlement_group  = "ENTITLEMENT_GROUP_B"
  monitor_id         = "...my_monitor_id..."
}