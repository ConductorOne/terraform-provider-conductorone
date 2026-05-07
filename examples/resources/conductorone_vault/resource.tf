resource "conductorone_vault" "my_vault" {
  description  = "...my_description..."
  display_name = "...my_display_name..."
  group_authz_vault = {
    # ...
  }
  magic_vault = {
    allow_unauthed_views = false
    allowed_views        = 5
  }
  owner_ids = [
    "..."
  ]
}