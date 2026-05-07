resource "conductorone_app_entitlement_owner_user" "my_app_entitlement_owner_user" {
  app_id         = "...my_app_id..."
  entitlement_id = "...my_entitlement_id..."
  role_slug      = "...my_role_slug..."
  user_ref = {
    id = "...my_id..."
  }
  user_ref_id = "...my_user_ref_id..."
}