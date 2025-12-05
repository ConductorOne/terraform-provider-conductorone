resource "conductorone_integration_couchdb" "couchdb" {
  app_id = conductorone_app.couchdb.id
  user_ids = [
    conductorone_user.admin.id
  ]
  couchdb_url      = "..."
  couchdb_username = "..."
  couchdb_password = "..."
}
