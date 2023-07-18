terraform {
  required_providers {
    conductorone = {
      source  = "registry.terraform.io/ConductorOne/conductorone"
      version = "1.0.0"
    }
  }
}

provider "conductorone" {
  server_url    = "<server_url>"
  client_id     = "<client_id>"
  client_secret = "<client_secret>"
}
