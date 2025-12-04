resource "conductorone_access_review" "my_access_review" {
  access_review_expand_mask = {
    paths = [
      "..."
    ]
  }
  access_review_service_delete_request = {
    # ...
  }
  completion_date = "2020-03-30T12:47:51.525Z"
  description     = "...my_description..."
  display_name    = "...my_display_name..."
  duplicate_from  = "...my_duplicate_from..."
  notification_config = {
    send_close     = true
    send_reminders = true
  }
  owner_ids = [
    "..."
  ]
  policy_id  = "...my_policy_id..."
  scope_type = "ACCESS_REVIEW_SCOPE_TYPE_UNSPECIFIED"
}