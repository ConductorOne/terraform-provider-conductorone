package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAppResource(t *testing.T) {
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
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("conductorone_app.test", "display_name", "automated-test"),
				),
			},
			{
				Config: providerConfig + `
				resource "conductorone_app" "test" {
					display_name = "automated-test changed"
					description = "this is a changed test description"
					identity_matching = "APP_USER_IDENTITY_MATCHING_DISPLAY_NAME"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("conductorone_app.test", "display_name", "automated-test changed"),
					resource.TestCheckResourceAttr("conductorone_app.test", "description", "this is a changed test description"),
					resource.TestCheckResourceAttr("conductorone_app.test", "identity_matching", "APP_USER_IDENTITY_MATCHING_DISPLAY_NAME"),
				),
			},
			{
				Config: providerConfig + `
				resource "conductorone_app" "test" {
					display_name = "automated-test"
			    }
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("conductorone_app.test", "display_name", "automated-test"),
					resource.TestCheckResourceAttr("conductorone_app.test", "description", "this is a changed test description"),
					resource.TestCheckResourceAttr("conductorone_app.test", "identity_matching", "APP_USER_IDENTITY_MATCHING_DISPLAY_NAME"),
				),
			},
		},
	})
}
