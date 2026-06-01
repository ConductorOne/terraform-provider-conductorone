// Package requestschemaplanmodifier holds hand-written plan modifiers for the
// request_schema resource. It lives outside the Speakeasy-generated
// planmodifiers packages so `make gen` does not overwrite it; the wiring into
// the generated resource schema is reapplied by a patch in patches/.
package requestschemaplanmodifier

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// nonNameChars mirrors the server's form-field name normalization
// (pkg/models/request_schema/create.go ValidateForm): strip every character
// that is not alphanumeric, underscore, or whitespace before spaces are
// collapsed to underscores.
var nonNameChars = regexp.MustCompile(`[^a-zA-Z0-9_\s]`)

// NormalizeFields returns a list plan modifier for the request_schema `fields`
// list that eliminates the two sources of perpetual diff on an unchanged
// schema:
//
//  1. Name churn. The server ignores the configured field name and regenerates
//     it as `<sanitized display_name>_<0-based occurrence index>` on every
//     write, so a configured "field1" perpetually diffs against the stored
//     "Field_One_0". When the configured display_name still normalizes to the
//     name already in state, the stored value is kept (no diff). When the
//     display_name changed — or on create — the name is left unknown so the
//     server can recompute it. The configured name is never overridden with a
//     freshly-computed value: Terraform forbids a provider from planning an
//     Optional+Computed value that matches neither config nor prior state.
//
//  2. Computed sub-attribute churn. The field's many Optional+Computed
//     sub-attributes (bool_field, string_field and its children, etc.) are
//     planned as unknown whenever any sibling changes, so they never settle.
//     Their prior-state value is copied into any planned-unknown leaf — the
//     value-level equivalent of UseStateForUnknown across the element tree.
//
// Both run at the attr.Value level (not via the generated Go structs) so that
// unknown nested values round-trip cleanly.
func NormalizeFields() planmodifier.List {
	return normalizeFields{}
}

type normalizeFields struct{}

func (m normalizeFields) Description(_ context.Context) string {
	return "Keeps server-normalized field names and computed sub-attributes stable so an unchanged schema reaches steady state."
}

func (m normalizeFields) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m normalizeFields) PlanModifyList(ctx context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
	if req.PlanValue.IsNull() || req.PlanValue.IsUnknown() {
		return
	}

	elemType, ok := req.PlanValue.ElementType(ctx).(types.ObjectType)
	if !ok {
		return
	}

	var stateElems []attr.Value
	if !req.StateValue.IsNull() && !req.StateValue.IsUnknown() {
		stateElems = req.StateValue.Elements()
	}

	planElems := req.PlanValue.Elements()
	newElems := make([]attr.Value, len(planElems))
	nameCounts := make(map[string]int)

	for i, pe := range planElems {
		planObj, ok := pe.(types.Object)
		if !ok {
			return
		}
		newAttrs := cloneAttrs(planObj.Attributes())

		var stateAttrs map[string]attr.Value
		if i < len(stateElems) {
			if stateObj, ok := stateElems[i].(types.Object); ok {
				stateAttrs = stateObj.Attributes()
			}
		}

		// (1) Preserve every planned-unknown computed leaf from prior state.
		if stateAttrs != nil {
			for k := range newAttrs {
				newAttrs[k] = preserveUnknown(newAttrs[k], stateAttrs[k])
			}
		}

		// (2) Reconcile the server-derived name. Done after the generic
		// preserve pass so it owns the final decision for `name`.
		newAttrs["name"] = m.reconcileName(newAttrs, stateAttrs, nameCounts)

		newObj, diags := types.ObjectValue(elemType.AttrTypes, newAttrs)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		newElems[i] = newObj
	}

	newList, diags := types.ListValue(elemType, newElems)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.PlanValue = newList
}

// reconcileName decides the planned value for a field's `name`, advancing
// nameCounts so duplicate base names get the same per-occurrence index the
// server assigns.
//
// Terraform's plan-consistency rules constrain the choices: an Optional+
// Computed value that is configured (known and non-null) MUST be planned
// exactly as configured, unless the planned value equals prior state. So the
// only safe deviation from config is "fall back to the value already in
// state". The server overriding a configured name at apply time is tolerated
// (computed attributes may change during apply); the resulting state then
// matches on the next plan.
func (m normalizeFields) reconcileName(planAttrs, stateAttrs map[string]attr.Value, nameCounts map[string]int) attr.Value {
	planName, _ := planAttrs["name"].(types.String)
	nameConfigured := !planName.IsNull() && !planName.IsUnknown()

	dn, ok := planAttrs["display_name"].(types.String)
	if !ok || dn.IsNull() || dn.IsUnknown() {
		// No statically-known display_name: nothing to derive. Keep whatever
		// was planned and don't perturb the occurrence count.
		return planName
	}

	base := nonNameChars.ReplaceAllString(dn.ValueString(), "")
	base = strings.ReplaceAll(base, " ", "_")
	nameCounts[base]++
	derived := fmt.Sprintf("%s_%d", base, nameCounts[base]-1)

	if stateAttrs != nil {
		if stateName, ok := stateAttrs["name"].(types.String); ok && !stateName.IsNull() && !stateName.IsUnknown() && stateName.ValueString() == derived {
			// display_name still normalizes to the stored name: keep it so
			// there is no diff (and so a configured-but-stale name like
			// "field1" doesn't churn).
			return stateName
		}
	}

	if nameConfigured {
		// Configured name that doesn't (yet) match state: must plan it as
		// configured. The server normalizes it on apply.
		return planName
	}
	// Unconfigured: let the server derive the name.
	return types.StringUnknown()
}

func cloneAttrs(in map[string]attr.Value) map[string]attr.Value {
	out := make(map[string]attr.Value, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}

// preserveUnknown returns the prior-state value when the planned value is
// unknown, and otherwise recurses into nested objects so that unknown leaves
// inside an otherwise-known object (e.g. an explicitly-configured
// `string_field = {}` whose children are computed) also fall back to state.
func preserveUnknown(plan, state attr.Value) attr.Value {
	if state == nil {
		return plan
	}
	if plan.IsUnknown() {
		return state
	}

	planObj, ok := plan.(types.Object)
	if !ok {
		return plan
	}
	stateObj, ok := state.(types.Object)
	if !ok {
		return plan
	}

	planAttrs := planObj.Attributes()
	stateAttrs := stateObj.Attributes()
	merged := make(map[string]attr.Value, len(planAttrs))
	for k, v := range planAttrs {
		merged[k] = preserveUnknown(v, stateAttrs[k])
	}

	obj, diags := types.ObjectValue(planObj.AttributeTypes(context.Background()), merged)
	if diags.HasError() {
		return plan
	}
	return obj
}
