---
page_title: "Provider: ConductorOne"
description: |-
  The ConductorOne provider is used to interact with the ConductorOne configuration plane.
---

# ConductorOne Provider

The ConductorOne provider allows you to configure important management objects including:
- Policies
- Integrations
- Entitlements
- Applications
- Access Profiles

Use the navigation to the left to read about the available resources.

## Example Usage

```terraform
terraform {
  required_providers {
    conductorone = {
      source  = "conductorone/conductorone"
      version = "1.4.4"
    }
  }
}

provider "conductorone" {
  # Configuration options
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `client_id` (String) ClientId for Auth
- `client_secret` (String) ClientSecret for Auth
- `server_url` (String) Server URL (defaults to https://{tenantDomain}.conductor.one)
- `tenant_domain` (String) Tenant Domain (derived from client_id if not provided)

## Limitations

### Secrets and Terraform state

Some of the resources that can be created with this provider, like `conductorone_integration_okta`,
can contain sensitive information, like API keys or passwords. These values are marked accordingly as _sensitive_
in the provider schema. This helps prevent the accidental leakage of these values in logs or other form of output.

It's important to remember that the values that constitute the "state" of those
resources will be stored in the [Terraform state](https://www.terraform.io/language/state) file.
This includes the "secrets", that will be part of the state file *unencrypted*.

Because of this, it's important to make sure that the state file is stored in a secure location. Better yet, it is possible
to inject secrets using something like the [vault provider](https://registry.terraform.io/providers/hashicorp/vault).


## Breaking changes

### 1.0.0

- The `request_catalog` resource and datasource are now the `access_profile` resource and datasource.
- There are also breaking changes to many of the datasources, as we switched to using search endpoints to give more options on how to filter the result. This may require you to update your Terraform configurations to use the proper fields for the datasources as some of the field names have changed. 
  The list of datasources that have been affected are:
    - `conductorone_app`
    - `conductorone_app_entitlement`
    - `conductorone_policy`
    - `conductorone_user`
    - `conductorone_webhook`
    - `conductorone_app_resource_type`
