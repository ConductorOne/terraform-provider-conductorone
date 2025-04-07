package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccPolicyDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Create a policy to query
			{
				Config: providerConfig + `
				resource "conductorone_policy" "test" {
					display_name                = "test-policy-data-source"
					description                 = "test policy for data source"
					policy_type                 = "POLICY_TYPE_CERTIFY"
					reassign_tasks_to_delegates = "false"
					policy_steps = {
						certify = {
							steps = [
								{
									approval = {
										allow_reassignment          = "true"
										require_reassignment_reason = "false"
										require_approval_reason     = "false"
										app_owner_approval = {
											allow_self_approval = "true"
										}
									}
								}
							]
						}
					}
				}
				`,
			},
			// Test fetching the policy by display name
			{
				Config: providerConfig + `
				resource "conductorone_policy" "test" {
					display_name                = "test-policy-data-source"
					description                 = "test policy for data source"
					policy_type                 = "POLICY_TYPE_CERTIFY"
					reassign_tasks_to_delegates = "false"
					policy_steps = {
						certify = {
							steps = [
								{
									approval = {
										allow_reassignment          = "true"
										require_reassignment_reason = "false"
										require_approval_reason     = "false"
										app_owner_approval = {
											allow_self_approval = "true"
										}
									}
								}
							]
						}
					}
				}

				data "conductorone_policy" "by_name" {
					display_name = conductorone_policy.test.display_name
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.conductorone_policy.by_name", "display_name", "test-policy-data-source"),
					resource.TestCheckResourceAttr("data.conductorone_policy.by_name", "description", "test policy for data source"),
					resource.TestCheckResourceAttr("data.conductorone_policy.by_name", "policy_type", "POLICY_TYPE_CERTIFY"),
					resource.TestCheckResourceAttrSet("data.conductorone_policy.by_name", "id"),
				),
			},
		},
	})
}
