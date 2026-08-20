data "conductorone_app_resource_owner_user" "my_app_resource_owner_user" {
  app_id           = "...my_app_id..."
  page_size        = 0
  page_token       = "...my_page_token..."
  resource_id      = "...my_resource_id..."
  resource_type_id = "...my_resource_type_id..."
  role_slug        = "...my_role_slug..."
}