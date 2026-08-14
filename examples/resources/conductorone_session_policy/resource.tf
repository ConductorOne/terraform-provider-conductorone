resource "conductorone_session_policy" "my_sessionpolicy" {
  access_token_ttl_seconds = 5
  continuous_rules = [
    {
      description = "...my_description..."
      id          = "...my_id..."
      match_cel   = "...my_match_cel..."
      mode        = "POLICY_RULE_MODE_DISABLED"
      session_policy_policy_outcome = {
        session_policy_allow = {
          floor_level = "AUTH_LEVEL_SINGLE_FACTOR"
        }
        session_policy_challenge_required = {
          types = [
            "CREDENTIAL_TYPE_RECOVERY_CODE"
          ]
        }
        session_policy_deny = {
          reason_admin = "...my_reason_admin..."
          reason_user  = "...my_reason_user..."
        }
        session_policy_enrollment_required = {
          credential_types = [
            "CREDENTIAL_TYPE_PASSKEY"
          ]
        }
        session_policy_step_up_required = {
          level           = "AUTH_LEVEL_NONE"
          max_age_seconds = 2
          types = [
            "CREDENTIAL_TYPE_UPSTREAM_IDP"
          ]
        }
      }
    }
  ]
  credential_durations = [
    {
      access_token_ttl_seconds     = 6
      credential_type              = "CREDENTIAL_TYPE_RECOVERY_CODE"
      max_session_duration_seconds = 3
    }
  ]
  display_name                    = "...my_display_name..."
  idle_timeout_seconds            = 7
  max_session_duration_seconds    = 6
  persistence                     = "PERSISTENCE_MODE_ALLOW_USER_CHOICE"
  priority                        = 3
  refresh_rotation_window_seconds = 3
  refresh_token_ttl_seconds       = 8
  rotate_refresh_on_use           = true
  session_policy_policy_outcome = {
    session_policy_allow = {
      floor_level = "AUTH_LEVEL_SINGLE_FACTOR"
    }
    session_policy_challenge_required = {
      types = [
        "CREDENTIAL_TYPE_PASSWORD"
      ]
    }
    session_policy_deny = {
      reason_admin = "...my_reason_admin..."
      reason_user  = "...my_reason_user..."
    }
    session_policy_enrollment_required = {
      credential_types = [
        "CREDENTIAL_TYPE_PASSWORD"
      ]
    }
    session_policy_step_up_required = {
      level           = "AUTH_LEVEL_NONE"
      max_age_seconds = 9
      types = [
        "CREDENTIAL_TYPE_DELEGATED_GOOGLE"
      ]
    }
  }
  session_policy_service_delete_request = {
    # ...
  }
  ssf_receiver_config = {
    enabled = false
    ssf_receiver_stream_ids = [
      "..."
    ]
  }
  ssf_transmitter_config = {
    enabled = false
    event_types = [
      "..."
    ]
    ssf_transmitter_stream_ids = [
      "..."
    ]
  }
}