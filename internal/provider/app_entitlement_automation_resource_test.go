package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAppEntitlementAutomationResource(t *testing.T) {
	const fixture = `
				resource "conductorone_app" "test" {
					display_name = "test-app-entitlement-automation"
					description  = "test app for entitlement automation"
				}

				resource "conductorone_app_resource_type" "test" {
					app_id        = conductorone_app.test.id
					display_name  = "test-group"
					resource_type = "GROUP"
				}

				resource "conductorone_app_resource" "test" {
					app_id               = conductorone_app.test.id
					display_name         = "test-resource"
					description          = "test resource for automation"
					app_resource_type_id = conductorone_app_resource_type.test.id
				}

				resource "conductorone_custom_app_entitlement" "test" {
					app_id               = conductorone_app.test.id
					app_resource_id      = conductorone_app_resource.test.id
					app_resource_type_id = conductorone_app_resource_type.test.id
					display_name         = "test-entitlement"
					description          = "test entitlement for automation"
					provision_policy = {
						connector_provision = {}
					}
					duration_unset = {}
				}
	`

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Create an automation with basic rule
			{
				Config: providerConfig + fixture + `
				resource "conductorone_app_entitlement_automation" "test" {
					app_id             = conductorone_app.test.id
					app_entitlement_id = conductorone_custom_app_entitlement.test.id
					display_name       = "test-automation"
					description        = "test automation with basic rule"
					app_entitlement_automation_rule_basic = {
						expression = "true"
					}
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("conductorone_app_entitlement_automation.test", "display_name", "test-automation"),
					resource.TestCheckResourceAttr("conductorone_app_entitlement_automation.test", "description", "test automation with basic rule"),
					resource.TestCheckResourceAttr("conductorone_app_entitlement_automation.test", "app_entitlement_automation_rule_basic.expression", "true"),
					resource.TestCheckResourceAttrPair("conductorone_app_entitlement_automation.test", "app_id", "conductorone_app.test", "id"),
					resource.TestCheckResourceAttrPair("conductorone_app_entitlement_automation.test", "app_entitlement_id", "conductorone_custom_app_entitlement.test", "id"),
				),
			},
			// Update to use CEL rule
			{
				Config: providerConfig + fixture + `
				resource "conductorone_app_entitlement_automation" "test" {
					app_id             = conductorone_app.test.id
					app_entitlement_id = conductorone_custom_app_entitlement.test.id
					display_name       = "test-automation-updated"
					description        = "test automation with CEL rule"
					app_entitlement_automation_rule_cel = {
						expression = "true"
					}
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("conductorone_app_entitlement_automation.test", "display_name", "test-automation-updated"),
					resource.TestCheckResourceAttr("conductorone_app_entitlement_automation.test", "description", "test automation with CEL rule"),
					resource.TestCheckResourceAttr("conductorone_app_entitlement_automation.test", "app_entitlement_automation_rule_cel.expression", "true"),
					resource.TestCheckResourceAttrPair("conductorone_app_entitlement_automation.test", "app_id", "conductorone_app.test", "id"),
					resource.TestCheckResourceAttrPair("conductorone_app_entitlement_automation.test", "app_entitlement_id", "conductorone_custom_app_entitlement.test", "id"),
				),
			},
			// Update to use entitlement rule
			{
				Config: providerConfig + fixture + `
				resource "conductorone_app_entitlement_automation" "test" {
					app_id             = conductorone_app.test.id
					app_entitlement_id = conductorone_custom_app_entitlement.test.id
					display_name       = "test-automation-entitlement"
					description        = "test automation with entitlement rule"
					app_entitlement_automation_rule_entitlement = {
						entitlement_refs = [
							{
								app_id = conductorone_app.test.id
								id     = conductorone_custom_app_entitlement.test.id
							}
						]
					}
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("conductorone_app_entitlement_automation.test", "display_name", "test-automation-entitlement"),
					resource.TestCheckResourceAttr("conductorone_app_entitlement_automation.test", "description", "test automation with entitlement rule"),
					resource.TestCheckResourceAttrPair("conductorone_app_entitlement_automation.test", "app_entitlement_automation_rule_entitlement.entitlement_refs.0.app_id", "conductorone_app.test", "id"),
					resource.TestCheckResourceAttrPair("conductorone_app_entitlement_automation.test", "app_entitlement_automation_rule_entitlement.entitlement_refs.0.id", "conductorone_custom_app_entitlement.test", "id"),
					resource.TestCheckResourceAttrPair("conductorone_app_entitlement_automation.test", "app_id", "conductorone_app.test", "id"),
					resource.TestCheckResourceAttrPair("conductorone_app_entitlement_automation.test", "app_entitlement_id", "conductorone_custom_app_entitlement.test", "id"),
				),
			},
		},
	})
}
