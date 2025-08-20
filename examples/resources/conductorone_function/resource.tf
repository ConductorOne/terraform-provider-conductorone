resource "conductorone_function" "my_function" {
  commit_message = "...my_commit_message..."
  description    = "...my_description..."
  display_name   = "...my_display_name..."
  function_type  = "FUNCTION_TYPE_ANY"
  functions_service_delete_function_request = {
    # ...
  }
  initial_content = {
    key = "value"
  }
}