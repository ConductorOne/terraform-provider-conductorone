package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAppResourceResource(t *testing.T) {
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

				resource "conductorone_app_resource" "test" {
					app_id = conductorone_app.test.id
					display_name = "automated-test"
					description = "this is a test description"
					app_resource_type_id = conductorone_app_resource_type.test.id
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("conductorone_app_resource.test", "display_name", "automated-test"),
					resource.TestCheckResourceAttr("conductorone_app_resource.test", "description", "this is a test description"),
					resource.TestCheckResourceAttrPair("conductorone_app_resource.test", "app_id", "conductorone_app.test", "id"),
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
					display_name  = "GROUP"
					app_id        = conductorone_app.test.id
					resource_type = "GROUP"
				}

				resource "conductorone_app_resource" "test" {
					app_id = conductorone_app.test.id
					display_name = "automated-test changed"
					description = "this is a changed test description"
					app_resource_type_id = conductorone_app_resource_type.test.id
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("conductorone_app_resource.test", "display_name", "automated-test changed"),
					resource.TestCheckResourceAttr("conductorone_app_resource.test", "description", "this is a changed test description"),
					resource.TestCheckResourceAttrPair("conductorone_app_resource.test", "app_id", "conductorone_app.test", "id"),
				),
			},
		},
	})
}
