resource "conductorone_app_resource_owner" "my_app_resource_owner" {
  app_id           = "...my_app_id..."
  resource_id      = "...my_resource_id..."
  resource_type_id = "...my_resource_type_id..."
  user_ids = [
    "..."
  ]
}