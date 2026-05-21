data "conductorone_connector_owner_user" "my_connector_owner_user" {
  app_id       = "...my_app_id..."
  connector_id = "...my_connector_id..."
  page_size    = 10
  page_token   = "...my_page_token..."
  role_slug    = "...my_role_slug..."
}