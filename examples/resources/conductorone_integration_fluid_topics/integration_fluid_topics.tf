resource "conductorone_integration_fluid_topics" "fluid_topics" {
  app_id = conductorone_app.fluid_topics.id
  user_ids = [
    conductorone_user.admin.id
  ]
  fluid_topics_domain    = "..."
  fluid_topics_api_token = "..."
}
