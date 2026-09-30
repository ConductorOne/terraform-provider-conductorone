resource "conductorone_access_review_template" "my_access_review_template" {
  access_review_duration = "...my_access_review_duration..."
  accuracy_issue_action  = "ACCURACY_ISSUE_ACTION_WAIT"
  annotations = {
    key = "value"
  }
  auto_close_campaign  = false
  auto_close_decision  = "CLOSE_DECISION_NO_ACTION"
  auto_generate_report = false
  auto_start_campaign  = true
  column_config = {
    columns = [
      "ACCESS_REVIEW_TASK_COLUMN_RISK_LEVEL"
    ]
    ordered_columns = [
      {
        app_user_attribute_key = "...my_app_user_attribute_key..."
        builtin                = "ACCESS_REVIEW_TASK_COLUMN_INSIGHTS"
      }
    ]
  }
  default_view                      = "ACCESS_REVIEW_VIEW_TYPE_UNSPECIFIED"
  description                       = "...my_description..."
  display_name                      = "...my_display_name..."
  exempt_certified_access_conflicts = true
  is_campaign_schedule_enabled      = false
  notification_config = {
    send_close     = true
    send_kickoff   = false
    send_reminders = false
  }
  owner_ids = [
    "..."
  ]
  policy_id = "...my_policy_id..."
  recurrence_rule = {
    end_date    = "2022-07-11T08:35:49.034Z"
    frequency   = "FREQUENCY_UNSPECIFIED"
    interval    = 10
    occurrences = 3
    start_date  = "2022-06-20T11:29:45.535Z"
  }
  review_instructions = "...my_review_instructions..."
  reviewer_attribute_config = {
    bindings = [
      {
        app_id        = "...my_app_id..."
        attribute_key = "...my_attribute_key..."
      }
    ]
  }
  scope = {
    account_cel_expression = {
      expression = "...my_expression..."
    }
    account_criteria = {
      account_domain = "APP_USER_DOMAIN_UNSPECIFIED"
      account_types = [
        "APP_USER_TYPE_UNSPECIFIED"
      ]
      app_user_statuses = [
        "APP_USER_STATUS_UNSPECIFIED"
      ]
      no_account_owner = false
    }
    all_access_conflicts = {
      # ...
    }
    all_accounts = {
      # ...
    }
    all_grants = {
      # ...
    }
    all_users = {
      # ...
    }
    app_access = {
      # ...
    }
    app_selection_criteria = {
      compliance_framework_attribute_value_ids = [
        "..."
      ]
      risk_level_attribute_value_ids = [
        "..."
      ]
    }
    cel_expression = {
      expression = "...my_expression..."
    }
    excluded_resource_type_selections = {
      # ...
    }
    excluded_specific_resources = {
      # ...
    }
    grants_by_criteria = {
      access_profile_filter = {
        excluded_access_profile_ids = [
          "..."
        ]
        filter_type = "ACCESS_PROFILE_FILTER_TYPE_EXCLUDE_SPECIFIC"
        included_access_profile_ids = [
          "..."
        ]
      }
      days_since_added     = "...my_days_since_added..."
      days_since_last_used = "...my_days_since_last_used..."
      days_since_reviewed  = "...my_days_since_reviewed..."
      grants_added_between = {
        end_date   = "2020-06-03T02:04:33.937Z"
        start_date = "2022-05-12T19:26:17.592Z"
      }
      source_filter = "GRANT_SOURCE_FILTER_INHERITED"
      type_filter   = "GRANT_FILTER_TYPE_PERMANENT"
    }
    principal_type_filter = "PRINCIPAL_TYPE_FILTER_USERS"
    resource_selection = {
      # ...
    }
    resource_type_selections = {
      # ...
    }
    scope_role_selection = {
      # ...
    }
    selected_users = {
      user_ids = [
        "..."
      ]
    }
    specific_access_conflicts = {
      # ...
    }
    specific_resources = {
      # ...
    }
    user_criteria = {
      group_app_entitlements_ref = [
        {
          app_id = "...my_app_id..."
          id     = "...my_id..."
        }
      ]
      manager_user_ids = [
        "..."
      ]
      multi_user_profile_attributes = {
        key = {
          values = [
            {
              value = "...my_value..."
            }
          ]
        }
      }
      user_status = [
        "DISABLED"
      ]
    }
  }
  scope_type = "ACCESS_REVIEW_SCOPE_TYPE_UNSPECIFIED"
  signature_config = {
    meaning_of_signature = "...my_meaning_of_signature..."
    require_signature    = true
    step_up_provider_id  = "...my_step_up_provider_id..."
    tsp_url              = "...my_tsp_url..."
  }
  use_policy_override = true
}