resource "conductorone_app" "my_app" {
  certify_policy_id = "...my_certify_policy_id..."
  description       = "...my_description..."
  display_name      = "...my_display_name..."
  grant_policy_id   = "...my_grant_policy_id..."
  monthly_cost_usd  = 48.12
  owners = [
    "...",
  ]
  revoke_policy_id = "...my_revoke_policy_id..."
}