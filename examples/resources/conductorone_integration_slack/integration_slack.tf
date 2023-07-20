resource "conductorone_integration_slack" "slack" {
  app_id = conductorone_app.slack.id
  user_ids = [
    conductorone_user.admin.id
  ]
  slack_api_key = "..."
}
