resource "conductorone_access_conflict" "my_access_conflict" {
  access_conflict_notification_config = {
    email_notifications = {
      enabled = false
      identity_user_ids = [
        "..."
      ]
    }
    slack_notifications = {
      channel_id   = "...my_channel_id..."
      channel_name = "...my_channel_name..."
      enabled      = false
    }
  }
  conflict_monitor_delete_request = {
    # ...
  }
  description  = "...my_description..."
  display_name = "...my_display_name..."
}