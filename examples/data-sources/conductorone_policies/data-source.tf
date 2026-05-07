data "conductorone_policies" "my_policies" {
  display_name = "...my_display_name..."
  exclude_policy_ids = [
    "..."
  ]
  include_deleted = false
  page_size       = 8
  page_token      = "...my_page_token..."
  policy_types = [
    "POLICY_TYPE_CERTIFY"
  ]
  query = "...my_query..."
  refs = [
    {
      id = "...my_id..."
    }
  ]
}