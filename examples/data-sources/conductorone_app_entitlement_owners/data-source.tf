data "conductorone_app_entitlement_owners" "my_appentitlementowners" {
  app_id         = "...my_app_id..."
  entitlement_id = "...my_entitlement_id..."
  page_size      = 2
  page_token     = "...my_page_token..."
}