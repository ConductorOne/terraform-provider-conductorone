resource "conductorone_app_entitlement_owner" "my_app_entitlement_owner" {
  app_id         = "...my_app_id..."
  entitlement_id = "...my_entitlement_id..."
  user_ids = [
    "..."
  ]
}