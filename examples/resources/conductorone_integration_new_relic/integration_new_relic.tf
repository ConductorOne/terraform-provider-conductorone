resource "conductorone_integration_new_relic" "new_relic" {
  app_id = conductorone_app.new_relic.id
  user_ids = [
    conductorone_user.admin.id
  ]
  newrelic_api_key = "..."
}
