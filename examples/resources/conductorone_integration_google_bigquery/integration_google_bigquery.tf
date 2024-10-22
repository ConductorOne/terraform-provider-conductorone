resource "conductorone_integration_google_bigquery" "google_bigquery" {
  app_id = conductorone_app.google_bigquery.id
  user_ids = [
    conductorone_user.admin.id
  ]
  credentials_json = "..."
}
