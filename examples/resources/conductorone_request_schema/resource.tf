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
      admin_config = {
        default_value_cel = "...my_default_value_cel..."
        show_to_user      = false
      }
      bool_field = {
        checkbox_field = {
          # ...
        }
        default_value = true
        rules = {
          const = true
        }
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
        number_field = {
          max_value = "...my_max_value..."
          min_value = "...my_min_value..."
          step      = "...my_step..."
        }
        placeholder = "...my_placeholder..."
        rules = {
          const        = "...my_const..."
          gt           = "...my_gt..."
          gte          = "...my_gte..."
          ignore_empty = false
          in = [
            "..."
          ]
          lt  = "...my_lt..."
          lte = "...my_lte..."
          not_in = [
            "..."
          ]
        }
      }
      name = "...my_name..."
      oauth2_field = {
        oauth2_field_view = {
          # ...
        }
      }
      read_only = true
      required  = false
      shared_config = {
        default_value_cel        = "...my_default_value_cel..."
        input_transformation_cel = "...my_input_transformation_cel..."
        lock_default_values      = true
      }
      string_field = {
        date_field = {
          default_to_today    = false
          max_date            = "...my_max_date..."
          max_days_from_today = 10
          min_date            = "...my_min_date..."
          min_days_from_today = 4
        }
        default_value = "...my_default_value..."
        password_field = {
          # ...
        }
        picker_field = {
          app_user_picker = {
            app_id = "...my_app_id..."
          }
          c1_user_picker = {
            exclude_user_ids = [
              "..."
            ]
            include_deactivated = true
            user_ids = [
              "..."
            ]
          }
          resource_picker = {
            app_id           = "...my_app_id..."
            resource_type_id = "...my_resource_type_id..."
          }
        }
        placeholder = "...my_placeholder..."
        rules = {
          address      = true
          const        = "...my_const..."
          contains     = "...my_contains..."
          email        = false
          hostname     = true
          ignore_empty = true
          in = [
            "..."
          ]
          ip           = true
          ipv4         = true
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
        text_field = {
          multiline = false
          suffix    = "...my_suffix..."
        }
      }
      string_map_field = {
        default_value = {
          key = "value"
        }
        rules = {
          is_required    = false
          validate_empty = false
        }
      }
      user_config = {
        input_transformation_cel = "...my_input_transformation_cel..."
      }
    }
  ]
  justification_visibility = "JUSTIFICATION_VISIBILITY_HIDE"
  name                     = "...my_name..."
}