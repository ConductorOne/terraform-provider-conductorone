resource "conductorone_integration_datadog_v2" "datadog_v2" {
  app_id = conductorone_app.datadog_v2.id
  user_ids = [
    conductorone_user.admin.id
  ]
  datadog_site            = "..."
  datadog_api_key         = "..."
  datadog_application_key = "..."
}
