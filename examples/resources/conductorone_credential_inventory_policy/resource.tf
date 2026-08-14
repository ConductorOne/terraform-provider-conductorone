resource "conductorone_credential_inventory_policy" "my_credentialinventorypolicy" {
  credential_inventory_policy_service_delete_request = {
    # ...
  }
  delegated_constraints = {
    google_enabled = false
    google_hosted_domains = [
      "..."
    ]
    microsoft_enabled = false
    microsoft_tenant_ids = [
      "..."
    ]
  }
  display_name = "...my_display_name..."
  email_otp_constraints = {
    code_length  = 7
    max_attempts = 10
    ttl_seconds  = 10
  }
  enabled_types = [
    "CREDENTIAL_TYPE_EMAIL_OTP"
  ]
  passkey_constraints = {
    allowed_aaguids = [
      "..."
    ]
    attestation               = "ATTESTATION_REQUIREMENT_INDIRECT"
    require_user_verification = true
  }
  password_constraints = {
    check_breached     = false
    history_depth      = 4
    min_length         = 5
    require_mixed_case = true
    require_number     = true
    require_symbol     = false
  }
  priority = 0
  totp_constraints = {
    code_length    = 0
    period_seconds = 6
    skew_tolerance = 4
  }
}