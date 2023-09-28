resource "conductorone_integration_buildkite" "buildkite" {
  app_id = conductorone_app.buildkite.id
  user_ids = [
   conductorone_user.admin.id
  ]
  api_token = "..."
  organization = "..."
  }
