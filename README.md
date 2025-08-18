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

<!-- Start Available Resources and Data Sources [operations] -->
## Available Resources and Data Sources

### Resources

* [conductorone_access_conflict](docs/resources/access_conflict.md)
* [conductorone_access_profile](docs/resources/access_profile.md)
* [conductorone_access_profile_requestable_entries](docs/resources/access_profile_requestable_entries.md)
* [conductorone_access_profile_visibility_bindings](docs/resources/access_profile_visibility_bindings.md)
* [conductorone_app](docs/resources/app.md)
* [conductorone_app_entitlement_automation](docs/resources/app_entitlement_automation.md)
* [conductorone_app_entitlement_proxy_binding](docs/resources/app_entitlement_proxy_binding.md)
* [conductorone_app_resource](docs/resources/app_resource.md)
* [conductorone_app_resource_type](docs/resources/app_resource_type.md)
* [conductorone_app_entitlement_owner](docs/resources/app_entitlement_owner.md)
* [conductorone_app_owner](docs/resources/app_owner.md)
* [conductorone_bundle_automation](docs/resources/bundle_automation.md)
* [conductorone_compliance_framework](docs/resources/compliance_framework.md)
* [conductorone_connector_credential](docs/resources/connector_credential.md)
* [conductorone_custom_app_entitlement](docs/resources/custom_app_entitlement.md)
* [conductorone_policy](docs/resources/policy.md)
* [conductorone_risk_level](docs/resources/risk_level.md)
* [conductorone_webhook](docs/resources/webhook.md)
### Data Sources

* [conductorone_access_profile](docs/data-sources/access_profile.md)
* [conductorone_app](docs/data-sources/app.md)
* [conductorone_app_entitlement](docs/data-sources/app_entitlement.md)
* [conductorone_app_entitlement_automation](docs/data-sources/app_entitlement_automation.md)
* [conductorone_app_entitlement_proxy_binding](docs/data-sources/app_entitlement_proxy_binding.md)
* [conductorone_app_entitlements](docs/data-sources/app_entitlements.md)
* [conductorone_app_resource](docs/data-sources/app_resource.md)
* [conductorone_app_resource_type](docs/data-sources/app_resource_type.md)
* [conductorone_app_resources](docs/data-sources/app_resources.md)
* [conductorone_aws_external_id](docs/data-sources/aws_external_id.md)
* [conductorone_bundle_automation](docs/data-sources/bundle_automation.md)
* [conductorone_compliance_framework](docs/data-sources/compliance_framework.md)
* [conductorone_connector_credential](docs/data-sources/connector_credential.md)
* [conductorone_policy](docs/data-sources/policy.md)
* [conductorone_risk_level](docs/data-sources/risk_level.md)
* [conductorone_user](docs/data-sources/user.md)
* [conductorone_webhook](docs/data-sources/webhook.md)
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
      version = "1.5.1"
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
