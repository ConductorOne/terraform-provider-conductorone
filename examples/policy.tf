terraform {
  required_providers {
    c1 = {
      source = "registry.terraform.io/speakeasy/conductorone"
      version = "1.0.0"
    }
  }
}

provider "c1" {
  server_url = ""
  client_id = ""
  client_secret = ""
}

resource "c1_terraform_policy" "plcy" {
  name = "Terraform Created Policy"
  description = "This policy was created via terraform"

}