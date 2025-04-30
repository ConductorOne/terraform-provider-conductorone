data "conductorone_app_entitlements" "tests" {
  app_ids = [
    "2w97z0DmRXQKy39hUOEbi5Bc3gP"
  ]
}

output "tests" {
  value = data.conductorone_app_entitlements.tests.list
}