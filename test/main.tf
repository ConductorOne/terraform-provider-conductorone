terraform {
  required_providers {
    conductorone = {
      source = "conductorone/conductorone"
    }
  }
}

variable "c1_server_url" {
  description = "ConductorOne server URL"
  type        = string
}

variable "c1_client_id" {
  description = "ConductorOne client ID"
  type        = string
}

variable "c1_client_secret" {
  description = "ConductorOne client secret"
  type        = string
  sensitive   = true
}

provider "conductorone" {
  server_url    = var.c1_server_url
  client_id     = var.c1_client_id
  client_secret = var.c1_client_secret
}


resource "conductorone_app" "my_app" {
  certify_policy_id                      = "20n6suDXRPLSiTnRLS8CURTCIq7"
  description                            = "my_description..."
  display_name                           = "mag terra"
  grant_policy_id                        = "1vPv2JswIqSI9M2DXN0eQhboegU"
  revoke_policy_id                       = "1xJnEz4VE1uT1xwHcFfMja0KF6Y"
  strict_access_entitlement_provisioning = true
}

resource "conductorone_app" "my_app2" {
  certify_policy_id                      = "31aWCTanmdcfgsJ7xAzZbWLhDIf"
  description                            = "my_description..."
  display_name                           = "mag terra custom policy"
  grant_policy_id                        = "31aWDRhcvFdtfBqZWOA0ubpLq58"
  revoke_policy_id                       = "31aWApBgA2vY4HU7UgTYc0sQRsX"
  strict_access_entitlement_provisioning = false
}

resource "conductorone_app_resource_type" "my_app_resource_type" {
  app_id        = conductorone_app.my_app.id
  display_name  = "custom res"
  resource_type = "CUSTOM"
}

resource "conductorone_app_resource" "my_app_resource" {
  app_id               = conductorone_app.my_app.id
  app_resource_type_id = conductorone_app_resource_type.my_app_resource_type.id
  description          = "...my_description..."
  display_name         = "tf custom res"
}

resource "conductorone_custom_app_entitlement" "my_custom_app_entitlement" {
  alias                = "my_alias"
  app_id               = conductorone_app.my_app.id
  app_resource_id      = conductorone_app_resource.my_app_resource.id
  app_resource_type_id = conductorone_app_resource_type.my_app_resource_type.id
  description    = "desc"
  display_name   = "custom ent"
  emergency_grant_enabled    = false
  duration_grant  = "5000s"
}

resource "conductorone_access_profile_requestable_entries" "my_access_profile_requestable_entries" {
  app_entitlements = [
    {
      app_id = conductorone_app.my_app.id
      id     = conductorone_custom_app_entitlement.my_custom_app_entitlement.id
    }
  ]
  catalog_id      = "31tvjMzbIYZCKfKiSS83y9aMTNZ"
  create_requests = false
}

resource "conductorone_app_entitlement_owner" "my_app_entitlement_owner" {
  app_id         = conductorone_app.my_app.id
  entitlement_id = conductorone_custom_app_entitlement.my_custom_app_entitlement.id
  user_ids = [
    "31lzDZ7FPFLp1bQbsxW2wCOM1Zc"
  ]
}