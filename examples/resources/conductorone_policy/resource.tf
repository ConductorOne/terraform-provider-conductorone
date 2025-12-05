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
          action = {
            action_target_automation = {
              automation_template_id = "...my_automation_template_id..."
            }
          }
          approval = {
            agent_approval = {
              agent_failure_action = "APPROVAL_AGENT_FAILURE_ACTION_REASSIGN_TO_USERS"
              agent_mode           = "APPROVAL_AGENT_MODE_FULL_CONTROL"
              agent_user_id        = "...my_agent_user_id..."
              instructions         = "...my_instructions..."
              policy_ids = [
                "..."
              ]
              reassign_to_user_ids = [
                "..."
              ]
            }
            allow_delegation   = true
            allow_reassignment = false
            allowed_reassignees = [
              "..."
            ]
            app_group_approval = {
              allow_self_approval = true
              app_group_id        = "...my_app_group_id..."
              app_id              = "...my_app_id..."
              fallback            = true
              fallback_group_ids = [
                {
                  app_entitlement_id = "...my_app_entitlement_id..."
                  app_id             = "...my_app_id..."
                }
              ]
              fallback_user_ids = [
                "..."
              ]
              is_group_fallback_enabled  = false
              require_distinct_approvers = false
            }
            app_owner_approval = {
              allow_self_approval        = true
              require_distinct_approvers = true
            }
            entitlement_owner_approval = {
              allow_self_approval = false
              fallback            = true
              fallback_user_ids = [
                "..."
              ]
              require_distinct_approvers = false
            }
            escalation = {
              escalation_comment = "...my_escalation_comment..."
              expiration         = "...my_expiration..."
              reassign_to_approvers = {
                approver_ids = [
                  "..."
                ]
              }
              replace_policy = {
                policy_id = "...my_policy_id..."
              }
            }
            escalation_enabled = false
            expression_approval = {
              allow_self_approval = true
              expressions = [
                "..."
              ]
              fallback = false
              fallback_user_ids = [
                "..."
              ]
              require_distinct_approvers = true
            }
            manager_approval = {
              allow_self_approval = true
              fallback            = true
              fallback_user_ids = [
                "..."
              ]
              require_distinct_approvers = true
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
              require_distinct_approvers = true
            }
            self_approval = {
              fallback = false
              fallback_user_ids = [
                "..."
              ]
            }
            user_approval = {
              allow_self_approval        = false
              require_distinct_approvers = false
              user_ids = [
                "..."
              ]
            }
            webhook_approval = {
              webhook_id = "...my_webhook_id..."
            }
          }
          form = "{ \"see\": \"documentation\" }"
          provision = {
            assigned = true
            provision_policy = {
              action_provision = {
                action_name  = "...my_action_name..."
                app_id       = "...my_app_id..."
                connector_id = "...my_connector_id..."
                display_name = "...my_display_name..."
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
            wait_duration = {
              duration = "...my_duration..."
            }
            wait_until_time = {
              hours    = 5
              minutes  = 8
              timezone = "...my_timezone..."
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