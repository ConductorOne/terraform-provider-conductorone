# Skill: Speakeasy Regeneration Audit

## Purpose

After running `make gen` or `speakeasy run`, audit the generated code for known regressions, missing fixes, and broken patterns. Fix them. Run tests to confirm.

## Trigger

User asks to:
- "Check the generated code"
- "Audit after regen"
- "Fix speakeasy output"
- "Verify regen"
- Or: any time `make gen` or `speakeasy run` was just executed

## Known Issues Registry

Each issue has: detection pattern, fix pattern, affected files, and a test strategy.

---

### Issue 1: Pagination Variable Shadowing

**Severity:** Critical - data sources return only the first page of results

**Detection:**
```
grep -r 'res, err := res\.Next()' internal/provider/*_data_source.go
```

If any matches: the pagination fix has regressed.

**Broken pattern:**
```go
for {
    res, err := res.Next()  // := creates new local variable, shadows outer res
    if err != nil { ... }
    if res == nil { break }
    // processes the SAME first page forever
}
```

**Correct pattern:**
```go
for {
    var err error
    res, err = res.Next()   // = assigns to outer res, advances pagination
    if err != nil { ... }
    if res == nil { break }
    // processes successive pages
}
```

**Affected files:** All `*_data_source.go` files with pagination loops.

**Fix:** In each match, replace `res, err := res.Next()` with a two-line `var err error` + `res, err = res.Next()`.

**Test:** `make verify-pagination` (already in Makefile). Also: any data source that returns >100 items would only return the first page with the broken pattern. A test could list apps and verify count matches expectations.

---

### Issue 2: Perpetual Diff on plan-only Fields (Missing `Computed: true`)

**Severity:** Critical - terraform plan shows changes on every run

**Detection:**
Scan resource schema definitions for fields that should be `Computed: true, Optional: true` but are only `Optional: true`. The affected fields use `x-speakeasy-terraform-plan-only: true` in the overlay:

```
grep -B5 'tfPlanOnly:"true"' internal/provider/custom_app_entitlement_resource.go
```

Then for each field name found, verify the schema block includes `Computed: true`:

```
# These schema fields MUST have Computed: true:
# - provision_policy (and all nested children)
# - duration_grant (and all nested children)
# - duration_unset (and all nested children)
# - All AppEntitlementAutomationRule* schemas
# - Approval schema
```

**Broken pattern:**
```go
"provision_policy": schema.SingleNestedAttribute{
    Optional: true,
    // Missing: Computed: true
    Attributes: map[string]schema.Attribute{ ... },
},
```

**Correct pattern:**
```go
"provision_policy": schema.SingleNestedAttribute{
    Computed: true,
    Optional: true,
    Attributes: map[string]schema.Attribute{ ... },
},
```

**Affected files:**
- `internal/provider/custom_app_entitlement_resource.go`
- Any resource with fields annotated `x-speakeasy-terraform-plan-only: true`

**Fix:** Add `Computed: true` to every schema attribute (including nested children) for plan-only fields.

**Test:** `make testacc` (acceptance tests) will catch this — the `TestAccBundleAutomationResource` test specifically fails with "refresh plan was not empty" when `Computed: true` is missing. Unit test: parse the schema and assert `Computed == true` for known plan-only fields.

---

### Issue 3: Nil Pointer on Nullable Nested Structs (ActorObjectPermissions)

**Severity:** High - terraform refresh crashes or leaves fields as "unknown"

**Detection:**
```
grep -rn 'resp\.\w*\.ActorObjectPermissions\.' internal/provider/*_sdk.go | \
  grep -v 'if.*!= nil'
```

Find lines that access `resp.Something.ActorObjectPermissions.Field` without a preceding nil check on `ActorObjectPermissions`. Alternatively, search for `_sdk.go` files that reference `ActorObjectPermissions` but lack an `else` branch setting fields to null:

```
# For each _sdk.go file referencing ActorObjectPermissions:
# Verify there is an else branch like:
#   } else {
#       r.Delete = types.BoolNull()
#       r.Edit = types.BoolNull()
#       r.Extra = nil
#       r.Read = types.BoolNull()
#   }
```

**Broken pattern:**
```go
if resp.ActorObjectPermissions != nil {
    r.Delete = types.BoolPointerValue(resp.ActorObjectPermissions.Delete)
    r.Edit = types.BoolPointerValue(resp.ActorObjectPermissions.Edit)
    // ... handle Extra map ...
    r.Read = types.BoolPointerValue(resp.ActorObjectPermissions.Read)
}
// NO else branch — fields stay "unknown" in terraform state
```

**Correct pattern:**
```go
if resp.ActorObjectPermissions != nil {
    r.Delete = types.BoolPointerValue(resp.ActorObjectPermissions.Delete)
    r.Edit = types.BoolPointerValue(resp.ActorObjectPermissions.Edit)
    if len(resp.ActorObjectPermissions.Extra) > 0 {
        r.Extra = make(map[string]types.Bool, len(resp.ActorObjectPermissions.Extra))
        for key, value := range resp.ActorObjectPermissions.Extra {
            r.Extra[key] = types.BoolValue(value)
        }
    } else {
        r.Extra = nil
    }
    r.Read = types.BoolPointerValue(resp.ActorObjectPermissions.Read)
} else {
    r.Delete = types.BoolNull()
    r.Edit = types.BoolNull()
    r.Extra = nil
    r.Read = types.BoolNull()
}
```

**Affected files:**
- `internal/provider/app_resource_resource_sdk.go`
- `internal/provider/app_resource_data_source_sdk.go`
- `internal/provider/app_entitlement_data_source_sdk.go`
- `internal/provider/custom_app_entitlement_resource_sdk.go`
- `internal/provider/access_review_resource_sdk.go`
- `internal/provider/access_review_data_source_sdk.go`

**Fix:** Add the `else` branch with explicit null assignments for every nullable nested struct access.

**Test:** Create or read a resource where the API returns `actor_object_permissions: null`. Verify terraform state has explicit null values (not "unknown"). The acceptance tests should catch this if the test tenant has resources without permissions set.

---

### Issue 4: int64 String Unmarshaling in SDK

**Severity:** High - API calls fail for resources with string-encoded int64 fields

**Detection:**
```
grep -rn 'integer:"string"' internal/sdk/models/
```

Then check `internal/sdk/utils/json.go` (or equivalent unmarshaling code) for whether `topLevel=true` bypasses the `integer:"string"` tag handling.

**Broken pattern:**
`UnmarshalJSON` with `topLevel=true` calls `json.Decoder.Decode()` directly, bypassing custom handling for `integer:"string"` struct tags. Fields like `Escalation.Expiration` receive `"3600"` (string) but expect `int64`.

**Correct pattern:**
The unmarshaling path must handle `integer:"string"` tags regardless of the `topLevel` flag. If the generated `utils/json.go` has a `topLevel` short-circuit, it needs to still process string-to-int64 conversion for tagged fields.

**Affected files:**
- `internal/sdk/utils/json.go` (or wherever `UnmarshalJSON` is defined)
- Any model with `integer:"string"` tagged fields

**Fix:** Ensure the top-level unmarshal path doesn't skip `integer:"string"` handling. This was fixed upstream but could regress.

**Test:** Unit test that unmarshals a JSON response containing a string-encoded int64 field and verifies it parses correctly:
```go
func TestStringEncodedInt64Unmarshal(t *testing.T) {
    input := `{"expiration": "3600"}`
    var result Escalation
    err := UnmarshalJSON([]byte(input), &result)
    require.NoError(t, err)
    assert.Equal(t, int64(3600), result.Expiration)
}
```

---

### Issue 5: Type Override Generates Nonexistent Types

**Severity:** Medium - compilation failure, caught by `make build`

**Detection:**
```
make build 2>&1 | grep "undefined:"
```

Or specifically: after regen, check `*_resource_sdk.go` files for references to types that don't exist in `internal/sdk/models/shared/`:

```
# Look for patterns like:
#   var config *shared.SomethingConfig
# where SomethingConfig doesn't exist as a type
grep -rn 'var .* \*shared\.\w*Config' internal/provider/*_sdk.go
```

Cross-reference against actual types in `internal/sdk/models/shared/`.

**Context:** When `x-speakeasy-type-override: any` is applied to a nested property, Speakeasy occasionally generates SDK conversion code referencing a fabricated type name (e.g., `AccountProvisionConfig` instead of `any`). This is intermittent — re-running generation sometimes fixes it.

**Fix:** Replace the fabricated type reference with `any` and use `json.Unmarshal` for the conversion, matching the pattern used for other `jsontypes.Normalized` fields.

**Test:** `make build` catches this immediately. No special test needed beyond compilation.

---

## Audit Procedure

Run these checks in order after any `make gen` or `speakeasy run`:

1. **Build first:** `make build` — catches type errors (Issue 5)
2. **Pagination check:** `make verify-pagination` — catches Issue 1
3. **Plan-only fields:** Grep for `tfPlanOnly:"true"` fields and verify `Computed: true` in their schema definitions — catches Issue 2
4. **Nil pointer else branches:** For each `*_sdk.go` file touching nullable nested structs, verify else branches exist — catches Issue 3
5. **String int64:** Check unmarshal code for top-level bypass — catches Issue 4
6. **Run tests:** `make test` then `make testacc` if credentials available
7. **Diff review:** `git diff` the generated files. Look for anything unexpected: removed fields, changed types, new unknown patterns.

## Writing Regression Tests

For each issue, a Go test can be added to `internal/provider/speakeasy_regen_test.go`:

```go
package provider

import (
    "os"
    "path/filepath"
    "regexp"
    "strings"
    "testing"
)

// TestPaginationUsesAssignment verifies that all data source pagination loops
// use assignment (=) not declaration (:=) for res.Next().
// Regression: older Speakeasy versions shadowed the loop variable.
func TestPaginationUsesAssignment(t *testing.T) {
    broken := regexp.MustCompile(`res,\s*err\s*:=\s*res\.Next\(\)`)
    files, _ := filepath.Glob("*_data_source.go")
    for _, f := range files {
        data, err := os.ReadFile(f)
        if err != nil {
            t.Fatalf("reading %s: %v", f, err)
        }
        if broken.Match(data) {
            t.Errorf("%s: pagination loop uses := (variable shadowing). Must use = assignment. See CLAUDE.md.", f)
        }
    }
}

// TestPlanOnlyFieldsAreComputed verifies that fields with tfPlanOnly:"true"
// have Computed: true in their schema definitions.
// Regression: Speakeasy 1.661.0+ drops Computed: true for plan-only fields.
func TestPlanOnlyFieldsAreComputed(t *testing.T) {
    // The custom_app_entitlement_resource.go defines schema for plan-only fields.
    // If Computed: true is missing, terraform shows perpetual diffs.
    data, err := os.ReadFile("custom_app_entitlement_resource.go")
    if err != nil {
        t.Fatalf("reading custom_app_entitlement_resource.go: %v", err)
    }
    content := string(data)

    // These field names must appear near a Computed: true in their schema block
    planOnlyFields := []string{
        "provision_policy",
        "duration_grant",
        "duration_unset",
    }

    for _, field := range planOnlyFields {
        // Find the schema definition for this field
        idx := strings.Index(content, `"`+field+`"`)
        if idx == -1 {
            t.Errorf("field %q not found in schema", field)
            continue
        }
        // Check that Computed: true appears within the next 200 chars
        window := content[idx:min(idx+200, len(content))]
        if !strings.Contains(window, "Computed:") {
            t.Errorf("field %q schema block missing Computed: true (perpetual diff regression)", field)
        }
    }
}
```

## Version Pin Enforcement

The Makefile already enforces the Speakeasy version pin. If a future version fixes the regressions, update:
1. `.speakeasy/workflow.yaml` — `speakeasyVersion: X.Y.Z`
2. Run `make gen`
3. Run this full audit
4. Run `make testacc`
5. If all pass, commit the version bump with a note about which issues were verified fixed
