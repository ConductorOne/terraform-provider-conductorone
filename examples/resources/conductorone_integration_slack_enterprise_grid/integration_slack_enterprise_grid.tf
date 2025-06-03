resource "conductorone_integration_slack_enterprise_grid" "slack_enterprise_grid" {
  app_id = conductorone_app.slack_enterprise_grid.id
  user_ids = [
    conductorone_user.admin.id
  ]
  slack_api_key            = "..."
  slack_api_enterprise_key = "..."
  use_gov_env              = false
}
