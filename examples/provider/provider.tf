terraform {
  required_providers {
    conductorone = {
      source  = "conductorone/conductorone"
      version = "1.8.0"
    }
  }
}

provider "conductorone" {
  bearer_auth   = "<YOUR_BEARER_AUTH>" # Required
  oauth         = "<YOUR_OAUTH>"       # Required
  server_url    = "..."                # Optional
  tenant_domain = "..."                # Optional
}