// Generated by tf-integration-gen. DO NOT EDIT.
package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func getIntegrationResources() []func() resource.Resource {
	return []func() resource.Resource{
		NewIntegrationAsanaResource,
		NewIntegrationAwsResource,
		NewIntegrationAzureAdResource,
		NewIntegrationBambooHrResource,
		NewIntegrationBuildkiteResource,
		NewIntegrationCloudamqpResource,
		NewIntegrationCloudflareResource,
		NewIntegrationCloudflareZeroTrustResource,
		NewIntegrationConfluenceResource,
		NewIntegrationCoupaResource,
		NewIntegrationDatadogResource,
		NewIntegrationDocusignResource,
		NewIntegrationDuoResource,
		NewIntegrationExpensifyResource,
		NewIntegrationGithubResource,
		NewIntegrationGithubEnterpriseResource,
		NewIntegrationGitlabResource,
		NewIntegrationGoogleCloudPlatformResource,
		NewIntegrationGoogleIdentityPlatformResource,
		NewIntegrationGoogleWorkspaceResource,
		NewIntegrationGustoResource,
		NewIntegrationJiraCloudResource,
		NewIntegrationJumpcloudResource,
		NewIntegrationLinearResource,
		NewIntegrationNetsuiteResource,
		NewIntegrationOktaResource,
		NewIntegrationOneloginResource,
		NewIntegrationOpsgenieResource,
		NewIntegrationPantherResource,
		NewIntegrationRampResource,
		NewIntegrationSalesforceResource,
		NewIntegrationSendgridResource,
		NewIntegrationSentryResource,
		NewIntegrationSlackResource,
		NewIntegrationSnowflakeResource,
		NewIntegrationTailscaleResource,
		NewIntegrationTwingateResource,
		NewIntegrationUkgResource,
		NewIntegrationZendeskResource,
	}
}