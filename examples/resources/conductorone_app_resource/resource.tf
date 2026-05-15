resource "conductorone_app_resource" "my_app_resource" {
  annotations = {
    key = "value"
  }
  app_id               = "...my_app_id..."
  app_resource_type_id = "...my_app_resource_type_id..."
  description          = "...my_description..."
  display_name         = "...my_display_name..."
  match_baton_id       = "...my_match_baton_id..."
}