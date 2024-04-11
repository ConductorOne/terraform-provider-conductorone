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
  client_id     = "expensive-griffin-21415@c1dev.mstanb.dev.ductone.com/pcc"
  client_secret = "secret-token:conductorone.com:v1:eyJrdHkiOiJPS1AiLCJjcnYiOiJFZDI1NTE5IiwieCI6InpOYm1sQlhEcWh0ODBvRDQ3NDBUbHZTaHBBQ1dyYWRQTkwtR0t5b3Z3cVkiLCJkIjoiU0dCbVByWlVNNkZTNHhUOVozMk1Wb0NjZThSMXFHMmZneEpicnNJdXU4QSJ9"
}
