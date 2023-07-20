resource "conductorone_integration_slack" "slack" {
  app_id        = conductorone_app.slack.id
  slack_api_key = "..."
}
