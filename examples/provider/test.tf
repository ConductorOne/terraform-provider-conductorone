# resource "conductorone_app_resource_owner" "my-owner" {
#    app_id           = "2zw2nqDA3Bkto4WxLhXxslPeZEZ"
#   resource_id      = "30ZP8q3AHstAucwXcHJPUmbGzw0"
#   resource_type_id = "30ZP8kcUmZjgzddpQhpb49nY7cY"
#   user_ids = [
#     # "310xbJQXqHdZmm23tP8VaTUMUfJ",
#     "310xbD4Gg0VdQtnykugwqCwTUnO"
#   ] 
# }



# import {
#   to = conductorone_access_profile_requestable_entries.hello
#   id = "32l8VIv97FIr8JEwxUTw0GYVnmy"
# }

# resource "conductorone_access_profile_requestable_entries" "hello" {
#   app_entitlements = [
#     {
#       app_id = "2zw2nqDA3Bkto4WxLhXxslPeZEZ"
#       id     = "30ZP8uMg4AHWtM8fQhDI85zuy9k"
#     }
#   ]
#   catalog_id      = "32l8VIv97FIr8JEwxUTw0GYVnmy"
#   create_requests = true
# }

resource "conductorone_request_schema" "my_request_schema" {
  description = "make a request with a forms"
  fields = [
    {
      bool_field = {
        bool_rules = {
          const = false
        }
        # checkbox_field = {
        #   # ...
        # }
        default_value = true
      }
      description  = "field 1 description"
      display_name = "field 1 display name"
      # string_field = {
      #   default_value = "field 2 default value"
      #   placeholder   = "...my_placeholder..."
      # }
      #     string_rules = {
      #       address      = false
      #       const        = "...my_const..."
      #       contains     = "...my_contains..."
      #       email        = true
      #       hostname     = false
      #       ignore_empty = false
      #       in = [
      #         "..."
      #       ]
      #       ip           = true
      #       ipv4         = false
      #       ipv6         = true
      #       len_bytes    = "...my_len_bytes..."
      #       length       = "...my_length..."
      #       max_bytes    = "...my_max_bytes..."
      #       max_len      = "...my_max_len..."
      #       min_bytes    = "...my_min_bytes..."
      #       min_len      = "...my_min_len..."
      #       not_contains = "...my_not_contains..."
      #       not_in = [
      #         "..."
      #       ]
      #       pattern          = "...my_pattern..."
      #       prefix           = "...my_prefix..."
      #       strict           = false
      #       suffix           = "...my_suffix..."
      #       uri              = true
      #       uri_ref          = false
      #       uuid             = false
      #       well_known_regex = "HTTP_HEADER_VALUE"
      #     }
      #     text_field = {
      #       multiline = false
      #     }
      #   }
    },
    # {
    #   string_field = {
    #     default_value = "field 2 default value"
    #     placeholder   = "this a a placeholder of field 2"
    #     text_field = {
    #       multiline = false
    #     }
    #   }

    #   display_name = "this is a string field"
    # }
  ]
  name = "Hero....rxyxdejj"
}

resource "conductorone_request_schema" "my_request_schemaA" {
  description = "make a request with a forms"
  fields = [
    {
      bool_field = {
        bool_rules = {
          const = false
        }
        default_value = true
      }
      description  = "field 1 description"
      display_name = "field 1 display name"
    },
  ]
  name = "Hero....2"
}

resource "conductorone_request_schema_entitlement_binding" "bindings" {
  request_schema_id = resource.conductorone_request_schema.my_request_schema.id
  app_entitlement_ref = {
    app_id = "339YwJp1pN7xGdFxCgAKgIGJ5EG"
    id     = "339ZZkNjUPh3r7F37LcYZBhqvmZ"
  }

}
resource "conductorone_request_schema_entitlement_binding" "bindings2" {
  request_schema_id = resource.conductorone_request_schema.my_request_schemaA.id
  app_entitlement_ref = {
    app_id = "339YwJp1pN7xGdFxCgAKgIGJ5EG"
    id     = "339ZZk1gDnfECKImKfpFTU85PeK"
  }
}