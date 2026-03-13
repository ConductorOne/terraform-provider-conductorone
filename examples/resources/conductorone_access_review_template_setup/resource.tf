resource "conductorone_access_review_template_setup" "my_access_review_template_setup" {
  access_review_scope_v2 = {
    account_criteria_scope = {
      account_domain = "APP_USER_DOMAIN_UNSPECIFIED"
      account_types = [
        "APP_USER_TYPE_SYSTEM_ACCOUNT"
      ]
      app_user_statuses = [
        "APP_USER_STATUS_DISABLED"
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
        filter_type = "ACCESS_PROFILE_FILTER_TYPE_EXCLUDE_ALL"
        included_access_profile_ids = [
          "..."
        ]
      }
      grants_added_between = {
        end_date   = "2022-01-25T08:02:55.357Z"
        start_date = "2022-09-13T06:49:21.124Z"
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
        "ENABLED"
      ]
    }
  }
  access_review_template_id = "...my_access_review_template_id..."
  access_review_template_setup_entitlement_expand_mask = {
    paths = [
      "..."
    ]
  }
  entitlements = [
    {
      app_entitlement_id = "...my_app_entitlement_id..."
      app_id             = "...my_app_id..."
    }
  ]
}