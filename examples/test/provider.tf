terraform {
  required_providers {
    conductorone = {
      source  = "conductorone/conductorone"
      version = "0.5.3"
    }
  }
}
provider "conductorone" {
  server_url    = "https://daemon.alee.d2.ductone.com:2443"
  client_id     = "outstanding-dragon-75366@daemon.alee.d2.ductone.com/pcc"
  client_secret = "secret-token:conductorone.com:v1:eyJrdHkiOiJPS1AiLCJjcnYiOiJFZDI1NTE5IiwieCI6IjNDa05ONlJJM01xTlhDOXNVSVZyVUdJYm9STHhFTnRJYkM0TGRUakYxc1kiLCJkIjoidGlqeTAxbHpOQzdRdnB3X3d5Q3NSTVIyZm5obkNfQk84OEtuUEM2aU5VdyJ9"
}