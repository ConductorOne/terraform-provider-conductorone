resource "conductorone_integration_linode" "linode" {
  app_id = conductorone_app.linode.id
  user_ids = [
    conductorone_user.admin.id
  ]
  linode_authorization_token = "..."
  linode_api_url             = "..."
}
