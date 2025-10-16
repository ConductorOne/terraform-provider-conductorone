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
          strict           = false
          suffix           = "...my_suffix..."
          uri              = true
          uri_ref          = false
          uuid             = false
          well_known_regex = "HTTP_HEADER_VALUE"
        }
        text_field = {
          multiline = false
        }
      }
    }
  ]
  name = "...my_name..."
}