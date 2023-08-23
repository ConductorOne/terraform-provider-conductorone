resource "conductorone_app_entitlement_owner" "okta_test_admin_myself" {
  entitlement_id = data.conductorone_app_entitlement.okta_test_admin.id
  app_id         = data.conductorone_app_entitlement.okta_test_admin.app_id
  user_id        = data.conductorone_user.my_user.id
}

