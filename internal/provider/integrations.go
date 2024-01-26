// Generated by tf-integration-gen. DO NOT EDIT.
package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func getIntegrationResources() []func() resource.Resource {
	return []func() resource.Resource{
		NewIntegrationAsanaResource,
		NewIntegrationAwsResource,
		NewIntegrationAwsV2Resource,
		NewIntegrationAzureAdResource,
		NewIntegrationBambooHrResource,
		NewIntegrationBitbucketResource,
		NewIntegrationBoxResource,
		NewIntegrationBroadcomSacResource,
		NewIntegrationBuildkiteResource,
		NewIntegrationCloudamqpResource,
		NewIntegrationCloudflareResource,
		NewIntegrationCloudflareZeroTrustResource,
		NewIntegrationConfluenceResource,
		NewIntegrationCoupaResource,
		NewIntegrationCrowdstrikeResource,
		NewIntegrationDatadogResource,
		NewIntegrationDocusignResource,
		NewIntegrationDuoResource,
		NewIntegrationExpensifyResource,
		NewIntegrationFastlyResource,
		NewIntegrationGithubResource,
		NewIntegrationGithubEnterpriseResource,
		NewIntegrationGitlabResource,
		NewIntegrationGoogleCloudPlatformResource,
		NewIntegrationGoogleIdentityPlatformResource,
		NewIntegrationGoogleWorkspaceResource,
		NewIntegrationGustoResource,
		NewIntegrationHubspotResource,
		NewIntegrationJamfResource,
		NewIntegrationJiraCloudResource,
		NewIntegrationJumpcloudResource,
		NewIntegrationLinearResource,
		NewIntegrationMicrosoft365Resource,
		NewIntegrationNetsuiteResource,
		NewIntegrationNewRelicResource,
		NewIntegrationOktaResource,
		NewIntegrationOneloginResource,
		NewIntegrationOpsgenieResource,
		NewIntegrationPagerdutyResource,
		NewIntegrationPantherResource,
		NewIntegrationRampResource,
		NewIntegrationSalesforceResource,
		NewIntegrationSegmentResource,
		NewIntegrationSendgridResource,
		NewIntegrationSentineloneResource,
		NewIntegrationSentryResource,
		NewIntegrationServicenowResource,
		NewIntegrationSlackResource,
		NewIntegrationSnipeItResource,
		NewIntegrationSnowflakeResource,
		NewIntegrationTableauResource,
		NewIntegrationTailscaleResource,
		NewIntegrationTwingateResource,
		NewIntegrationUkgResource,
		NewIntegrationXeroResource,
		NewIntegrationXsoarResource,
		NewIntegrationZendeskResource,
		NewIntegrationZoomResource,
	}
}
