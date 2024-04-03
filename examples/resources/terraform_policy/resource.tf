resource "terraform_policy" "my_policy" {
  delete_policy_request       = {}
  description                 = "...my_description..."
  display_name                = "...my_display_name..."
  policy_type                 = "POLICY_TYPE_ACCESS_REQUEST"
  reassign_tasks_to_delegates = false
}