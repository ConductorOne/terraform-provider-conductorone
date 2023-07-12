terraform {
  required_providers {
    conductorone = {
      source = "registry.terraform.io/ConductorOne/conductorone"
      version = "1.0.0"
    }
  }
}

provider "conductorone" {
  server_url = "<server_url>"
  client_id = "<client_id>"
  client_secret = "<client_secret>"
}

resource "conductorone_policy" "new_policy" {
  display_name = "Terraform Created Policy"
  description = "New description"
  policy_type = "POLICY_TYPE_CERTIFY"
  reassign_tasks_to_delegates = "false"
  policy_steps = {
    certify = {
      steps = [
        {
          approval = {
            allow_reassignment = "true"
            require_reassignment_reason = "false"
            assigned = "false"
            require_approval_reason = "false"
            app_group_approval = {
              allow_self_approval = "true"
              app_group_id = "..."
              app_id = "..."
              fallback = "true"
              fallback_user_ids = [
                "..."
              ]
            }
          }
        }
      ]
    }
  }
}
