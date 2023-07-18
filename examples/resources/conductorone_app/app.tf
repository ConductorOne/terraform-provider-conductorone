resource "conductorone_app" "test_new_app" {
  display_name = "Terraform Created App"
  owners = [
    data.conductorone_user.my_user.id
  ]
  certify_policy_id = data.conductorone_policy.default_review_policy.id
  grant_policy_id   = data.conductorone_policy.default_request_policy.id
  revoke_policy_id  = data.conductorone_policy.default_revoke_policy.id
  monthly_cost_usd  = "10"
}
