data "conductorone_session_policy_user" "my_sessionpolicyuser" {
  id         = "...my_id..."
  page_size  = 10
  page_token = "...my_page_token..."
  query      = "...my_query..."
  source     = "EFFECTIVE_SESSION_POLICY_SOURCE_FILTER_UNSPECIFIED"
}