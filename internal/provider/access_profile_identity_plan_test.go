package provider

import (
	"context"
	"testing"

	frameworkprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// Exercise the registered resource through the public planning protocol without
// configuring credentials or invoking Read/Apply. The control removes only the
// identity modifier from the same schema to reproduce the original behavior.
// Downstream propagation is explicit, not a Terraform Core graph test.
type accessProfileWithoutIdentityModifier struct {
	AccessProfileResource
}

func (r *accessProfileWithoutIdentityModifier) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.AccessProfileResource.Schema(ctx, req, resp)
	id, ok := resp.Schema.Attributes["id"].(schema.StringAttribute)
	if !ok {
		resp.Diagnostics.AddError("Unexpected ID schema", "The identity control requires a string ID attribute.")
		return
	}
	id.PlanModifiers = nil
	resp.Schema.Attributes["id"] = id
}

type accessProfileIdentityControlProvider struct {
	frameworkprovider.Provider
}

func (p *accessProfileIdentityControlProvider) Resources(ctx context.Context) []func() resource.Resource {
	resources := p.Provider.Resources(ctx)
	for i, constructor := range resources {
		if _, ok := constructor().(*AccessProfileResource); ok {
			resources[i] = func() resource.Resource {
				return &accessProfileWithoutIdentityModifier{}
			}
		}
	}
	return resources
}

var _ frameworkprovider.Provider = &accessProfileIdentityControlProvider{}

func TestAccessProfileMetadataIdentityPlan(t *testing.T) {
	ctx := context.Background()
	const profileName = "conductorone_access_profile"
	const bindingName = "conductorone_access_profile_visibility_bindings"
	const syntheticID = "synthetic-profile-existing"
	variants := []struct {
		name      string
		unpatched bool
	}{
		{name: "registered"},
		{name: "unpatched_control", unpatched: true},
	}
	for _, variant := range variants {
		t.Run(variant.name, func(t *testing.T) {
			provider := New("test")()
			if variant.unpatched {
				provider = &accessProfileIdentityControlProvider{Provider: provider}
			}
			server := providerserver.NewProtocol6(provider)()
			schemas, err := server.GetProviderSchema(ctx, &tfprotov6.GetProviderSchemaRequest{})
			if err != nil {
				t.Fatal(err)
			}
			identityProposalDiagnostics(t, schemas.Diagnostics)
			profileType, ok := schemas.ResourceSchemas[profileName].ValueType().(tftypes.Object)
			if !ok {
				t.Fatal("access profile schema is not an object")
			}
			bindingType, ok := schemas.ResourceSchemas[bindingName].ValueType().(tftypes.Object)
			if !ok {
				t.Fatal("visibility binding schema is not an object")
			}
			stateValues := map[string]interface{}{
				"id": syntheticID, "display_name": "Synthetic profile", "description": "before",
				"published": false, "updated_at": "2026-01-01T00:00:00Z",
				"type": "REQUEST_CATALOG_TYPE_CATALOG",
			}
			configValues := map[string]interface{}{
				"display_name": "Synthetic profile", "description": "before", "published": false,
				"type": "REQUEST_CATALOG_TYPE_CATALOG",
			}
			prior := identityProposalObject(profileType, stateValues)
			plan := func(t *testing.T, name string, typ tftypes.Object, state, config, proposed tftypes.Value) (tftypes.Value, *tfprotov6.PlanResourceChangeResponse) {
				t.Helper()
				response, err := server.PlanResourceChange(ctx, &tfprotov6.PlanResourceChangeRequest{
					TypeName: name, PriorState: identityProposalDynamic(t, state),
					Config: identityProposalDynamic(t, config), ProposedNewState: identityProposalDynamic(t, proposed),
				})
				if err != nil {
					t.Fatal(err)
				}
				identityProposalDiagnostics(t, response.Diagnostics)
				value, err := response.PlannedState.Unmarshal(typ)
				if err != nil {
					t.Fatal(err)
				}
				return value, response
			}
			for _, change := range []struct {
				name  string
				field string
				value interface{}
			}{
				{name: "description_update", field: "description", value: "after"},
				{name: "publication_update", field: "published", value: true},
			} {
				t.Run(change.name, func(t *testing.T) {
					config := identityProposalCopy(configValues)
					proposed := identityProposalCopy(stateValues)
					config[change.field], proposed[change.field] = change.value, change.value
					planned, response := plan(t, profileName, profileType, prior, identityProposalObject(profileType, config), identityProposalObject(profileType, proposed))
					attrs := identityProposalAttributes(t, planned)
					wantID := tftypes.NewValue(tftypes.String, tftypes.UnknownValue)
					if !variant.unpatched {
						wantID = tftypes.NewValue(tftypes.String, syntheticID)
					}
					identityProposalEqual(t, "id", attrs["id"], wantID)
					identityProposalEqual(t, change.field, attrs[change.field], tftypes.NewValue(profileType.AttributeTypes[change.field], change.value))
					identityProposalEqual(t, "updated_at", attrs["updated_at"], tftypes.NewValue(tftypes.String, tftypes.UnknownValue))
					if len(response.RequiresReplace) != 0 {
						t.Fatalf("metadata update requires profile replacement: %v", response.RequiresReplace)
					}

					// Exercise the existing merge/conversion path with this actual plan:
					// unknown id already routes Update to prior identity, independently
					// of whether downstream references are known during planning.
					var schemaResponse resource.SchemaResponse
					(&AccessProfileResource{}).Schema(ctx, resource.SchemaRequest{}, &schemaResponse)
					var model *AccessProfileResourceModel
					updateResponse := resource.UpdateResponse{}
					merge(ctx, resource.UpdateRequest{
						State: tfsdk.State{Raw: prior, Schema: schemaResponse.Schema},
						Plan:  tfsdk.Plan{Raw: planned, Schema: schemaResponse.Schema},
					}, &updateResponse, &model)
					if updateResponse.Diagnostics.HasError() {
						t.Fatal(updateResponse.Diagnostics)
					}
					request, diags := model.ToOperationsC1APIRequestcatalogV1RequestCatalogManagementServiceUpdateRequest(ctx)
					if diags.HasError() || request.ID != syntheticID {
						t.Fatalf("Update conversion identity = %v; diagnostics = %v", request, diags)
					}
					// Simulate a fresh API-computed timestamp, then the generated
					// post-response refreshPlan step. It must not restore old mutable data.
					model.UpdatedAt = types.StringValue("2026-01-02T00:00:00Z")
					var object types.Object
					if diags := (tfsdk.Plan{Raw: planned, Schema: schemaResponse.Schema}).Get(ctx, &object); diags.HasError() {
						t.Fatal(diags)
					}
					if diags := refreshPlan(ctx, object, &model); diags.HasError() {
						t.Fatal(diags)
					}
					if model.UpdatedAt.ValueString() != "2026-01-02T00:00:00Z" {
						t.Fatalf("mutable timestamp frozen: %v", model.UpdatedAt)
					}

					// One static synthetic audience entry; only the interpolated
					// catalog_id changes in knowledge, not identity or membership.
					listType, ok := bindingType.AttributeTypes["access_entitlements"].(tftypes.List)
					if !ok {
						t.Fatal("access_entitlements schema is not a list")
					}
					entryType, ok := listType.ElementType.(tftypes.Object)
					if !ok {
						t.Fatal("access_entitlements element schema is not an object")
					}
					entries := []tftypes.Value{identityProposalObject(entryType, map[string]interface{}{"app_id": "synthetic-app", "id": "synthetic-audience"})}
					bindingState := identityProposalObject(bindingType, map[string]interface{}{"catalog_id": syntheticID, "access_entitlements": entries})
					bindingInput := identityProposalObject(bindingType, map[string]interface{}{"catalog_id": attrs["id"], "access_entitlements": entries})
					bindingPlan, bindingResponse := plan(t, bindingName, bindingType, bindingState, bindingInput, bindingInput)
					if variant.unpatched {
						if len(bindingResponse.RequiresReplace) != 1 || !bindingResponse.RequiresReplace[0].Equal(tftypes.NewAttributePath().WithAttributeName("catalog_id")) {
							t.Fatalf("unknown catalog_id replacement paths = %v", bindingResponse.RequiresReplace)
						}
					} else if len(bindingResponse.RequiresReplace) != 0 || !bindingPlan.Equal(bindingState) {
						t.Fatalf("unchanged binding not preserved: %v, replacement paths %v", bindingPlan, bindingResponse.RequiresReplace)
					}
					t.Logf("profile id=%s; updated_at=%s; binding replacement paths=%v", attrs["id"], attrs["updated_at"], bindingResponse.RequiresReplace)
				})
			}

			for _, condition := range []string{"create", "missing_after_refresh", "unknown_prior_id", "null_prior_id", "refreshed_identity", "replace_create_pass", "type_replacement", "destroy", "no_change"} {
				t.Run(condition, func(t *testing.T) {
					state := prior
					config := identityProposalCopy(configValues)
					proposed := identityProposalCopy(stateValues)
					config["description"], proposed["description"] = "after", "after"
					wantID := tftypes.NewValue(tftypes.String, tftypes.UnknownValue)
					switch condition {
					case "create", "missing_after_refresh", "replace_create_pass":
						state = tftypes.NewValue(profileType, nil)
						proposed = identityProposalCopy(config)
						proposed["id"] = tftypes.UnknownValue
					case "unknown_prior_id", "null_prior_id":
						priorValues := identityProposalCopy(stateValues)
						priorValues["id"] = tftypes.UnknownValue
						if condition == "null_prior_id" {
							priorValues["id"] = nil
						}
						state = identityProposalObject(profileType, priorValues)
						proposed["id"] = priorValues["id"]
					case "refreshed_identity":
						refreshed := identityProposalCopy(stateValues)
						refreshed["id"], proposed["id"] = "synthetic-refreshed", "synthetic-refreshed"
						state = identityProposalObject(profileType, refreshed)
						if !variant.unpatched {
							wantID = tftypes.NewValue(tftypes.String, "synthetic-refreshed")
						}
					case "type_replacement":
						config["type"], proposed["type"] = "REQUEST_CATALOG_TYPE_BUNDLE", "REQUEST_CATALOG_TYPE_BUNDLE"
						if !variant.unpatched {
							wantID = tftypes.NewValue(tftypes.String, syntheticID)
						}
					case "no_change":
						config, proposed = identityProposalCopy(configValues), identityProposalCopy(stateValues)
						wantID = tftypes.NewValue(tftypes.String, syntheticID)
					}
					configRaw, proposedRaw := identityProposalObject(profileType, config), identityProposalObject(profileType, proposed)
					if condition == "destroy" {
						configRaw, proposedRaw = tftypes.NewValue(profileType, nil), tftypes.NewValue(profileType, nil)
					}
					planned, response := plan(t, profileName, profileType, state, configRaw, proposedRaw)
					if condition == "destroy" {
						if !planned.IsNull() {
							t.Fatalf("destroy plan resurrected: %v", planned)
						}
						return
					}
					identityProposalEqual(t, "id", identityProposalAttributes(t, planned)["id"], wantID)
					if condition == "type_replacement" {
						if len(response.RequiresReplace) != 1 || !response.RequiresReplace[0].Equal(tftypes.NewAttributePath().WithAttributeName("type")) {
							t.Fatalf("type replacement lost: %v", response.RequiresReplace)
						}
					} else if len(response.RequiresReplace) != 0 {
						t.Fatalf("unexpected replacement: %v", response.RequiresReplace)
					}
				})
			}
		})
	}
}

func identityProposalObject(typ tftypes.Object, values map[string]interface{}) tftypes.Value {
	attrs := make(map[string]tftypes.Value, len(typ.AttributeTypes))
	for name, attrType := range typ.AttributeTypes {
		if value, ok := values[name].(tftypes.Value); ok {
			attrs[name] = value
		} else {
			attrs[name] = tftypes.NewValue(attrType, values[name])
		}
	}
	return tftypes.NewValue(typ, attrs)
}

func identityProposalCopy(values map[string]interface{}) map[string]interface{} {
	cloned := make(map[string]interface{}, len(values))
	for name, value := range values {
		cloned[name] = value
	}
	return cloned
}

func identityProposalDynamic(t *testing.T, value tftypes.Value) *tfprotov6.DynamicValue {
	t.Helper()
	dynamic, err := tfprotov6.NewDynamicValue(value.Type(), value)
	if err != nil {
		t.Fatal(err)
	}
	return &dynamic
}

func identityProposalAttributes(t *testing.T, value tftypes.Value) map[string]tftypes.Value {
	t.Helper()
	var attrs map[string]tftypes.Value
	if err := value.As(&attrs); err != nil {
		t.Fatal(err)
	}
	return attrs
}

func identityProposalEqual(t *testing.T, name string, got, want tftypes.Value) {
	t.Helper()
	if !got.Equal(want) {
		t.Fatalf("%s = %s; want %s", name, got, want)
	}
}

func identityProposalDiagnostics(t *testing.T, diagnostics []*tfprotov6.Diagnostic) {
	t.Helper()
	for _, diagnostic := range diagnostics {
		if diagnostic.Severity == tfprotov6.DiagnosticSeverityError {
			t.Fatalf("protocol diagnostic: %s: %s", diagnostic.Summary, diagnostic.Detail)
		}
	}
}
