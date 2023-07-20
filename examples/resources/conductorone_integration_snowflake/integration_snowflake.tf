resource "conductorone_integration_snowflake" "snowflake" {
  app_id              = conductorone_app.snowflake.id
  snowflake_account   = "..."
  snowflake_username  = "..."
  snowflake_password  = "..."
  snowflake_user_role = "..."
}
