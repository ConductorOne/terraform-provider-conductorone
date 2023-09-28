resource "conductorone_integration_datadog" "datadog" {
  app_id = conductorone_app.datadog.id
  user_ids = [
   conductorone_user.admin.id
  ]
  datadog_site = "..."
  datadog_api_key = "..."
  datadog_application_key = "..."
  }
