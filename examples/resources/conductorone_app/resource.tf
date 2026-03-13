resource "conductorone_app" "my_app" {
  certify_policy_id                      = "...my_certify_policy_id..."
  description                            = "...my_description..."
  display_name                           = "...my_display_name..."
  grant_policy_id                        = "...my_grant_policy_id..."
  identity_matching                      = "APP_USER_IDENTITY_MATCHING_CUSTOM"
  instructions                           = "...my_instructions..."
  monthly_cost_usd                       = 1
  revoke_policy_id                       = "...my_revoke_policy_id..."
  strict_access_entitlement_provisioning = true
}