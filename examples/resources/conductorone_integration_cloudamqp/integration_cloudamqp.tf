resource "conductorone_integration_cloudamqp" "cloudamqp" {
  app_id            = conductorone_app.cloudamqp.id
  cloudamqp_api_key = "..."
}
