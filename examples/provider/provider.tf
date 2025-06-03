terraform {
  required_providers {
    conductorone = {
      source  = "conductorone/conductorone"
      version = "1.0.7"
    }
  }
}
provider "conductorone" {
  server_url    = "https://c1dev.bsu.d2.ductone.com:2443"
  client_id     = "aggressive-minotaur-86298@c1dev.bsu.d2.ductone.com/pcc"
  client_secret = "secret-token:conductorone.com:v1:eyJrdHkiOiJPS1AiLCJjcnYiOiJFZDI1NTE5IiwieCI6IlNQbVpFVXlfN2Q2cWdfZXc5MmM4UHZZTkZsd0d0UEE2d0lxUGVnek1OU28iLCJkIjoiYXZ5SGZtdnBwS241QUY3bHl4RkU0V3FQeVptcnhfWGk5VEFGU2xqbWRTSSJ9"
}
