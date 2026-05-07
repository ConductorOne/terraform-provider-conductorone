data "conductorone_app_owner_entitlement" "my_app_owner_entitlement" {
  app_id     = "...my_app_id..."
  page_size  = 9
  page_token = "...my_page_token..."
  role_slug  = "...my_role_slug..."
}