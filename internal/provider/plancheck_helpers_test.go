package provider

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

// ExpectEmptyPlanWithDriftLog is a self-diagnosing drop-in replacement for
// plancheck.ExpectEmptyPlan. It asserts the same invariant (no output or
// resource changes) and fails the step identically, BUT on a non-empty plan it
// additionally logs — via t.Log — the exact attribute PATHS that drifted
// (Before != After) for each resource change whose actions include "update".
//
// It logs PATHS + action ONLY, never the before/after VALUES, so it cannot leak
// config secrets into CI logs. It emits nothing on an empty (green) plan, so it
// is zero-downside to leave wired at every ExpectEmptyPlan call site
// permanently.
//
// Why permanent: CI's acceptance suite is the only place the live ConductorOne
// API drives the post-apply refresh plan, and CI only ever prints
// "...has planned action(s): [update]" without naming the attribute. A
// self-diagnosing ExpectEmptyPlan means we never again have to guess which field
// drifts or reproduce CI drift locally — the failing run names the path itself.
//
// This file is hand-maintained and matched by `.genignore` (/internal/provider/
// *_test.go), so `speakeasy run` never overwrites it.
func ExpectEmptyPlanWithDriftLog(t *testing.T) plancheck.PlanCheck {
	t.Helper()
	return expectEmptyPlanWithDriftLog{t: t}
}

var _ plancheck.PlanCheck = expectEmptyPlanWithDriftLog{}

type expectEmptyPlanWithDriftLog struct {
	t *testing.T
}

// CheckPlan mirrors plancheck.ExpectEmptyPlan's assertion, then appends a
// path-level drift diagnostic for any updated resource.
func (e expectEmptyPlanWithDriftLog) CheckPlan(ctx context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	var result []error

	for output, change := range req.Plan.OutputChanges {
		if !change.Actions.NoOp() {
			result = append(result, fmt.Errorf("expected empty plan, but output %q has planned action(s): %v", output, change.Actions))
		}
	}

	for _, rc := range req.Plan.ResourceChanges {
		if rc.Change == nil || rc.Change.Actions.NoOp() {
			continue
		}
		result = append(result, fmt.Errorf("expected empty plan, but %s has planned action(s): %v", rc.Address, rc.Change.Actions))

		if !rc.Change.Actions.Update() {
			continue
		}
		entries := driftReport(rc.Change.Before, rc.Change.After, rc.Change.AfterUnknown)
		if len(entries) == 0 {
			e.t.Logf("DRIFT [%s] update planned but no concrete Before/After leaf differences detected; "+
				"inspect the raw plan", rc.Address)
			continue
		}
		for _, en := range entries {
			// classification + path + kinds only — values are intentionally
			// redacted so this is safe to emit in CI logs regardless of config.
			//
			// TRIGGER  = genuine config/state value mismatch — the field actually
			//            driving the [update]. This is what a fix must target.
			// collateral = known-after-apply: a Computed field that goes unknown
			//            only because an update is already planned. Fixing the
			//            trigger(s) resolves these automatically; do NOT chase them.
			class := "TRIGGER(value-mismatch)"
			if en.knownAfterApply {
				class = "collateral(known-after-apply)"
			}
			e.t.Logf("DRIFT [%s] %s @ %s (before=%s after=%s; values redacted)",
				rc.Address, class, en.path, en.beforeKind, en.afterKind)
		}
	}

	resp.Error = errors.Join(result...)
}

// driftEntry is one differing attribute, classified by whether the After value
// is unknown (known-after-apply, i.e. collateral cascade from a planned update)
// versus a genuine value mismatch (the trigger). Kinds describe the JSON shape
// (object/array/string/number/bool/null) without ever exposing the value.
type driftEntry struct {
	path            string
	knownAfterApply bool
	beforeKind      string
	afterKind       string
}

// driftReport deep-walks two decoded-JSON values (map[string]interface{},
// []interface{}, or scalars, as produced by `terraform show -json`) plus the
// after_unknown structure, and returns the sorted attribute entries that differ.
// Paths use dotted keys and [i] index notation, e.g.
// policy_steps.certify.steps[0].approval.escalation.
func driftReport(before, after, afterUnknown interface{}) []driftEntry {
	var entries []driftEntry
	walkDrift("", before, after, afterUnknown, &entries)
	sort.Slice(entries, func(i, j int) bool { return entries[i].path < entries[j].path })
	return entries
}

// driftPaths is a thin path-only view over driftReport, retained for the
// behavioral unit test.
func driftPaths(before, after interface{}) []string {
	entries := driftReport(before, after, nil)
	if len(entries) == 0 {
		return nil
	}
	paths := make([]string, 0, len(entries))
	for _, e := range entries {
		paths = append(paths, e.path)
	}
	return paths
}

func walkDrift(path string, before, after, afterUnknown interface{}, out *[]driftEntry) {
	// Identical subtrees (including both nil) contribute no drift.
	if reflect.DeepEqual(before, after) {
		return
	}

	switch b := before.(type) {
	case map[string]interface{}:
		a, ok := after.(map[string]interface{})
		if !ok {
			*out = append(*out, classifyLeaf(path, before, after, afterUnknown))
			return
		}
		auMap, _ := afterUnknown.(map[string]interface{})
		for _, k := range unionKeys(b, a) {
			var auChild interface{}
			if auMap != nil {
				auChild = auMap[k]
			}
			walkDrift(joinKey(path, k), b[k], a[k], auChild, out)
		}
	case []interface{}:
		a, ok := after.([]interface{})
		if !ok {
			*out = append(*out, classifyLeaf(path, before, after, afterUnknown))
			return
		}
		auArr, _ := afterUnknown.([]interface{})
		n := len(b)
		if len(a) > n {
			n = len(a)
		}
		for i := 0; i < n; i++ {
			var bv, av, auv interface{}
			if i < len(b) {
				bv = b[i]
			}
			if i < len(a) {
				av = a[i]
			}
			if auArr != nil && i < len(auArr) {
				auv = auArr[i]
			}
			walkDrift(fmt.Sprintf("%s[%d]", path, i), bv, av, auv, out)
		}
	default:
		// Scalar, nil-vs-scalar, or type mismatch: a differing leaf.
		*out = append(*out, classifyLeaf(path, before, after, afterUnknown))
	}
}

// classifyLeaf builds a driftEntry for a single differing node. A node is
// known-after-apply when its after_unknown is literally true (the whole subtree
// is unknown because an update is planned).
func classifyLeaf(path string, before, after, afterUnknown interface{}) driftEntry {
	known := false
	if b, ok := afterUnknown.(bool); ok && b {
		known = true
	}
	return driftEntry{
		path:            leafPath(path),
		knownAfterApply: known,
		beforeKind:      kindOf(before),
		afterKind:       kindOf(after),
	}
}

// kindOf names the JSON shape of a decoded value without exposing the value.
func kindOf(v interface{}) string {
	switch v.(type) {
	case nil:
		return "null"
	case bool:
		return "bool"
	case float64:
		return "number"
	case string:
		return "string"
	case map[string]interface{}:
		return "object"
	case []interface{}:
		return "array"
	default:
		return fmt.Sprintf("%T", v)
	}
}

func leafPath(path string) string {
	if path == "" {
		return "(root)"
	}
	return path
}

func joinKey(path, key string) string {
	if path == "" {
		return key
	}
	return path + "." + key
}

func unionKeys(a, b map[string]interface{}) []string {
	seen := make(map[string]struct{}, len(a)+len(b))
	for k := range a {
		seen[k] = struct{}{}
	}
	for k := range b {
		seen[k] = struct{}{}
	}
	keys := make([]string, 0, len(seen))
	for k := range seen {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// TestDriftPaths locks the deep-diff that names the drifting attributes. It uses
// the post-apply-refresh shape we expect from the C1 API: a nested
// policy_steps > steps[] > approval subtree where a server-populated scalar
// (Before) faces a config-null (After) leaf.
func TestDriftPaths(t *testing.T) {
	tests := []struct {
		name          string
		before, after interface{}
		want          []string
	}{
		{
			name:   "identical maps -> no drift",
			before: map[string]interface{}{"a": "x", "b": float64(1)},
			after:  map[string]interface{}{"a": "x", "b": float64(1)},
			want:   nil,
		},
		{
			name:   "server scalar vs config null is named",
			before: map[string]interface{}{"escalation": map[string]interface{}{"expiration": "0", "escalation_comment": ""}},
			after:  map[string]interface{}{"escalation": map[string]interface{}{"expiration": nil, "escalation_comment": nil}},
			want:   []string{"escalation.escalation_comment", "escalation.expiration"},
		},
		{
			name: "nested list index path",
			before: map[string]interface{}{"policy_steps": map[string]interface{}{"certify": map[string]interface{}{
				"steps": []interface{}{map[string]interface{}{"approval": map[string]interface{}{"allowed_reassignees": []interface{}{}}}},
			}}},
			after: map[string]interface{}{"policy_steps": map[string]interface{}{"certify": map[string]interface{}{
				"steps": []interface{}{map[string]interface{}{"approval": map[string]interface{}{"allowed_reassignees": nil}}},
			}}},
			want: []string{"policy_steps.certify.steps[0].approval.allowed_reassignees"},
		},
		{
			name:   "differing scalar leaf",
			before: map[string]interface{}{"x": "a"},
			after:  map[string]interface{}{"x": "b"},
			want:   []string{"x"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := driftPaths(tc.before, tc.after)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("driftPaths() = %v, want %v", got, tc.want)
			}
		})
	}
}

// TestDriftReportClassifies locks the trigger-vs-collateral classification: a
// field whose after_unknown is true is known-after-apply (collateral), while a
// concrete before/after mismatch with no unknown marker is the trigger.
func TestDriftReportClassifies(t *testing.T) {
	before := map[string]interface{}{
		"escalation": map[string]interface{}{"skip_step": map[string]interface{}{}}, // configured object
		"id":         "pol-123",                                                     // server-assigned
		"created_at": "2026-01-01T00:00:00Z",
	}
	after := map[string]interface{}{
		"escalation": nil, // genuine mismatch — config object vs planned null
		"id":         nil, // unknown-after-apply
		"created_at": nil, // unknown-after-apply
	}
	afterUnknown := map[string]interface{}{
		"id":         true,
		"created_at": true,
		// escalation absent / not unknown -> genuine value mismatch
	}

	entries := driftReport(before, after, afterUnknown)
	got := map[string]bool{} // path -> knownAfterApply
	kinds := map[string]string{}
	for _, e := range entries {
		got[e.path] = e.knownAfterApply
		kinds[e.path] = e.beforeKind + "->" + e.afterKind
	}

	if known, ok := got["escalation"]; !ok || known {
		t.Errorf("escalation must be a TRIGGER (knownAfterApply=false), got entry=%v present=%v", known, ok)
	}
	if kinds["escalation"] != "object->null" {
		t.Errorf("escalation kinds = %q, want object->null", kinds["escalation"])
	}
	for _, p := range []string{"id", "created_at"} {
		if known, ok := got[p]; !ok || !known {
			t.Errorf("%s must be collateral (knownAfterApply=true), got entry=%v present=%v", p, known, ok)
		}
	}
}
