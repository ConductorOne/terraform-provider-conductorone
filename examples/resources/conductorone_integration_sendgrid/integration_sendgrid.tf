resource "conductorone_integration_sendgrid" "sendgrid" {
  app_id = conductorone_app.sendgrid.id
  user_ids = [
   conductorone_user.admin.id
  ]
  sendgrid_api_key = "..."
  }
