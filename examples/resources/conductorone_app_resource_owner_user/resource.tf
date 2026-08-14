resource "conductorone_app_resource_owner_user" "my_app_resource_owner_user" {
  app_id           = "...my_app_id..."
  resource_id      = "...my_resource_id..."
  resource_type_id = "...my_resource_type_id..."
  role_slug        = "...my_role_slug..."
  user_ref = {
    id = "...my_id..."
  }
  user_ref_id = "...my_user_ref_id..."
}