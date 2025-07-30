resource "conductorone_integration_oracle_cloud_infrastructure" "oracle_cloud_infrastructure" {
  app_id = conductorone_app.oracle_cloud_infrastructure.id
  user_ids = [
    conductorone_user.admin.id
  ]
  tenancy_ocid = "..."
  user_ocid    = "..."
  region       = "..."
  fingerprint  = "..."
  private_key  = "..."
  sync_secrets = false
}
