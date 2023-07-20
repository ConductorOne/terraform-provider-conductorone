resource "conductorone_integration_ukg" "ukg" {
  app_id = conductorone_app.ukg.id
  user_ids = [
    conductorone_user.admin.id
  ]
  ukg_customer_api_key = "..."
  ukg_username         = "..."
  ukg_password         = "..."
  ukg_service_endpoint = "..."
}
