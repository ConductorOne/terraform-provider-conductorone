package typeconvert

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestBoolPointerOrFalse(t *testing.T) {
	t.Run("nil returns false", func(t *testing.T) {
		got := BoolPointerOrFalse(nil)
		if got.IsNull() {
			t.Error("expected non-null, got null")
		}
		if got.ValueBool() != false {
			t.Errorf("expected false, got %v", got.ValueBool())
		}
	})

	t.Run("false returns false", func(t *testing.T) {
		v := false
		got := BoolPointerOrFalse(&v)
		if got.IsNull() {
			t.Error("expected non-null, got null")
		}
		if got.ValueBool() != false {
			t.Errorf("expected false, got %v", got.ValueBool())
		}
	})

	t.Run("true returns true", func(t *testing.T) {
		v := true
		got := BoolPointerOrFalse(&v)
		if got.IsNull() {
			t.Error("expected non-null, got null")
		}
		if got.ValueBool() != true {
			t.Errorf("expected true, got %v", got.ValueBool())
		}
	})

	t.Run("nil is equal to false pointer", func(t *testing.T) {
		nilResult := BoolPointerOrFalse(nil)
		v := false
		falseResult := BoolPointerOrFalse(&v)
		if !nilResult.Equal(falseResult) {
			t.Error("nil and false should produce equal types.Bool values")
		}
	})

	t.Run("nil does not match BoolPointerValue behavior", func(t *testing.T) {
		nilResult := BoolPointerOrFalse(nil)
		oldResult := types.BoolPointerValue(nil)
		if nilResult.Equal(oldResult) {
			t.Error("nil helper should differ from BoolPointerValue(nil) which returns null")
		}
	})
}

func TestStringPointerOrEmpty(t *testing.T) {
	t.Run("nil returns empty string", func(t *testing.T) {
		got := StringPointerOrEmpty(nil)
		if got.IsNull() {
			t.Error("expected non-null, got null")
		}
		if got.ValueString() != "" {
			t.Errorf("expected empty string, got %q", got.ValueString())
		}
	})

	t.Run("empty returns empty string", func(t *testing.T) {
		v := ""
		got := StringPointerOrEmpty(&v)
		if got.IsNull() {
			t.Error("expected non-null, got null")
		}
		if got.ValueString() != "" {
			t.Errorf("expected empty string, got %q", got.ValueString())
		}
	})

	t.Run("non-empty returns value", func(t *testing.T) {
		v := "some-id"
		got := StringPointerOrEmpty(&v)
		if got.IsNull() {
			t.Error("expected non-null, got null")
		}
		if got.ValueString() != "some-id" {
			t.Errorf("expected 'some-id', got %q", got.ValueString())
		}
	})

	t.Run("nil is equal to empty pointer", func(t *testing.T) {
		nilResult := StringPointerOrEmpty(nil)
		v := ""
		emptyResult := StringPointerOrEmpty(&v)
		if !nilResult.Equal(emptyResult) {
			t.Error("nil and empty string should produce equal types.String values")
		}
	})
}
