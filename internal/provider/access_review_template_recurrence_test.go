package provider

import (
	"context"
	"reflect"
	"testing"

	"github.com/conductorone/terraform-provider-conductorone/internal/validators/stringvalidators"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func TestAccessReviewTemplateRecurrenceFrequencyRemainsOptional(t *testing.T) {
	resp := &resource.SchemaResponse{}
	NewAccessReviewTemplateResource().Schema(context.Background(), resource.SchemaRequest{}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("building access review template schema: %v", resp.Diagnostics)
	}

	recurrenceRule, ok := resp.Schema.Attributes["recurrence_rule"].(schema.SingleNestedAttribute)
	if !ok {
		t.Fatalf("recurrence_rule has schema type %T, want schema.SingleNestedAttribute", resp.Schema.Attributes["recurrence_rule"])
	}
	frequency, ok := recurrenceRule.Attributes["frequency"].(schema.StringAttribute)
	if !ok {
		t.Fatalf("recurrence_rule.frequency has schema type %T, want schema.StringAttribute", recurrenceRule.Attributes["frequency"])
	}
	if !frequency.Optional || !frequency.Computed {
		t.Fatalf("recurrence_rule.frequency must remain Optional and Computed, got Optional=%t Computed=%t", frequency.Optional, frequency.Computed)
	}

	notNullType := reflect.TypeOf(stringvalidators.NotNull())
	for _, validator := range frequency.Validators {
		if reflect.TypeOf(validator) == notNullType {
			t.Fatal("recurrence_rule.frequency must accept omission for v1.5.0 configuration compatibility")
		}
	}
}
