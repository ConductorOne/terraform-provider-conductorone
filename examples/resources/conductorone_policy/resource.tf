resource "conductorone_policy" "my_policy" {
  description  = "...my_description..."
  display_name = "...my_display_name..."
  policy_steps = {
    "Applications" = {
      steps = [
        {
          accept = {}
          approval = {
            allow_reassignment = true
            app_group_approval = {
              allow_self_approval = true
              app_group_id        = "...my_app_group_id..."
              app_id              = "...my_app_id..."
              fallback            = false
              fallback_user_ids = [
                "...",
              ]
            }
            app_owner_approval = {
              allow_self_approval = false
            }
            assigned = false
            entitlement_owner_approval = {
              allow_self_approval = true
              fallback            = false
              fallback_user_ids = [
                "...",
              ]
            }
            expression_approval = {
              allow_self_approval = false
              assigned_user_ids = [
                "...",
              ]
              expressions = [
                "...",
              ]
              fallback = true
              fallback_user_ids = [
                "...",
              ]
            }
            manager_approval = {
              allow_self_approval = true
              assigned_user_ids = [
                "...",
              ]
              fallback = true
              fallback_user_ids = [
                "...",
              ]
            }
            require_approval_reason     = false
            require_reassignment_reason = false
            self_approval = {
              assigned_user_ids = [
                "...",
              ]
              fallback = true
              fallback_user_ids = [
                "...",
              ]
            }
            user_approval = {
              allow_self_approval = false
              user_ids = [
                "...",
              ]
            }
          }
          provision = {
            assigned = true
            provision_policy = {
              connector_provision = {}
              delegated_provision = {
                app_id         = "...my_app_id..."
                entitlement_id = "...my_entitlement_id..."
                implicit       = true
              }
              manual_provision = {
                instructions = "...my_instructions..."
                user_ids = [
                  "...",
                ]
              }
            }
            provision_target = {
              app_entitlement_id = "...my_app_entitlement_id..."
              app_id             = "...my_app_id..."
              app_user_id        = "...my_app_user_id..."
              grant_duration     = "...my_grant_duration..."
            }
          }
          reject = {}
        },
      ]
    }
    "Metrics" = {
      steps = [
        {
          accept = {}
          approval = {
            allow_reassignment = false
            app_group_approval = {
              allow_self_approval = true
              app_group_id        = "...my_app_group_id..."
              app_id              = "...my_app_id..."
              fallback            = true
              fallback_user_ids = [
                "...",
              ]
            }
            app_owner_approval = {
              allow_self_approval = false
            }
            assigned = false
            entitlement_owner_approval = {
              allow_self_approval = true
              fallback            = true
              fallback_user_ids = [
                "...",
              ]
            }
            expression_approval = {
              allow_self_approval = false
              assigned_user_ids = [
                "...",
              ]
              expressions = [
                "...",
              ]
              fallback = true
              fallback_user_ids = [
                "...",
              ]
            }
            manager_approval = {
              allow_self_approval = false
              assigned_user_ids = [
                "...",
              ]
              fallback = true
              fallback_user_ids = [
                "...",
              ]
            }
            require_approval_reason     = true
            require_reassignment_reason = true
            self_approval = {
              assigned_user_ids = [
                "...",
              ]
              fallback = false
              fallback_user_ids = [
                "...",
              ]
            }
            user_approval = {
              allow_self_approval = true
              user_ids = [
                "...",
              ]
            }
          }
          provision = {
            assigned = false
            provision_policy = {
              connector_provision = {}
              delegated_provision = {
                app_id         = "...my_app_id..."
                entitlement_id = "...my_entitlement_id..."
                implicit       = false
              }
              manual_provision = {
                instructions = "...my_instructions..."
                user_ids = [
                  "...",
                ]
              }
            }
            provision_target = {
              app_entitlement_id = "...my_app_entitlement_id..."
              app_id             = "...my_app_id..."
              app_user_id        = "...my_app_user_id..."
              grant_duration     = "...my_grant_duration..."
            }
          }
          reject = {}
        },
      ]
    }
  }
  policy_type = "POLICY_TYPE_CERTIFY"
  post_actions = [
    {
      certify_remediate_immediately = true
    },
  ]
  reassign_tasks_to_delegates = false
}