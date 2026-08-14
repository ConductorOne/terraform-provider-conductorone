resource "conductorone_sign_in_policy" "my_signinpolicy" {
  allowed_mfa_types = [
    "CREDENTIAL_TYPE_DELEGATED_GOOGLE"
  ]
  allowed_primary_types = [
    "CREDENTIAL_TYPE_UNSPECIFIED"
  ]
  display_name = "...my_display_name..."
  policy_outcome = {
    allow = {
      floor_level = "AUTH_LEVEL_UNSPECIFIED"
    }
    challenge_required = {
      types = [
        "CREDENTIAL_TYPE_EMAIL_OTP"
      ]
    }
    deny = {
      reason_admin = "...my_reason_admin..."
      reason_user  = "...my_reason_user..."
    }
    enrollment_required = {
      credential_types = [
        "CREDENTIAL_TYPE_UNSPECIFIED"
      ]
    }
    step_up_required = {
      level           = "AUTH_LEVEL_NONE"
      max_age_seconds = 3
      types = [
        "CREDENTIAL_TYPE_DELEGATED_GOOGLE"
      ]
    }
  }
  priority = 6
  rules = [
    {
      description = "...my_description..."
      id          = "...my_id..."
      match_cel   = "...my_match_cel..."
      mode        = "POLICY_RULE_MODE_DISABLED"
      policy_outcome = {
        allow = {
          floor_level = "AUTH_LEVEL_UNSPECIFIED"
        }
        challenge_required = {
          types = [
            "CREDENTIAL_TYPE_PASSKEY"
          ]
        }
        deny = {
          reason_admin = "...my_reason_admin..."
          reason_user  = "...my_reason_user..."
        }
        enrollment_required = {
          credential_types = [
            "CREDENTIAL_TYPE_PASSKEY"
          ]
        }
        step_up_required = {
          level           = "AUTH_LEVEL_NONE"
          max_age_seconds = 5
          types = [
            "CREDENTIAL_TYPE_TOTP"
          ]
        }
      }
    }
  ]
  sign_in_policy_service_delete_request = {
    # ...
  }
}