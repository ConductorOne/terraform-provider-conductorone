resource "conductorone_policy" "new_policy" {
  display_name                = "Terraform Created Policy"
  description                 = "New description"
  policy_type                 = "POLICY_TYPE_CERTIFY"
  reassign_tasks_to_delegates = "false"
  policy_steps = {
    certify = {
      steps = [
        {
          approval = {
            allow_reassignment          = "true"
            require_reassignment_reason = "false"
            assigned                    = "false"
            require_approval_reason     = "false"
            app_group_approval = {
              allow_self_approval = "true"
              app_group_id        = data.conductorone_app_entitlement.okta_everyone.id
              app_id              = data.conductorone_app_entitlement.okta_everyone.app_id
              fallback            = "true"
              fallback_user_ids = [
                data.conductorone_user.my_user.id
              ]
            }
          }
        }
      ]
    }
  }
}
