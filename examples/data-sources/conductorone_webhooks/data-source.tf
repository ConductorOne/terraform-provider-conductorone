data "conductorone_webhooks" "my_webhooks" {
  page_size  = 2
  page_token = "...my_page_token..."
  query      = "...my_query..."
  refs = [
    {
      id = "...my_id..."
    }
  ]
}