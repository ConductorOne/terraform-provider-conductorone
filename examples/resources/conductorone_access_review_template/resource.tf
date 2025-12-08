resource "conductorone_access_review_template" "my_access_review_template" {
  access_review_duration = "...my_access_review_duration..."
  description            = "...my_description..."
  display_name           = "...my_display_name..."
  notification_config = {
    send_close     = true
    send_reminders = false
  }
  owner_ids = [
    "..."
  ]
  policy_id  = "...my_policy_id..."
  scope_type = "ACCESS_REVIEW_SCOPE_TYPE_UNSPECIFIED"
}