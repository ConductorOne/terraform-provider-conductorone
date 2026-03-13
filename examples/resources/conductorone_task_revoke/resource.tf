resource "conductorone_task_revoke" "my_taskrevoke" {
  app_entitlement_id = "...my_app_entitlement_id..."
  app_id             = "...my_app_id..."
  app_user_id        = "...my_app_user_id..."
  description        = "...my_description..."
  identity_user_id   = "...my_identity_user_id..."
}