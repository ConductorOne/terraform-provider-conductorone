terraform {
  required_providers {
    conductorone = {
      source  = "registry.terraform.io/ConductorOne/conductorone"
      version = "1.0.0"
    }
  }
}

provider "conductorone" {
  server_url    = "https://mstanb.mstanb.dev.ductone.com:2443"
  client_id     = "obnoxious-gnome-50189@mstanb.mstanb.dev.ductone.com/pcc"
  client_secret = "secret-token:conductorone.com:v1:eyJrdHkiOiJPS1AiLCJjcnYiOiJFZDI1NTE5IiwieCI6IjRnUVhFY21IeUJ6U1BRVTR6RDRRcThtX3hvamt1NFNYdUpqd0hLZzNjbkUiLCJkIjoiUnJ4T2pLeGZ6NHd4VGd1emFKbU1zSmV6ZVQ0Wlc2QmxTUU13RDVxd3BWMCJ9"
}
