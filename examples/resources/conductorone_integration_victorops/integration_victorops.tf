resource "conductorone_integration_victorops" "victorops" {
  app_id = conductorone_app.victorops.id
  user_ids = [
    conductorone_user.admin.id
  ]
  victorops_api_id  = "..."
  victorops_api_key = "..."
}
