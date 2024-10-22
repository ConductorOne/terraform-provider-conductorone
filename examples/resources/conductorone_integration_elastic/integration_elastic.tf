resource "conductorone_integration_elastic" "elastic" {
  app_id = conductorone_app.elastic.id
  user_ids = [
    conductorone_user.admin.id
  ]
  elastic_api_key             = "..."
  elastic_deployment_api_key  = "..."
  elastic_deployment_endpoint = "..."
  elastic_organization_id     = "..."
}
