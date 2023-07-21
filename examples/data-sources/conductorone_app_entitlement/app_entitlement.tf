data "conductorone_app_entitlement" "okta_everyone" {
  alias = "okta_everyone"
}

data "conductorone_app_entitlement" "okta_test_admin" {
  app_id = "<app_id>"
  id     = "<entitlement_id>"
}

data "conductorone_app_entitlement" "okta_administrators" {
  alias = "okta_test_admin"
}
