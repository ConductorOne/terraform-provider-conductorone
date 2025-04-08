data "conductorone_app" "my_app" {
  app_ids = [
    "..."
  ]
  display_name = "...my_display_name..."
  exclude_app_ids = [
    "..."
  ]
  only_directories = true
  query            = "...my_query..."
}