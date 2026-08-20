resource "conductorone_directory" "my_directory" {
  all = {
    # ...
  }
  app_id = "...my_app_id..."
  cel_expression = {
    expression = "...my_expression..."
  }
  merge_config = {
    match_cases = [
      {
        app_user_key_cel = "...my_app_user_key_cel..."
        user_key_cel     = "...my_user_key_cel..."
      }
    ]
  }
}