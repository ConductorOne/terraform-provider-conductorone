package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccPolicyResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
				resource "conductorone_policy" "test" {
					display_name                = "Terraform Created Certify Policy"
					description                 = "New description"
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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("conductorone_policy.test", "display_name", "Terraform Created Certify Policy"),
					resource.TestCheckResourceAttr("conductorone_policy.test", "description", "New description"),
					resource.TestCheckResourceAttr("conductorone_policy.test", "policy_type", "POLICY_TYPE_CERTIFY"),
					resource.TestCheckResourceAttr("conductorone_policy.test", "reassign_tasks_to_delegates", "false"),
					resource.TestCheckResourceAttr("conductorone_policy.test", "policy_steps.certify.steps.0.approval.allow_reassignment", "true"),
					resource.TestCheckResourceAttr("conductorone_policy.test", "policy_steps.certify.steps.0.approval.require_reassignment_reason", "false"),
					resource.TestCheckResourceAttr("conductorone_policy.test", "policy_steps.certify.steps.0.approval.require_approval_reason", "false"),
					resource.TestCheckResourceAttr("conductorone_policy.test", "policy_steps.certify.steps.0.approval.app_owner_approval.allow_self_approval", "true"),
				),
			},
			{
				Config: providerConfig + `
				resource "conductorone_policy" "test" {
					display_name                = "Terraform Created Certify Policy"
					description                 = "now manager approval"
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
										manager_approval = {
											fallback = "false"
										}
									}
								}
							]
						}
					}
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("conductorone_policy.test", "display_name", "Terraform Created Certify Policy"),
					resource.TestCheckResourceAttr("conductorone_policy.test", "description", "now manager approval"),
					resource.TestCheckResourceAttr("conductorone_policy.test", "policy_type", "POLICY_TYPE_CERTIFY"),
					resource.TestCheckResourceAttr("conductorone_policy.test", "reassign_tasks_to_delegates", "false"),
					resource.TestCheckResourceAttr("conductorone_policy.test", "policy_steps.certify.steps.0.approval.allow_reassignment", "true"),
					resource.TestCheckResourceAttr("conductorone_policy.test", "policy_steps.certify.steps.0.approval.require_reassignment_reason", "false"),
					resource.TestCheckResourceAttr("conductorone_policy.test", "policy_steps.certify.steps.0.approval.require_approval_reason", "false"),
					resource.TestCheckResourceAttr("conductorone_policy.test", "policy_steps.certify.steps.0.approval.manager_approval.fallback", "false"),
				),
			},
		},
	})
}
