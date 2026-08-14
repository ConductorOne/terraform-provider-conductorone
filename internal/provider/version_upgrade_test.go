package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

// upgradeFromVersion is the last published provider release — the version a
// practitioner upgrades FROM when they take the build under test, so it is the
// statefile shape the upgrade has to survive.
//
// Bump this on every release. A stale constant still passes while testing an
// upgrade nobody performs any more.
const upgradeFromVersion = "1.4.3"

// providerNamespaceEnvVar overrides the namespace terraform-plugin-testing
// stamps into the generated required_providers block for an in-tree provider
// factory. It defaults to "hashicorp", which addresses the provider under test
// as registry.terraform.io/hashicorp/conductorone while the ExternalProviders
// step wrote its statefile under registry.terraform.io/conductorone/conductorone.
// Terraform 1.9 silently re-addresses the state in that case rather than
// failing, so the test still passes — on an upgrade path no practitioner walks.
// Setting the namespace keeps both steps on the real published address.
const providerNamespaceEnvVar = "TF_ACC_PROVIDER_NAMESPACE"

// upgradeBaseConfig is applied by both the released provider and the provider
// under test, byte for byte. Any difference between the two configs would show
// up as a plan diff attributable to the config rather than to the provider.
//
// deprovisioner_policy is absent here: it is Computed-only in
// upgradeFromVersion, so a config the released provider must accept cannot
// carry it. Its statefile value comes from the API, which is exactly the
// pre-upgrade state whose survival is under test.
//
// Both entitlements set duration_grant. Omitting every duration_* attribute
// leaves the API returning duration_unset = {} against a null config, which
// plans an unrelated perpetual update (issue #229, still open) and would mask
// the upgrade signal this test exists to isolate. Deleting the duration_grant
// lines is a way to reproduce #229, not a way to widen this test.
const upgradeBaseConfig = `
resource "conductorone_app" "upgrade_app" {
  display_name      = "tf-upgrade-test-app"
  description       = "provider version upgrade acceptance test"
  identity_matching = "APP_USER_IDENTITY_MATCHING_STRICT"
}

resource "conductorone_app_resource_type" "upgrade_resource_type" {
  display_name  = "GROUP"
  app_id        = conductorone_app.upgrade_app.id
  resource_type = "GROUP"
}

resource "conductorone_app_resource" "upgrade_resource" {
  app_id               = conductorone_app.upgrade_app.id
  display_name         = "tf-upgrade-test-resource"
  description          = "provider version upgrade acceptance test"
  app_resource_type_id = conductorone_app_resource_type.upgrade_resource_type.id
}

resource "conductorone_webhook" "upgrade_webhook" {
  display_name = "tf-upgrade-test-webhook"
  description  = "provider version upgrade acceptance test"
  url          = "https://www.example.com"
}

resource "conductorone_custom_app_entitlement" "upgrade_webhook_provision" {
  app_id               = conductorone_app.upgrade_app.id
  app_resource_id      = conductorone_app_resource.upgrade_resource.id
  app_resource_type_id = conductorone_app_resource_type.upgrade_resource_type.id
  display_name         = "tf-upgrade webhook provision"
  alias                = "tf_upgrade_webhook_provision"
  slug                 = "tf-upgrade-webhook"
  description          = "provider version upgrade acceptance test"
  purpose              = "APP_ENTITLEMENT_PURPOSE_VALUE_ASSIGNMENT"
  duration_grant       = "3601s"
  provision_policy = {
    webhook_provision = {
      webhook_id = conductorone_webhook.upgrade_webhook.id
    }
  }
}
`

// upgradeConnectorEntitlementConfig is split out so the deprovisioner_policy
// step can re-declare the same entitlement with the block added and nothing
// else changed.
const upgradeConnectorEntitlementConfig = `
resource "conductorone_custom_app_entitlement" "upgrade_connector_provision" {
  app_id               = conductorone_app.upgrade_app.id
  app_resource_id      = conductorone_app_resource.upgrade_resource.id
  app_resource_type_id = conductorone_app_resource_type.upgrade_resource_type.id
  display_name         = "tf-upgrade connector provision"
  alias                = "tf_upgrade_connector_provision"
  slug                 = "tf-upgrade-connector"
  description          = "provider version upgrade acceptance test"
  purpose              = "APP_ENTITLEMENT_PURPOSE_VALUE_ASSIGNMENT"
  duration_grant       = "3602s"
  provision_policy = {
    connector_provision = {}
  }
}
`

// upgradeConnectorEntitlementWithDeprovisionerConfig is the same entitlement
// with deprovisioner_policy written from configuration, which only the provider
// under test can accept.
//
// connector_provision.default_behavior is the arm the API accepts for a
// non-access entitlement; delete_account is rejected with "cannot set delete
// account connector deprovisioner policy for non-access entitlement".
const upgradeConnectorEntitlementWithDeprovisionerConfig = `
resource "conductorone_custom_app_entitlement" "upgrade_connector_provision" {
  app_id               = conductorone_app.upgrade_app.id
  app_resource_id      = conductorone_app_resource.upgrade_resource.id
  app_resource_type_id = conductorone_app_resource_type.upgrade_resource_type.id
  display_name         = "tf-upgrade connector provision"
  alias                = "tf_upgrade_connector_provision"
  slug                 = "tf-upgrade-connector"
  description          = "provider version upgrade acceptance test"
  purpose              = "APP_ENTITLEMENT_PURPOSE_VALUE_ASSIGNMENT"
  duration_grant       = "3602s"
  provision_policy = {
    connector_provision = {}
  }
  deprovisioner_policy = {
    connector_provision = {
      default_behavior = {}
    }
  }
}
`

// TestAccCustomAppEntitlementVersionUpgrade applies a config with the last
// published provider release, then re-plans the identical config with the
// provider under test and requires an empty plan.
//
// This is the only test in the repo that reads a statefile written by a
// different provider version. Every other acceptance test creates its state
// with the build under test, so none of them can see a schema change that reads
// a previously-written statefile differently — the failure mode behind the
// perpetual-drift (#229) and import-churn (#236) bugs.
//
// custom_app_entitlement.deprovisioner_policy is the surface under test: it
// moves from Computed-only to Computed+Optional, which changes how the plan
// reconciles a null config against a statefile value.
func TestAccCustomAppEntitlementVersionUpgrade(t *testing.T) {
	t.Setenv(providerNamespaceEnvVar, "conductorone")

	upgradeConfig := providerConfig + upgradeBaseConfig + upgradeConnectorEntitlementConfig
	deprovisionerConfig := providerConfig + upgradeBaseConfig + upgradeConnectorEntitlementWithDeprovisionerConfig

	entitlements := []string{
		"conductorone_custom_app_entitlement.upgrade_webhook_provision",
		"conductorone_custom_app_entitlement.upgrade_connector_provision",
	}

	releasedProviderOnly := map[string]resource.ExternalProvider{
		"conductorone": {
			Source:            "conductorone/conductorone",
			VersionConstraint: upgradeFromVersion,
		},
	}

	emptyPlanAfterApply := resource.ConfigPlanChecks{
		PostApplyPreRefresh:  []plancheck.PlanCheck{ExpectEmptyPlanWithDriftLog(t)},
		PostApplyPostRefresh: []plancheck.PlanCheck{ExpectEmptyPlanWithDriftLog(t)},
	}

	resource.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			// Written by the released provider. deprovisioner_policy is
			// Computed-only here, so whatever lands in state came from the API.
			{
				ExternalProviders: releasedProviderOnly,
				Config:            upgradeConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(entitlements[0], "provision_policy.webhook_provision.webhook_id",
						"conductorone_webhook.upgrade_webhook", "id"),
					resource.TestCheckResourceAttrSet(entitlements[1], "provision_policy.connector_provision.%"),
				),
			},
			// The upgrade assertion: the first plan the provider under test
			// makes against the released provider's state must be a no-op.
			{
				ProtoV6ProviderFactories: testAccProviderFactories,
				Config:                   upgradeConfig,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply:             []plancheck.PlanCheck{ExpectEmptyPlanWithDriftLog(t)},
					PostApplyPreRefresh:  emptyPlanAfterApply.PostApplyPreRefresh,
					PostApplyPostRefresh: emptyPlanAfterApply.PostApplyPostRefresh,
				},
			},
			// The capability the upgrade unlocks: deprovisioner_policy written
			// from configuration has to round-trip, or it becomes the next
			// perpetual diff.
			{
				ProtoV6ProviderFactories: testAccProviderFactories,
				Config:                   deprovisionerConfig,
				ConfigPlanChecks:         emptyPlanAfterApply,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(entitlements[1],
						"deprovisioner_policy.connector_provision.default_behavior.%"),
				),
			},
		},
	})
}
