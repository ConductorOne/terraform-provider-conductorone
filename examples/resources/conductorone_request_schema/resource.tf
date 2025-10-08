resource "conductorone_request_schema" "my_request_schema" {
  description = "...my_description..."
  fields = [
    {
      bool_field = {
        bool_rules = {
          const = false
        }
        checkbox_field = {
          # ...
        }
        default_value = true
      }
      description  = "...my_description..."
      display_name = "...my_display_name..."
      int64_field = {
        default_value = "...my_default_value..."
        int64_rules = {
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
        number_field = {
          max_value = "...my_max_value..."
          min_value = "...my_min_value..."
          step      = "...my_step..."
        }
        placeholder = "...my_placeholder..."
      }
      name = "...my_name..."
      string_field = {
        default_value = "...my_default_value..."
        password_field = {
          # ...
        }
        placeholder = "...my_placeholder..."
        select_field = {
          options = [
            {
              display_name = "...my_display_name..."
              value        = "...my_value..."
            }
          ]
        }
        string_rules = "{ \"see\": \"documentation\" }"
        text_field = {
          multiline = false
        }
      }
      string_slice_field = {
        chips_field = {
          # ...
        }
        default_values = [
          "..."
        ]
        placeholder = "...my_placeholder..."
        repeated_rules = {
          field_rules  = "{ \"see\": \"documentation\" }"
          ignore_empty = true
          max_items    = "...my_max_items..."
          min_items    = "...my_min_items..."
          unique       = true
        }
      }
    }
  ]
  name              = "...my_name..."
  request_schema_id = "...my_request_schema_id..."
  request_schema_service_delete_request = {
    # ...
  }
}