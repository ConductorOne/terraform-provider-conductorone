resource "conductorone_integration_datadog" "datadog" {
  app_id                  = conductorone_app.datadog.id
  datadog_site            = "..."
  datadog_api_key         = "..."
  datadog_application_key = "..."
}
