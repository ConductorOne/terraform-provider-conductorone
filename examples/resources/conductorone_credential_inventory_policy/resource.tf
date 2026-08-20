resource "conductorone_credential_inventory_policy" "my_credentialinventorypolicy" {
  credential_inventory_policy_service_delete_request = {
    # ...
  }
  delegated = {
    google_enabled = true
    google_hosted_domains = [
      "..."
    ]
    microsoft_enabled = true
    microsoft_tenant_ids = [
      "..."
    ]
  }
  display_name = "...my_display_name..."
  email_otp = {
    code_length  = 8
    max_attempts = 0
    ttl_seconds  = 0
  }
  enabled_types = [
    "CREDENTIAL_TYPE_EMAIL_OTP"
  ]
  passkey = {
    allowed_aaguids = [
      "..."
    ]
    attestation               = "ATTESTATION_REQUIREMENT_UNSPECIFIED"
    require_user_verification = false
  }
  password = {
    check_breached     = false
    history_depth      = 1
    min_length         = 2
    require_mixed_case = false
    require_number     = false
    require_symbol     = false
  }
  priority = 0
  totp = {
    code_length    = 5
    period_seconds = 6
    skew_tolerance = 0
  }
}