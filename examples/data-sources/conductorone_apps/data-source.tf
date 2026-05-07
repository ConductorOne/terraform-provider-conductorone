data "conductorone_apps" "my_apps" {
  app_ids = [
    "..."
  ]
  display_name = "...my_display_name..."
  exclude_app_ids = [
    "..."
  ]
  only_directories = false
  page_size        = 5
  page_token       = "...my_page_token..."
  policy_refs = [
    {
      id = "...my_id..."
    }
  ]
  query = "...my_query..."
}