resource "conductorone_request_schema" "my_request_schema" {
  description = "...my_description..."
  field_groups = [
    {
      default      = false
      display_name = "...my_display_name..."
      fields = [
        "..."
      ]
      help_text = "...my_help_text..."
      name      = "...my_name..."
    }
  ]
  field_relationships = [
    {
      at_least_one = {
        # ...
      }
      dependent_on = {
        dependency_field_names = [
          "..."
        ]
      }
      field_names = [
        "..."
      ]
      mutually_exclusive = {
        # ...
      }
      required_together = {
        # ...
      }
    }
  ]
  fields = [
    {
      admin_provider_config = {
        default_value_cel = "...my_default_value_cel..."
        show_to_user      = false
      }
      bool_field = {
        bool_rules = {
          const = false
        }
        checkbox_field = {
          # ...
        }
        default_value = true
        toggle_field = {
          # ...
        }
      }
      description  = "...my_description..."
      display_name = "...my_display_name..."
      file_field = {
        accepted_file_types = [
          "..."
        ]
        file_input_field = {
          # ...
        }
        max_file_size = "...my_max_file_size..."
      }
      int64_field = {
        default_value = "...my_default_value..."
        int64_rules = {
          const        = "...my_const..."
          gt           = "...my_gt..."
          gte          = "...my_gte..."
          ignore_empty = true
          in = [
            "..."
          ]
          lt  = "...my_lt..."
          lte = "...my_lte..."
          not_in = [
            "..."
          ]
        }
        number_field = {
          max_value = "...my_max_value..."
          min_value = "...my_min_value..."
          step      = "...my_step..."
        }
        placeholder = "...my_placeholder..."
      }
      name = "...my_name..."
      oauth2_field = {
        oauth2_field_view = {
          # ...
        }
      }
      required = false
      shared_provider_config = {
        default_value_cel        = "...my_default_value_cel..."
        input_transformation_cel = "...my_input_transformation_cel..."
        lock_default_values      = false
      }
      string_field = {
        default_value = "...my_default_value..."
        password_field = {
          # ...
        }
        picker_field = {
          app_resource_filter = {
            app_id           = "...my_app_id..."
            resource_type_id = "...my_resource_type_id..."
          }
          app_user_filter = {
            app_id = "...my_app_id..."
          }
          c1_user_filter = {
            # ...
          }
        }
        placeholder = "...my_placeholder..."
        select_field = {
          options = [
            {
              description  = "...my_description..."
              display_name = "...my_display_name..."
              value        = "...my_value..."
            }
          ]
          type = "SELECT_TYPE_RADIO"
        }
        string_rules = {
          address      = false
          const        = "...my_const..."
          contains     = "...my_contains..."
          email        = true
          hostname     = false
          ignore_empty = false
          in = [
            "..."
          ]
          ip           = true
          ipv4         = false
          ipv6         = true
          len_bytes    = "...my_len_bytes..."
          length       = "...my_length..."
          max_bytes    = "...my_max_bytes..."
          max_len      = "...my_max_len..."
          min_bytes    = "...my_min_bytes..."
          min_len      = "...my_min_len..."
          not_contains = "...my_not_contains..."
          not_in = [
            "..."
          ]
          pattern          = "...my_pattern..."
          prefix           = "...my_prefix..."
          strict           = true
          suffix           = "...my_suffix..."
          uri              = false
          uri_ref          = true
          uuid             = false
          well_known_regex = "HTTP_HEADER_NAME"
        }
        text_field = {
          multiline = false
          suffix    = "...my_suffix..."
        }
      }
      string_map_field = {
        default_value = {
          key = "value"
        }
        string_map_rules = {
          is_required    = false
          validate_empty = false
        }
      }
      user_provider_config = {
        input_transformation_cel = "...my_input_transformation_cel..."
      }
    }
  ]
  justification_visibility = "JUSTIFICATION_VISIBILITY_HIDE"
  name                     = "...my_name..."
}