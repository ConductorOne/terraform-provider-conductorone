resource "conductorone_integration_expensify" "expensify" {
  app_id = conductorone_app.expensify.id
  user_ids = [
   conductorone_user.admin.id
  ]
  expensify_user_id = "..."
  expensify_user_secret = "..."
  }
