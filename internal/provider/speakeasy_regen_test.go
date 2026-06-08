package provider

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// TestPaginationUsesAssignment verifies that all data source pagination loops
// use assignment (=) not declaration (:=) for res.Next().
//
// Regression: older Speakeasy versions shadowed the loop variable with :=,
// causing pagination to re-fetch the first page forever.
func TestPaginationUsesAssignment(t *testing.T) {
	broken := regexp.MustCompile(`res,\s*err\s*:=\s*res\.Next\(\)`)
	files, err := filepath.Glob("*_data_source.go")
	if err != nil {
		t.Fatalf("globbing data source files: %v", err)
	}
	if len(files) == 0 {
		t.Fatal("no *_data_source.go files found — run test from internal/provider/")
	}
	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			t.Fatalf("reading %s: %v", f, err)
		}
		if broken.Match(data) {
			t.Errorf("%s: pagination loop uses := (variable shadowing), must use = assignment", f)
		}
	}
}

// TestPlanOnlyFieldsAreComputed verifies that fields annotated
// x-speakeasy-terraform-plan-only: true preserve plan-only behavior in
// their generated schema blocks.
//
// Background: Speakeasy 1.661.0–1.677.x dropped `Computed: true` on these
// fields, causing perpetual diffs. Upstream fix shipped in v1.678.1 via
// speakeasy-api/speakeasy#1770, which switched the implementation to an
// attribute-defined plan modifier (UseConfigValue) instead of `Computed: true`.
// Either pattern is acceptable for our purposes — both prevent the perpetual
// diff. The test fails only if neither is present.
func TestPlanOnlyFieldsAreComputed(t *testing.T) {
	data, err := os.ReadFile("custom_app_entitlement_resource.go")
	if err != nil {
		t.Fatalf("reading custom_app_entitlement_resource.go: %v", err)
	}
	content := string(data)

	// Fields annotated with x-speakeasy-terraform-plan-only: true in overlay.yaml.
	planOnlyFields := []string{
		"provision_policy",
		"duration_grant",
		"duration_unset",
	}

	for _, field := range planOnlyFields {
		// Match the schema definition specifically: "field_name": schema.
		// This avoids matching struct tags like `tfsdk:"field_name"`.
		schemaPattern := fmt.Sprintf(`"%s": schema.`, field)
		idx := strings.Index(content, schemaPattern)
		if idx == -1 {
			continue
		}
		// Either marker should appear within the next 300 chars of the schema definition.
		end := min(idx+300, len(content))
		window := content[idx:end]
		hasComputed := strings.Contains(window, "Computed:")
		hasUseConfigValue := strings.Contains(window, "UseConfigValue()")
		if !hasComputed && !hasUseConfigValue {
			t.Errorf("field %q: schema block missing both Computed: true and UseConfigValue() plan modifier (perpetual diff regression)", field)
		}
	}
}

// TestNullableNestedStructsHaveElseBranch verifies that SDK conversion code
// for nullable nested structs includes else branches that set fields to null.
//
// Regression: Speakeasy doesn't generate else branches for nil pointer structs,
// leaving terraform state fields as "unknown" instead of null.
func TestNullableNestedStructsHaveElseBranch(t *testing.T) {
	// Files known to reference ActorObjectPermissions
	sdkFiles := []string{
		"app_resource_resource_sdk.go",
		"app_resource_data_source_sdk.go",
		"app_entitlement_data_source_sdk.go",
		"custom_app_entitlement_resource_sdk.go",
		"access_review_resource_sdk.go",
		"access_review_data_source_sdk.go",
	}

	for _, f := range sdkFiles {
		data, err := os.ReadFile(f)
		if err != nil {
			// File may not exist in all versions
			continue
		}
		content := string(data)

		if !strings.Contains(content, "ActorObjectPermissions") {
			continue
		}

		// If the file accesses ActorObjectPermissions fields, it must also
		// contain BoolNull() calls for the else branch.
		if !strings.Contains(content, "BoolNull()") {
			t.Errorf("%s: accesses ActorObjectPermissions but has no BoolNull() else branch", f)
		}
	}
}

// TestNestedDeletedAtSchemaAttribute verifies that resource schemas nesting
// entity-root types (AppEntitlement, User, Directory) declare the deleted_at
// attribute in their nested SingleNestedAttribute block.
//
// Regression: Speakeasy's x-speakeasy-soft-delete-property handling correctly
// strips deleted_at from the resource schema when the field sits on the entity
// root, but on nested uses of the same shared struct it leaves the field on
// the Go struct while omitting the schema attribute. The resulting
// struct/schema mismatch crashes terraform plan/apply with
// "Mismatch between struct and object type: Struct defines fields not found
// in object: deleted_at". Patched by patches/02-deleted-at-on-nested-resources.patch.
func TestNestedDeletedAtSchemaAttribute(t *testing.T) {
	// Resource files that nest a *tfTypes.{AppEntitlement,User,Directory}
	// and therefore need deleted_at restored after every regen.
	resourceFiles := []string{
		"app_entitlement_owner_entitlement_resource.go",
		"app_owner_entitlement_resource.go",
		"app_entitlement_owner_user_resource.go",
		"app_owner_user_resource.go",
		"connector_owner_entitlement_resource.go",
		"connector_owner_user_resource.go",
		"directory_resource.go",
	}

	for _, f := range resourceFiles {
		data, err := os.ReadFile(f)
		if err != nil {
			t.Errorf("reading %s: %v", f, err)
			continue
		}
		if !strings.Contains(string(data), `"deleted_at":`) {
			t.Errorf("%s: nested schema missing \"deleted_at\" attribute (IGA-1774 regression); apply patches/02-deleted-at-on-nested-resources.patch", f)
		}
	}
}

// TestAppEntitlementSDKMirrorsManualProvisionSchema guards against schema/SDK
// drift on the hand-written conductorone_app_entitlement resource.
//
// Both app_entitlement_resource.go (schema) and app_entitlement_resource_sdk.go
// (Terraform <-> shared-types conversion) are hand-maintained — there is no
// upstream OAS binding that drives them. When commit 2b29cade added
// provisioner_assignment to the manual_provision schemas without updating the
// SDK, user-configured values were silently dropped on the write path and
// state never reflected the API on the read path. The build was green, every
// pre-existing test passed, and only a manual smoke test caught the perpetual
// diff and dropped writes.
//
// This tripwire checks that ProvisionerAssignment handling is present at all
// four required sites (write-provision, write-deprovision, read-provision,
// read-deprovision). It does NOT prove the field-by-field mapping is correct;
// it proves the four sites exist at all. If you intentionally remove
// provisioner_assignment from the schema, delete this test in the same change.
func TestAppEntitlementSDKMirrorsManualProvisionSchema(t *testing.T) {
	data, err := os.ReadFile("app_entitlement_resource_sdk.go")
	if err != nil {
		t.Fatalf("reading app_entitlement_resource_sdk.go: %v", err)
	}
	content := string(data)

	// Write side: Terraform plan -> *shared.ManualProvision. Both policies
	// must read ProvisionerAssignment from the plan and wire it into the
	// shared.ManualProvision{} struct literal sent to the API.
	writeMarkers := map[string]string{
		"r.ProvisionPolicy.ManualProvision.ProvisionerAssignment != nil":     "ProvisionPolicy write-side reads ProvisionerAssignment from plan",
		"r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment != nil": "DeprovisionerPolicy write-side reads ProvisionerAssignment from plan",
		"ProvisionerAssignment: provisionerAssignment,":                      "ProvisionPolicy ManualProvision struct literal wires ProvisionerAssignment",
		"ProvisionerAssignment: provisionerAssignment1,":                     "DeprovisionerPolicy ManualProvision struct literal wires ProvisionerAssignment",
	}
	for marker, what := range writeMarkers {
		if !strings.Contains(content, marker) {
			t.Errorf("write-side mapping missing (%s): %q not found. Schema declares provisioner_assignment but plan->API conversion does not handle it; user-configured values will be silently dropped on apply.", what, marker)
		}
	}

	// Read side: API response -> Terraform state. Both policies must
	// hydrate the tfTypes.ProvisionerAssignment from the response.
	readMarkers := map[string]string{
		"r.ProvisionPolicy.ManualProvision.ProvisionerAssignment = &tfTypes.ProvisionerAssignment{}":     "ProvisionPolicy read-side hydrates state from response",
		"r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment = &tfTypes.ProvisionerAssignment{}": "DeprovisionerPolicy read-side hydrates state from response",
	}
	for marker, what := range readMarkers {
		if !strings.Contains(content, marker) {
			t.Errorf("read-side mapping missing (%s): %q not found. Schema declares provisioner_assignment but API->state conversion does not populate it; refresh will null the field and produce a perpetual diff.", what, marker)
		}
	}
}

// TestPolicySDKUsesBoolPointerOrFalse verifies that the policy resource and
// data source SDK files use typeconvert.BoolPointerOrFalse instead of
// types.BoolPointerValue. When protobuf omits a zero-valued optional bool,
// BoolPointerValue returns null while state holds false, causing perpetual
// drift on every plan.
func TestPolicySDKUsesBoolPointerOrFalse(t *testing.T) {
	files := []string{
		"policy_resource_sdk.go",
		"policy_data_source_sdk.go",
	}
	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			t.Fatalf("reading %s: %v", f, err)
		}
		content := string(data)
		if strings.Contains(content, "types.BoolPointerValue(") {
			t.Errorf("%s: still uses types.BoolPointerValue — must use typeconvert.BoolPointerOrFalse to prevent null/false drift", f)
		}
		if !strings.Contains(content, "typeconvert.BoolPointerOrFalse(") {
			t.Errorf("%s: does not use typeconvert.BoolPointerOrFalse — approval-block bools will drift false→null on every refresh", f)
		}
		// RequiresStepUpProviderID is the lone optional string with the same
		// zero-value/null drift; guard its coalescing helper too. We can't
		// assert absence of types.StringPointerValue here — it's correct for
		// genuinely-nullable IDs/paths in these files.
		if !strings.Contains(content, "typeconvert.StringPointerOrEmpty(") {
			t.Errorf("%s: does not use typeconvert.StringPointerOrEmpty — requires_step_up_provider_id will drift \"\"→null on every refresh", f)
		}
	}
}
