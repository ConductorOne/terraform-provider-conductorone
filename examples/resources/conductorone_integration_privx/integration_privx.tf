resource "conductorone_integration_privx" "privx" {
  app_id = conductorone_app.privx.id
  user_ids = [
    conductorone_user.admin.id
  ]
  privx_group_oauth = {
    privx_base_url            = "..."
    privx_oauth_client_id     = "..."
    privx_oauth_client_secret = "..."
  }
  privx_group_client_secret = {
    privx_base_url      = "..."
    privx_client_id     = "..."
    privx_client_secret = "..."
  }
}
