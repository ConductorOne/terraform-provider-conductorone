# Manage an existing app entitlement (typically one synced by a connector).
# Import the entitlement first (see "Import" below), then attach policies and
# behavior overrides like the example shows.
resource "conductorone_app_entitlement" "my_app_entitlement" {
  app_id = "...my_app_id..."
  id     = "...my_entitlement_id..."

  description      = "...my_description..."
  display_name     = "...my_display_name..."
  grant_policy_id  = "...my_grant_policy_id..."
  revoke_policy_id = "...my_revoke_policy_id..."
  duration_grant   = "3600s"

  emergency_grant_enabled   = false
  emergency_grant_policy_id = "...my_emergency_grant_policy_id..."

  provision_policy = {
    connector_provision = {
      default_behavior = {}
    }
  }

  deprovisioner_policy = {
    manual_provision = {
      instructions = "...my_instructions..."
      user_ids     = ["...my_user_id..."]
    }
  }
}
