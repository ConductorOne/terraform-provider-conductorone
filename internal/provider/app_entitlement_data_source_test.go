package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAppEntitlementDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Create an app with a custom entitlement to query
			{
				Config: providerConfig + `
				resource "conductorone_app" "test" {
					display_name = "test-app-entitlement-ds"
					description = "test app for entitlement data source"
				}

				resource "conductorone_app_resource_type" "test" {
					display_name  = "GROUP"
					app_id       = conductorone_app.test.id
					resource_type = "GROUP"
				}

				resource "conductorone_app_resource" "test" {
					app_id = conductorone_app.test.id
					display_name = "test-resource"
					description = "test resource for entitlement data source"
					app_resource_type_id = conductorone_app_resource_type.test.id
				}

				resource "conductorone_risk_level" "high" {
					value = "Extra high"
				}

				resource "conductorone_compliance_framework" "soc2" {
					value = "SOC 2"
				}

				resource "conductorone_custom_app_entitlement" "test" {
					app_id               = conductorone_app.test.id
					app_resource_id      = conductorone_app_resource.test.id
					app_resource_type_id = conductorone_app_resource_type.test.id
					display_name         = "test-entitlement"
					description         = "test entitlement for data source"
					alias               = "test-alias"
					provision_policy = {
						connector_provision = {}
					}
					risk_level_value_id = resource.conductorone_risk_level.high.id
					compliance_framework_value_ids = [resource.conductorone_compliance_framework.soc2.id]
					purpose = "APP_ENTITLEMENT_PURPOSE_VALUE_ASSIGNMENT"
					duration_unset = {}
				}
				`,
			},
			// Test fetching the entitlement by alias
			{
				Config: providerConfig + `
				resource "conductorone_app" "test" {
					display_name = "test-app-entitlement-ds"
					description = "test app for entitlement data source"
				}

				resource "conductorone_app_resource_type" "test" {
					display_name  = "GROUP"
					app_id       = conductorone_app.test.id
					resource_type = "GROUP"
				}

				resource "conductorone_app_resource" "test" {
					app_id = conductorone_app.test.id
					display_name = "test-resource"
					description = "test resource for entitlement data source"
					app_resource_type_id = conductorone_app_resource_type.test.id
				}

				resource "conductorone_risk_level" "high" {
					value = "Extra high"
				}

				resource "conductorone_compliance_framework" "soc2" {
					value = "SOC 2"
				}

				resource "conductorone_custom_app_entitlement" "test" {
					app_id               = conductorone_app.test.id
					app_resource_id      = conductorone_app_resource.test.id
					app_resource_type_id = conductorone_app_resource_type.test.id
					display_name         = "test-entitlement"
					description         = "test entitlement for data source"
					alias               = "test-alias"
					provision_policy = {
						connector_provision = {}
					}
					risk_level_value_id = resource.conductorone_risk_level.high.id
					compliance_framework_value_ids = [resource.conductorone_compliance_framework.soc2.id]
					purpose = "APP_ENTITLEMENT_PURPOSE_VALUE_ASSIGNMENT"
					duration_unset = {}
				}

				data "conductorone_app_entitlement" "by_alias" {
					alias = conductorone_custom_app_entitlement.test.alias
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.conductorone_app_entitlement.by_alias", "display_name", "test-entitlement"),
					resource.TestCheckResourceAttr("data.conductorone_app_entitlement.by_alias", "description", "test entitlement for data source"),
					resource.TestCheckResourceAttr("data.conductorone_app_entitlement.by_alias", "alias", "test-alias"),
					resource.TestCheckResourceAttrSet("data.conductorone_app_entitlement.by_alias", "id"),
					resource.TestCheckResourceAttrPair("data.conductorone_app_entitlement.by_alias", "app_id", "conductorone_app.test", "id"),
					resource.TestCheckResourceAttrPair("data.conductorone_app_entitlement.by_alias", "app_resource_id", "conductorone_app_resource.test", "id"),
					resource.TestCheckResourceAttrPair("data.conductorone_app_entitlement.by_alias", "app_resource_type_id", "conductorone_app_resource_type.test", "id"),
				),
			},
		},
	})
}
