resource "conductorone_integration_okta_ciam" "okta_ciam" {
  app_id = conductorone_app.okta_ciam.id
  user_ids = [
    conductorone_user.admin.id
  ]
  okta_ciam_domain        = "..."
  okta_ciam_api_token     = "..."
  okta_ciam_email_domains = "..."
}
