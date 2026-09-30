resource "conductorone_access_review_setup" "my_access_review_setup" {
  access_review_id = "...my_access_review_id..."
  entitlements = [
    {
      app_entitlement_id = "...my_app_entitlement_id..."
      app_id             = "...my_app_id..."
    }
  ]
  expand_mask = {
    paths = [
      "..."
    ]
  }
  scope_v2 = {
    account_cel_expression = {
      expression = "...my_expression..."
    }
    account_criteria = {
      account_domain = "APP_USER_DOMAIN_UNSPECIFIED"
      account_types = [
        "APP_USER_TYPE_USER"
      ]
      app_user_statuses = [
        "APP_USER_STATUS_DISABLED"
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
        filter_type = "ACCESS_PROFILE_FILTER_TYPE_EXCLUDE_ALL"
        included_access_profile_ids = [
          "..."
        ]
      }
      days_since_added     = "...my_days_since_added..."
      days_since_last_used = "...my_days_since_last_used..."
      days_since_reviewed  = "...my_days_since_reviewed..."
      grants_added_between = {
        end_date   = "2022-09-19T00:56:10.984Z"
        start_date = "2021-05-13T23:08:00.233Z"
      }
      source_filter = "GRANT_SOURCE_FILTER_DIRECT"
      type_filter   = "GRANT_FILTER_TYPE_UNSPECIFIED"
    }
    principal_type_filter = "PRINCIPAL_TYPE_FILTER_RESOURCES"
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
}