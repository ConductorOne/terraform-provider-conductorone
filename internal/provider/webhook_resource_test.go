package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccWebhookResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
				resource "conductorone_webhook" "test" {
					display_name = "automated-test"
					description = "this is a test description"
					url = "https://example.com"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("conductorone_webhook.test", "display_name", "automated-test"),
					resource.TestCheckResourceAttr("conductorone_webhook.test", "description", "this is a test description"),
					resource.TestCheckResourceAttr("conductorone_webhook.test", "url", "https://example.com"),
				),
			},
			{
				Config: providerConfig + `
				resource "conductorone_webhook" "test" {
					display_name = "automated-test changed"
					description = "this is a changed test description"
					url = "https://examplechanged.com"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("conductorone_webhook.test", "display_name", "automated-test changed"),
					resource.TestCheckResourceAttr("conductorone_webhook.test", "description", "this is a changed test description"),
					resource.TestCheckResourceAttr("conductorone_webhook.test", "url", "https://examplechanged.com"),
				),
			},
		},
	})
}
