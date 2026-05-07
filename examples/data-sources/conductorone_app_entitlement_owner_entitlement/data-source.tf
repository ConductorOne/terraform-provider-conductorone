data "conductorone_app_entitlement_owner_entitlement" "my_app_entitlement_owner_entitlement" {
  app_id         = "...my_app_id..."
  entitlement_id = "...my_entitlement_id..."
  page_size      = 3
  page_token     = "...my_page_token..."
  role_slug      = "...my_role_slug..."
}