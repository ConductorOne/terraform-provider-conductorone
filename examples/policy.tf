terraform {
  required_providers {
    terraform = {
      source = "registry.terraform.io/speakeasy/conductorone"
      version = "1.0.0"
    }
  }
}

provider "terraform" {
  server_url = "https://mstanb.sandbox.ductone.com"
  client_id = ""
  client_secret = ""
}

resource "terraform_policy" "plcy" {
  display_name = "Terraform Created Policy"
  description = "This policy was created via terraform"
  policy_type = "POLICY_TYPE_CERTIFY"
}