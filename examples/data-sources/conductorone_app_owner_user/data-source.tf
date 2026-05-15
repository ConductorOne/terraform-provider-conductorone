data "conductorone_app_owner_user" "my_app_owner_user" {
  app_id     = "...my_app_id..."
  page_size  = 4
  page_token = "...my_page_token..."
  role_slug  = "...my_role_slug..."
}