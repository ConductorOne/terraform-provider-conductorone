package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAppResourceDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Create an app resource to query
			{
				Config: providerConfig + `
				resource "conductorone_app" "test" {
					display_name = "test-app-resource-ds"
					description = "test app for app resource data source"
				}

				resource "conductorone_app_resource_type" "test" {
					display_name  = "GROUP"
					app_id       = conductorone_app.test.id
					resource_type = "GROUP"
				}

				resource "conductorone_app_resource" "test" {
					app_id = conductorone_app.test.id
					display_name = "test-app-resource"
					description = "test app resource for data source"
					app_resource_type_id = conductorone_app_resource_type.test.id
				}
				`,
			},
			// Test fetching the app resource by ID
			{
				Config: providerConfig + `
				resource "conductorone_app" "test" {
					display_name = "test-app-resource-ds"
					description = "test app for app resource data source"
				}

				resource "conductorone_app_resource_type" "test" {
					display_name  = "GROUP"
					app_id       = conductorone_app.test.id
					resource_type = "GROUP"
				}

				resource "conductorone_app_resource" "test" {
					app_id = conductorone_app.test.id
					display_name = "test-app-resource"
					description = "test app resource for data source"
					app_resource_type_id = conductorone_app_resource_type.test.id
				}

				data "conductorone_app_resource" "by_id" {
					app_id = conductorone_app.test.id
					app_resource_type_id = conductorone_app_resource_type.test.id
					id = conductorone_app_resource.test.id
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.conductorone_app_resource.by_id", "display_name", "test-app-resource"),
					resource.TestCheckResourceAttr("data.conductorone_app_resource.by_id", "description", "test app resource for data source"),
					resource.TestCheckResourceAttrSet("data.conductorone_app_resource.by_id", "id"),
					resource.TestCheckResourceAttrPair("data.conductorone_app_resource.by_id", "app_id", "conductorone_app.test", "id"),
					resource.TestCheckResourceAttrPair("data.conductorone_app_resource.by_id", "app_resource_type_id", "conductorone_app_resource_type.test", "id"),
				),
			},
		},
	})
}
