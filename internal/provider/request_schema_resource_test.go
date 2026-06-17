package provider

import (
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
// already matched the read-back never converged. The fix marks each via
// x-speakeasy-param-suppress-computed-diff in overlay.yaml:
//
//	created_at, justification_visibility, field_groups, field_relationships,
//	and the nested fields[].name.
//
// The config below mirrors the original repro (a single bool_field with
// default_value) and intentionally leaves every server-derived attribute
// UNSET — including the nested fields[].name, which the API derives and
// normalizes server-side. That is what makes this a real convergence test:
// those attributes are populated only by the server, so an empty plan can hold
// only if their SuppressDiff modifiers propagate the prior state value when the
// planned value is unknown, instead of re-planning as "(known after apply)".
// (Setting fields[].name explicitly would not exercise the fix — SuppressDiff
// is a no-op for an explicit config value, and the server's normalized name
// would then perpetually mismatch the literal.) The empty-plan assertion runs
// on BOTH the pre-refresh plan (state vs config) and the post-refresh plan
// (live read-back vs config).
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
