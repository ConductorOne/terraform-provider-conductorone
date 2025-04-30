resource "conductorone_integration_trello" "trello" {
  app_id = conductorone_app.trello.id
  user_ids = [
    conductorone_user.admin.id
  ]
  api_key       = "..."
  api_token     = "..."
  organizations = ["..."]
}
