resource "conductorone_integration_docusign" "docusign" {
  app_id = conductorone_app.docusign.id
  user_ids = [
   conductorone_user.admin.id
  ]
  account_id = "..."
  }
