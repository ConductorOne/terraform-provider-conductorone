terraform {
  required_providers {
    conductorone = {
      source  = "conductorone/conductorone"
      version = "1.6.26"
    }
  }
}

provider "conductorone" {
  server_url    = "https://c1dev.bsu.d2.ductone.com:2443"
  client_id     = "united-line-79019@c1dev.bsu.d2.ductone.com/pcc"
  client_secret = "secret-token:conductorone.com:v1:eyJrdHkiOiJPS1AiLCJjcnYiOiJFZDI1NTE5IiwieCI6IkM0bkxDVERLUE1MdnNzNjBDZ1h0bGNQZE9QbXFLeXRMZmE0czZveC01UjQiLCJkIjoiSXVDNURKOTZ6dEEtY0dHTXRyc2s0dll6enY3Z3NjOFRuVHJnb3d1OU05NCJ9"
}