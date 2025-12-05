data "conductorone_user" "my_user" {
  departments = [
    "..."
  ]
  email = "...my_email..."
  exclude_ids = [
    "..."
  ]
  exclude_types = [
    "USER_TYPE_HUMAN"
  ]
  ids = [
    "..."
  ]
  job_titles = [
    "..."
  ]
  manager_ids = [
    "..."
  ]
  page_size  = 10
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