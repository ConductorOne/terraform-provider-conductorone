data "conductorone_app_resource_owner_entitlement" "my_app_resource_owner_entitlement" {
  app_id           = "...my_app_id..."
  page_size        = 9
  page_token       = "...my_page_token..."
  resource_id      = "...my_resource_id..."
  resource_type_id = "...my_resource_type_id..."
  role_slug        = "...my_role_slug..."
}