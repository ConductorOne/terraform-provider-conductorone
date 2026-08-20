resource "conductorone_recovery_policy" "my_recoverypolicy" {
  allowed_recovery_types = [
    "CREDENTIAL_TYPE_EMAIL_OTP"
  ]
  display_name            = "...my_display_name..."
  min_recovery_auth_level = "AUTH_LEVEL_NONE"
  priority                = 6
  recovery_policy_service_delete_request = {
    # ...
  }
  revoke_on_recovery = false
}