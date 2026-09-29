resource "conductorone_function" "my_function" {
  browser_enabled = false
  commit_message  = "...my_commit_message..."
  description     = "...my_description..."
  display_name    = "...my_display_name..."
  function_type   = "FUNCTION_TYPE_CONNECTOR"
  functions_service_delete_function_request = {
    # ...
  }
  initial_content = {
    key = "value"
  }
}