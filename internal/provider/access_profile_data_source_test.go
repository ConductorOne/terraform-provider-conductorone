package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAccessProfileDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Create an access profile to query
			{
				Config: providerConfig + `
				resource "conductorone_access_profile" "test" {
					display_name = "test-access-profile-data-source"
					description = "test access profile for data source"
					published = true
					request_bundle = true
					visible_to_everyone = true
					enrollment_behavior = "REQUEST_CATALOG_ENROLLMENT_BEHAVIOR_BYPASS_ENTITLEMENT_REQUEST_POLICY"
					unenrollment_behavior = "REQUEST_CATALOG_UNENROLLMENT_BEHAVIOR_LEAVE_ACCESS_AS_IS"
					unenrollment_entitlement_behavior = "REQUEST_CATALOG_UNENROLLMENT_ENTITLEMENT_BEHAVIOR_BYPASS"
				}
				`,
			},
			// Test fetching the access profile by ID
			{
				Config: providerConfig + `
				resource "conductorone_access_profile" "test" {
					display_name = "test-access-profile-data-source"
					description = "test access profile for data source"
					published = true
					request_bundle = true
					visible_to_everyone = true
					enrollment_behavior = "REQUEST_CATALOG_ENROLLMENT_BEHAVIOR_BYPASS_ENTITLEMENT_REQUEST_POLICY"
					unenrollment_behavior = "REQUEST_CATALOG_UNENROLLMENT_BEHAVIOR_LEAVE_ACCESS_AS_IS"
					unenrollment_entitlement_behavior = "REQUEST_CATALOG_UNENROLLMENT_ENTITLEMENT_BEHAVIOR_BYPASS"
				}

				data "conductorone_access_profile" "by_id" {
					id = conductorone_access_profile.test.id
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.conductorone_access_profile.by_id", "display_name", "test-access-profile-data-source"),
					resource.TestCheckResourceAttr("data.conductorone_access_profile.by_id", "description", "test access profile for data source"),
					resource.TestCheckResourceAttr("data.conductorone_access_profile.by_id", "published", "true"),
					resource.TestCheckResourceAttr("data.conductorone_access_profile.by_id", "request_bundle", "true"),
					resource.TestCheckResourceAttr("data.conductorone_access_profile.by_id", "visible_to_everyone", "true"),
					resource.TestCheckResourceAttr("data.conductorone_access_profile.by_id", "enrollment_behavior", "REQUEST_CATALOG_ENROLLMENT_BEHAVIOR_BYPASS_ENTITLEMENT_REQUEST_POLICY"),
					resource.TestCheckResourceAttr("data.conductorone_access_profile.by_id", "unenrollment_behavior", "REQUEST_CATALOG_UNENROLLMENT_BEHAVIOR_LEAVE_ACCESS_AS_IS"),
					resource.TestCheckResourceAttr("data.conductorone_access_profile.by_id", "unenrollment_entitlement_behavior", "REQUEST_CATALOG_UNENROLLMENT_ENTITLEMENT_BEHAVIOR_BYPASS"),
					resource.TestCheckResourceAttrSet("data.conductorone_access_profile.by_id", "id"),
					resource.TestCheckResourceAttrSet("data.conductorone_access_profile.by_id", "created_by_user_id"),
				),
			},
		},
	})
}
