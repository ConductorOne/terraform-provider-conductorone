resource "conductorone_automation" "my_automation" {
  annotations = {
    key = "value"
  }
  app_id = "...my_app_id..."
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
        resource_type_id = "...my_resource_type_id..."
      }
      connector_create_account = {
        connector_ref = {
          app_id = "...my_app_id..."
          id     = "...my_id..."
        }
        password_cel = "...my_password_cel..."
        user_id_cel  = "...my_user_id_cel..."
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
        exclusion_criteria = {
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
        exclusion_list = {
          excluded_app_entitlement_refs = [
            {
              app_id = "...my_app_id..."
              id     = "...my_id..."
            }
          ]
        }
        exclusion_list_cel = {
          excluded_app_entitlement_refs_cel = "...my_excluded_app_entitlement_refs_cel..."
        }
        exclusion_none = {
          # ...
        }
        grant_source_filter = "GRANT_SOURCE_FILTER_UNSPECIFIED"
        inclusion_access_only = {
          # ...
        }
        inclusion_all = {
          # ...
        }
        inclusion_criteria = {
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
        inclusion_list = {
          app_entitlement_refs = [
            {
              app_id = "...my_app_id..."
              id     = "...my_id..."
            }
          ]
        }
        inclusion_list_cel = {
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
        password_policy_id = "...my_password_policy_id..."
        policy = {
          custom_characters          = "...my_custom_characters..."
          excluded_characters        = "...my_excluded_characters..."
          max_character_count        = 8
          min_character_count        = 9
          no_restrictions            = true
          require_lowercase          = false
          require_numbers            = true
          require_special_characters = true
          require_uppercase          = true
        }
      }
      grant_entitlements = {
        exclusion_criteria = {
          excluded_app_ids = [
            "..."
          ]
          excluded_compliance_framework_ids = [
            "..."
          ]
          excluded_risk_level_ids = [
            "..."
          ]
        }
        exclusion_list = {
          excluded_app_entitlement_refs = [
            {
              app_id = "...my_app_id..."
              id     = "...my_id..."
            }
          ]
        }
        exclusion_list_cel = {
          excluded_app_entitlement_refs_cel = "...my_excluded_app_entitlement_refs_cel..."
        }
        exclusion_none = {
          # ...
        }
        inclusion_criteria = {
          app_ids = [
            "..."
          ]
          compliance_framework_ids = [
            "..."
          ]
          risk_level_ids = [
            "..."
          ]
        }
        inclusion_list = {
          app_entitlement_refs = [
            {
              app_id = "...my_app_id..."
              id     = "...my_id..."
            }
          ]
        }
        inclusion_list_cel = {
          app_entitlement_refs_cel = "...my_app_entitlement_refs_cel..."
        }
        use_subject_user = true
        user_id_cel      = "...my_user_id_cel..."
        user_ref = {
          id = "...my_id..."
        }
      }
      remove_from_delegation = {
        replacement_user_id_cel = "...my_replacement_user_id_cel..."
        replacement_user_ref = {
          id = "...my_id..."
        }
        use_subject_user = true
        user_id_cel      = "...my_user_id_cel..."
        user_ref = {
          id = "...my_id..."
        }
      }
      run_automation = {
        automation_template_id_cel = "...my_automation_template_id_cel..."
        automation_template_ref = {
          id = "...my_id..."
        }
        context = {
          context = {
            # ...
          }
        }
      }
      send_email = {
        body             = "...my_body..."
        email            = "...my_email..."
        email_cel        = "...my_email_cel..."
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
        channel_is_id    = true
        channel_name     = "...my_channel_name..."
        channel_name_cel = "...my_channel_name_cel..."
        use_subject_user = false
        user_ids_cel     = "...my_user_ids_cel..."
        user_refs = [
          {
            id = "...my_id..."
          }
        ]
      }
      set_credential = {
        account_id_cel = "...my_account_id_cel..."
        connector_ref = {
          app_id = "...my_app_id..."
          id     = "...my_id..."
        }
        password_cel = "...my_password_cel..."
      }
      skip_if_true_cel  = "...my_skip_if_true_cel..."
      step_display_name = "...my_step_display_name..."
      step_name         = "...my_step_name..."
      store_credential = {
        app_id_cel          = "...my_app_id_cel..."
        auth_type           = "STORE_CREDENTIAL_AUTH_TYPE_VERIFY_EMAIL"
        credential_cel      = "...my_credential_cel..."
        expiry              = "...my_expiry..."
        label_cel           = "...my_label_cel..."
        max_views           = 9
        recipient_cel       = "...my_recipient_cel..."
        recipient_email_cel = "...my_recipient_email_cel..."
        ttl                 = "...my_ttl..."
        vault_type          = "STORE_CREDENTIAL_VAULT_TYPE_UNSPECIFIED"
      }
      task_action = {
        close = {
          use_subject_user = false
          user_id_cel      = "...my_user_id_cel..."
          user_ref = {
            id = "...my_id..."
          }
        }
        reassign = {
          assignee_user_id_cel = "...my_assignee_user_id_cel..."
          assignee_user_ref = {
            id = "...my_id..."
          }
          subject_user_id_cel = "...my_subject_user_id_cel..."
          subject_user_ref = {
            id = "...my_id..."
          }
          use_subject_user = false
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
  automations_delete_automation_request = {
    # ...
  }
  circuit_breaker_max    = 3
  circuit_breaker_period = "CIRCUIT_BREAKER_PERIOD_MONTH"
  context = {
    context = {
      # ...
    }
  }
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
        resource_type_id = "...my_resource_type_id..."
      }
      connector_create_account = {
        connector_ref = {
          app_id = "...my_app_id..."
          id     = "...my_id..."
        }
        password_cel = "...my_password_cel..."
        user_id_cel  = "...my_user_id_cel..."
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
        exclusion_criteria = {
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
        exclusion_list = {
          excluded_app_entitlement_refs = [
            {
              app_id = "...my_app_id..."
              id     = "...my_id..."
            }
          ]
        }
        exclusion_list_cel = {
          excluded_app_entitlement_refs_cel = "...my_excluded_app_entitlement_refs_cel..."
        }
        exclusion_none = {
          # ...
        }
        grant_source_filter = "GRANT_SOURCE_FILTER_UNSPECIFIED"
        inclusion_access_only = {
          # ...
        }
        inclusion_all = {
          # ...
        }
        inclusion_criteria = {
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
        inclusion_list = {
          app_entitlement_refs = [
            {
              app_id = "...my_app_id..."
              id     = "...my_id..."
            }
          ]
        }
        inclusion_list_cel = {
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
        password_policy_id = "...my_password_policy_id..."
        policy = {
          custom_characters          = "...my_custom_characters..."
          excluded_characters        = "...my_excluded_characters..."
          max_character_count        = 6
          min_character_count        = 6
          no_restrictions            = false
          require_lowercase          = true
          require_numbers            = true
          require_special_characters = true
          require_uppercase          = false
        }
      }
      grant_entitlements = {
        exclusion_criteria = {
          excluded_app_ids = [
            "..."
          ]
          excluded_compliance_framework_ids = [
            "..."
          ]
          excluded_risk_level_ids = [
            "..."
          ]
        }
        exclusion_list = {
          excluded_app_entitlement_refs = [
            {
              app_id = "...my_app_id..."
              id     = "...my_id..."
            }
          ]
        }
        exclusion_list_cel = {
          excluded_app_entitlement_refs_cel = "...my_excluded_app_entitlement_refs_cel..."
        }
        exclusion_none = {
          # ...
        }
        inclusion_criteria = {
          app_ids = [
            "..."
          ]
          compliance_framework_ids = [
            "..."
          ]
          risk_level_ids = [
            "..."
          ]
        }
        inclusion_list = {
          app_entitlement_refs = [
            {
              app_id = "...my_app_id..."
              id     = "...my_id..."
            }
          ]
        }
        inclusion_list_cel = {
          app_entitlement_refs_cel = "...my_app_entitlement_refs_cel..."
        }
        use_subject_user = true
        user_id_cel      = "...my_user_id_cel..."
        user_ref = {
          id = "...my_id..."
        }
      }
      remove_from_delegation = {
        replacement_user_id_cel = "...my_replacement_user_id_cel..."
        replacement_user_ref = {
          id = "...my_id..."
        }
        use_subject_user = true
        user_id_cel      = "...my_user_id_cel..."
        user_ref = {
          id = "...my_id..."
        }
      }
      run_automation = {
        automation_template_id_cel = "...my_automation_template_id_cel..."
        automation_template_ref = {
          id = "...my_id..."
        }
        context = {
          context = {
            # ...
          }
        }
      }
      send_email = {
        body             = "...my_body..."
        email            = "...my_email..."
        email_cel        = "...my_email_cel..."
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
        channel_is_id    = true
        channel_name     = "...my_channel_name..."
        channel_name_cel = "...my_channel_name_cel..."
        use_subject_user = true
        user_ids_cel     = "...my_user_ids_cel..."
        user_refs = [
          {
            id = "...my_id..."
          }
        ]
      }
      set_credential = {
        account_id_cel = "...my_account_id_cel..."
        connector_ref = {
          app_id = "...my_app_id..."
          id     = "...my_id..."
        }
        password_cel = "...my_password_cel..."
      }
      skip_if_true_cel  = "...my_skip_if_true_cel..."
      step_display_name = "...my_step_display_name..."
      step_name         = "...my_step_name..."
      store_credential = {
        app_id_cel          = "...my_app_id_cel..."
        auth_type           = "STORE_CREDENTIAL_AUTH_TYPE_VERIFY_EMAIL"
        credential_cel      = "...my_credential_cel..."
        expiry              = "...my_expiry..."
        label_cel           = "...my_label_cel..."
        max_views           = 2
        recipient_cel       = "...my_recipient_cel..."
        recipient_email_cel = "...my_recipient_email_cel..."
        ttl                 = "...my_ttl..."
        vault_type          = "STORE_CREDENTIAL_VAULT_TYPE_PAPER_VAULT"
      }
      task_action = {
        close = {
          use_subject_user = false
          user_id_cel      = "...my_user_id_cel..."
          user_ref = {
            id = "...my_id..."
          }
        }
        reassign = {
          assignee_user_id_cel = "...my_assignee_user_id_cel..."
          assignee_user_ref = {
            id = "...my_id..."
          }
          subject_user_id_cel = "...my_subject_user_id_cel..."
          subject_user_ref = {
            id = "...my_id..."
          }
          use_subject_user = true
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
      access_conflict = {
        all_conflict_monitors = true
        conflict_monitor_refs = {
          conflict_monitor_refs = [
            {
              id = "...my_id..."
            }
          ]
        }
      }
      app_user_created = {
        app_id     = "...my_app_id..."
        app_id_cel = "...my_app_id_cel..."
        condition  = "...my_condition..."
      }
      app_user_updated = {
        app_id     = "...my_app_id..."
        app_id_cel = "...my_app_id_cel..."
        condition  = "...my_condition..."
      }
      grant_deleted = {
        grant_trigger_filter = {
          account_filter = {
            account_type = "APP_USER_TYPE_UNSPECIFIED"
          }
          grant_filter = {
            grant_filter_type        = "GRANT_FILTER_TYPE_PERMANENT"
            grant_justification_type = "GRANT_JUSTIFICATION_TYPE_CONDUCTOR_ONE"
            grant_source_filter      = "GRANT_SOURCE_FILTER_INHERITED"
          }
          inclusion_all = {
            # ...
          }
          inclusion_criteria = {
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
          inclusion_list = {
            app_entitlement_refs = [
              {
                app_id = "...my_app_id..."
                id     = "...my_id..."
              }
            ]
          }
          inclusion_list_cel = {
            app_entitlement_refs_cel = "...my_app_entitlement_refs_cel..."
          }
        }
      }
      grant_found = {
        grant_trigger_filter = {
          account_filter = {
            account_type = "APP_USER_TYPE_UNSPECIFIED"
          }
          grant_filter = {
            grant_filter_type        = "GRANT_FILTER_TYPE_TEMPORARY"
            grant_justification_type = "GRANT_JUSTIFICATION_TYPE_CONDUCTOR_ONE"
            grant_source_filter      = "GRANT_SOURCE_FILTER_INHERITED"
          }
          inclusion_all = {
            # ...
          }
          inclusion_criteria = {
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
          inclusion_list = {
            app_entitlement_refs = [
              {
                app_id = "...my_app_id..."
                id     = "...my_id..."
              }
            ]
          }
          inclusion_list_cel = {
            app_entitlement_refs_cel = "...my_app_entitlement_refs_cel..."
          }
        }
      }
      schedule = {
        advanced         = true
        condition        = "...my_condition..."
        cron_spec        = "...my_cron_spec..."
        skip_if_true_cel = "...my_skip_if_true_cel..."
        start            = "2021-04-30T19:00:50.170Z"
        timezone         = "...my_timezone..."
      }
      schedule_app_user = {
        app_id    = "...my_app_id..."
        condition = "...my_condition..."
        cron_spec = "...my_cron_spec..."
        start     = "2022-08-16T18:01:29.335Z"
        timezone  = "...my_timezone..."
      }
      schedule_no_user = {
        advanced  = false
        cron_spec = "...my_cron_spec..."
        start     = "2022-05-09T07:31:59.841Z"
        timezone  = "...my_timezone..."
      }
      usage_based_revocation = {
        app_id     = "...my_app_id..."
        enabled_at = "2021-05-11T12:49:25.031Z"
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
        include_users_with_no_activity = true
        run_delayed = {
          cold_start_delay_days = 9
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
        unused_for_days = 1
      }
      user_created = {
        condition = "...my_condition..."
      }
      user_profile_change = {
        condition = "...my_condition..."
      }
      webhook = {
        capability_url = {
          # ...
        }
        hmac = {
          # ...
        }
        jwt = {
          jwks_url = "...my_jwks_url..."
        }
        listener_id = "...my_listener_id..."
      }
    }
  ]
  enabled  = false
  is_draft = true
  triggers = [
    {
      access_conflict = {
        all_conflict_monitors = false
        conflict_monitor_refs = {
          conflict_monitor_refs = [
            {
              id = "...my_id..."
            }
          ]
        }
      }
      app_user_created = {
        app_id     = "...my_app_id..."
        app_id_cel = "...my_app_id_cel..."
        condition  = "...my_condition..."
      }
      app_user_updated = {
        app_id     = "...my_app_id..."
        app_id_cel = "...my_app_id_cel..."
        condition  = "...my_condition..."
      }
      grant_deleted = {
        grant_trigger_filter = {
          account_filter = {
            account_type = "APP_USER_TYPE_SERVICE_ACCOUNT"
          }
          grant_filter = {
            grant_filter_type        = "GRANT_FILTER_TYPE_UNSPECIFIED"
            grant_justification_type = "GRANT_JUSTIFICATION_TYPE_DIRECT"
            grant_source_filter      = "GRANT_SOURCE_FILTER_INHERITED"
          }
          inclusion_all = {
            # ...
          }
          inclusion_criteria = {
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
          inclusion_list = {
            app_entitlement_refs = [
              {
                app_id = "...my_app_id..."
                id     = "...my_id..."
              }
            ]
          }
          inclusion_list_cel = {
            app_entitlement_refs_cel = "...my_app_entitlement_refs_cel..."
          }
        }
      }
      grant_found = {
        grant_trigger_filter = {
          account_filter = {
            account_type = "APP_USER_TYPE_SERVICE_ACCOUNT"
          }
          grant_filter = {
            grant_filter_type        = "GRANT_FILTER_TYPE_TEMPORARY"
            grant_justification_type = "GRANT_JUSTIFICATION_TYPE_CONDUCTOR_ONE"
            grant_source_filter      = "GRANT_SOURCE_FILTER_INHERITED"
          }
          inclusion_all = {
            # ...
          }
          inclusion_criteria = {
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
          inclusion_list = {
            app_entitlement_refs = [
              {
                app_id = "...my_app_id..."
                id     = "...my_id..."
              }
            ]
          }
          inclusion_list_cel = {
            app_entitlement_refs_cel = "...my_app_entitlement_refs_cel..."
          }
        }
      }
      schedule = {
        advanced         = true
        condition        = "...my_condition..."
        cron_spec        = "...my_cron_spec..."
        skip_if_true_cel = "...my_skip_if_true_cel..."
        start            = "2022-09-11T15:40:46.002Z"
        timezone         = "...my_timezone..."
      }
      schedule_app_user = {
        app_id    = "...my_app_id..."
        condition = "...my_condition..."
        cron_spec = "...my_cron_spec..."
        start     = "2021-01-14T19:18:02.966Z"
        timezone  = "...my_timezone..."
      }
      schedule_no_user = {
        advanced  = false
        cron_spec = "...my_cron_spec..."
        start     = "2022-01-24T14:53:52.603Z"
        timezone  = "...my_timezone..."
      }
      usage_based_revocation = {
        app_id     = "...my_app_id..."
        enabled_at = "2022-08-29T08:08:12.512Z"
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
        include_users_with_no_activity = true
        run_delayed = {
          cold_start_delay_days = 0
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
      user_created = {
        condition = "...my_condition..."
      }
      user_profile_change = {
        condition = "...my_condition..."
      }
      webhook = {
        capability_url = {
          # ...
        }
        hmac = {
          # ...
        }
        jwt = {
          jwks_url = "...my_jwks_url..."
        }
        listener_id = "...my_listener_id..."
      }
    }
  ]
}