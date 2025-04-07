package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAppResourceTypeDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Create an app resource type to query
			{
				Config: providerConfig + `
				resource "conductorone_app" "test" {
					display_name = "test-app-resource-type-ds"
					description = "test app for app resource type data source"
				}

				resource "conductorone_app_resource_type" "test" {
					display_name  = "GROUP"
					app_id       = conductorone_app.test.id
					resource_type = "GROUP"
				}
				`,
			},
			// Test searching for the app resource type by display name
			{
				Config: providerConfig + `
				resource "conductorone_app" "test" {
					display_name = "test-app-resource-type-ds"
					description = "test app for app resource type data source"
				}

				resource "conductorone_app_resource_type" "test" {
					display_name  = "GROUP"
					app_id       = conductorone_app.test.id
					resource_type = "GROUP"
				}

				data "conductorone_app_resource_type" "by_name" {
					display_name = "GROUP"
					app_ids = [conductorone_app.test.id]
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.conductorone_app_resource_type.by_name", "display_name", "GROUP"),
					resource.TestCheckResourceAttrSet("data.conductorone_app_resource_type.by_name", "id"),
					resource.TestCheckResourceAttrPair("data.conductorone_app_resource_type.by_name", "app_id", "conductorone_app.test", "id"),
				),
			},
		},
	})
}
