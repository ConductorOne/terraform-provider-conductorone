resource "conductorone_automation" "my_automation" {
  app_id = "...my_app_id..."
  automation_context = {
    context = {
      # ...
    }
  }
  automation_steps = [
    {
      account_lifecycle_action = {
        account_in_context = {
          # ...
        }
        account_ref = {
          account_id_cel = "...my_account_id_cel..."
        }
        action_name = "...my_action_name..."
        connector_ref = {
          app_id = "...my_app_id..."
          id     = "...my_id..."
        }
      }
      call_function = {
        args = {
          key = "value"
        }
        function_id = "...my_function_id..."
      }
      connector_action = {
        action_name = "...my_action_name..."
        args_template = {
          # ...
        }
        connector_ref = {
          app_id = "...my_app_id..."
          id     = "...my_id..."
        }
      }
      connector_create_account = {
        connector_ref = {
          app_id = "...my_app_id..."
          id     = "...my_id..."
        }
        user_id_cel = "...my_user_id_cel..."
        user_properties = {
          display_name_cel      = "...my_display_name_cel..."
          email_cel             = "...my_email_cel..."
          profile_attribute_cel = "...my_profile_attribute_cel..."
          username_cel          = "...my_username_cel..."
        }
      }
      create_access_review = {
        access_review_template_cel = "...my_access_review_template_cel..."
        access_review_template_id  = "...my_access_review_template_id..."
        campaign_name              = "...my_campaign_name..."
        use_subject_user           = false
        user_ids_cel               = "...my_user_ids_cel..."
        user_refs = [
          {
            id = "...my_id..."
          }
        ]
      }
      create_revoke_tasks = {
        app_entitlement_refs = [
          {
            app_id = "...my_app_id..."
            id     = "...my_id..."
          }
        ]
        app_entitlement_refs_cel = "...my_app_entitlement_refs_cel..."
        excluded_app_entitlement_refs = [
          {
            app_id = "...my_app_id..."
            id     = "...my_id..."
          }
        ]
        excluded_app_entitlement_refs_cel = "...my_excluded_app_entitlement_refs_cel..."
        revoke_all                        = false
        use_subject_user                  = false
        user_id_cel                       = "...my_user_id_cel..."
        user_ref = {
          id = "...my_id..."
        }
      }
      create_revoke_tasks_v2 = {
        entitlement_exclusion_criteria = {
          excluded_app_ids = [
            "..."
          ]
          excluded_compliance_framework_ids = [
            "..."
          ]
          excluded_resource_type_ids = [
            "..."
          ]
          excluded_risk_level_ids = [
            "..."
          ]
        }
        entitlement_exclusion_list = {
          excluded_app_entitlement_refs = [
            {
              app_id = "...my_app_id..."
              id     = "...my_id..."
            }
          ]
        }
        entitlement_exclusion_list_cel = {
          excluded_app_entitlement_refs_cel = "...my_excluded_app_entitlement_refs_cel..."
        }
        entitlement_exclusion_none = {
          # ...
        }
        entitlement_inclusion_all = {
          # ...
        }
        entitlement_inclusion_criteria = {
          app_ids = [
            "..."
          ]
          compliance_framework_ids = [
            "..."
          ]
          resource_type_ids = [
            "..."
          ]
          risk_level_ids = [
            "..."
          ]
        }
        entitlement_inclusion_list = {
          app_entitlement_refs = [
            {
              app_id = "...my_app_id..."
              id     = "...my_id..."
            }
          ]
        }
        entitlement_inclusion_list_cel = {
          app_entitlement_refs_cel = "...my_app_entitlement_refs_cel..."
        }
        use_subject_user = false
        user_id_cel      = "...my_user_id_cel..."
        user_ref = {
          id = "...my_id..."
        }
      }
      evaluate_expressions = {
        expressions = [
          {
            expression_cel = "...my_expression_cel..."
            is_secret      = false
            key            = "...my_key..."
          }
        ]
      }
      generate_password = {
        # ...
      }
      grant_entitlements = {
        app_entitlement_refs = [
          {
            app_id = "...my_app_id..."
            id     = "...my_id..."
          }
        ]
        app_entitlement_refs_cel = "...my_app_entitlement_refs_cel..."
        use_subject_user         = true
        user_id_cel              = "...my_user_id_cel..."
        user_ref = {
          id = "...my_id..."
        }
      }
      remove_from_delegation = {
        replacement_user_id_cel = "...my_replacement_user_id_cel..."
        use_subject_user        = true
        user_id_cel             = "...my_user_id_cel..."
        user_ref = {
          id = "...my_id..."
        }
        user_ref1 = {
          id = "...my_id..."
        }
      }
      run_automation = {
        automation_context = {
          context = {
            # ...
          }
        }
        automation_template_id_cel = "...my_automation_template_id_cel..."
        automation_template_ref = {
          id = "...my_id..."
        }
      }
      send_email = {
        body             = "...my_body..."
        subject          = "...my_subject..."
        title            = "...my_title..."
        use_subject_user = true
        user_ids_cel     = "...my_user_ids_cel..."
        user_refs = [
          {
            id = "...my_id..."
          }
        ]
      }
      send_slack_message = {
        body             = "...my_body..."
        channel_name     = "...my_channel_name..."
        channel_name_cel = "...my_channel_name_cel..."
      }
      skip_if_true_cel  = "...my_skip_if_true_cel..."
      step_display_name = "...my_step_display_name..."
      step_name         = "...my_step_name..."
      task_action = {
        close_action = {
          use_subject_user = false
          user_id_cel      = "...my_user_id_cel..."
          user_ref = {
            id = "...my_id..."
          }
        }
        reassign_action = {
          assignee_user_id_cel = "...my_assignee_user_id_cel..."
          subject_user_id_cel  = "...my_subject_user_id_cel..."
          use_subject_user     = false
          user_ref = {
            id = "...my_id..."
          }
          user_ref1 = {
            id = "...my_id..."
          }
        }
        task_types = [
          "TASK_TYPE_REVOKE"
        ]
        task_user_relation = "TASK_USER_RELATION_ASSIGNEE"
      }
      unenroll_from_all_access_profiles = {
        catalog_ids = [
          "..."
        ]
        catalog_ids_cel  = "...my_catalog_ids_cel..."
        use_subject_user = false
        user_ids_cel     = "...my_user_ids_cel..."
        user_refs = [
          {
            id = "...my_id..."
          }
        ]
      }
      update_user = {
        use_subject_user = false
        user_id_cel      = "...my_user_id_cel..."
        user_ref = {
          id = "...my_id..."
        }
        user_status_cel  = "...my_user_status_cel..."
        user_status_enum = "DELETED"
      }
      wait_for_duration = {
        duration = "...my_duration..."
      }
      webhook = {
        payload = {
          # ...
        }
        webhook_id     = "...my_webhook_id..."
        webhook_id_cel = "...my_webhook_id_cel..."
      }
    }
  ]
  description  = "...my_description..."
  display_name = "...my_display_name..."
  draft_automation_steps = [
    {
      account_lifecycle_action = {
        account_in_context = {
          # ...
        }
        account_ref = {
          account_id_cel = "...my_account_id_cel..."
        }
        action_name = "...my_action_name..."
        connector_ref = {
          app_id = "...my_app_id..."
          id     = "...my_id..."
        }
      }
      call_function = {
        args = {
          key = "value"
        }
        function_id = "...my_function_id..."
      }
      connector_action = {
        action_name = "...my_action_name..."
        args_template = {
          # ...
        }
        connector_ref = {
          app_id = "...my_app_id..."
          id     = "...my_id..."
        }
      }
      connector_create_account = {
        connector_ref = {
          app_id = "...my_app_id..."
          id     = "...my_id..."
        }
        user_id_cel = "...my_user_id_cel..."
        user_properties = {
          display_name_cel      = "...my_display_name_cel..."
          email_cel             = "...my_email_cel..."
          profile_attribute_cel = "...my_profile_attribute_cel..."
          username_cel          = "...my_username_cel..."
        }
      }
      create_access_review = {
        access_review_template_cel = "...my_access_review_template_cel..."
        access_review_template_id  = "...my_access_review_template_id..."
        campaign_name              = "...my_campaign_name..."
        use_subject_user           = false
        user_ids_cel               = "...my_user_ids_cel..."
        user_refs = [
          {
            id = "...my_id..."
          }
        ]
      }
      create_revoke_tasks = {
        app_entitlement_refs = [
          {
            app_id = "...my_app_id..."
            id     = "...my_id..."
          }
        ]
        app_entitlement_refs_cel = "...my_app_entitlement_refs_cel..."
        excluded_app_entitlement_refs = [
          {
            app_id = "...my_app_id..."
            id     = "...my_id..."
          }
        ]
        excluded_app_entitlement_refs_cel = "...my_excluded_app_entitlement_refs_cel..."
        revoke_all                        = true
        use_subject_user                  = true
        user_id_cel                       = "...my_user_id_cel..."
        user_ref = {
          id = "...my_id..."
        }
      }
      create_revoke_tasks_v2 = {
        entitlement_exclusion_criteria = {
          excluded_app_ids = [
            "..."
          ]
          excluded_compliance_framework_ids = [
            "..."
          ]
          excluded_resource_type_ids = [
            "..."
          ]
          excluded_risk_level_ids = [
            "..."
          ]
        }
        entitlement_exclusion_list = {
          excluded_app_entitlement_refs = [
            {
              app_id = "...my_app_id..."
              id     = "...my_id..."
            }
          ]
        }
        entitlement_exclusion_list_cel = {
          excluded_app_entitlement_refs_cel = "...my_excluded_app_entitlement_refs_cel..."
        }
        entitlement_exclusion_none = {
          # ...
        }
        entitlement_inclusion_all = {
          # ...
        }
        entitlement_inclusion_criteria = {
          app_ids = [
            "..."
          ]
          compliance_framework_ids = [
            "..."
          ]
          resource_type_ids = [
            "..."
          ]
          risk_level_ids = [
            "..."
          ]
        }
        entitlement_inclusion_list = {
          app_entitlement_refs = [
            {
              app_id = "...my_app_id..."
              id     = "...my_id..."
            }
          ]
        }
        entitlement_inclusion_list_cel = {
          app_entitlement_refs_cel = "...my_app_entitlement_refs_cel..."
        }
        use_subject_user = false
        user_id_cel      = "...my_user_id_cel..."
        user_ref = {
          id = "...my_id..."
        }
      }
      evaluate_expressions = {
        expressions = [
          {
            expression_cel = "...my_expression_cel..."
            is_secret      = true
            key            = "...my_key..."
          }
        ]
      }
      generate_password = {
        # ...
      }
      grant_entitlements = {
        app_entitlement_refs = [
          {
            app_id = "...my_app_id..."
            id     = "...my_id..."
          }
        ]
        app_entitlement_refs_cel = "...my_app_entitlement_refs_cel..."
        use_subject_user         = true
        user_id_cel              = "...my_user_id_cel..."
        user_ref = {
          id = "...my_id..."
        }
      }
      remove_from_delegation = {
        replacement_user_id_cel = "...my_replacement_user_id_cel..."
        use_subject_user        = true
        user_id_cel             = "...my_user_id_cel..."
        user_ref = {
          id = "...my_id..."
        }
        user_ref1 = {
          id = "...my_id..."
        }
      }
      run_automation = {
        automation_context = {
          context = {
            # ...
          }
        }
        automation_template_id_cel = "...my_automation_template_id_cel..."
        automation_template_ref = {
          id = "...my_id..."
        }
      }
      send_email = {
        body             = "...my_body..."
        subject          = "...my_subject..."
        title            = "...my_title..."
        use_subject_user = true
        user_ids_cel     = "...my_user_ids_cel..."
        user_refs = [
          {
            id = "...my_id..."
          }
        ]
      }
      send_slack_message = {
        body             = "...my_body..."
        channel_name     = "...my_channel_name..."
        channel_name_cel = "...my_channel_name_cel..."
      }
      skip_if_true_cel  = "...my_skip_if_true_cel..."
      step_display_name = "...my_step_display_name..."
      step_name         = "...my_step_name..."
      task_action = {
        close_action = {
          use_subject_user = false
          user_id_cel      = "...my_user_id_cel..."
          user_ref = {
            id = "...my_id..."
          }
        }
        reassign_action = {
          assignee_user_id_cel = "...my_assignee_user_id_cel..."
          subject_user_id_cel  = "...my_subject_user_id_cel..."
          use_subject_user     = false
          user_ref = {
            id = "...my_id..."
          }
          user_ref1 = {
            id = "...my_id..."
          }
        }
        task_types = [
          "TASK_TYPE_UNSPECIFIED"
        ]
        task_user_relation = "TASK_USER_RELATION_UNSPECIFIED"
      }
      unenroll_from_all_access_profiles = {
        catalog_ids = [
          "..."
        ]
        catalog_ids_cel  = "...my_catalog_ids_cel..."
        use_subject_user = true
        user_ids_cel     = "...my_user_ids_cel..."
        user_refs = [
          {
            id = "...my_id..."
          }
        ]
      }
      update_user = {
        use_subject_user = false
        user_id_cel      = "...my_user_id_cel..."
        user_ref = {
          id = "...my_id..."
        }
        user_status_cel  = "...my_user_status_cel..."
        user_status_enum = "UNKNOWN"
      }
      wait_for_duration = {
        duration = "...my_duration..."
      }
      webhook = {
        payload = {
          # ...
        }
        webhook_id     = "...my_webhook_id..."
        webhook_id_cel = "...my_webhook_id_cel..."
      }
    }
  ]
  draft_triggers = [
    {
      access_conflict_trigger = {
        all_conflict_monitors = true
        conflict_monitor_refs = {
          conflict_monitor_refs = [
            {
              id = "...my_id..."
            }
          ]
        }
      }
      app_user_created_trigger = {
        app_id     = "...my_app_id..."
        app_id_cel = "...my_app_id_cel..."
        condition  = "...my_condition..."
      }
      app_user_updated_trigger = {
        app_id     = "...my_app_id..."
        app_id_cel = "...my_app_id_cel..."
        condition  = "...my_condition..."
      }
      form_trigger = "{ \"see\": \"documentation\" }"
      grant_deleted_trigger = {
        grant_trigger_filter = {
          account_filter = {
            account_type = "APP_USER_TYPE_UNSPECIFIED"
          }
          entitlement_filter = {
            app_entitlement_refs = [
              {
                app_id = "...my_app_id..."
                id     = "...my_id..."
              }
            ]
            app_entitlement_refs_cel = "...my_app_entitlement_refs_cel..."
            app_id                   = "...my_app_id..."
          }
          grant_filter = {
            grant_filter_type        = "GRANT_FILTER_TYPE_TEMPORARY"
            grant_justification_type = "GRANT_JUSTIFICATION_TYPE_UNSPECIFIED"
            grant_source_filter      = "GRANT_SOURCE_FILTER_DIRECT"
          }
        }
      }
      grant_found_trigger = {
        grant_trigger_filter = {
          account_filter = {
            account_type = "APP_USER_TYPE_UNSPECIFIED"
          }
          entitlement_filter = {
            app_entitlement_refs = [
              {
                app_id = "...my_app_id..."
                id     = "...my_id..."
              }
            ]
            app_entitlement_refs_cel = "...my_app_entitlement_refs_cel..."
            app_id                   = "...my_app_id..."
          }
          grant_filter = {
            grant_filter_type        = "GRANT_FILTER_TYPE_TEMPORARY"
            grant_justification_type = "GRANT_JUSTIFICATION_TYPE_DIRECT"
            grant_source_filter      = "GRANT_SOURCE_FILTER_INHERITED"
          }
        }
      }
      manual_automation_trigger = {
        # ...
      }
      schedule_trigger = {
        advanced         = false
        condition        = "...my_condition..."
        cron_spec        = "...my_cron_spec..."
        skip_if_true_cel = "...my_skip_if_true_cel..."
        start            = "2022-11-05T06:51:06.048Z"
        timezone         = "...my_timezone..."
      }
      schedule_trigger_app_user = {
        app_id    = "...my_app_id..."
        condition = "...my_condition..."
        cron_spec = "...my_cron_spec..."
        start     = "2022-01-16T12:05:57.834Z"
        timezone  = "...my_timezone..."
      }
      usage_based_revocation_trigger = {
        app_id     = "...my_app_id..."
        enabled_at = "2021-08-03T22:56:50.306Z"
        excluded_group_refs = [
          {
            app_id = "...my_app_id..."
            id     = "...my_id..."
          }
        ]
        excluded_user_refs = [
          {
            id = "...my_id..."
          }
        ]
        include_users_with_no_activity = false
        run_delayed = {
          cold_start_delay_days = 8
        }
        run_immediately = {
          # ...
        }
        targeted_app_user_types = [
          "APP_USER_TYPE_UNSPECIFIED"
        ]
        targeted_entitlement_refs = [
          {
            app_id = "...my_app_id..."
            id     = "...my_id..."
          }
        ]
        unused_for_days = 8
      }
      user_created_trigger = {
        condition = "...my_condition..."
      }
      user_profile_change_trigger = {
        condition = "...my_condition..."
      }
      webhook_automation_trigger = {
        listener_id = "...my_listener_id..."
        webhook_listener_auth_hmac = {
          # ...
        }
        webhook_listener_auth_jwt = {
          jwks_url = "...my_jwks_url..."
        }
      }
    }
  ]
  enabled  = false
  is_draft = true
  triggers = [
    {
      access_conflict_trigger = {
        all_conflict_monitors = false
        conflict_monitor_refs = {
          conflict_monitor_refs = [
            {
              id = "...my_id..."
            }
          ]
        }
      }
      app_user_created_trigger = {
        app_id     = "...my_app_id..."
        app_id_cel = "...my_app_id_cel..."
        condition  = "...my_condition..."
      }
      app_user_updated_trigger = {
        app_id     = "...my_app_id..."
        app_id_cel = "...my_app_id_cel..."
        condition  = "...my_condition..."
      }
      form_trigger = "{ \"see\": \"documentation\" }"
      grant_deleted_trigger = {
        grant_trigger_filter = {
          account_filter = {
            account_type = "APP_USER_TYPE_SERVICE_ACCOUNT"
          }
          entitlement_filter = {
            app_entitlement_refs = [
              {
                app_id = "...my_app_id..."
                id     = "...my_id..."
              }
            ]
            app_entitlement_refs_cel = "...my_app_entitlement_refs_cel..."
            app_id                   = "...my_app_id..."
          }
          grant_filter = {
            grant_filter_type        = "GRANT_FILTER_TYPE_PERMANENT"
            grant_justification_type = "GRANT_JUSTIFICATION_TYPE_CONDUCTOR_ONE"
            grant_source_filter      = "GRANT_SOURCE_FILTER_INHERITED"
          }
        }
      }
      grant_found_trigger = {
        grant_trigger_filter = {
          account_filter = {
            account_type = "APP_USER_TYPE_UNSPECIFIED"
          }
          entitlement_filter = {
            app_entitlement_refs = [
              {
                app_id = "...my_app_id..."
                id     = "...my_id..."
              }
            ]
            app_entitlement_refs_cel = "...my_app_entitlement_refs_cel..."
            app_id                   = "...my_app_id..."
          }
          grant_filter = {
            grant_filter_type        = "GRANT_FILTER_TYPE_PERMANENT"
            grant_justification_type = "GRANT_JUSTIFICATION_TYPE_UNSPECIFIED"
            grant_source_filter      = "GRANT_SOURCE_FILTER_INHERITED"
          }
        }
      }
      manual_automation_trigger = {
        # ...
      }
      schedule_trigger = {
        advanced         = false
        condition        = "...my_condition..."
        cron_spec        = "...my_cron_spec..."
        skip_if_true_cel = "...my_skip_if_true_cel..."
        start            = "2022-01-25T09:55:20.150Z"
        timezone         = "...my_timezone..."
      }
      schedule_trigger_app_user = {
        app_id    = "...my_app_id..."
        condition = "...my_condition..."
        cron_spec = "...my_cron_spec..."
        start     = "2021-11-11T16:22:31.994Z"
        timezone  = "...my_timezone..."
      }
      usage_based_revocation_trigger = {
        app_id     = "...my_app_id..."
        enabled_at = "2022-01-14T07:38:09.864Z"
        excluded_group_refs = [
          {
            app_id = "...my_app_id..."
            id     = "...my_id..."
          }
        ]
        excluded_user_refs = [
          {
            id = "...my_id..."
          }
        ]
        include_users_with_no_activity = false
        run_delayed = {
          cold_start_delay_days = 1
        }
        run_immediately = {
          # ...
        }
        targeted_app_user_types = [
          "APP_USER_TYPE_USER"
        ]
        targeted_entitlement_refs = [
          {
            app_id = "...my_app_id..."
            id     = "...my_id..."
          }
        ]
        unused_for_days = 0
      }
      user_created_trigger = {
        condition = "...my_condition..."
      }
      user_profile_change_trigger = {
        condition = "...my_condition..."
      }
      webhook_automation_trigger = {
        listener_id = "...my_listener_id..."
        webhook_listener_auth_hmac = {
          # ...
        }
        webhook_listener_auth_jwt = {
          jwks_url = "...my_jwks_url..."
        }
      }
    }
  ]
}