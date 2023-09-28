resource "conductorone_integration_cloudamqp" "cloudamqp" {
  app_id = conductorone_app.cloudamqp.id
  user_ids = [
   conductorone_user.admin.id
  ]
  cloudamqp_api_key = "..."
  }
