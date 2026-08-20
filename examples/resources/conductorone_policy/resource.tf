resource "conductorone_policy" "my_policy" {
  annotations = {
    key = "value"
  }
  baseline_policy_id = "...my_baseline_policy_id..."
  description        = "...my_description..."
  display_name       = "...my_display_name..."
  policy_steps = {
    key = {
      steps = [
        {
          accept = {
            accept_message = "...my_accept_message..."
          }
          action = {
            automation = {
              automation_template_id = "...my_automation_template_id..."
            }
            baton_resource_action = {
              baton_resource_action_id = "...my_baton_resource_action_id..."
            }
            client_id_approval = {
              # ...
            }
          }
          approval = {
            agent = {
              agent_failure_action = "APPROVAL_AGENT_FAILURE_ACTION_REASSIGN_TO_USERS"
              agent_mode           = "APPROVAL_AGENT_MODE_COMMENT_ONLY"
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
            app_owners = {
              allow_self_approval        = false
              require_distinct_approvers = false
            }
            entitlement_owners = {
              allow_self_approval = true
              fallback            = false
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
            escalation = {
              cancel_ticket = {
                # ...
              }
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
              skip_step = {
                # ...
              }
            }
            escalation_enabled = false
            expression = {
              allow_self_approval = false
              expressions = [
                "..."
              ]
              fallback = true
              fallback_group_ids = [
                {
                  app_entitlement_id = "...my_app_entitlement_id..."
                  app_id             = "...my_app_id..."
                }
              ]
              fallback_user_ids = [
                "..."
              ]
              is_group_fallback_enabled  = true
              require_distinct_approvers = false
            }
            group = {
              allow_self_approval = true
              app_group_id        = "...my_app_group_id..."
              app_id              = "...my_app_id..."
              fallback            = false
              fallback_group_ids = [
                {
                  app_entitlement_id = "...my_app_entitlement_id..."
                  app_id             = "...my_app_id..."
                }
              ]
              fallback_user_ids = [
                "..."
              ]
              is_group_fallback_enabled  = true
              require_distinct_approvers = false
            }
            manager = {
              allow_self_approval = false
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
              is_group_fallback_enabled  = true
              require_distinct_approvers = true
            }
            require_approval_reason      = true
            require_denial_reason        = true
            require_reassignment_reason  = true
            requires_step_up_provider_id = "...my_requires_step_up_provider_id..."
            resource_owners = {
              allow_self_approval = true
              fallback            = false
              fallback_group_ids = [
                {
                  app_entitlement_id = "...my_app_entitlement_id..."
                  app_id             = "...my_app_id..."
                }
              ]
              fallback_user_ids = [
                "..."
              ]
              is_group_fallback_enabled  = true
              require_distinct_approvers = true
            }
            self = {
              fallback = false
              fallback_group_ids = [
                {
                  app_entitlement_id = "...my_app_entitlement_id..."
                  app_id             = "...my_app_id..."
                }
              ]
              fallback_user_ids = [
                "..."
              ]
              is_group_fallback_enabled = true
            }
            users = {
              allow_self_approval        = false
              require_distinct_approvers = false
              user_ids = [
                "..."
              ]
            }
            webhook = {
              webhook_id = "...my_webhook_id..."
            }
          }
          form = "{ \"see\": \"documentation\" }"
          provision = {
            assigned = true
            provision_policy = {
              action = {
                action_name  = "...my_action_name..."
                app_id       = "...my_app_id..."
                connector_id = "...my_connector_id..."
                display_name = "...my_display_name..."
              }
              connector = {
                account = {
                  config       = "{ \"see\": \"documentation\" }"
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
              delegated = {
                app_id         = "...my_app_id..."
                entitlement_id = "...my_entitlement_id..."
              }
              device_placement = {
                vault_boundary_id = "...my_vault_boundary_id..."
              }
              external_ticket = {
                app_id                                = "...my_app_id..."
                connector_id                          = "...my_connector_id..."
                external_ticket_provisioner_config_id = "...my_external_ticket_provisioner_config_id..."
                instructions                          = "...my_instructions..."
              }
              manual = {
                assignee = {
                  app_owners = {
                    allow_reassignment = true
                    fallback_user_ids = [
                      "..."
                    ]
                  }
                  entitlement_owners = {
                    allow_reassignment = true
                    fallback_user_ids = [
                      "..."
                    ]
                  }
                  expression = {
                    allow_reassignment = false
                    expressions = [
                      "..."
                    ]
                    fallback_user_ids = [
                      "..."
                    ]
                  }
                  group = {
                    allow_reassignment = true
                    app_group_id       = "...my_app_group_id..."
                    app_id             = "...my_app_id..."
                    fallback_user_ids = [
                      "..."
                    ]
                  }
                  manager = {
                    allow_reassignment = true
                    fallback_user_ids = [
                      "..."
                    ]
                  }
                  users = {
                    allow_reassignment = false
                    user_ids = [
                      "..."
                    ]
                  }
                }
                instructions = "...my_instructions..."
                user_ids = [
                  "..."
                ]
              }
              multi_step = "{ \"see\": \"documentation\" }"
              unconfigured = {
                # ...
              }
              webhook = {
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
            condition = {
              condition = "...my_condition..."
            }
            duration = {
              duration = "...my_duration..."
            }
            name             = "...my_name..."
            timeout_duration = "...my_timeout_duration..."
            until_time = {
              hours    = 8
              minutes  = 9
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
      policy_id  = "...my_policy_id..."
      policy_key = "...my_policy_key..."
      step_key   = "...my_step_key..."
    }
  ]
  scope = {
    app_entitlement_id = "...my_app_entitlement_id..."
    app_id             = "...my_app_id..."
    slot               = "POLICY_SCOPE_SLOT_EMERGENCY"
  }
}