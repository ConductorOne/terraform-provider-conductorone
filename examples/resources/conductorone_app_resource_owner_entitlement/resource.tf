resource "conductorone_app_resource_owner_entitlement" "my_app_resource_owner_entitlement" {
  app_entitlement_ref = {
    app_id = "...my_app_id..."
    id     = "...my_id..."
  }
  app_entitlement_ref_app_id = "...my_app_entitlement_ref_app_id..."
  app_entitlement_ref_id     = "...my_app_entitlement_ref_id..."
  app_id                     = "...my_app_id..."
  resource_id                = "...my_resource_id..."
  resource_type_id           = "...my_resource_type_id..."
  role_slug                  = "...my_role_slug..."
}