resource "conductorone_access_review_template" "my_access_review_template" {
  access_review_column_config = {
    columns = [
      "ACCESS_REVIEW_TASK_COLUMN_RESOURCE_CHILDREN"
    ]
  }
  access_review_duration = "...my_access_review_duration..."
  access_review_scope_v2 = {
    account_criteria_scope = {
      account_domain = "APP_USER_DOMAIN_UNSPECIFIED"
      account_types = [
        "APP_USER_TYPE_USER"
      ]
      app_user_statuses = [
        "APP_USER_STATUS_DELETED"
      ]
      no_account_owner = false
    }
    all_access_conflicts_scope = {
      # ...
    }
    all_accounts_scope = {
      # ...
    }
    all_grants_scope = {
      # ...
    }
    all_users_scope = {
      # ...
    }
    app_selection_criteria_scope = {
      compliance_framework_attribute_value_ids = [
        "..."
      ]
      risk_level_attribute_value_ids = [
        "..."
      ]
    }
    application_access_scope = {
      # ...
    }
    cel_expression_scope = {
      expression = "...my_expression..."
    }
    cel_expression_scope1 = {
      expression = "...my_expression..."
    }
    grants_by_criteria_scope = {
      days_since_added     = "...my_days_since_added..."
      days_since_last_used = "...my_days_since_last_used..."
      days_since_reviewed  = "...my_days_since_reviewed..."
      grant_access_profile_filter = {
        excluded_access_profile_ids = [
          "..."
        ]
        filter_type = "ACCESS_PROFILE_FILTER_TYPE_INCLUDE_ALL"
        included_access_profile_ids = [
          "..."
        ]
      }
      grants_added_between = {
        end_date   = "2022-10-31T16:44:21.565Z"
        start_date = "2022-03-24T07:54:08.302Z"
      }
      source_filter = "GRANT_SOURCE_FILTER_INHERITED"
      type_filter   = "GRANT_FILTER_TYPE_PERMANENT"
    }
    resource_selection_scope = {
      # ...
    }
    resource_type_selection_scope = {
      # ...
    }
    selected_users_scope = {
      user_ids = [
        "..."
      ]
    }
    specific_access_conflicts_scope = {
      # ...
    }
    specific_resources_scope = {
      # ...
    }
    user_criteria_scope = {
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
        "ENABLED"
      ]
    }
  }
  accuracy_issue_action             = "ACCURACY_ISSUE_ACTION_WAIT"
  auto_close_campaign               = false
  auto_close_decision               = "CLOSE_DECISION_NO_ACTION"
  auto_generate_report              = false
  auto_start_campaign               = true
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
  review_signature_config = {
    meaning_of_signature = "...my_meaning_of_signature..."
    require_signature    = true
    step_up_provider_id  = "...my_step_up_provider_id..."
    tsp_url              = "...my_tsp_url..."
  }
  scope_type          = "ACCESS_REVIEW_SCOPE_TYPE_UNSPECIFIED"
  use_policy_override = true
}