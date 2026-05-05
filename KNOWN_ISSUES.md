# Known Issues

## Speakeasy `x-speakeasy-terraform-plan-only` regression — FIXED upstream

### Status

**Fixed upstream** in [speakeasy-api/speakeasy#1770](https://github.com/speakeasy-api/speakeasy/pull/1770) (merged 2025-12-17). The fix converts `x-speakeasy-terraform-plan-only` from a fragile schema-level flag into an attribute-defined plan modifier. Any Speakeasy release after that merge includes it. The provider currently pins **1.761.10**, which has the fix. The original "do not upgrade" warning that lived here applied to versions 1.658.1–1.661.0 only.

### Background

For history: late 2025 versions silently dropped `Computed: true` on fields tagged `x-speakeasy-terraform-plan-only`, producing perpetual diffs on resources like `conductorone_custom_app_entitlement` (visible as `provision_policy`, `duration_grant`, `duration_unset` flapping on every refresh). Robert Chiniquy hit this in Dec 2025, defensively pinned 1.658.1, and the pin held until 1.761.10 in PR #199 (2026-04-30). The pin was applied 32 days before the upstream fix landed; nobody re-tested after the fix shipped, so the pin outlived the bug by ~5 months.

## Speakeasy nullable-nested-struct flatten regression — STILL PRESENT

### Summary

When a nullable nested struct in the OpenAPI spec is "flattened" into top-level Computed attributes on a Terraform resource (e.g. `ActorObjectPermissions { delete, edit, read, extra }` → top-level `delete`, `edit`, `read`, `extra` on `conductorone_app_resource`), Speakeasy generates SDK conversion code that handles the non-nil case but **omits the `else` branch that nulls the fields** when the parent struct is nil:

```go
// Generated (broken):
if resp.ActorObjectPermissions != nil {
    r.Delete = types.BoolPointerValue(resp.ActorObjectPermissions.Delete)
    r.Edit  = types.BoolPointerValue(resp.ActorObjectPermissions.Edit)
    r.Read  = types.BoolPointerValue(resp.ActorObjectPermissions.Read)
}
// missing: } else { r.Delete = types.BoolNull() ... }
```

When the API returns `nil` for the parent, those fields stay at their zero `types.Bool{}` (which is `Unknown`, not `Null`). The terraform-plugin-framework rejects Unknown after Refresh with `Provider returned invalid result object after apply`.

### Affected Files (current count: 6)

- `internal/provider/app_resource_resource_sdk.go`
- `internal/provider/app_resource_data_source_sdk.go`
- `internal/provider/app_entitlement_data_source_sdk.go`
- `internal/provider/custom_app_entitlement_resource_sdk.go`
- `internal/provider/access_review_resource_sdk.go`
- `internal/provider/access_review_data_source_sdk.go`

### The Patch

After every Speakeasy regen, append the missing `else` branch in each of the 6 files:

```go
} else {
    r.Delete = types.BoolNull()
    r.Edit = types.BoolNull()
    r.Extra = nil
    r.Read = types.BoolNull()
}
```

In `app_resource_resource_sdk.go` there is also a `RefreshFromSharedCreateManuallyManagedAppResourceResponse` variant: the parent type (`AppResource`, not `AppResourceView`) lacks `ActorObjectPermissions` entirely, so the four null assignments must be appended unconditionally at the end of the `if resp != nil` block.

### Tripwire

`internal/provider/speakeasy_regen_test.go` runs `TestNullableNestedStructsHaveElseBranch`, which scans the 6 files for `BoolNull()`. If a regen ships without the patch, this test fails — catch it before merging the regen PR.

### Status

Filed upstream as [speakeasy-api/speakeasy#2031](https://github.com/speakeasy-api/speakeasy/issues/2031) on 2026-04-30 with a minimal repro and our reproduction history. The generator template lives in the private `speakeasy-api/openapi-generation` repo, so the public CLI repo is the only place we can file. Robert previously applied the same hand-patch in [PR #184](https://github.com/ConductorOne/terraform-provider-conductorone/pull/184) (2026-02-26) without filing upstream; the patch ritual continues until the upstream issue is fixed.

We confirmed `terraform.respectRequiredFields: true` in `gen.yaml` does NOT fix this — that flag improves nullable handling for ordinary nested struct fields but does not affect the flatten-promotion case. Documented in the upstream issue.
