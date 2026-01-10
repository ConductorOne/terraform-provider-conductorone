# Known Issues

## Speakeasy Version Sensitivity

### Summary

The Terraform provider is sensitive to the exact Speakeasy version used for code generation. Version **1.658.1** is pinned in `.speakeasy/workflow.yaml` because newer versions (tested through 1.661.0) introduce a regression causing perpetual diffs in Terraform state.

### Affected Fields

Fields marked with `x-speakeasy-terraform-plan-only: true` in the overlay are affected:

| Field | Correct (v1.658.1) | Broken (v1.661.0+) |
|-------|--------------------|--------------------|
| `provision_policy` | `Computed: true, Optional: true` | `Optional: true` only |
| `duration_grant` | `Computed: true, Optional: true` | `Optional: true` only |
| `duration_unset` | `Computed: true, Optional: true` | `Optional: true` only |

### Symptom

When using code generated with a newer Speakeasy version, acceptance tests fail with non-empty refresh plans:

```
TestAccBundleAutomationResource
    bundleautomation_resource_test.go:10: Step 1/1 error: After applying this test step, the refresh plan was not empty.
    
    - provision_policy = {
        - unconfigured_provision = {} -> null
      } -> null
```

Users will see perpetual diffs on every `terraform plan` for resources like `conductorone_custom_app_entitlement`, even when no changes were made.

### Root Cause

The `x-speakeasy-terraform-plan-only: true` annotation tells Speakeasy that a field should only be managed through the Terraform plan, not refreshed from the API. In version 1.658.1, Speakeasy correctly generates these fields as `Computed: true, Optional: true`. In version 1.661.0+, the `Computed: true` is omitted.

Without `Computed: true`:
- The API returns default values (e.g., `provision_policy: { unconfigured_provision: {} }`)
- These values are stored in Terraform state
- But the user's config doesn't specify `provision_policy`
- Terraform sees state != config and plans a change on every refresh

### The Overlay Is Correct

The fix is NOT in the OpenAPI spec or overlay. The overlay already has the correct annotation:

```yaml
- target: $["components"]["schemas"]["c1.api.policy.v1.ProvisionPolicy"]
  update:
    x-speakeasy-terraform-plan-only: true
```

This is a Speakeasy bug in how it interprets `x-speakeasy-terraform-plan-only`.

### Prevention

The Makefile enforces the pinned Speakeasy version before allowing generation. If you see a version mismatch error:

1. Install the correct version: check `.speakeasy/workflow.yaml` for `speakeasyVersion`
2. Or use `speakeasy run` which respects the workflow.yaml version pin

### History

- **Dec 2025**: Issue discovered, version pinned to 1.658.1 (commit c4503477)
- **Jan 2026**: Regression reintroduced when `make gen` was run with locally-installed 1.661.0

### Status

Upstream Speakeasy bug. Not fixed as of January 2026. Do not upgrade `speakeasyVersion` in workflow.yaml until confirmed fixed.
