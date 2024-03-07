resource "conductorone_app_entitlement" "okta_test_admin" {
  id     = data.conductorone_app_entitlement.okta_test_admin.id
  app_id = data.conductorone_app_entitlement.okta_test_admin.app_id
  alias  = "okta_test_admin"
  provision_policy = {
    manual_provision = {
      instructions = "Please contact the IT department to request this entitlement."
      user_ids = [
        data.conductorone_user.my_user.id
      ]
    }
  }

  duration_grant      = "3600s"
  risk_level_value_id = conductorone_risk_level.test_risk_level.id
  compliance_framework_value_ids = [
    data.conductorone_compliance_framework.soc2.id
  ]
}
