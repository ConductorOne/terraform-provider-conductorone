data "conductorone_policy" "default_request_policy" {
  display_name = "App Owner Request Policy"
}

data "conductorone_policy" "default_revoke_policy" {
  display_name = "Default Revoke Policy"
}

data "conductorone_policy" "default_review_policy" {
  display_name = "App Owner Review Policy"
}
