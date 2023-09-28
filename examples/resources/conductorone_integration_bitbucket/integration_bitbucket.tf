resource "conductorone_integration_bitbucket" "bitbucket" {
  app_id = conductorone_app.bitbucket.id
  user_ids = [
   conductorone_user.admin.id
  ]
  bitbucket_consumer_key = "..."
  bitbucket_consumer_secret = "..."
  }
