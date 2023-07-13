terraform {
  required_providers {
    conductorone = {
      source = "registry.terraform.io/ConductorOne/conductorone"
      version = "1.0.0"
    }
  }
}

provider "conductorone" {
  server_url = "<server_url>"
  client_id = "<client_id>"
  client_secret = "<client_secret>"
}

resource "conductorone_catalog" "test_catalog" {
    display_name = "terraform created catalog"
    description = "terraform test"
    owner_ids = [
        "..."
    ]
    visible_to_everyone = "false"
    published = "true"
}

resource "conductorone_catalog_visibility_bindings" "test_bindings" {
    catalog_id = conductorone_catalog.test_catalog.id
    access_entitlements = [
        {
            app_id = "..."
            id = "..."
        }
    ]
}

resource "conductorone_catalog_requestable_entries" "test_entries" {
    catalog_id = conductorone_catalog.test_catalog.id
    app_entitlements = [
        {
            app_id = "..."
            id = "..."
        },
        {
            app_id = "..."
            id = "..."
        }
    ]
}