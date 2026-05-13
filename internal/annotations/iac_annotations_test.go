package annotations

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestPlanModifier_DefaultsApplied(t *testing.T) {
	t.Setenv("TF_WORKSPACE", "")
	t.Setenv("OPENTOFU", "")
	t.Setenv("TF_FORK", "")

	pv, _ := types.MapValueFrom(context.Background(), types.StringType, map[string]string{})
	req := planmodifier.MapRequest{PlanValue: pv}
	var resp planmodifier.MapResponse

	PlanModifier().PlanModifyMap(context.Background(), req, &resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected diagnostics: %v", resp.Diagnostics)
	}

	got := mapValueToGoMap(t, resp.PlanValue)
	if got["managed_by"] != "terraform" {
		t.Errorf("managed_by: got %q, want %q", got["managed_by"], "terraform")
	}
	if _, present := got["iac_workspace"]; present {
		t.Errorf("iac_workspace should not be set when TF_WORKSPACE is empty; got %q", got["iac_workspace"])
	}
	if _, present := got["iac_tool_version"]; present {
		t.Errorf("iac_tool_version should not be auto-set; got %q", got["iac_tool_version"])
	}
}

func TestPlanModifier_UserValueWins(t *testing.T) {
	t.Setenv("OPENTOFU", "")
	t.Setenv("TF_FORK", "")

	pv, _ := types.MapValueFrom(context.Background(), types.StringType, map[string]string{
		"managed_by":  "opentofu",
		"team":        "platform",
		"criticality": "high",
	})
	req := planmodifier.MapRequest{PlanValue: pv}
	var resp planmodifier.MapResponse

	PlanModifier().PlanModifyMap(context.Background(), req, &resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected diagnostics: %v", resp.Diagnostics)
	}

	got := mapValueToGoMap(t, resp.PlanValue)
	if got["managed_by"] != "opentofu" {
		t.Errorf("managed_by overridden by default: got %q, want %q", got["managed_by"], "opentofu")
	}
	if got["team"] != "platform" {
		t.Errorf("user key not preserved: team=%q", got["team"])
	}
	if got["criticality"] != "high" {
		t.Errorf("user key not preserved: criticality=%q", got["criticality"])
	}
}

func TestPlanModifier_UserVersionPreserved(t *testing.T) {
	t.Setenv("OPENTOFU", "")
	t.Setenv("TF_FORK", "")

	pv, _ := types.MapValueFrom(context.Background(), types.StringType, map[string]string{
		"iac_tool_version": "1.8.2",
	})
	req := planmodifier.MapRequest{PlanValue: pv}
	var resp planmodifier.MapResponse

	PlanModifier().PlanModifyMap(context.Background(), req, &resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected diagnostics: %v", resp.Diagnostics)
	}

	got := mapValueToGoMap(t, resp.PlanValue)
	if got["iac_tool_version"] != "1.8.2" {
		t.Errorf("user-set iac_tool_version not preserved: got %q, want %q", got["iac_tool_version"], "1.8.2")
	}
}

func TestPlanModifier_OpenTofuDetectionViaEnv(t *testing.T) {
	t.Setenv("OPENTOFU", "true")
	if !isOpenTofu() {
		t.Error("OPENTOFU=true should be detected as OpenTofu")
	}
	if got := iacToolName(); got != "opentofu" {
		t.Errorf("iacToolName: got %q, want %q", got, "opentofu")
	}
}

func TestPlanModifier_OpenTofuDetectionViaTFFork(t *testing.T) {
	t.Setenv("OPENTOFU", "")
	t.Setenv("TF_FORK", "opentofu")
	if !isOpenTofu() {
		t.Error("TF_FORK=opentofu should be detected as OpenTofu")
	}
}

func TestPlanModifier_TerraformByDefault(t *testing.T) {
	t.Setenv("OPENTOFU", "")
	t.Setenv("TF_FORK", "")
	if isOpenTofu() {
		t.Error("no env signals: should default to Terraform")
	}
	if got := iacToolName(); got != "terraform" {
		t.Errorf("iacToolName default: got %q, want %q", got, "terraform")
	}
}

func TestPlanModifier_WorkspaceFromEnv(t *testing.T) {
	t.Setenv("TF_WORKSPACE", "prod/c1-config")
	t.Setenv("OPENTOFU", "")
	t.Setenv("TF_FORK", "")

	pv, _ := types.MapValueFrom(context.Background(), types.StringType, map[string]string{})
	req := planmodifier.MapRequest{PlanValue: pv}
	var resp planmodifier.MapResponse

	PlanModifier().PlanModifyMap(context.Background(), req, &resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected diagnostics: %v", resp.Diagnostics)
	}

	got := mapValueToGoMap(t, resp.PlanValue)
	if got["iac_workspace"] != "prod/c1-config" {
		t.Errorf("iac_workspace from TF_WORKSPACE: got %q, want %q", got["iac_workspace"], "prod/c1-config")
	}
}

func TestPlanModifier_NullPlanValueHandled(t *testing.T) {
	// A null plan value (e.g., the field isn't in the .tf at all) should
	// still produce a populated map with defaults applied.
	t.Setenv("TF_WORKSPACE", "")
	t.Setenv("OPENTOFU", "")
	t.Setenv("TF_FORK", "")

	req := planmodifier.MapRequest{PlanValue: types.MapNull(types.StringType)}
	var resp planmodifier.MapResponse

	PlanModifier().PlanModifyMap(context.Background(), req, &resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected diagnostics: %v", resp.Diagnostics)
	}

	got := mapValueToGoMap(t, resp.PlanValue)
	if got["managed_by"] != "terraform" {
		t.Errorf("managed_by: got %q, want %q", got["managed_by"], "terraform")
	}
}

func TestPlanModifier_DescriptionsNonEmpty(t *testing.T) {
	m, ok := PlanModifier().(iacAnnotationsModifier)
	if !ok {
		t.Fatalf("PlanModifier() returned unexpected type %T", PlanModifier())
	}
	if m.Description(context.Background()) == "" {
		t.Error("Description must not be empty")
	}
	if m.MarkdownDescription(context.Background()) == "" {
		t.Error("MarkdownDescription must not be empty")
	}
}

func TestSetIfMissing(t *testing.T) {
	tests := []struct {
		name     string
		initial  map[string]string
		k, v     string
		expected map[string]string
	}{
		{
			name:     "writes when missing",
			initial:  map[string]string{},
			k:        "managed_by",
			v:        "terraform",
			expected: map[string]string{"managed_by": "terraform"},
		},
		{
			name:     "preserves when present",
			initial:  map[string]string{"managed_by": "opentofu"},
			k:        "managed_by",
			v:        "terraform",
			expected: map[string]string{"managed_by": "opentofu"},
		},
		{
			name:     "no-op when value empty",
			initial:  map[string]string{},
			k:        "iac_workspace",
			v:        "",
			expected: map[string]string{},
		},
		{
			name:     "preserves other keys",
			initial:  map[string]string{"team": "platform"},
			k:        "managed_by",
			v:        "terraform",
			expected: map[string]string{"team": "platform", "managed_by": "terraform"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setIfMissing(tt.initial, tt.k, tt.v)
			if !mapsEqual(tt.initial, tt.expected) {
				t.Errorf("got %v, want %v", tt.initial, tt.expected)
			}
		})
	}
}

// --- helpers ---

func mapValueToGoMap(t *testing.T, mv types.Map) map[string]string {
	t.Helper()
	out := map[string]string{}
	elements := mv.Elements()
	for k, v := range elements {
		sv, ok := v.(types.String)
		if !ok {
			t.Fatalf("element %q is not a string: %T", k, v)
		}
		out[k] = sv.ValueString()
	}
	return out
}

func mapsEqual(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

// Verify the planmodifier.Map interface is satisfied at compile time.
var _ planmodifier.Map = (*iacAnnotationsModifier)(nil)
var _ planmodifier.Map = PlanModifier()

// attr.Value is imported by tests to make sure the public types work.
var _ attr.Value = types.StringValue("")
