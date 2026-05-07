data "conductorone_app_entitlement_owner_user" "my_app_entitlement_owner_user" {
  app_id         = "...my_app_id..."
  entitlement_id = "...my_entitlement_id..."
  page_size      = 7
  page_token     = "...my_page_token..."
  role_slug      = "...my_role_slug..."
}