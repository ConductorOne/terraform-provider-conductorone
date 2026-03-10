resource "conductorone_access_review" "my_access_review" {
  access_review_scope_v2 = {
    account_criteria_scope = {
      account_domain = "APP_USER_DOMAIN_EXTERNAL"
      account_types = [
        "APP_USER_TYPE_SERVICE_ACCOUNT"
      ]
      app_user_statuses = [
        "APP_USER_STATUS_ENABLED"
      ]
      no_account_owner = true
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
        filter_type = "ACCESS_PROFILE_FILTER_TYPE_INCLUDE_SPECIFIC"
        included_access_profile_ids = [
          "..."
        ]
      }
      grants_added_between = {
        end_date   = "2022-07-05T14:23:47.903Z"
        start_date = "2021-11-22T05:25:02.885Z"
      }
      source_filter = "GRANT_SOURCE_FILTER_UNSPECIFIED"
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
        "DISABLED"
      ]
    }
  }
  completion_date = "2020-03-30T12:47:51.525Z"
  description     = "...my_description..."
  display_name    = "...my_display_name..."
  duplicate_from  = "...my_duplicate_from..."
  notification_config = {
    send_close     = true
    send_kickoff   = true
    send_reminders = true
  }
  owner_ids = [
    "..."
  ]
  policy_id  = "...my_policy_id..."
  scope_type = "ACCESS_REVIEW_SCOPE_TYPE_UNSPECIFIED"
}