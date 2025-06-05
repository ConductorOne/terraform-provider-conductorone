resource "conductorone_integration_contentful" "contentful" {
  app_id = conductorone_app.contentful.id
  user_ids = [
    conductorone_user.admin.id
  ]
  contentful_token           = "..."
  contentful_organization_id = "..."
}
