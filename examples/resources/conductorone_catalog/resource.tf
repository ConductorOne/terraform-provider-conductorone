resource "conductorone_catalog" "my_catalog" {
  description  = "...my_description..."
  display_name = "...my_display_name..."
  published    = false
  request_catalog_expand_mask = {
    paths = [
      "...",
    ]
  }
  visible_to_everyone = false
}