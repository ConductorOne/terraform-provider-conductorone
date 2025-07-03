data "conductorone_app" "my_app" {
  app_ids = [
    "..."
  ]
  display_name = "...my_display_name..."
  exclude_app_ids = [
    "..."
  ]
  only_directories = true
  page_size        = 6
  page_token       = "...my_page_token..."
  query            = "...my_query..."
}