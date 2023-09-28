resource "conductorone_integration_bamboo_hr" "bamboo_hr" {
  app_id = conductorone_app.bamboo_hr.id
  user_ids = [
   conductorone_user.admin.id
  ]
  company_domain = "..."
  api_key = "..."
  }
