resource "conductorone_access_review_template" "my_access_review_template" {
  access_review_duration = "...my_access_review_duration..."
  description            = "...my_description..."
  display_name           = "...my_display_name..."
  owner_ids = [
    "..."
  ]
  policy_id = "...my_policy_id..."
}