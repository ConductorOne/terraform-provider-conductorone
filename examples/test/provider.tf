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
  client_id     = "good-minotaur-71841@mstanb.mstanb.dev.ductone.com/pcc"
  client_secret = "secret-token:conductorone.com:v1:eyJrdHkiOiJPS1AiLCJjcnYiOiJFZDI1NTE5IiwieCI6IjNsQmo5TVJFbllXU3ExbEozdHluSVN5NXBSQ0F1ZTlNeWhpUkVCQzdCLVUiLCJkIjoiNHV5OUVnR1RkeVBqNGJiRU5SZXBzVk8zdzFBNF9TeVhQNmY5YThYTl9pWSJ9"
}
