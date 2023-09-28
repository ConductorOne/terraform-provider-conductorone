resource "conductorone_integration_xsoar" "xsoar" {
  app_id = conductorone_app.xsoar.id
  user_ids = [
   conductorone_user.admin.id
  ]
  api_url = "..."
  token = "..."
  }
