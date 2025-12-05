resource "conductorone_integration_successfactors" "successfactors" {
  app_id = conductorone_app.successfactors.id
  user_ids = [
    conductorone_user.admin.id
  ]
  company_id      = "..."
  cid             = "..."
  public_key      = "..."
  private_key     = "..."
  instance_url    = "..."
  issuer_url      = "..."
  subject_name_id = "..."
  saml_api_key    = "..."
}
