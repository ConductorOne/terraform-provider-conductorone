---
page_title: "Upgrading to v2.0.0"
subcategory: "Upgrade Guides"
description: |-
  Migration guide for the ConductorOne Terraform provider v2.0.0 release.
---

# Upgrading to v2.0.0

v2.0.0 renames a small number of attributes on existing resources. Any
HCL config using the old name will fail at plan time with
`An argument named "<old>" is not expected here.` No resources or data
sources were removed.

The full release-by-release changelog ships with each tag on the
[Releases page](https://github.com/ConductorOne/terraform-provider-conductorone/releases).

~> **Prerelease.** v2.0.0-alpha.1 is an alpha. Pin it explicitly; the
Terraform Registry will not auto-resolve prereleases.

## Pinning the prerelease

```hcl
terraform {
  required_providers {
    conductorone = {
      source  = "ConductorOne/conductorone"
      version = "2.0.0-alpha.1"
    }
  }
}
```

Then run `terraform init -upgrade`.

## Renamed attributes

| Resource | v1.0.40 | v2.0.0 |
| --- | --- | --- |
| `conductorone_access_conflict` | `notification_config` | `access_conflict_notification_config` |
| `conductorone_access_review` | `notification_config` | `access_review_notification_config` |
| `conductorone_access_review_template` | `notification_config` | `access_review_notification_config` |
| `conductorone_policy` | `form` | `policy_form` |
| `conductorone_automation` | `task_action` | `automations_task_action` |
| `conductorone_automation` | `webhook` | `automations_webhook` |

The `form` → `policy_form` rename on `conductorone_policy` applies twice:
as a top-level attribute and as a nested field inside
`policy_steps.steps.action.policy_form` (the approval-step form payload).
Both need to be renamed in your HCL. The same rename applies to the
`conductorone_policy` and `conductorone_policies` data sources.

## Migration recipe

For each affected resource, update the HCL key. The values do not need
to change. Example for `conductorone_access_review`:

```hcl
# Before (v1.0.40)
resource "conductorone_access_review" "quarterly" {
  display_name = "Q3 access review"

  notification_config = {
    notify_users_of_completed_task = true
  }
}

# After (v2.0.0)
resource "conductorone_access_review" "quarterly" {
  display_name = "Q3 access review"

  access_review_notification_config = {
    notify_users_of_completed_task = true
  }
}
```

Then upgrade the provider and align state with the v2 schema:

```sh
terraform init -upgrade
terraform apply -refresh-only    # rewrites state under the v2 schema
terraform plan                   # should be a no-op
```

If `terraform plan` is non-empty after the refresh, inspect the diff —
that's a real intent change and not a v2 migration artifact.

## Rolling back

Downgrade by repinning to `1.0.40` and running
`terraform init -upgrade`. Any HCL that adopted the new attribute names
needs to be reverted alongside the provider downgrade.
