data "conductorone_webhook" "my_webhook" {
  page_size  = 3
  page_token = "...my_page_token..."
  query      = "...my_query..."
  refs = [
    {
      id = "...my_id..."
    }
  ]
}