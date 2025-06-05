resource "conductorone_integration_pandadoc" "pandadoc" {
  app_id = conductorone_app.pandadoc.id
  user_ids = [
    conductorone_user.admin.id
  ]
  api_key       = "..."
  europe_domain = false
}
