resource "conductorone_directory" "my_directory" {
  app_id = "...my_app_id..."
  directory_account_filter_all = {
    # ...
  }
  directory_account_filter_cel = {
    expression = "...my_expression..."
  }
  directory_merge_config = {
    match_cases = [
      {
        app_user_key_cel = "...my_app_user_key_cel..."
        user_key_cel     = "...my_user_key_cel..."
      }
    ]
  }
}