resource "conductorone_integration_jamf" "jamf" {
  app_id = conductorone_app.jamf.id
  user_ids = [
   conductorone_user.admin.id
  ]
  jamf_instance_url = "..."
  jamf_username = "..."
  jamf_password = "..."
  }
