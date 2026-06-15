package provider

import (
	"context"
	"reflect"
	"testing"

	tfTypes "github.com/conductorone/terraform-provider-conductorone/internal/provider/types"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// TestRequestSchemaFormFieldRoundTrips exercises the full Terraform <-> API
// conversion for every FormField union variant. For each case we build a
// resource model, convert it to the shared API model (ToSharedRequestSchemaForm), convert
// the result back (RefreshFromSharedRequestSchemaForm), and assert the field survives the
// round-trip unchanged.
//
// This covers the IGA-743 customer repro (a bool_field with default_value), the
// the all-variants-null workaround (a field with all union variants left null), and one field of
// each remaining union variant — including the new variants added at v1.3.0
// (oauth2_field, *_provider_config, form_string_map_field, read_only, required)
// that the genignored conversion previously dropped.
func TestRequestSchemaFormFieldRoundTrips(t *testing.T) {
	ctx := context.Background()

	cases := map[string]tfTypes.FormField{
		// IGA-743 customer repro.
		"bool_field": {
			Name:        types.StringValue("approved"),
			Description: types.StringValue("Approve the request"),
			DisplayName: types.StringValue("Approved"),
			BoolField:   &tfTypes.BoolField{DefaultValue: types.BoolValue(true)},
		},
		// All-variants-null workaround: every union variant explicitly null, only metadata set.
		"all_variants_null": {
			Name:        types.StringValue("justification"),
			Description: types.StringValue("Why do you need this?"),
			DisplayName: types.StringValue("Justification"),
		},
		"form_string_field": {
			Name: types.StringValue("reason"),
			FormStringField: &tfTypes.FormStringField{
				DefaultValue: types.StringValue("n/a"),
				Placeholder:  types.StringValue("enter a reason"),
			},
		},
		"form_string_field_picker": {
			Name: types.StringValue("owner"),
			FormStringField: &tfTypes.FormStringField{
				PickerField: &tfTypes.PickerField{
					AppResourceFilter: &tfTypes.AppResourceFilter{
						AppID:          types.StringValue("app-123"),
						ResourceTypeID: types.StringValue("rt-456"),
					},
					AppUserFilter: &tfTypes.AppUserFilter{AppID: types.StringValue("app-789")},
					C1UserFilter:  &tfTypes.C1UserFilter{},
				},
			},
		},
		"int64_field": {
			Name:       types.StringValue("count"),
			Int64Field: &tfTypes.Int64Field{DefaultValue: types.StringValue("7")},
		},
		"file_field": {
			Name:      types.StringValue("attachment"),
			FileField: &tfTypes.FileField{MaxFileSize: types.StringValue("1048576")},
		},
		"form_string_map_field": {
			Name: types.StringValue("labels"),
			FormStringMapField: &tfTypes.FormStringMapField{
				DefaultValue:   map[string]types.String{"team": types.StringValue("platform")},
				StringMapRules: &tfTypes.StringMapRules{IsRequired: types.BoolValue(true), ValidateEmpty: types.BoolValue(false)},
			},
		},
		"oauth2_field": {
			Name:        types.StringValue("oauth"),
			Oauth2Field: &tfTypes.Oauth2Field{Oauth2FieldView: &tfTypes.Oauth2FieldView{}},
		},
		"admin_provider_config": {
			Name:                types.StringValue("admin_cfg"),
			AdminProviderConfig: &tfTypes.AdminProviderConfig{DefaultValueCel: types.StringValue("true"), ShowToUser: types.BoolValue(true)},
		},
		"shared_provider_config": {
			Name: types.StringValue("shared_cfg"),
			SharedProviderConfig: &tfTypes.SharedProviderConfig{
				DefaultValueCel:        types.StringValue("a"),
				InputTransformationCel: types.StringValue("b"),
				LockDefaultValues:      types.BoolValue(true),
			},
		},
		"user_provider_config": {
			Name:               types.StringValue("user_cfg"),
			UserProviderConfig: &tfTypes.UserProviderConfig{InputTransformationCel: types.StringValue("c")},
		},
		"read_only_and_required": {
			Name:     types.StringValue("ro_req"),
			ReadOnly: types.BoolValue(true),
			Required: types.BoolValue(true),
		},
	}

	for name, field := range cases {
		t.Run(name, func(t *testing.T) {
			in := &RequestSchemaResourceModel{
				Name:   types.StringValue("schema"),
				Fields: []tfTypes.FormField{field},
			}

			apiForm, diags := in.ToSharedRequestSchemaForm(ctx)
			if diags.HasError() {
				t.Fatalf("ToSharedRequestSchemaForm: %v", diags)
			}

			out := &RequestSchemaResourceModel{}
			if d := out.RefreshFromSharedRequestSchemaForm(ctx, apiForm); d.HasError() {
				t.Fatalf("RefreshFromSharedRequestSchemaForm: %v", d)
			}

			if len(out.Fields) != 1 {
				t.Fatalf("expected 1 field after round-trip, got %d", len(out.Fields))
			}
			if !reflect.DeepEqual(in.Fields[0], out.Fields[0]) {
				t.Errorf("field did not round-trip\n in: %#v\nout: %#v", in.Fields[0], out.Fields[0])
			}
		})
	}
}
