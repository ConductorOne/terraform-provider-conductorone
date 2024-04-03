resource "terraform_catalog" "my_catalog" {
  description    = "...my_description..."
  display_name   = "...my_display_name..."
  published      = true
  request_bundle = false
  request_catalog_expand_mask = {
    paths = [
      "...",
    ]
  }
  request_catalog_management_service_delete_request = {}
  visible_to_everyone                               = false
}