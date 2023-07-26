terraform {
  required_providers {
    conductorone = {
      source  = "registry.terraform.io/ConductorOne/conductorone"
      version = "1.0.0"
    }
  }
}

provider "conductorone" {
  server_url    = "https://c1dev.mstanb.dev.ductone.com:2443"
  client_id     = "elated-merman-44242@c1dev.mstanb.dev.ductone.com/pcc"
  client_secret = "secret-token:conductorone.com:v1:eyJrdHkiOiJPS1AiLCJjcnYiOiJFZDI1NTE5IiwieCI6ImMyZDd3eEwxNTJTb1ROSUF6LUE2WjRUTVZJV3ZhSXh5WDZNTkZ2dEk3Q3MiLCJkIjoiSERhekprNGJrbDExQ2ROQWhVd3V1VXRiTzYyd2dReG04dktxcVhrV2pSYyJ9"
}
