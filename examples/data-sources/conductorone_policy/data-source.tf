data "conductorone_policy" "my_policy" {
  display_name = "...my_display_name..."
  exclude_policy_ids = [
    "..."
  ]
  include_deleted = false
  page_size       = 7
  page_token      = "...my_page_token..."
  policy_types = [
    "POLICY_TYPE_ACCESS_REQUEST"
  ]
  query = "...my_query..."
  refs = [
    {
      id = "...my_id..."
    }
  ]
}