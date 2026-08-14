data "conductorone_sign_in_policy" "my_signinpolicy" {
  page_size  = 7
  page_token = "...my_page_token..."
  query      = "...my_query..."
  refs = [
    {
      id = "...my_id..."
    }
  ]
}