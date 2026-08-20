resource "conductorone_session_policy" "my_sessionpolicy" {
  access_token_ttl_seconds = 5
  continuous_default_outcome = {
    allow = {
      floor_level = "AUTH_LEVEL_NONE"
    }
    challenge_required = {
      types = [
        "CREDENTIAL_TYPE_DELEGATED_MICROSOFT"
      ]
    }
    deny = {
      reason_admin = "...my_reason_admin..."
      reason_user  = "...my_reason_user..."
    }
    enrollment_required = {
      credential_types = [
        "CREDENTIAL_TYPE_RECOVERY_CODE"
      ]
    }
    step_up_required = {
      level           = "AUTH_LEVEL_UNSPECIFIED"
      max_age_seconds = 6
      types = [
        "CREDENTIAL_TYPE_PASSWORD"
      ]
    }
  }
  continuous_rules = [
    {
      description = "...my_description..."
      id          = "...my_id..."
      match_cel   = "...my_match_cel..."
      mode        = "POLICY_RULE_MODE_DISABLED"
      outcome = {
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
            "CREDENTIAL_TYPE_RECOVERY_CODE"
          ]
        }
        step_up_required = {
          level           = "AUTH_LEVEL_PHRH"
          max_age_seconds = 7
          types = [
            "CREDENTIAL_TYPE_PASSWORD"
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
  session_policy_service_delete_request = {
    # ...
  }
  ssf_receive = {
    enabled = true
    ssf_receiver_stream_ids = [
      "..."
    ]
  }
  ssf_transmit = {
    enabled = true
    event_types = [
      "..."
    ]
    ssf_transmitter_stream_ids = [
      "..."
    ]
  }
}