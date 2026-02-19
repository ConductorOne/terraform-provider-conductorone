# ConductorOne Terraform Provider

## Usage

[See the docs for usage information](./docs)

<!-- Start SDK Example Usage -->
## Testing the provider locally

Should you want to validate a change locally, the `--debug` flag allows you to execute the provider against a terraform instance locally.

This also allows for debuggers (e.g. delve) to be attached to the provider.

```sh
go run main.go --debug
# Copy the TF_REATTACH_PROVIDERS env var
# In a new terminal
cd examples/your-example
TF_REATTACH_PROVIDERS=... terraform init
TF_REATTACH_PROVIDERS=... terraform apply
```
<!-- End SDK Example Usage -->

<!-- Start Authentication [security] -->
## Authentication

The provider supports three authentication methods, evaluated in this order:

### 1. Static Access Token (highest priority)

Set `CONDUCTORONE_ACCESS_TOKEN` as an environment variable. This is typically set by the
[conductorone/oidc-token-action](https://github.com/conductorone/oidc-token-action) in GitHub Actions.

```yaml
# GitHub Actions example
steps:
  - uses: conductorone/oidc-token-action@v1
    with:
      audience: acme.conductorone.com
      client_id: clever-fox@acme.conductorone.com/wfe
  - run: terraform apply  # CONDUCTORONE_ACCESS_TOKEN is set automatically
```

### 2. OIDC Token / Workload Federation

Exchange an external OIDC token for a ConductorOne access token. This is the recommended
method for HCP Terraform and other CI/CD environments.

```hcl
provider "conductorone" {
  oidc_token = var.tfc_conductorone_token  # or auto-detected from env
  client_id  = "clever-fox@acme.conductorone.com/wfe"
}
```

The `oidc_token` value is resolved from (in order):
1. The `oidc_token` provider attribute
2. `CONDUCTORONE_OIDC_TOKEN` environment variable
3. `TFC_WORKLOAD_IDENTITY_TOKEN` environment variable (HCP Terraform auto-detection)

#### HCP Terraform Setup

1. Set workspace environment variable: `TFC_WORKLOAD_IDENTITY_AUDIENCE = acme.conductorone.com`
2. HCP Terraform auto-generates `TFC_WORKLOAD_IDENTITY_TOKEN`
3. Configure the provider with just `client_id`:

```hcl
provider "conductorone" {
  client_id = "clever-fox@acme.conductorone.com/wfe"
  # oidc_token auto-detected from TFC_WORKLOAD_IDENTITY_TOKEN
}
```

### 3. Client Credentials (Ed25519)

Traditional authentication with a client ID and secret:

```hcl
provider "conductorone" {
  client_id     = "client-id@tenant.conductorone.com/app"
  client_secret = "v1:base64-encoded-jwk..."
}
```

Or via environment variables: `CONDUCTORONE_CLIENT_ID` and `CONDUCTORONE_CLIENT_SECRET`.

### Environment Variables

| Variable | Description |
|---|---|
| `CONDUCTORONE_ACCESS_TOKEN` | Pre-exchanged bearer token (highest priority) |
| `CONDUCTORONE_OIDC_TOKEN` | Raw OIDC JWT for token exchange |
| `TFC_WORKLOAD_IDENTITY_TOKEN` | HCP Terraform workload identity token (auto-detected) |
| `CONDUCTORONE_CLIENT_ID` | Client ID for Ed25519 or workload federation auth |
| `CONDUCTORONE_CLIENT_SECRET` | Ed25519 private key in `v1:base64-jwk` format |
| `CONDUCTORONE_TENANT_DOMAIN` | Tenant domain override |
| `CONDUCTORONE_SERVER_URL` | Full server URL override |
<!-- End Authentication [security] -->

<!-- Start Available Resources and Data Sources [operations] -->
## Available Resources and Data Sources

### Resources

* [conductorone_access_review](docs/resources/access_review.md)
* [conductorone_access_review_template](docs/resources/access_review_template.md)
* [conductorone_access_conflict](docs/resources/access_conflict.md)
* [conductorone_access_profile](docs/resources/access_profile.md)
* [conductorone_access_profile_requestable_entries](docs/resources/access_profile_requestable_entries.md)
* [conductorone_access_profile_requestable_entry](docs/resources/access_profile_requestable_entry.md)
* [conductorone_access_profile_visibility_bindings](docs/resources/access_profile_visibility_bindings.md)
* [conductorone_app](docs/resources/app.md)
* [conductorone_app_entitlement_automation](docs/resources/app_entitlement_automation.md)
* [conductorone_app_entitlement_proxy_binding](docs/resources/app_entitlement_proxy_binding.md)
* [conductorone_app_resource](docs/resources/app_resource.md)
* [conductorone_app_resource_type](docs/resources/app_resource_type.md)
* [conductorone_app_entitlement_monitor_binding](docs/resources/app_entitlement_monitor_binding.md)
* [conductorone_app_entitlement_owner](docs/resources/app_entitlement_owner.md)
* [conductorone_app_owner](docs/resources/app_owner.md)
* [conductorone_app_resource_owner](docs/resources/app_resource_owner.md)
* [conductorone_automation](docs/resources/automation.md)
* [conductorone_bundle_automation](docs/resources/bundle_automation.md)
* [conductorone_compliance_framework](docs/resources/compliance_framework.md)
* [conductorone_connector_credential](docs/resources/connector_credential.md)
* [conductorone_custom_app_entitlement](docs/resources/custom_app_entitlement.md)
* [conductorone_directory](docs/resources/directory.md)
* [conductorone_function](docs/resources/function.md)
* [conductorone_function_tag](docs/resources/function_tag.md)
* [conductorone_policy](docs/resources/policy.md)
* [conductorone_request_schema_entitlement_binding](docs/resources/request_schema_entitlement_binding.md)
* [conductorone_request_schema](docs/resources/request_schema.md)
* [conductorone_risk_level](docs/resources/risk_level.md)
* [conductorone_task_grant](docs/resources/task_grant.md)
* [conductorone_task_offboarding](docs/resources/task_offboarding.md)
* [conductorone_task_revoke](docs/resources/task_revoke.md)
* [conductorone_vault](docs/resources/vault.md)
* [conductorone_webhook](docs/resources/webhook.md)
### Data Sources

* [conductorone_access_review](docs/data-sources/access_review.md)
* [conductorone_access_review_template](docs/data-sources/access_review_template.md)
* [conductorone_access_reviews](docs/data-sources/access_reviews.md)
* [conductorone_access_profile](docs/data-sources/access_profile.md)
* [conductorone_access_profile_requestable_entry](docs/data-sources/access_profile_requestable_entry.md)
* [conductorone_app](docs/data-sources/app.md)
* [conductorone_app_entitlement](docs/data-sources/app_entitlement.md)
* [conductorone_app_entitlement_automation](docs/data-sources/app_entitlement_automation.md)
* [conductorone_app_entitlement_proxy_binding](docs/data-sources/app_entitlement_proxy_binding.md)
* [conductorone_app_entitlements](docs/data-sources/app_entitlements.md)
* [conductorone_app_resource](docs/data-sources/app_resource.md)
* [conductorone_app_resource_type](docs/data-sources/app_resource_type.md)
* [conductorone_app_resource_types](docs/data-sources/app_resource_types.md)
* [conductorone_app_resources](docs/data-sources/app_resources.md)
* [conductorone_app_entitlement_monitor_binding](docs/data-sources/app_entitlement_monitor_binding.md)
* [conductorone_app_entitlement_owners](docs/data-sources/app_entitlement_owners.md)
* [conductorone_app_entitlement_users](docs/data-sources/app_entitlement_users.md)
* [conductorone_app_owners](docs/data-sources/app_owners.md)
* [conductorone_app_resource_owners](docs/data-sources/app_resource_owners.md)
* [conductorone_apps](docs/data-sources/apps.md)
* [conductorone_aws_external_id](docs/data-sources/aws_external_id.md)
* [conductorone_bundle_automation](docs/data-sources/bundle_automation.md)
* [conductorone_compliance_framework](docs/data-sources/compliance_framework.md)
* [conductorone_compliance_frameworks](docs/data-sources/compliance_frameworks.md)
* [conductorone_connector_credential](docs/data-sources/connector_credential.md)
* [conductorone_directories](docs/data-sources/directories.md)
* [conductorone_directory](docs/data-sources/directory.md)
* [conductorone_function](docs/data-sources/function.md)
* [conductorone_function_tag](docs/data-sources/function_tag.md)
* [conductorone_org_domains](docs/data-sources/org_domains.md)
* [conductorone_policies](docs/data-sources/policies.md)
* [conductorone_policy](docs/data-sources/policy.md)
* [conductorone_request_catalogs](docs/data-sources/request_catalogs.md)
* [conductorone_request_schema_entitlement_binding](docs/data-sources/request_schema_entitlement_binding.md)
* [conductorone_request_schema](docs/data-sources/request_schema.md)
* [conductorone_risk_level](docs/data-sources/risk_level.md)
* [conductorone_risk_levels](docs/data-sources/risk_levels.md)
* [conductorone_role](docs/data-sources/role.md)
* [conductorone_roles](docs/data-sources/roles.md)
* [conductorone_user](docs/data-sources/user.md)
* [conductorone_users](docs/data-sources/users.md)
* [conductorone_webhook](docs/data-sources/webhook.md)
* [conductorone_webhooks](docs/data-sources/webhooks.md)
<!-- End Available Resources and Data Sources [operations] -->

Terraform allows you to use local provider builds by setting a `dev_overrides` block in a configuration file called `.terraformrc`. This block overrides all other configured installation methods.

Terraform searches for the `.terraformrc` file in your home directory and applies any configuration settings you set.

```
provider_installation {

  dev_overrides {
      "registry.terraform.io/terraform/scaffolding" = "<PATH>"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```

Your `<PATH>` may vary depending on how your Go environment variables are configured. Execute `go env GOBIN` to set it, then set the `<PATH>` to the value returned. If nothing is returned, set it to the default location, `$HOME/go/bin`.

## Versions with breaking changes
### [v1.0.0](https://github.com/ConductorOne/terraform-provider-conductorone/releases/tag/v1.0.0)
- This version introduces a breaking change to the `request_catalog` resource and datasource. The `request_catalog` resource and datasource are now the `access_profile` resource and datasource. 
- There are also breaking changes to many of the datasources, as we switched to using search endpoints to give more options on how to filter the result. This may require you to update your Terraform configurations to use the proper fields for the datasources as some of the field names have changed. 

### [v0.4.0](https://github.com/ConductorOne/terraform-provider-conductorone/releases/tag/v0.4.0)
- This version fixes issues with importing app entitlements. The SDK and the provider schema did not match so imports were broken, the provider was changed to match the SDK which introduced a breaking change. the `max_grant_druation` object was removed and the internal `duration_unset` and `grant_unset` fields have been moved out.

### [v0.3.0](https://github.com/ConductorOne/terraform-provider-conductorone/releases/tag/v0.3.0)
- This version fixes issue where you weren't able to update app owners on the app resource by breaking out app owners into it's own resource. This will break any Terraform configurations that specify appOwners on the app resource. To use the new version of the provider you will need to use the new app_owners resource to configure the owners of an app, and no longer specify owners on the app resource.

### [v0.2.0](https://github.com/ConductorOne/terraform-provider-conductorone/releases/tag/v0.2.0)
- This version introduces a breaking change to the `app_entitlement` datasource. The `app_entitlement` datasource now nests the `grant_duration` and `grant_unset` inside the `max_grant_duration` field.

### [v0.1.0](https://github.com/ConductorOne/terraform-provider-conductorone/releases/tag/v0.1.0)
- This version introduces a breaking change to the `app_entitlement_owners` resource. The `app_entitlement_owners` resource now requires a list of strings called `user_ids`, which is used to set the owners
on the resource. The list of `user_ids` will replace any existing owners on the app entitlement with the new list of owners. 

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)

<!-- Start Summary [summary] -->
## Summary

ConductorOne API: The ConductorOne API is a HTTP API for managing ConductorOne resources.
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [ConductorOne Terraform Provider](#conductorone-terraform-provider)
  * [Usage](#usage)
  * [Testing the provider locally](#testing-the-provider-locally)
  * [Authentication](#authentication)
  * [Available Resources and Data Sources](#available-resources-and-data-sources)
  * [Versions with breaking changes](#versions-with-breaking-changes)
  * [Contributions](#contributions)
  * [Installation](#installation)
  * [Testing the provider locally](#testing-the-provider-locally-1)

<!-- End Table of Contents [toc] -->

<!-- Start Installation [installation] -->
## Installation

To install this provider, copy and paste this code into your Terraform configuration. Then, run `terraform init`.

```hcl
terraform {
  required_providers {
    conductorone = {
      source  = "conductorone/conductorone"
      version = "1.7.11"
    }
  }
}

provider "conductorone" {
  # Configuration options
}
```
<!-- End Installation [installation] -->

<!-- Start Testing the provider locally [usage] -->
## Testing the provider locally

#### Local Provider

Should you want to validate a change locally, the `--debug` flag allows you to execute the provider against a terraform instance locally.

This also allows for debuggers (e.g. delve) to be attached to the provider.

```sh
go run main.go --debug
# Copy the TF_REATTACH_PROVIDERS env var
# In a new terminal
cd examples/your-example
TF_REATTACH_PROVIDERS=... terraform init
TF_REATTACH_PROVIDERS=... terraform apply
```

#### Compiled Provider

Terraform allows you to use local provider builds by setting a `dev_overrides` block in a configuration file called `.terraformrc`. This block overrides all other configured installation methods.

1. Execute `go build` to construct a binary called `terraform-provider-conductorone`
2. Ensure that the `.terraformrc` file is configured with a `dev_overrides` section such that your local copy of terraform can see the provider binary

Terraform searches for the `.terraformrc` file in your home directory and applies any configuration settings you set.

```
provider_installation {

  dev_overrides {
      "registry.terraform.io/conductorone/conductorone" = "<PATH>"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```
<!-- End Testing the provider locally [usage] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->
