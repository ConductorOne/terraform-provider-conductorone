resource "conductorone_custom_app_entitlement" "custom_app_entitlement" {
  app_id               = "<app_id>"
  app_resource_id      = "<app_resource_id>"
  app_resource_type_id = "<app_resource_type_id>"
  display_name         = "Test new entitlement"
  alias                = "terraform_test_entitlement"
  provision_policy = {
    manual_provision = {
      instructions = "Please contact the IT department to request this entitlement."
      user_ids = [
        data.conductorone_user.my_user.id
      ]
    }
  }
  risk_level_value_id               = data.conductorone_risk_level.high.id
  slug                              = "test slug"
  description                       = "test description"
  compliance_framework_value_ids    = [data.conductorone_compliance_framework.soc2.id]
  override_access_requests_defaults = true
  certify_policy_id                 = "<certify_policy_id>"
  grant_policy_id                   = "<grant_policy_id>"
  purpose                           = "APP_ENTITLEMENT_PURPOSE_VALUE_ASSIGNMENT"
  revoke_policy_id                  = "<revoke_policy_id>"
  duration_grant                    = "3601s"
}
