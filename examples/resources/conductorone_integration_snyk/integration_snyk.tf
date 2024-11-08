resource "conductorone_integration_snyk" "snyk" {
  app_id = conductorone_app.snyk.id
  user_ids = [
    conductorone_user.admin.id
  ]
  snyk_api_token = "..."
  snyk_group_id  = "..."
  snyk_org_ids   = ["..."]
}
