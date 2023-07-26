data "conductorone_app_entitlement" "okta_everyone" {
  alias = "okta_everyone"
}

data "conductorone_app_entitlement" "okta_test_admin" {
  app_id = "2P4WqESCtljFQ46vSDX9Cred22S"
  id     = "2P4X9LmUZDYWV09KAc8UW1MAlXJ"
}

data "conductorone_app_entitlement" "okta_administrators" {
  alias = "okta_test_admin"
}
