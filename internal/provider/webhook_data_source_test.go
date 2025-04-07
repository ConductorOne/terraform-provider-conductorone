package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccWebhookDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Create a webhook to query
			{
				Config: providerConfig + `
				resource "conductorone_webhook" "test" {
					display_name = "test-webhook-data-source"
					description = "test webhook for data source"
					url = "https://example.com/webhook"
				}
				`,
			},
			// Test fetching the webhook by display name
			{
				Config: providerConfig + `
				resource "conductorone_webhook" "test" {
					display_name = "test-webhook-data-source"
					description = "test webhook for data source"
					url = "https://example.com/webhook"
				}

				data "conductorone_webhook" "by_name" {
					query = conductorone_webhook.test.display_name
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.conductorone_webhook.by_name", "display_name", "test-webhook-data-source"),
					resource.TestCheckResourceAttr("data.conductorone_webhook.by_name", "description", "test webhook for data source"),
					resource.TestCheckResourceAttr("data.conductorone_webhook.by_name", "url", "https://example.com/webhook"),
					resource.TestCheckResourceAttrSet("data.conductorone_webhook.by_name", "id"),
				),
			},
		},
	})
}
