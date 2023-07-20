resource "conductorone_integration_expensify" "expensify" {
  app_id                = conductorone_app.expensify.id
  expensify_user_id     = "..."
  expensify_user_secret = "..."
}
