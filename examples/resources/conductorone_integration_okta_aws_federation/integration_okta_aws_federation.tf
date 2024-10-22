resource "conductorone_integration_okta_aws_federation" "okta_aws_federation" {
  app_id = conductorone_app.okta_aws_federation.id
  user_ids = [
    conductorone_user.admin.id
  ]
  okta_aws_federation_domain          = "..."
  okta_aws_federation_api_token       = "..."
  okta_aws_federation_aws_okta_app_id = "..."
}
