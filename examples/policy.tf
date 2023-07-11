terraform {
  required_providers {
    terraform = {
      source = "registry.terraform.io/speakeasy/conductorone"
      version = "1.0.0"
    }
  }
}

provider "terraform" {
  server_url = "<server_url>"
  client_id = "<client_id>"
  client_secret = "<client_secret>"
}

resource "terraform_policy" "new_policy" {
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
              app_group_id = "2P4X9MMlOTlQ80l1umHhpuukkWW"
              app_id = "2P4WqESCtljFQ46vSDX9Cred22S"
              fallback = "true"
              fallback_user_ids = [
                "2P4VnGXgZv5PDKQ9zEJnjEXNBxY"
              ]
            }
          }
        }
      ]
    }
  }
}

resource "terraform_policy" "prov_policy" {
  display_name = "Terraform Created Prov Policy"
  description = "New description"
  policy_type = "POLICY_TYPE_GRANT"
  reassign_tasks_to_delegates = "false"
  policy_steps = {
    grant = {
      steps = [
        {
          provision = {
            assigned = "false"
            provision_policy = {
              delegated_provision = {
                app_id = "2P4WqESCtljFQ46vSDX9Cred22S"
                entitlement_id = "2P4X9MMlOTlQ80l1umHhpuukkWW"
              }
            }
          }
        }
      ]
    }
  }
}