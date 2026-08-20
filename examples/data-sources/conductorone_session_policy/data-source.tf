data "conductorone_session_policy" "my_sessionpolicy" {
  page_size  = 0
  page_token = "...my_page_token..."
  query      = "...my_query..."
  refs = [
    {
      id = "...my_id..."
    }
  ]
}