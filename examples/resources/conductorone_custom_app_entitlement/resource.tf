resource "conductorone_custom_app_entitlement" "my_custom_app_entitlement" {
  alias                = "...my_alias..."
  app_id               = "...my_app_id..."
  app_resource_id      = "...my_app_resource_id..."
  app_resource_type_id = "...my_app_resource_type_id..."
  certify_policy_id    = "...my_certify_policy_id..."
  compliance_framework_value_ids = [
    "..."
  ]
  description    = "...my_description..."
  display_name   = "...my_display_name..."
  duration_grant = "...my_duration_grant..."
  duration_unset = {
    # ...
  }
  emergency_grant_enabled           = false
  emergency_grant_policy_id         = "...my_emergency_grant_policy_id..."
  grant_policy_id                   = "...my_grant_policy_id..."
  match_baton_id                    = "...my_match_baton_id..."
  override_access_requests_defaults = true
  provision_policy = {
    action_provision = {
      action_name  = "...my_action_name..."
      app_id       = "...my_app_id..."
      connector_id = "...my_connector_id..."
    }
    connector_provision = {
      account_provision = {
        config = {
          # ...
        }
        connector_id = "...my_connector_id..."
        do_not_save = {
          # ...
        }
        save_to_vault = {
          vault_ids = [
            "..."
          ]
        }
        schema_id = "...my_schema_id..."
      }
      default_behavior = {
        connector_id = "...my_connector_id..."
      }
      delete_account = {
        connector_id = "...my_connector_id..."
      }
    }
    delegated_provision = {
      app_id         = "...my_app_id..."
      entitlement_id = "...my_entitlement_id..."
    }
    external_ticket_provision = {
      app_id                                = "...my_app_id..."
      connector_id                          = "...my_connector_id..."
      external_ticket_provisioner_config_id = "...my_external_ticket_provisioner_config_id..."
      instructions                          = "...my_instructions..."
    }
    manual_provision = {
      instructions = "...my_instructions..."
      user_ids = [
        "..."
      ]
    }
    multi_step = "{ \"see\": \"documentation\" }"
    unconfigured_provision = {
      # ...
    }
    webhook_provision = {
      webhook_id = "...my_webhook_id..."
    }
  }
  purpose             = "APP_ENTITLEMENT_PURPOSE_VALUE_UNSPECIFIED"
  revoke_policy_id    = "...my_revoke_policy_id..."
  risk_level_value_id = "...my_risk_level_value_id..."
  slug                = "...my_slug..."
}