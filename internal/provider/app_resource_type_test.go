package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAppResourceTypeResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
				resource "conductorone_app" "test" {
					display_name = "automated-test"
					description = "this is a test description"
					identity_matching = "APP_USER_IDENTITY_MATCHING_STRICT"
				}

				resource "conductorone_app_resource_type" "test" {
					display_name  = "GROUP"
					app_id        = conductorone_app.test.id
					resource_type = "GROUP"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("conductorone_app_resource_type.test", "display_name", "GROUP"),
					resource.TestCheckResourceAttrPair("conductorone_app_resource_type.test", "app_id", "conductorone_app.test", "id"),
					resource.TestCheckResourceAttr("conductorone_app_resource_type.test", "resource_type", "GROUP"),
				),
			},
			{
				Config: providerConfig + `
				resource "conductorone_app" "test" {
					display_name = "automated-test"
					description = "this is a test description"
					identity_matching = "APP_USER_IDENTITY_MATCHING_STRICT"
				}

				resource "conductorone_app_resource_type" "test" {
					display_name  = "ROLE"
					app_id        = conductorone_app.test.id
					resource_type = "ROLE"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("conductorone_app_resource_type.test", "display_name", "ROLE"),
					resource.TestCheckResourceAttr("conductorone_app_resource_type.test", "resource_type", "ROLE"),
					resource.TestCheckResourceAttrPair("conductorone_app_resource_type.test", "app_id", "conductorone_app.test", "id"),
				),
			},
		},
	})
}
