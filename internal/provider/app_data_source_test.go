package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAppDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Create two apps to query
			{
				Config: providerConfig + `
				resource "conductorone_app" "test1" {
					display_name = "test-app-data-source-1"
					description = "test app 1 for data source"
				}
				resource "conductorone_app" "test2" {
					display_name = "test-app-data-source-2"
					description = "test app 2 for data source"
				}
				`,
			},
			// Test fetching the first app by display name
			{
				Config: providerConfig + `
				resource "conductorone_app" "test1" {
					display_name = "test-app-data-source-1"
					description = "test app 1 for data source"
				}
				resource "conductorone_app" "test2" {
					display_name = "test-app-data-source-2"
					description = "test app 2 for data source"
				}
				
				data "conductorone_app" "by_name" {
					display_name = conductorone_app.test1.display_name
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.conductorone_app.by_name", "display_name", "test-app-data-source-1"),
					resource.TestCheckResourceAttr("data.conductorone_app.by_name", "description", "test app 1 for data source"),
					resource.TestCheckResourceAttrSet("data.conductorone_app.by_name", "id"),
				),
			},
			// Test fetching the second app by ID
			{
				Config: providerConfig + `
				resource "conductorone_app" "test1" {
					display_name = "test-app-data-source-1"
					description = "test app 1 for data source"
				}
				resource "conductorone_app" "test2" {
					display_name = "test-app-data-source-2"
					description = "test app 2 for data source"
				}
				
				data "conductorone_app" "by_id" {
					app_ids = [conductorone_app.test2.id]
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.conductorone_app.by_id", "display_name", "test-app-data-source-2"),
					resource.TestCheckResourceAttr("data.conductorone_app.by_id", "description", "test app 2 for data source"),
					resource.TestCheckResourceAttrSet("data.conductorone_app.by_id", "id"),
				),
			},
		},
	})
}
