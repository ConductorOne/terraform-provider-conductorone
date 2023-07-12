terraform {
  required_providers {
    conductorone = {
      source = "registry.terraform.io/ConductorOne/conductorone"
      version = "1.0.0"
    }
  }
}

provider "conductorone" {
  server_url = "https://c1dev.mstanb.dev.ductone.com:2443"
  client_id = "elated-merman-44242@c1dev.mstanb.dev.ductone.com/pcc"
  client_secret = "secret-token:conductorone.com:v1:eyJrdHkiOiJPS1AiLCJjcnYiOiJFZDI1NTE5IiwieCI6ImMyZDd3eEwxNTJTb1ROSUF6LUE2WjRUTVZJV3ZhSXh5WDZNTkZ2dEk3Q3MiLCJkIjoiSERhekprNGJrbDExQ2ROQWhVd3V1VXRiTzYyd2dReG04dktxcVhrV2pSYyJ9"
}

resource "conductorone_catalog" "test_catalog" {
    display_name = "terraform created catalog"
    description = "terraform test"
    owner_ids = [
        "2P4VnGXgZv5PDKQ9zEJnjEXNBxY"
    ]
    visible_to_everyone = "false"
    published = "true"
}

resource "conductorone_catalog_visibility_bindings" "test_bindings" {
    catalog_id = conductorone_catalog.test_catalog.id
    access_entitlements = [
        {
            app_id = "2P4WqESCtljFQ46vSDX9Cred22S"
            id = "2P4X9MMlOTlQ80l1umHhpuukkWW"
        }
    ]
}

resource "conductorone_catalog_requestable_entries" "test_entries" {
    catalog_id = conductorone_catalog.test_catalog.id
    app_entitlements = [
        {
            app_id = "2P4WqESCtljFQ46vSDX9Cred22S"
            id = "2P4X9LmUZDYWV09KAc8UW1MAlXJ"
        },
        {
            app_id = "2P4WqESCtljFQ46vSDX9Cred22S"
            id = "2PvPCYCC65J1zu5pPhYmyR5eNKc"
        }
    ]
}