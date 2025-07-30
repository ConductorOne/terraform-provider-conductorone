resource "conductorone_integration_zuper" "zuper" {
  app_id = conductorone_app.zuper.id
  user_ids = [
    conductorone_user.admin.id
  ]
  api_url = "..."
  api_key = "..."
}
