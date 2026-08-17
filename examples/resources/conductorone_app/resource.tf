resource "conductorone_app" "my_app" {
  annotations = {
    key = "value"
  }
  app_entitlement_owner_refs = [
    {
      app_id = "...my_app_id..."
      id     = "...my_id..."
    }
  ]
  certify_policy_id = "...my_certify_policy_id..."
  description       = "...my_description..."
  display_name      = "...my_display_name..."
  grant_policy_id   = "...my_grant_policy_id..."
  idempotency_key   = "terraform/workspace/conductorone_app.my_app"
  identity_matching = "APP_USER_IDENTITY_MATCHING_CUSTOM"
  instructions      = "...my_instructions..."
  match_baton_ref = {
    app_id       = "...my_app_id..."
    connector_id = "...my_connector_id..."
    external_id  = "app::0oa123"
  }
  monthly_cost_usd                       = 1
  revoke_policy_id                       = "...my_revoke_policy_id..."
  strict_access_entitlement_provisioning = true
}