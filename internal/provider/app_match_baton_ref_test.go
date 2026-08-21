package provider

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
)

func TestAppMatchBatonRefRequiresReplacement(t *testing.T) {
	resp := &resource.SchemaResponse{}
	NewAppResource().Schema(context.Background(), resource.SchemaRequest{}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("building app resource schema: %v", resp.Diagnostics)
	}

	attr, ok := resp.Schema.Attributes["match_baton_ref"].(schema.SingleNestedAttribute)
	if !ok {
		t.Fatalf("match_baton_ref has type %T, want schema.SingleNestedAttribute", resp.Schema.Attributes["match_baton_ref"])
	}

	want := reflect.TypeOf(objectplanmodifier.RequiresReplaceIfConfigured())
	for _, modifier := range attr.PlanModifiers {
		if reflect.TypeOf(modifier) == want {
			return
		}
	}
	t.Fatalf("match_baton_ref does not require resource replacement when configured: %#v", attr.PlanModifiers)
}
