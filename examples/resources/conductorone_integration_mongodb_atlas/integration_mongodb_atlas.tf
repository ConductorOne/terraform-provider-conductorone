resource "conductorone_integration_mongodb_atlas" "mongodb_atlas" {
  app_id = conductorone_app.mongodb_atlas.id
  user_ids = [
    conductorone_user.admin.id
  ]
  mongodbatlas_public_key  = "..."
  mongodbatlas_private_key = "..."
}
