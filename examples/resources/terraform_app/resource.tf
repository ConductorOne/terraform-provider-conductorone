resource "terraform_app" "my_app" {
  certify_policy_id  = "...my_certify_policy_id..."
  delete_app_request = {}
  description        = "...my_description..."
  display_name       = "...my_display_name..."
  grant_policy_id    = "...my_grant_policy_id..."
  monthly_cost_usd   = 7
  owners = [
    "...",
  ]
  revoke_policy_id = "...my_revoke_policy_id..."
}