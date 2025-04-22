package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccBundleAutomationResource(t *testing.T) {
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
				resource "conductorone_custom_app_entitlement" "test" {
					app_id               = conductorone_app.test_new_app.id
					app_resource_id      = conductorone_app_resource.test_new_app_resource.id
					app_resource_type_id = conductorone_app_resource_type.custom_role_resource_type.id
					display_name         = "Admin"
					alias                = "tf_test_admin_role"
					slug                           = "member"
					description                    = "Terraform generated admin role"
					purpose                        = "APP_ENTITLEMENT_PURPOSE_VALUE_ASSIGNMENT"
					duration_grant                 = "3601s"
				}
				resource "conductorone_access_profile" "test" {
					display_name = "automated-test"
					description = "this is a test description"
					enrollment_behavior = "REQUEST_CATALOG_ENROLLMENT_BEHAVIOR_BYPASS_ENTITLEMENT_REQUEST_POLICY"
					published = true
					request_bundle = true
					unenrollment_behavior = "REQUEST_CATALOG_UNENROLLMENT_BEHAVIOR_LEAVE_ACCESS_AS_IS"
					unenrollment_entitlement_behavior = "REQUEST_CATALOG_UNENROLLMENT_ENTITLEMENT_BEHAVIOR_BYPASS"
					visible_to_everyone = true
				}
				resource "conductorone_bundle_automation" "test" {
				   bundle_automation_rule_entitlement = {
				   		entitlement_refs = [
							{
								app_id = conductorone_app.test_new_app.id
								id = conductorone_custom_app_entitlement.test.id
							}
						]
				   }
				   create_tasks = false
				   enabled = true
				   request_catalog_id = conductorone_access_profile.test.id
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("conductorone_access_profile.test", "display_name", "automated-test"),
					resource.TestCheckResourceAttr("conductorone_access_profile.test", "description", "this is a test description"),
					resource.TestCheckResourceAttr("conductorone_access_profile.test", "enrollment_behavior", "REQUEST_CATALOG_ENROLLMENT_BEHAVIOR_BYPASS_ENTITLEMENT_REQUEST_POLICY"),
					resource.TestCheckResourceAttr("conductorone_access_profile.test", "published", "true"),
					resource.TestCheckResourceAttr("conductorone_access_profile.test", "request_bundle", "true"),
					resource.TestCheckResourceAttr("conductorone_access_profile.test", "unenrollment_behavior", "REQUEST_CATALOG_UNENROLLMENT_BEHAVIOR_LEAVE_ACCESS_AS_IS"),
					resource.TestCheckResourceAttr("conductorone_access_profile.test", "visible_to_everyone", "true"),
					resource.TestCheckResourceAttr("conductorone_access_profile.test", "unenrollment_entitlement_behavior", "REQUEST_CATALOG_UNENROLLMENT_ENTITLEMENT_BEHAVIOR_BYPASS"),
					resource.TestCheckResourceAttr("conductorone_bundle_automation.test", "enabled", "true"),
					resource.TestCheckResourceAttr("conductorone_bundle_automation.test", "create_tasks", "false"),
					resource.TestCheckResourceAttrPair("conductorone_bundle_automation.test", "request_catalog_id", "conductorone_access_profile.test", "id"),
					resource.TestCheckResourceAttrPair("conductorone_bundle_automation.test", "bundle_automation_rule_entitlement.entitlement_refs.0.app_id", "conductorone_app.test_new_app", "id"),
					resource.TestCheckResourceAttrPair("conductorone_bundle_automation.test", "bundle_automation_rule_entitlement.entitlement_refs.0.id", "conductorone_custom_app_entitlement.test", "id"),
				),
			},
		},
	})
}
