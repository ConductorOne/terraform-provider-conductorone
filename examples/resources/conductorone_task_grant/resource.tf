resource "conductorone_task_grant" "my_taskgrant" {
  app_entitlement_id = "...my_app_entitlement_id..."
  app_id             = "...my_app_id..."
  app_user_id        = "...my_app_user_id..."
  description        = "...my_description..."
  emergency_access   = true
  grant_duration     = "...my_grant_duration..."
  identity_user_id   = "...my_identity_user_id..."
  request_data = {
    # ...
  }
  task_grant_source = {
    conversation_id = "...my_conversation_id..."
    external_url    = "...my_external_url..."
    integration_id  = "...my_integration_id..."
    is_extension    = true
    request_id      = "...my_request_id..."
  }
}