data "conductorone_policy" "default_request_policy" {
    display_name = "App Owner Request Policy"
}

data "conductorone_policy" "default_revoke_policy" {
    display_name = "Default Revoke Policy"
}

data "conductorone_policy" "default_review_policy" {
    display_name = "App Owner Review Policy"
}

resource "conductorone_app" "test_new_app" {
    display_name = "Terraform App"
    owners = [
        "..."
    ]
    certify_policy_id = data.conductorone_policy.default_review_policy.id
    grant_policy_id = data.conductorone_policy.default_request_policy.id
    revoke_policy_id = data.conductorone_policy.default_revoke_policy.id
    monthly_cost_usd = "10"
}