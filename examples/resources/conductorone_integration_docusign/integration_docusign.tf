resource "conductorone_integration_docusign" "docusign" {
  app_id     = conductorone_app.docusign.id
  account_id = "..."
}
