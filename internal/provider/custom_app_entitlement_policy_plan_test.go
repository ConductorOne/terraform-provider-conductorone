package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestUnchangedPolicyAliasPlan(t *testing.T) {
	obj := tftypes.Object{AttributeTypes: map[string]tftypes.Type{"name": tftypes.String, "computed": tftypes.String}}
	makeValue := func(name, computed any) tftypes.Value {
		return tftypes.NewValue(obj, map[string]tftypes.Value{
			"name": tftypes.NewValue(tftypes.String, name), "computed": tftypes.NewValue(tftypes.String, computed),
		})
	}
	state := makeValue("original", "server")
	for _, tt := range []struct {
		name                string
		config, plan, state tftypes.Value
		want                bool
	}{
		{"unknown computed only", makeValue("original", nil), makeValue("original", tftypes.UnknownValue), state, true},
		{"configured change", makeValue("changed", nil), makeValue("changed", tftypes.UnknownValue), state, false},
		{"unknown configuration", makeValue(tftypes.UnknownValue, nil), makeValue(tftypes.UnknownValue, tftypes.UnknownValue), state, false},
		{"known removal", makeValue(nil, nil), makeValue(nil, tftypes.UnknownValue), state, false},
		{"unknown prior value", makeValue("original", nil), makeValue("original", tftypes.UnknownValue), makeValue("original", tftypes.UnknownValue), false},
		{"destroy", makeValue("original", nil), tftypes.NewValue(obj, nil), state, false},
		{"create", makeValue("original", nil), makeValue("original", tftypes.UnknownValue), tftypes.NewValue(obj, nil), false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := unchangedPolicyAliasPlan(tt.config, tt.plan, tt.state)
			if ok != tt.want {
				t.Fatalf("restored=%t want %t", ok, tt.want)
			}
			if ok && !got.Equal(tt.state) {
				t.Fatal("restored plan differs from prior state")
			}
			if !ok && !got.Equal(tt.plan) {
				t.Fatal("changed a non-restorable plan")
			}
		})
	}
}
