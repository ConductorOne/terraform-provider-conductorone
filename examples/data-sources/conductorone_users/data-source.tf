data "conductorone_users" "my_users" {
  email = "...my_email..."
  exclude_ids = [
    "..."
  ]
  exclude_types = [
    "USER_TYPE_SYSTEM"
  ]
  ids = [
    "..."
  ]
  page_size  = 5
  page_token = "...my_page_token..."
  query      = "...my_query..."
  refs = [
    {
      id = "...my_id..."
    }
  ]
  role_ids = [
    "..."
  ]
  user_statuses = [
    "DELETED"
  ]
}