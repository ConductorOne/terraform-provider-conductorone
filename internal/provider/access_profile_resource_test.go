package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccRequestAccessProfileResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
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
				),
			},
			{
				Config: providerConfig + `
				resource "conductorone_access_profile" "test" {
					display_name = ""
					description = ""
					enrollment_behavior = "REQUEST_CATALOG_ENROLLMENT_BEHAVIOR_ENFORCE_ENTITLEMENT_REQUEST_POLICY"
					published = false
					request_bundle = false
					unenrollment_behavior = "REQUEST_CATALOG_UNENROLLMENT_BEHAVIOR_LEAVE_ACCESS_AS_IS"
					unenrollment_entitlement_behavior = "REQUEST_CATALOG_UNENROLLMENT_ENTITLEMENT_BEHAVIOR_BYPASS"
					visible_to_everyone = false
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("conductorone_access_profile.test", "display_name", "automated-test changed"),
					resource.TestCheckResourceAttr("conductorone_access_profile.test", "description", "this is a changed test description"),
					resource.TestCheckResourceAttr("conductorone_access_profile.test", "enrollment_behavior", "REQUEST_CATALOG_ENROLLMENT_BEHAVIOR_ENFORCE_ENTITLEMENT_REQUEST_POLICY"),
					resource.TestCheckResourceAttr("conductorone_access_profile.test", "published", "false"),
					resource.TestCheckResourceAttr("conductorone_access_profile.test", "request_bundle", "false"),
					resource.TestCheckResourceAttr("conductorone_access_profile.test", "unenrollment_behavior", "REQUEST_CATALOG_UNENROLLMENT_BEHAVIOR_REVOKE_UNJUSTIFIED"),
					resource.TestCheckResourceAttr("conductorone_access_profile.test", "visible_to_everyone", "false"),
					resource.TestCheckResourceAttr("conductorone_access_profile.test", "unenrollment_entitlement_behavior", "REQUEST_CATALOG_UNENROLLMENT_ENTITLEMENT_BEHAVIOR_ENFORCE"),
				),
			},
		},
	})
}
