data "conductorone_user" "my_user" {
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
  query = "...my_query..."
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