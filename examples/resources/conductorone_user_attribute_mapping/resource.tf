resource "conductorone_user_attribute_mapping" "my_userattributemapping" {
  attribute_type = {
    app = {
      app_id = "...my_app_id..."
    }
    app_user_profile = {
      app_id                         = "...my_app_id..."
      app_user_profile_attribute_key = "...my_app_user_profile_attribute_key..."
    }
    cel_expression = {
      cel_expression             = "...my_cel_expression..."
      user_profile_attribute_key = "...my_user_profile_attribute_key..."
    }
    custom_app_user_profile = {
      app_id                         = "...my_app_id..."
      app_user_profile_attribute_key = "...my_app_user_profile_attribute_key..."
      user_profile_attribute_key     = "...my_user_profile_attribute_key..."
    }
    display_name = "...my_display_name..."
    id           = "...my_id..."
  }
  expand_mask = {
    paths = [
      "..."
    ]
  }
  fallbacks = [
    {
      app = {
        app_id = "...my_app_id..."
      }
      app_user_profile = {
        app_id                         = "...my_app_id..."
        app_user_profile_attribute_key = "...my_app_user_profile_attribute_key..."
      }
      cel_expression = {
        cel_expression             = "...my_cel_expression..."
        user_profile_attribute_key = "...my_user_profile_attribute_key..."
      }
    }
  ]
  profile_type_ids = [
    "..."
  ]
  user_attribute_management_service_delete_request = {
    # ...
  }
}