resource "conductorone_sign_in_policy" "my_signinpolicy" {
  allowed_mfa_types = [
    "CREDENTIAL_TYPE_DELEGATED_GOOGLE"
  ]
  allowed_primary_types = [
    "CREDENTIAL_TYPE_UNSPECIFIED"
  ]
  default_outcome = {
    allow = {
      floor_level = "AUTH_LEVEL_UNSPECIFIED"
    }
    challenge_required = {
      types = [
        "CREDENTIAL_TYPE_DELEGATED_GOOGLE"
      ]
    }
    deny = {
      reason_admin = "...my_reason_admin..."
      reason_user  = "...my_reason_user..."
    }
    enrollment_required = {
      credential_types = [
        "CREDENTIAL_TYPE_TOTP"
      ]
    }
    step_up_required = {
      level           = "AUTH_LEVEL_UNSPECIFIED"
      max_age_seconds = 4
      types = [
        "CREDENTIAL_TYPE_EMAIL_OTP"
      ]
    }
  }
  display_name = "...my_display_name..."
  priority     = 6
  rules = [
    {
      description = "...my_description..."
      id          = "...my_id..."
      match_cel   = "...my_match_cel..."
      mode        = "POLICY_RULE_MODE_DISABLED"
      outcome = {
        allow = {
          floor_level = "AUTH_LEVEL_PHR"
        }
        challenge_required = {
          types = [
            "CREDENTIAL_TYPE_DELEGATED_GOOGLE"
          ]
        }
        deny = {
          reason_admin = "...my_reason_admin..."
          reason_user  = "...my_reason_user..."
        }
        enrollment_required = {
          credential_types = [
            "CREDENTIAL_TYPE_TOTP"
          ]
        }
        step_up_required = {
          level           = "AUTH_LEVEL_PHR"
          max_age_seconds = 3
          types = [
            "CREDENTIAL_TYPE_RECOVERY_CODE"
          ]
        }
      }
    }
  ]
  sign_in_policy_service_delete_request = {
    # ...
  }
}