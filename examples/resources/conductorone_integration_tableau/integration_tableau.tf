resource "conductorone_integration_tableau" "tableau" {
  app_id = conductorone_app.tableau.id
  user_ids = [
    conductorone_user.admin.id
  ]
  tableau_site_id             = "..."
  tableau_server_path         = "..."
  tableau_access_token_name   = "..."
  tableau_access_token_secret = "..."
}
