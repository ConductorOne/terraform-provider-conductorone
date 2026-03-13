data "conductorone_app_owners" "my_appowners" {
  app_id     = "...my_app_id..."
  page_size  = 5
  page_token = "...my_page_token..."
}