---
name: release
description: Spec-driven release workflow for this Speakeasy-generated Terraform provider. Documents the regen sequence and the four hand-maintained edges Speakeasy will not handle — provider.go registration, doc-emission ordering, README install-snippet version lag, and commit segmentation. Trigger on `/release`, "cut a release", "regen for vX.Y.Z", "bump version", or any time a Speakeasy regen is about to ship.
---

# Release workflow

The provider is Speakeasy-generated from `openapi.yaml`. Most of the release is mechanical, but four places require hand intervention. Skipping any of them ships a broken or misleading release.

## Canonical sequence

| Step | Command | Manual? | Notes |
|------|---------|---------|-------|
| 1 | edit `gen.yaml` → `version: X.Y.Z` | **yes** | Drives Speakeasy versioning. Commit alone. |
| 2 | `make gen` | no | Pulls fresh OAS from insulator, applies overlays, regenerates SDK + provider. Commit alone. |
| 3 | register new resources/data sources in `internal/provider/provider.go` | **yes** | See [Edge 1](#edge-1). Commit alone. |
| 4 | `make generate` (tfplugindocs) | no | Emits `docs/**/*.md`. Must run *after* step 3. Commit alone. |
| 5 | patch README install snippet to `X.Y.Z` | **yes** | See [Edge 3](#edge-3). Commit alone. |
| 6 | `make pre-commit` | no | Final gate: build + lint + test + generate + pagination check. |
| 7 | open PR → merge → tag `vX.Y.Z` | manual | Tag triggers the published-release Speakeasy template. |

**Segment commits** (per repo `CLAUDE.md`): manual edits before codegen; codegen alone; tfplugindocs alone. Eases rebases when you have to re-run a step.

---

## Sharp edges

### Edge 1 — `internal/provider/provider.go` is outside Speakeasy

Speakeasy will produce new `*_resource.go` / `*_data_source.go` files for any new resource family it discovers in the spec, but it will **not** add the `NewXResource` / `NewXDataSource` constructors to the registration slices in `provider.go`. The build still passes; only `terraform plan` reveals that the resource is unreachable.

**Detect missing registrations** after step 2:

```bash
# Constructors that exist in the package
sg --lang go --pattern 'func $NAME() resource.Resource { $$$ }' internal/provider/ \
  | rg -o 'func \w+' | sort -u > /tmp/ctors_resource.txt
sg --lang go --pattern 'func $NAME() datasource.DataSource { $$$ }' internal/provider/ \
  | rg -o 'func \w+' | sort -u > /tmp/ctors_datasource.txt

# Constructors that are wired in provider.go
rg -o 'New\w+(?:Resource|DataSource)' internal/provider/provider.go | sort -u > /tmp/registered.txt

# Unregistered = present in package but missing from provider.go
diff /tmp/ctors_resource.txt /tmp/registered.txt
diff /tmp/ctors_datasource.txt /tmp/registered.txt
```

Or, blunter:

```bash
for fn in $(rg -o '^func (New\w+(?:Resource|DataSource))' internal/provider/ -r '$1' --no-filename | sort -u); do
  rg -q "\\b${fn}\\b" internal/provider/provider.go || echo "MISSING: $fn"
done
```

Add the missing constructors to the right slice in `provider.go`. Group with similar resources (owner-family next to existing owner-family entries, etc.).

> **Watch-out:** `provider.go` has *two* slices — `Resources()` around line 178 and `DataSources()` around line 220. A new resource family typically needs entries in both.

### Edge 2 — tfplugindocs only emits docs for *registered* resources

`make generate` runs tfplugindocs, which walks the provider schema by instantiating the registered resources. Resources present in the package but missing from `provider.go` get **no docs**. This means the doc-regen has to run *after* the wire-up, not before.

If you ran step 4 before step 3, just re-run step 4 — there is no harm. But it's a wasted commit cycle.

### Edge 3 — README install snippet version is driven by the latest git tag, not `gen.yaml`

The Speakeasy template that renders the `<!-- Start Installation [installation] -->` block in `README.md` pulls the version from the most recent **published release tag**, not from `gen.yaml`. So after step 1 bumps `gen.yaml` to X.Y.Z, the README still shows the *previous* version until `vX.Y.Z` is tagged on GitHub.

This is backwards: from "1.4.0 merged to main" until "v1.4.0 tagged", `README.md` documents 1.3.0 as the current install target — exactly when users would be looking it up.

**Workaround:** hand-edit the line during step 5:

```bash
sd 'version = "[0-9]+\.[0-9]+\.[0-9]+"' "version = \"$(yq '.generation.sdkClassName // .generation.version // .' gen.yaml | head -1)\"" README.md
# Or just:
sd 'version = "1\.3\.0"' 'version = "1.4.0"' README.md
```

The patch lives inside the Speakeasy-managed block, so it **will be clobbered on the next `make gen`** unless the upstream template is fixed. Re-apply on every release until that happens.

### Edge 4 — Pagination fix verification

Older Speakeasy versions emitted variable-shadowed pagination loops (`res, err := res.Next()` instead of `res, err = res.Next()`), silently breaking multi-page iteration in data sources. There is a make-time check (`make pre-commit` prints `Pagination pattern OK`) and a Go test (`TestPaginationUsesAssignment`) — both must pass before merging.

If either fails after a regen, see `KNOWN_ISSUES.md` and the `speakeasy-regen-audit` skill in this same directory for the manual fixup pattern.

---

## Verification before PR

```bash
make pre-commit                             # build + lint + test + generate
git status                                  # should be clean — generate produced no untracked files
git log --format='%s' main..HEAD            # commits should be one-purpose each
grep -E '^      version = "[0-9.]+"' README.md  # must match gen.yaml
```

If `make generate` left untracked or modified files, commit them — that means you didn't run step 4 to completion.

## PR description template

```markdown
## Summary
Release vX.Y.Z: regen against the current insulator OpenAPI spec, plus the manual wire-up needed to expose <N> new resource families.

### New HCL surface
- `conductorone_<name>` — resource + data source
- ...

### New SDK-only surface (no HCL exposure this release)
- `internal/sdk/<file>.go` — <one-liner>
- ...

### Other notable spec changes
- <model> — additional fields
- ...

## Test plan
- [ ] `make pre-commit` passes locally
- [ ] Smoke-test `terraform plan` against the new resources in a scratch tenant
- [ ] Pagination check `Pagination pattern OK` in pre-commit output
```

## Why this skill exists

Shipping v1.4.0 hit edges 1, 2, and 3 in sequence. Each one fails silently — the build is green, the tests are green, the PR looks fine, and the broken thing only surfaces when someone tries to use the new release. Future agents reading this: do not skip the manual steps because the build passes.
