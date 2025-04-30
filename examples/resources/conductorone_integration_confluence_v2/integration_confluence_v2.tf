resource "conductorone_integration_confluence_v2" "confluence_v2" {
  app_id = conductorone_app.confluence_v2.id
  user_ids = [
    conductorone_user.admin.id
  ]
  domain_url           = "..."
  username             = "..."
  api_key              = "..."
  skip_personal_spaces = false
}
