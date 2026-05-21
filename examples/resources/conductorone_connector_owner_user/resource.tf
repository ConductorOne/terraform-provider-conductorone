resource "conductorone_connector_owner_user" "my_connector_owner_user" {
  app_id       = "...my_app_id..."
  connector_id = "...my_connector_id..."
  role_slug    = "...my_role_slug..."
  user_ref = {
    id = "...my_id..."
  }
  user_ref_id = "...my_user_ref_id..."
}