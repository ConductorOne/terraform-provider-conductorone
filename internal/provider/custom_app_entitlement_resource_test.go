package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCustomAppEntitlementResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
				resource "conductorone_app" "test_new_app" {
					display_name = "automated-test"
					description = "this is a test description"
					identity_matching = "APP_USER_IDENTITY_MATCHING_STRICT"
				}

				resource "conductorone_app_resource_type" "custom_role_resource_type" {
					display_name  = "GROUP"
					app_id        = conductorone_app.test_new_app.id
					resource_type = "GROUP"
				}

				resource "conductorone_app_resource" "test_new_app_resource" {
					app_id = conductorone_app.test_new_app.id
					display_name = "automated-test"
					description = "this is a test description"
					app_resource_type_id = conductorone_app_resource_type.custom_role_resource_type.id
				}

				resource "conductorone_risk_level" "high" {
					value = "Extra high"
				}

				resource "conductorone_compliance_framework" "soc2" {
					value = "SOC 2"
				}

				resource "conductorone_webhook" "new_webhook" {
					display_name = "Terraform Created Webhook"
					description = "New description"
					url = "https://www.example.com"
				}

				resource "conductorone_custom_app_entitlement" "test" {
					app_id               = conductorone_app.test_new_app.id
					app_resource_id      = conductorone_app_resource.test_new_app_resource.id
					app_resource_type_id = conductorone_app_resource_type.custom_role_resource_type.id
					display_name         = "Admin"
					alias                = "tf_test_admin_role"
					provision_policy = {
						webhook_provision = {
     						webhook_id = conductorone_webhook.new_webhook.id
    					}
					}
					risk_level_value_id            = resource.conductorone_risk_level.high.id
					slug                           = "member"
					description                    = "Terraform generated admin role"
					compliance_framework_value_ids = [resource.conductorone_compliance_framework.soc2.id]
					purpose                        = "APP_ENTITLEMENT_PURPOSE_VALUE_ASSIGNMENT"
					duration_grant                 = "3601s"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("conductorone_custom_app_entitlement.test", "app_id",
						"conductorone_app.test_new_app", "id"),
					resource.TestCheckResourceAttrPair("conductorone_custom_app_entitlement.test", "app_resource_id",
						"conductorone_app_resource.test_new_app_resource", "id"),
					resource.TestCheckResourceAttrPair("conductorone_custom_app_entitlement.test", "app_resource_type_id",
						"conductorone_app_resource_type.custom_role_resource_type", "id"),
					resource.TestCheckResourceAttr("conductorone_custom_app_entitlement.test", "display_name", "Admin"),
					resource.TestCheckResourceAttr("conductorone_custom_app_entitlement.test", "alias", "tf_test_admin_role"),
					resource.TestCheckResourceAttrPair("conductorone_custom_app_entitlement.test", "provision_policy.webhook_provision.webhook_id",
						"conductorone_webhook.new_webhook", "id"),
					resource.TestCheckResourceAttrPair("conductorone_custom_app_entitlement.test", "risk_level_value_id",
						"conductorone_risk_level.high", "id"),
					resource.TestCheckResourceAttr("conductorone_custom_app_entitlement.test", "slug", "member"),
					resource.TestCheckResourceAttr("conductorone_custom_app_entitlement.test", "description", "Terraform generated admin role"),
					resource.TestCheckResourceAttr("conductorone_custom_app_entitlement.test", "compliance_framework_value_ids.#", "1"),
					resource.TestCheckResourceAttrPair("conductorone_custom_app_entitlement.test", "compliance_framework_value_ids.0",
						"conductorone_compliance_framework.soc2", "id"),
					resource.TestCheckResourceAttr("conductorone_custom_app_entitlement.test", "purpose", "APP_ENTITLEMENT_PURPOSE_VALUE_ASSIGNMENT"),
					resource.TestCheckResourceAttr("conductorone_custom_app_entitlement.test", "duration_grant", "3601s"),
				),
			},
			{
				Config: providerConfig + `
				resource "conductorone_app" "test_new_app" {
					display_name = "automated-test"
					description = "this is a test description"
					identity_matching = "APP_USER_IDENTITY_MATCHING_STRICT"
				}

				resource "conductorone_app_resource_type" "custom_role_resource_type" {
					display_name  = "GROUP"
					app_id        = conductorone_app.test_new_app.id
					resource_type = "GROUP"
				}

				resource "conductorone_app_resource" "test_new_app_resource" {
					app_id = conductorone_app.test_new_app.id
					display_name = "automated-test"
					description = "this is a test description"
					app_resource_type_id = conductorone_app_resource_type.custom_role_resource_type.id
				}

				resource "conductorone_risk_level" "high" {
					value = "Extra high"
				}

				resource "conductorone_compliance_framework" "soc2" {
					value = "SOC 2"
				}

				resource "conductorone_custom_app_entitlement" "test" {
					app_id               = conductorone_app.test_new_app.id
					app_resource_id      = conductorone_app_resource.test_new_app_resource.id
					app_resource_type_id = conductorone_app_resource_type.custom_role_resource_type.id
					display_name         = "Admin changed"
					alias                = "tf_test_admin_role"
					provision_policy = {
						connector_provision = {}
					}
					risk_level_value_id            = resource.conductorone_risk_level.high.id
					slug                           = "member"
					description                    = "Terraform generated admin role changed"
					compliance_framework_value_ids = [resource.conductorone_compliance_framework.soc2.id]
					purpose                        = "APP_ENTITLEMENT_PURPOSE_VALUE_ASSIGNMENT"
					duration_unset				  = {}
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("conductorone_custom_app_entitlement.test", "app_id",
						"conductorone_app.test_new_app", "id"),
					resource.TestCheckResourceAttrPair("conductorone_custom_app_entitlement.test", "app_resource_id",
						"conductorone_app_resource.test_new_app_resource", "id"),
					resource.TestCheckResourceAttrPair("conductorone_custom_app_entitlement.test", "app_resource_type_id",
						"conductorone_app_resource_type.custom_role_resource_type", "id"),
					resource.TestCheckResourceAttr("conductorone_custom_app_entitlement.test", "display_name", "Admin changed"),
					resource.TestCheckResourceAttr("conductorone_custom_app_entitlement.test", "alias", "tf_test_admin_role"),
					resource.TestCheckResourceAttrSet("conductorone_custom_app_entitlement.test", "provision_policy.connector_provision.%"),

					resource.TestCheckResourceAttrPair("conductorone_custom_app_entitlement.test", "risk_level_value_id",
						"conductorone_risk_level.high", "id"),
					resource.TestCheckResourceAttr("conductorone_custom_app_entitlement.test", "slug", "member"),
					resource.TestCheckResourceAttr("conductorone_custom_app_entitlement.test", "description", "Terraform generated admin role changed"),
					resource.TestCheckResourceAttr("conductorone_custom_app_entitlement.test", "compliance_framework_value_ids.#", "1"),
					resource.TestCheckResourceAttrPair("conductorone_custom_app_entitlement.test", "compliance_framework_value_ids.0",
						"conductorone_compliance_framework.soc2", "id"),
					resource.TestCheckResourceAttr("conductorone_custom_app_entitlement.test", "purpose", "APP_ENTITLEMENT_PURPOSE_VALUE_ASSIGNMENT"),

					// Check if "duration_grant" is null or unset
					resource.TestCheckNoResourceAttr("conductorone_custom_app_entitlement.test", "duration_grant"),
				),
			},
		},
	})
}
