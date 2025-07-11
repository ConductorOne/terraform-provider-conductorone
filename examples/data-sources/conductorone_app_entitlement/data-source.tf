data "conductorone_app_entitlement" "my_app_entitlement" {
  access_review_id = "...my_access_review_id..."
  alias            = "...my_alias..."
  app_ids = [
    "..."
  ]
  app_user_ids = [
    "..."
  ]
  compliance_framework_ids = [
    "..."
  ]
  display_name = "...my_display_name..."
  exclude_app_ids = [
    "..."
  ]
  exclude_app_user_ids = [
    "..."
  ]
  exclude_resource_type_ids = [
    "..."
  ]
  include_deleted = true
  is_automated    = false
  membership_type = [
    "APP_ENTITLEMENT_MEMBERSHIP_TYPE_OWNER"
  ]
  only_get_expiring = true
  page_size         = 5
  page_token        = "...my_page_token..."
  query             = "...my_query..."
  refs = [
    {
      app_id = "...my_app_id..."
      id     = "...my_id..."
    }
  ]
  resource_ids = [
    "..."
  ]
  resource_trait_ids = [
    "..."
  ]
  resource_type_ids = [
    "..."
  ]
  risk_level_ids = [
    "..."
  ]
  source_connector_id = "...my_source_connector_id..."
}