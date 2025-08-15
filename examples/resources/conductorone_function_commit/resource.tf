resource "conductorone_function_commit" "my_functioncommit" {
  commit_message = "...my_commit_message..."
  content = {
    key = "value"
  }
  function_id = "...my_function_id..."
}