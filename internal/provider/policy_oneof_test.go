package provider

import (
	"slices"
	"testing"
)

const multiStepJSON = `{"provisionSteps":[{"delegated":{"appId":"app2","entitlementId":"ent2"}}]}`

// testOneofArms returns arms backed by set, which records each arm's state.
func testOneofArms(set map[string]bool) []oneofArm {
	var arms []oneofArm
	for _, name := range []string{"a", "b", "c"} {
		arms = append(arms, oneofArm{
			name:  name,
			isSet: func() bool { return set[name] },
			clear: func() { set[name] = false },
		})
	}
	return arms
}

func TestSelectedOneofArm(t *testing.T) {
	tests := []struct {
		name    string
		set     map[string]bool
		want    string
		wantErr bool
	}{
		{"no arms", nil, "", false},
		{"none set", map[string]bool{}, "", false},
		{"one set", map[string]bool{"b": true}, "b", false},
		{"two set", map[string]bool{"a": true, "c": true}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var arms []oneofArm
			if tt.set != nil {
				arms = testOneofArms(tt.set)
			}
			got, diags := selectedOneofArm("provision_policy", arms)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
			if diags.HasError() != tt.wantErr {
				t.Errorf("HasError() = %v, want %v: %v", diags.HasError(), tt.wantErr, diags)
			}
		})
	}
}

func TestKeepOnlyOneofArm(t *testing.T) {
	set := map[string]bool{"a": true, "b": true, "c": true}
	keepOnlyOneofArm(testOneofArms(set), "b")
	var kept []string
	for name, on := range set {
		if on {
			kept = append(kept, name)
		}
	}
	if !slices.Equal(kept, []string{"b"}) {
		t.Errorf("kept %v, want [b]", kept)
	}

	set = map[string]bool{"a": true, "b": true}
	keepOnlyOneofArm(testOneofArms(set), "")
	if !set["a"] || !set["b"] {
		t.Errorf("empty keep cleared arms: %v", set)
	}

	keepOnlyOneofArm(nil, "b")
}
