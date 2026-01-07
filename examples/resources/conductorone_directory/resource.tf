resource "conductorone_directory" "my_directory" {
  app_id = "...my_app_id..."
  directory_account_filter_all = {
    # ...
  }
  directory_account_filter_cel = {
    expression = "...my_expression..."
  }
}