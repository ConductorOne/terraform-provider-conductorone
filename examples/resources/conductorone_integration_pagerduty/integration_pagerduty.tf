resource "conductorone_integration_pagerduty" "pagerduty" {
  app_id = conductorone_app.pagerduty.id
  user_ids = [
   conductorone_user.admin.id
  ]
  pagerduty_api_token = "..."
  }
