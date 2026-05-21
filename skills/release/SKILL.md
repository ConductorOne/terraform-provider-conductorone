---
name: release
description: Spec-driven release workflow for this Speakeasy-generated Terraform provider. Documents the regen sequence and the six hand-maintained edges Speakeasy will not handle — provider.go registration, doc-emission ordering, README install-snippet version lag, breaking-changes log, post-merge tagging by a human, and post-tag release-note polish. Trigger on `/release`, "cut a release", "regen for vX.Y.Z", "bump version", or any time a Speakeasy regen is about to ship.
---

# Release workflow

The provider is Speakeasy-generated from `openapi.yaml`. Most of the release is mechanical, but **six places require hand intervention**. Skipping any of them ships a broken, misleading, or silently undocumented release.

## Canonical sequence

| Step | Command | Manual? | Notes |
|------|---------|---------|-------|
| 1 | edit `gen.yaml` → `version: X.Y.Z` | **yes** | Drives Speakeasy versioning. Commit alone. |
| 2 | `make gen` | no | Pulls fresh OAS from insulator, applies overlays, regenerates SDK + provider. Commit alone. |
| 3 | register new resources/data sources in `internal/provider/provider.go` | **yes** | See [Edge 1](#edge-1). Commit alone. |
| 4 | if release has breaking changes, prepend a `### X.Y.Z` block to `## Breaking changes` in `templates/index.md.tmpl` | **yes** | See [Edge 5](#edge-5). Commit alone. |
| 5 | `make generate` (tfplugindocs) | no | Emits `docs/**/*.md` including the rendered `index.md` from the template above. Must run *after* steps 3 and 4. Commit alone. |
| 6 | patch README install snippet to `X.Y.Z` | **yes** | See [Edge 3](#edge-3). Commit alone. |
| 7 | `make pre-commit` | no | Final gate: build + lint + test + generate + pagination check. |
| 8 | open PR → merge to `main` | manual | |
| 9 | **human** cuts the tag `vX.Y.Z` via the GitHub Releases UI | **human, not agent** | See [Edge 6](#edge-6). |
| 10 | agent offers to polish the release notes to match prior-release style | **yes, post-tag** | See [Edge 6](#edge-6). |

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

### Edge 5 — Breaking-changes log lives in a template, not in `docs/`

The canonical breaking-changes log is the `## Breaking changes` section of `templates/index.md.tmpl`. The rendered `docs/index.md` is generated from it by `make generate`, so **edit the template, never the rendered output** — and run `make generate` afterwards so the rendered file stays in sync.

If the release contains any user-visible breaking change (renamed attribute, removed attribute, type change, behavioral semantic change), prepend a new section to that log. Follow the prior-release format exactly:

```markdown
### X.Y.Z

- `conductorone_<resource>`: <one-sentence description of the break>. <Optional migration sentence: what the user must do.>

  | Resource | vPREV attribute | vX.Y.Z+ attribute |
  |---|---|---|
  | `conductorone_<resource>` | `<old>` | `<new>` |
```

The table is only needed when renames are involved; pure removals or behavior-only changes can omit it. See the existing v1.3.0, v1.1.1, and v1.0.0 entries for the style template, and the v1.1.0 "do not use" entry for how to document a botched release.

Skip this step entirely for additive-only releases. If unsure whether a change qualifies as breaking, default to documenting it — false positives are cheap, false negatives strand users.

### Edge 6 — Tagging is human-only; release-note polish is post-tag

The provider's GitHub release tags are cut **by a human via the GitHub Releases UI**, not by an agent and not by a CI workflow. The human typically writes one or two short bullets as starter notes. Do not attempt to cut the tag, push a tag, or call `gh release create` as part of the release workflow.

**After** the human has tagged and saved an initial release, the agent should offer to polish the notes to match the prior-release house style. The established structure (see `gh release view v1.3.0`):

```markdown
## Highlights

### <Feature 1 short title>

<2–4 sentence narrative on what shipped and why a user cares. Tables/limits/constraints
inline if they're tight enough to fit.>

### New API surface

- `<Service.Method>` — <one-line description>
- ...

### Schema changes

- **Breaking:** `conductorone_<resource>` <what changed>. <Migration note.>
- `conductorone_<resource>`: <additive change>
- ...

### Maintenance

- <dependency bumps, SDK version bump, hygiene>

## What's Changed
* <auto-generated by GitHub from PR titles>

**Full Changelog**: <auto-generated compare link>
```

Anything marked **Breaking** here must already appear in `templates/index.md.tmpl` (Edge 5) — the GH release notes are a summary; the template is the canonical log.

When offering to polish, propose the rewrite as a diff against the human's draft rather than overwriting their narrative. They may have framing the agent doesn't have context for.

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

Shipping v1.4.0 hit edges 1, 2, and 3 in sequence. Each one fails silently — the build is green, the tests are green, the PR looks fine, and the broken thing only surfaces when someone tries to use the new release. Edges 5 and 6 don't break the build at all; they just produce a release whose docs and notes don't match the code. Future agents reading this: do not skip the manual steps because the build passes, and do not try to cut the tag — that's a human gate.
