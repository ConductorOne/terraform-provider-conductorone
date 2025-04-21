resource "conductorone_integration_airbyte" "airbyte" {
  app_id = conductorone_app.airbyte.id
  user_ids = [
    conductorone_user.admin.id
  ]
  airbyte_hostname      = "..."
  airbyte_client_id     = "..."
  airbyte_client_secret = "..."
}
