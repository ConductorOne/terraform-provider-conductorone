data "conductorone_app_entitlement_users" "my_appentitlementusers" {
  app_entitlement_id = "...my_app_entitlement_id..."
  app_id             = "...my_app_id..."
  page_size          = 3
  page_token         = "...my_page_token..."
}