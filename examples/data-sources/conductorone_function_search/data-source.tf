data "conductorone_function_search" "my_functionsearch" {
  function_types = [
    "FUNCTION_TYPE_UNSPECIFIED"
  ]
  page_size  = 10
  page_token = "...my_page_token..."
  query      = "...my_query..."
}