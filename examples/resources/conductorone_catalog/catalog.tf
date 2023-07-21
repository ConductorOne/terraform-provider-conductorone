resource "conductorone_catalog" "test_catalog" {
  display_name        = "terraform created catalog"
  description         = "terraform test"
  visible_to_everyone = "false"
  published           = "true"
}
