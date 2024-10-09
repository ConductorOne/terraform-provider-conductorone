resource "conductorone_custom_app_entitlement" "okta_test_admin" {
  app_id               = "2mDhyLPiZch9LFjuQHb9ucAezW0"
  app_resource_id      = "2myx9VP6aMbcYd0E2d3uOmFp4gH"
  app_resource_type_id = "2myx9Y3ov04FiALGuPceBhIl4zO"
  display_name         = "Test new entitlement"
  alias                = "okta_test_admin"
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
  certify_policy_id                 = "2ZmHrHKTozPmh7RNcDqBUcSJlpl"
  grant_policy_id                   = "2ZmHrLxLq74Q8IYwoBo79Pcka69"
  purpose                           = "APP_ENTITLEMENT_PURPOSE_VALUE_ASSIGNMENT"
  revoke_policy_id                  = "2n7u6jYKbF9EqzojTqPwSsmuN3S"
  duration_grant                    = "3601s"
}
