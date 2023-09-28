resource "conductorone_integration_broadcom_sac" "broadcom_sac" {
  app_id = conductorone_app.broadcom_sac.id
  user_ids = [
   conductorone_user.admin.id
  ]
  username = "..."
  password = "..."
  tenant = "..."
  }
