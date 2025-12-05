resource "conductorone_integration_docusign_v2" "docusign_v2" {
  app_id = conductorone_app.docusign_v2.id
  user_ids = [
    conductorone_user.admin.id
  ]
  use_demo_environment = false
  docusign_account_id  = "..."
}
