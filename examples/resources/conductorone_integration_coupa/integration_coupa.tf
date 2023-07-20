resource "conductorone_integration_coupa" "coupa" {
  app_id                                 = conductorone_app.coupa.id
  coupa_domain                           = "..."
  oauth2_client_cred_grant_client_id     = "..."
  oauth2_client_cred_grant_client_secret = "..."
}
