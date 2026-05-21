---
name: upgrading
description: Upgrade a Terraform configuration to a newer version of the ConductorOne provider safely — read release notes, apply attribute migrations, validate with plan, and recover from known issues
---

# upgrading

Use this skill when a user wants to bump the `ConductorOne/conductorone` provider version in their Terraform / OpenTofu configuration and needs to handle breaking changes, attribute renames, or known issues across the v1.0.0 → current range.

This skill is for **consumers of the provider** (people writing HCL against ConductorOne). For maintainer-side regen/release workflow, see `CLAUDE.md` in this repo.

## Workflow

### 1. Identify current and target versions

Find the user's pinned version:

```bash
grep -rE "ConductorOne/conductorone" --include="*.tf" --include="*.tofu"
# or check .terraform.lock.hcl
grep -A2 'ConductorOne/conductorone' .terraform.lock.hcl
```

Identify the target version (latest stable unless the user specifies one):

```bash
gh release list --repo ConductorOne/terraform-provider-conductorone --limit 5
```

### 2. Read the breaking-changes log

The canonical list lives in `docs/index.md` under `## Breaking changes` (rendered from `templates/index.md.tmpl`).

For releases that span multiple versions, read the release notes for every version in the path:

```bash
for v in <list of versions between current and target>; do
  echo "===== $v ====="
  gh release view $v --repo ConductorOne/terraform-provider-conductorone --json body -q .body
done
```

Releases that **always** require attention when crossing them:

- **v1.0.0** — `request_catalog` → `access_profile`; datasource field renames for `app`, `app_entitlement`, `policy`, `user`, `webhook`, `app_resource_type`.
- **v1.1.0** — **skip this version**. It was a broken release with unintended attribute renames. Upgrade `v1.0.x` → `v1.1.1` directly.
- **v1.1.1** — `conductorone_access_conflict.notification_config` → `access_conflict_notification_config`.
- **v1.3.0** — `conductorone_access_profile.grant_policy_id` removed. Move grant policy management to `AppEntitlement.grant_policy_id` in the catalog.

### 3. Apply HCL migrations

For each breaking change in scope, search for usages and apply renames:

```bash
# Example: v1.1.1 access_conflict rename
grep -rn "notification_config" --include="*.tf" | grep -B5 conductorone_access_conflict
```

Prefer `terraform fmt` after edits. If a block is removed (e.g. `grant_policy_id` in v1.3.0), check whether the field needs to migrate to a different resource or just be deleted.

### 4. Recover from v1.1.0 if the user adopted it

If `terraform state list` or HCL contains the v1.1.0 incorrect names, revert them before bumping to v1.1.1+:

| Resource | v1.1.0 (incorrect) | v1.1.1+ (correct) |
|---|---|---|
| `conductorone_policy` (`policy_steps.steps[]`) | `policy_form` | `form` |
| `conductorone_access_review` | `access_review_notification_config` | `notification_config` |
| `conductorone_access_review_template` | `access_review_notification_config` | `notification_config` |
| `conductorone_automation` (`automation_steps[]`, `draft_automation_steps[]`) | `automations_task_action` | `task_action` |
| `conductorone_automation` (`automation_steps[]`, `draft_automation_steps[]`) | `automations_webhook` | `webhook` |
| `conductorone_request_schema` (`form.fields[]`) | `form_string_field` | `string_field` |
| `conductorone_request_schema` (`form.fields[]`) | `form_string_map_field` | `string_map_field` |

### 5. Bump the version constraint

Update both the `required_providers` block and the lock file:

```hcl
terraform {
  required_providers {
    conductorone = {
      source  = "ConductorOne/conductorone"
      version = "~> <target>"
    }
  }
}
```

```bash
terraform init -upgrade
```

### 6. Validate

Run a plan against a non-production workspace first:

```bash
terraform validate
terraform plan -out=upgrade.plan
```

Inspect the plan carefully. **Expected churn after a major bump** typically includes new computed attributes appearing (e.g. `annotations` in v1.3.0+). **Unexpected churn** (resources being destroyed/recreated, unrelated attribute changes) usually signals a missed migration — go back to step 3.

### 7. Watch for known issues

#### Approval-block plan-only drift (v1.1.1+)

If the user explicitly sets any of these bool fields to `false`:

- `allow_delegation`, `assigned`, `escalation_enabled`, `require_denial_reason`, `allowed_reassignees`
- `app_owner_approval.require_distinct_approvers`

…on a `conductorone_policy` Approval block, Terraform will show a perpetual `false -> null` diff. **Workaround:** remove the explicit `false`; let the field default. Tracked upstream at [speakeasy-api/speakeasy#2031](https://github.com/speakeasy-api/speakeasy/issues/2031).

#### `deleted_at` schema mismatch (v1.2.0 only)

`terraform apply` crashes with `Mismatch between struct and object type: Struct defines fields not found in object: deleted_at` on the four new app-owner resources. **Fix:** upgrade straight to v1.2.1 or later — skip v1.2.0.

### 8. Apply

Only after a clean plan in a lower environment:

```bash
terraform apply upgrade.plan
```

If apply fails partway, do not roll back the provider version — the new schema is already in state. Fix forward.

## Quick reference: safe upgrade paths

| From | Recommended target | Skip |
|---|---|---|
| `v0.4.x` | `v1.3.0` (via `v1.0.0` migration first) | `v1.1.0` |
| `v1.0.x` | `v1.3.0` | `v1.1.0` |
| `v1.1.0` | `v1.1.1` then `v1.3.0` | — |
| `v1.1.1` / `v1.2.0` | `v1.3.0` | `v1.2.0` (use `v1.2.1`) if not yet adopted |
| `v1.2.1` | `v1.3.0` | — |

## When to escalate

Open an issue at https://github.com/ConductorOne/terraform-provider-conductorone/issues when:

- The plan shows unexplained drift after applying every documented migration.
- An attribute is missing from the new version with no documented replacement.
- The provider panics during `plan` or `apply`.

Include `terraform version`, the source and target provider versions, and a minimal reproducing HCL snippet.
