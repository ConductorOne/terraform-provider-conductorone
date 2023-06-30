terraform {
  required_providers {
    c1 = {
      source = "registry.terraform.io/speakeasy/conductorone"
      version = "1.0.0"
    }
  }
}

provider "c1" {
  server_url = "https://c1dev.logan.dev.ductone.com:2443"
  client_id = "gifted-goblin-51243@logan.conductor.one/pcc"
  client_secret = "secret-token:conductorone.com:v1:eyJrdHkiOiJPS1AiLCJjcnYiOiJFZDI1NTE5IiwieCI6IkZGeWdVRkNBYU43VU04bkgwRG1ZcklPejFxYjFuRjMwMmRzYmVLWUNiRUEiLCJkIjoibkt1alhPMkQybExFT0NFVnlPc01FeHk4Y0tvLUdEVG9CMXIzUEtIVGpaZyJ9"
}

resource "c1_terraform_policy" "plcy" {
  name = "Terraform Created Policy"
  description = "This policy was created via terraform"

}