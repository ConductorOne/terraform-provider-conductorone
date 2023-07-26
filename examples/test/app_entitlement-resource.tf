resource "conductorone_app_entitlement" "okta_test_admin" {
  id     = data.conductorone_app_entitlement.okta_test_admin.id
  app_id = data.conductorone_app_entitlement.okta_test_admin.app_id
  alias  = "okta_test_admin"

  certify_policy_id = conductorone_policy.new_policy.id
}
