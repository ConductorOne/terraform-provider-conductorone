resource "conductorone_policy" "my_policy" {
  description  = "...my_description..."
  display_name = "...my_display_name..."
  policy_steps = {
    key = {
      steps = [
        {
          accept = {
            accept_message = "...my_accept_message..."
          }
          approval = {
            agent_approval = {
              agent_user_id = "...my_agent_user_id..."
              instructions  = "...my_instructions..."
              policy_ids = [
                "..."
              ]
            }
            allow_reassignment = false
            allowed_reassignees = [
              "..."
            ]
            app_group_approval = {
              allow_self_approval = true
              app_group_id        = "...my_app_group_id..."
              app_id              = "...my_app_id..."
              fallback            = true
              fallback_user_ids = [
                "..."
              ]
            }
            app_owner_approval = {
              allow_self_approval = true
            }
            entitlement_owner_approval = {
              allow_self_approval = false
              fallback            = true
              fallback_user_ids = [
                "..."
              ]
            }
            expression_approval = {
              allow_self_approval = true
              expressions = [
                "..."
              ]
              fallback = false
              fallback_user_ids = [
                "..."
              ]
            }
            manager_approval = {
              allow_self_approval = true
              fallback            = true
              fallback_user_ids = [
                "..."
              ]
            }
            require_approval_reason      = true
            require_denial_reason        = true
            require_reassignment_reason  = true
            requires_step_up_provider_id = "...my_requires_step_up_provider_id..."
            resource_owner_approval = {
              allow_self_approval = true
              fallback            = false
              fallback_user_ids = [
                "..."
              ]
            }
            self_approval = {
              fallback = false
              fallback_user_ids = [
                "..."
              ]
            }
            user_approval = {
              allow_self_approval = false
              user_ids = [
                "..."
              ]
            }
            webhook_approval = {
              webhook_id = "...my_webhook_id..."
            }
          }
          provision = {
            assigned = true
            provision_policy = {
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
            provision_target = {
              app_entitlement_id = "...my_app_entitlement_id..."
              app_id             = "...my_app_id..."
              app_user_id        = "...my_app_user_id..."
              grant_duration     = "...my_grant_duration..."
            }
          }
          reject = {
            reject_message = "...my_reject_message..."
          }
          wait = {
            comment_on_first_wait = "...my_comment_on_first_wait..."
            comment_on_timeout    = "...my_comment_on_timeout..."
            name                  = "...my_name..."
            timeout_duration      = "...my_timeout_duration..."
            wait_condition = {
              condition = "...my_condition..."
            }
          }
        }
      ]
    }
  }
  policy_type = "POLICY_TYPE_UNSPECIFIED"
  post_actions = [
    {
      certify_remediate_immediately = false
    }
  ]
  reassign_tasks_to_delegates = false
  rules = [
    {
      condition  = "...my_condition..."
      policy_key = "...my_policy_key..."
    }
  ]
}