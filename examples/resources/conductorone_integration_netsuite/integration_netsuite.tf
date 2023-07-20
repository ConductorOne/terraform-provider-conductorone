resource "conductorone_integration_netsuite" "netsuite" {
  app_id = conductorone_app.netsuite.id
  user_ids = [
    conductorone_user.admin.id
  ]
  netsuite_account_id      = "..."
  netsuite_consumer_key    = "..."
  netsuite_consumer_secret = "..."
  netsuite_token_key       = "..."
  netsuite_token_secret    = "..."
}
