terraform {
  required_providers {
    conductorone = {
      source  = "conductorone/conductorone"
      version = "2.0.0-alpha.1"
    }
  }
}

provider "conductorone" {
  bearer_auth   = "<YOUR_BEARER_AUTH>" # Required
  oauth         = "<YOUR_OAUTH>"       # Required
  server_url    = "..."                # Optional
  tenant_domain = "..."                # Optional
}