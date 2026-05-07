data "conductorone_app_resource_owners" "my_appresourceowners" {
  app_id           = "...my_app_id..."
  page_size        = 4
  page_token       = "...my_page_token..."
  resource_id      = "...my_resource_id..."
  resource_type_id = "...my_resource_type_id..."
}