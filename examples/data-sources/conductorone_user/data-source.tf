data "conductorone_user" "my_user" {
  delegate_status = "DELEGATE_STATUS_UNSPECIFIED"
  delegated_user_ids = [
    "..."
  ]
  departments = [
    "..."
  ]
  email = "...my_email..."
  exclude_ids = [
    "..."
  ]
  exclude_origins = [
    "USER_ORIGIN_SYSTEM"
  ]
  exclude_types = [
    "USER_TYPE_HUMAN"
  ]
  ids = [
    "..."
  ]
  is_delegate = false
  job_titles = [
    "..."
  ]
  manager_ids = [
    "..."
  ]
  origins = [
    "USER_ORIGIN_LOCAL"
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