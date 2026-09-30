data "conductorone_function" "my_function" {
  function_types = [
    "FUNCTION_TYPE_UNSPECIFIED"
  ]
  page_size  = 7
  page_token = "...my_page_token..."
  query      = "...my_query..."
}