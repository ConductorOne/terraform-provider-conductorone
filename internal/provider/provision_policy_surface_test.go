package provider

import (
	"context"
	"sort"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// collectPolicyAttrs walks a Terraform type and records the dotted path of every
// attribute named provision_policy or provisioner_policy.
func collectPolicyAttrs(prefix string, typ tftypes.Type, out map[string]struct{}) {
	switch t := typ.(type) {
	case tftypes.Object:
		for name, attrType := range t.AttributeTypes {
			path := name
			if prefix != "" {
				path = prefix + "." + name
			}
			if name == "provision_policy" || name == "provisioner_policy" {
				out[path] = struct{}{}
			}
			collectPolicyAttrs(path, attrType, out)
		}
	case tftypes.List:
		collectPolicyAttrs(prefix, t.ElementType, out)
	case tftypes.Set:
		collectPolicyAttrs(prefix, t.ElementType, out)
	case tftypes.Map:
		collectPolicyAttrs(prefix, t.ElementType, out)
	}
}

func policyAttrPaths(schema *tfprotov6.Schema) []string {
	out := map[string]struct{}{}
	collectPolicyAttrs("", schema.ValueType(), out)
	paths := make([]string, 0, len(out))
	for p := range out {
		paths = append(paths, p)
	}
	sort.Strings(paths)
	return paths
}

// TestProvisionPolicyPublicSchemaSurface pins the reviewed public schema surface
// for provision policies after the IGA-4347 overlay.
//
// The overlay renames the shared c1.api.app.v1.AppEntitlement property
// provisionerPolicy to provisionPolicy. Speakeasy resolves that rename at the
// component, so every other entity that embeds AppEntitlement also renames its
// computed attribute. That collateral is the compatibility cost of the overlay
// approach and it is pinned here: a future regen that widens or narrows the
// affected set fails this test instead of shipping silently.
//
// Only live, registered types are listed. There is no
// conductorone_app_resource_owner_entitlement type in the provider schema (its
// constructors are generated but unregistered — IGA-1741), so its renamed
// generated files are not part of the public surface.
func TestProvisionPolicyPublicSchemaSurface(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	ps := providerserver.NewProtocol6(New("test")())()
	schemaResp, err := ps.GetProviderSchema(ctx, &tfprotov6.GetProviderSchemaRequest{})
	if err != nil {
		t.Fatalf("GetProviderSchema: %s", err)
	}

	want := map[string][]string{
		// The IGA-4347 target: one configurable policy attribute; the old
		// computed provisioner_policy is unified away.
		"resource:conductorone_custom_app_entitlement": {"provision_policy"},

		// Hand-written resource, unchanged by the fix (it already unified the
		// mapping in Go). Its data source still exposes the read-side name.
		"resource:conductorone_app_entitlement":   {"provision_policy"},
		"datasource:conductorone_app_entitlement": {"provisioner_policy"},

		// Collateral, computed-only rename on AppEntitlement consumers.
		"resource:conductorone_app_entitlement_owner_entitlement":   {"app_entitlement.provision_policy"},
		"datasource:conductorone_app_entitlement_owner_entitlement": {"app_entitlement.provision_policy"},
		"resource:conductorone_app_owner_entitlement":               {"app_entitlement.provision_policy"},
		"datasource:conductorone_app_owner_entitlement":             {"app_entitlement.provision_policy"},
		"resource:conductorone_connector_owner_entitlement":         {"app_entitlement.provision_policy"},
		"datasource:conductorone_connector_owner_entitlement":       {"app_entitlement.provision_policy"},
		"datasource:conductorone_app_entitlements":                  {"list.app_entitlement.provision_policy"},

		// Policy resources already used this name; listed so a rename of the
		// policy step subtree is caught here too.
		"resource:conductorone_policy":   {"policy_steps.steps.provision.provision_policy"},
		"datasource:conductorone_policy": {"policy_steps.steps.provision.provision_policy"},
		"datasource:conductorone_policies": {
			"list.policy_steps.steps.provision.provision_policy",
		},
	}

	for key := range want {
		kind, name := splitTypeKey(key)
		if kind == "resource" {
			if _, ok := schemaResp.ResourceSchemas[name]; !ok {
				t.Fatalf("expected resource %q not found in provider schema", name)
			}
		} else if _, ok := schemaResp.DataSourceSchemas[name]; !ok {
			t.Fatalf("expected data source %q not found in provider schema", name)
		}
	}

	for key, expected := range want {
		kind, name := splitTypeKey(key)
		var schema *tfprotov6.Schema
		if kind == "resource" {
			schema = schemaResp.ResourceSchemas[name]
		} else {
			schema = schemaResp.DataSourceSchemas[name]
		}
		got := policyAttrPaths(schema)
		if len(got) != len(expected) {
			t.Errorf("%s policy attribute surface changed\n got: %v\nwant: %v", key, got, expected)
			continue
		}
		for i := range got {
			if got[i] != expected[i] {
				t.Errorf("%s policy attribute surface changed\n got: %v\nwant: %v", key, got, expected)
				break
			}
		}
	}
}

func splitTypeKey(key string) (kind, name string) {
	for _, prefix := range []string{"resource:", "datasource:"} {
		if len(key) > len(prefix) && key[:len(prefix)] == prefix {
			return prefix[:len(prefix)-1], key[len(prefix):]
		}
	}
	return "", key
}
