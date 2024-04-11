resource "conductorone_app_entitlement" "okta_test_admin" {
  id     = "287oY0rG4UirjDNFEYguMBvxyim"
  app_id = "2dFEyrKk0BYLpOIw8T1vaDpmTX5"
  alias  = "terraform_ent"
  provision_policy = {
    webhook_provision = {
      webhook_id = data.conductorone_webhook.test_webhook.id
    }
  }
}
