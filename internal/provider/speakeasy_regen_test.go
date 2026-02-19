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

// TestPlanOnlyFieldsAreComputed verifies that fields with tfPlanOnly:"true"
// have Computed: true in their schema definitions.
//
// Regression: Speakeasy 1.661.0+ drops Computed: true for plan-only fields,
// causing perpetual diffs on every terraform plan.
func TestPlanOnlyFieldsAreComputed(t *testing.T) {
	data, err := os.ReadFile("custom_app_entitlement_resource.go")
	if err != nil {
		t.Fatalf("reading custom_app_entitlement_resource.go: %v", err)
	}
	content := string(data)

	// Fields annotated with x-speakeasy-terraform-plan-only: true in overlay.yaml.
	// Their schema blocks must include Computed: true.
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
		// Computed: true should appear within the next 200 chars of the schema definition.
		end := min(idx+200, len(content))
		window := content[idx:end]
		if !strings.Contains(window, "Computed:") {
			t.Errorf("field %q: schema block missing Computed: true (perpetual diff regression)", field)
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
