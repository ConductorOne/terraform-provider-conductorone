# ConductorOne Terraform Provider

## Usage

[See the docs for usage information](./docs)

<!-- Start SDK Installation -->
## SDK Installation

To install this provider, copy and paste this code into your Terraform configuration. Then, run `terraform init`.

```hcl
terraform {
  required_providers {
    conductorone = {
      source  = "ConductorOne/conductorone"
      version = ">=0.0.1"
    }
  }
}

provider "conductorone" {
  server_url    = "<server_url>"
  client_id     = "<client_id>"
  client_secret = "<client_secret>"
}
```
<!-- End SDK Installation -->

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

<!-- Start SDK Available Operations -->

<!-- End SDK Available Operations -->

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
