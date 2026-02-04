resource "conductorone_access_review_template" "my_access_review_template" {
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
      # ...
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
  description  = "...my_description..."
  display_name = "...my_display_name..."
  notification_config = {
    send_close     = true
    send_kickoff   = false
    send_reminders = false
  }
  owner_ids = [
    "..."
  ]
  policy_id  = "...my_policy_id..."
  scope_type = "ACCESS_REVIEW_SCOPE_TYPE_UNSPECIFIED"
}