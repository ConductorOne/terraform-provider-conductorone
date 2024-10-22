resource "conductorone_integration_snowflake_v2" "snowflake_v2" {
  app_id = conductorone_app.snowflake_v2.id
  user_ids = [
    conductorone_user.admin.id
  ]
  snowflake_account_url = "..."
  snowflake_account_id  = "..."
  snowflake_user_id     = "..."
  snowflake_private_key = "..."
}
