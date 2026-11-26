package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

// TestAccRequestSchemaResource is the IGA-743 convergence acceptance test: it
// proves against a live ConductorOne API that a conductorone_request_schema
// reaches an EMPTY post-apply plan. This is the live counterpart to the
// schema-level unit assertion in request_schema_convergence_test.go, which can
// only confirm the SuppressDiff plan modifiers are wired — not that they
// actually settle the plan against real server-derived values.
//
// The crash fix on this branch stopped the "Value Conversion Error", but
// several server-derived computed attributes still re-planned as
// "(known after apply)" on every plan after the first apply, so a config that
// already matched the read-back never converged. The generated schema retains
// suppress-computed-diff modifiers for created_at, justification_visibility,
// field_groups and field_relationships. The resource-specific wrapper makes
// nested fields[].name computed-only and preserves prior names only while the
// complete ordered display-label sequence is known and unchanged.
//
// The config below mirrors the original repro (a single bool_field with
// default_value) and intentionally leaves every server-derived attribute
// UNSET — including the nested fields[].name, which the API derives and
// normalizes server-side. That is what makes this a real convergence test:
// those attributes are populated only by the server. The empty-plan assertion
// runs on BOTH the pre-refresh plan (state vs config) and the post-refresh plan
// (live read-back vs config). Explicit nested names are rejected before apply.
//
// ExpectEmptyPlanWithDriftLog is used instead of plancheck.ExpectEmptyPlan so
// that, on a non-empty plan, the CI log names the exact attribute path(s) that
// drifted (paths only, never values) rather than the bare
// "...has planned action(s): [update]".
func TestAccRequestSchemaResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
				resource "conductorone_request_schema" "test" {
					name = "tf-acc-request-schema-iga-743"
					fields = [
						{
							display_name = "Approved"
							description  = "Approve the request"
							bool_field = {
								default_value = true
							}
						}
					]
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("conductorone_request_schema.test", "name", "tf-acc-request-schema-iga-743"),
					resource.TestCheckResourceAttr("conductorone_request_schema.test", "fields.0.display_name", "Approved"),
					resource.TestCheckResourceAttr("conductorone_request_schema.test", "fields.0.bool_field.default_value", "true"),
					// Server-derived; left unset in config. Presence in state proves
					// the read-back populated them and the plan still converged.
					resource.TestCheckResourceAttrSet("conductorone_request_schema.test", "id"),
					resource.TestCheckResourceAttrSet("conductorone_request_schema.test", "created_at"),
				),
				// The convergence assertion: no diff on either the pre-refresh
				// plan or the post-refresh plan immediately after apply.
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPreRefresh: []plancheck.PlanCheck{
						ExpectEmptyPlanWithDriftLog(t),
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						ExpectEmptyPlanWithDriftLog(t),
					},
				},
			},
		},
	})
}

// This is the ticket's string-field flow, not proof supplied by the bool fixture:
// valid views, server-owned duplicate names, import, rename and reorder converge.
func TestAccRequestSchemaStringViewsAndNames(t *testing.T) {
	const address = "conductorone_request_schema.test"
	const label = "Retool Site or URL 0"
	var resourceID string
	step := func(first, second, firstName, secondName string, explicitDefaults bool) resource.TestStep {
		fieldDefaults, stringDefaults := "", ""
		if explicitDefaults {
			// Import reads server defaults; make them explicit for full-state
			// equality instead of ignoring differences in the import assertion.
			fieldDefaults = `description = ""
      read_only = false
      required = false`
			stringDefaults = `default_value = ""
        placeholder = ""`
		}
		return resource.TestStep{
			Config: providerConfig + fmt.Sprintf(`
resource "conductorone_request_schema" "test" {
  name = "tf-acc-request-schema-iga-3623"
  fields = [
    {
      display_name = %q
      %s
      string_field = {
        %s
        text_field = { multiline = false }
      }
    },
    {
      display_name = %q
      %s
      string_field = {
        %s
        text_field = { multiline = false }
      }
    }
  ]
}
`, first, fieldDefaults, stringDefaults, second, fieldDefaults, stringDefaults),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(address, "id"),
				resource.TestCheckResourceAttrWith(address, "id", func(value string) error {
					if resourceID == "" {
						resourceID = value
					}
					if value != resourceID {
						return fmt.Errorf("request schema was replaced: ID changed")
					}
					return nil
				}),
				resource.TestCheckResourceAttr(address, "fields.0.name", firstName),
				resource.TestCheckResourceAttr(address, "fields.1.name", secondName),
				resource.TestCheckResourceAttr(address, "fields.0.string_field.text_field.multiline", "false"),
				resource.TestCheckResourceAttr(address, "fields.1.string_field.text_field.multiline", "false"),
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PostApplyPreRefresh:  []plancheck.PlanCheck{ExpectEmptyPlanWithDriftLog(t)},
				PostApplyPostRefresh: []plancheck.PlanCheck{ExpectEmptyPlanWithDriftLog(t)},
			},
		}
	}
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Keep the ticket's minimal configuration and empty-plan assertions.
			step(label, label, "Retool_Site_or_URL_0_0", "Retool_Site_or_URL_0_1", false),
			// Compare every imported attribute after explicit default coverage.
			step(label, label, "Retool_Site_or_URL_0_0", "Retool_Site_or_URL_0_1", true),
			{ResourceName: address, ImportState: true, ImportStateVerify: true},
			step("Other URL", label, "Other_URL_0", "Retool_Site_or_URL_0_0", true),
			step(label, "Other URL", "Retool_Site_or_URL_0_0", "Other_URL_0", true),
		},
	})
}
