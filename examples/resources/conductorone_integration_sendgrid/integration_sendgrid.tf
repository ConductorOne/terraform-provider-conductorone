resource "conductorone_integration_sendgrid" "sendgrid" {
  app_id           = conductorone_app.sendgrid.id
  sendgrid_api_key = "..."
}
