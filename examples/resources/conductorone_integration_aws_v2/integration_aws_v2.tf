resource "conductorone_integration_aws_v2" "aws_v2" {
  app_id = conductorone_app.aws_v2.id
  user_ids = [
    conductorone_user.admin.id
  ]
  aws_external_id           = "..."
  aws_role_arn              = "..."
  aws_orgs_enable           = false
  aws_sso_enable            = false
  aws_sso_region            = "..."
  aws_sso_scim_enable       = false
  aws_sso_scim_endpoint     = "..."
  aws_sso_scim_access_token = "..."
  aws_sync_secrets          = false
}
