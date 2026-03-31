package provider

import (
	"context"
	"fmt"
	"os"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &ConductoroneProvider{}

type ConductoroneProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// TerraformProviderModel describes the provider data model.
type ConductoroneProviderModel struct {
	ServerURL    types.String `tfsdk:"server_url"`
	ClientID     types.String `tfsdk:"client_id"`
	ClientSecret types.String `tfsdk:"client_secret"`
	TenantDomain types.String `tfsdk:"tenant_domain"`
	OIDCToken    types.String `tfsdk:"oidc_token"`
}

func (p *ConductoroneProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "conductorone"
	resp.Version = p.version
}

func (p *ConductoroneProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "ConductorOne API: The ConductorOne API is a HTTP API for managing ConductorOne resources.\n\n" +
			"## Authentication\n\n" +
			"The provider supports multiple authentication methods. When more than one is configured, " +
			"the following priority order applies (highest priority first):\n\n" +
			"1. **`CONDUCTORONE_ACCESS_TOKEN`** env var -- static bearer token (overrides all other auth)\n" +
			"2. **`oidc_token`** attribute / **`CONDUCTORONE_OIDC_TOKEN`** / **`TFC_WORKLOAD_IDENTITY_TOKEN`** env -- workload federation (RFC 8693 token exchange)\n" +
			"3. **`client_id`** + **`client_secret`** attributes / **`CONDUCTORONE_CLIENT_ID`** + **`CONDUCTORONE_CLIENT_SECRET`** env -- API credentials (Ed25519 JWT assertion)\n\n" +
			"For each attribute, the provider attribute value takes precedence over the corresponding environment variable.",
		Attributes: map[string]schema.Attribute{
			"server_url": schema.StringAttribute{
				MarkdownDescription: "Server URL (defaults to `https://{tenant_domain}.conductor.one`). " +
					"Can also be set via `CONDUCTORONE_SERVER_URL` env var.",
				Optional: true,
			},
			"client_id": schema.StringAttribute{
				MarkdownDescription: "Client ID for authentication. Required for OIDC and credential-based auth. " +
					"Can also be set via `CONDUCTORONE_CLIENT_ID` env var.",
				Optional: true,
			},
			"client_secret": schema.StringAttribute{
				MarkdownDescription: "Client secret for credential-based authentication (Ed25519 JWT assertion). " +
					"Can also be set via `CONDUCTORONE_CLIENT_SECRET` env var.",
				Optional:  true,
				Sensitive: true,
			},
			"tenant_domain": schema.StringAttribute{
				MarkdownDescription: "Tenant domain (derived from `client_id` if not provided). " +
					"Can also be set via `CONDUCTORONE_TENANT_DOMAIN` env var.",
				Optional: true,
			},
			"oidc_token": schema.StringAttribute{
				MarkdownDescription: "OIDC token for workload federation authentication (RFC 8693 token exchange). " +
					"Auto-detected from `CONDUCTORONE_OIDC_TOKEN` or `TFC_WORKLOAD_IDENTITY_TOKEN` environment variables.",
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func (p *ConductoroneProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data ConductoroneProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Resolve each attribute: provider config takes precedence over env var.
	serverURL := data.ServerURL.ValueString()
	if serverURL == "" {
		serverURL = os.Getenv(sdk.EnvServerURL)
	}
	clientID := data.ClientID.ValueString()
	if clientID == "" {
		clientID = os.Getenv(sdk.EnvClientID)
	}
	clientSecret := data.ClientSecret.ValueString()
	if clientSecret == "" {
		clientSecret = os.Getenv(sdk.EnvClientSecret)
	}
	tenantDomain := data.TenantDomain.ValueString()
	if tenantDomain == "" {
		tenantDomain = os.Getenv(sdk.EnvTenantDomain)
	}

	// OIDC token fallback chain: attribute > CONDUCTORONE_OIDC_TOKEN > TFC_WORKLOAD_IDENTITY_TOKEN
	oidcToken := data.OIDCToken.ValueString()
	if oidcToken == "" {
		oidcToken = os.Getenv(sdk.EnvOIDCToken)
	}
	if oidcToken == "" {
		oidcToken = os.Getenv(sdk.EnvTFCWorkloadIdentityToken)
	}

	// Build tenant/server option.
	optStr := ""
	if tenantDomain != "" {
		optStr = tenantDomain
	} else if serverURL != "" {
		optStr = serverURL
	}

	var opts []sdk.CustomSDKOption
	if optStr != "" {
		opt, err := sdk.WithTenantCustom(optStr)
		if err != nil {
			resp.Diagnostics.AddError("Invalid tenant/server configuration", err.Error())
			return
		}
		opts = append(opts, opt)
	}

	// Auth priority (highest wins):
	//   1. CONDUCTORONE_ACCESS_TOKEN env var -- static bearer token
	//   2. oidc_token / CONDUCTORONE_OIDC_TOKEN / TFC_WORKLOAD_IDENTITY_TOKEN -- workload federation
	//   3. client_id + client_secret -- Ed25519 JWT assertion
	var client *sdk.ConductoroneAPI
	var err error

	if accessToken := os.Getenv(sdk.EnvAccessToken); accessToken != "" {
		// Priority 1: static bearer token (escape hatch / CI override)
		client, err = sdk.NewWithAccessToken(ctx, accessToken, clientID, opts...)
	} else if oidcToken != "" {
		// Priority 2: workload federation (RFC 8693 token exchange)
		if clientID == "" {
			resp.Diagnostics.AddError(
				"Missing client_id",
				fmt.Sprintf("client_id is required when using OIDC token authentication (oidc_token, %s, or %s)", sdk.EnvOIDCToken, sdk.EnvTFCWorkloadIdentityToken),
			)
			return
		}
		client, err = sdk.NewWithOIDCToken(ctx, oidcToken, clientID, opts...)
	} else {
		// Priority 3: client credentials (Ed25519 JWT assertion)
		client, err = sdk.NewWithCredentials(ctx, &sdk.ClientCredentials{
			ClientID:     clientID,
			ClientSecret: clientSecret,
		}, opts...)
	}
	if err != nil {
		resp.Diagnostics.AddError("Failed to create SDK Client", err.Error())
		return
	}

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *ConductoroneProvider) Resources(ctx context.Context) []func() resource.Resource {
	resources := []func() resource.Resource{
		NewAccessProfileResource,
		NewAccessProfileRequestableEntriesResource,
		NewAccessProfileRequestableEntryResource,
		NewAccessProfileVisibilityBindingsResource,
		NewAccessReviewResource,
		NewAccessReviewSetupResource,
		NewAccessReviewTemplateResource,
		NewAccessReviewTemplateSetupResource,
		NewAppResource,
		NewAppEntitlementResource,
		NewAppEntitlementProxyBindingResource,
		NewAppResourceResource,
		NewAppResourceTypeResource,
		NewAppEntitlementOwnerResource,
		NewAppOwnerResource,
		NewConnectorCredentialResource,
		NewCustomAppEntitlementResource,
		NewPolicyResource,
		NewWebhookResource,
		NewComplianceFrameworkResource,
		NewRiskLevelResource,
		NewAppEntitlementAutomationResource,
		NewBundleAutomationResource,
		NewAccessConflictResource,
		NewAutomationResource,
		NewAppResourceOwnerResource,
		NewRequestSchemaEntitlementBindingResource,
		NewRequestSchemaResource,
		NewVaultResource,
		NewAppEntitlementMonitorBindingResource,
		NewRequestCatalogResource,
	}

	resources = append(resources, getIntegrationResources()...)

	return resources
}

func (p *ConductoroneProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewAccessReviewSetupDataSource,
		NewAccessReviewTemplateSetupDataSource,
		NewAccessProfileDataSource,
		NewAppDataSource,
		NewAppEntitlementProxyBindingDataSource,
		NewAppResourceDataSource,
		NewAppResourceTypeDataSource,
		NewAwsExternalIDDataSource,
		NewBundleAutomationDataSource,
		NewConnectorCredentialDataSource,
		NewAppEntitlementDataSource,
		NewPolicyDataSource,
		NewWebhookDataSource,
		NewComplianceFrameworkDataSource,
		NewRiskLevelDataSource,
		NewUserDataSource,
		NewAppEntitlementAutomationDataSource,
		NewAppEntitlementsDataSource,
		NewAppResourcesDataSource,
		NewAppsDataSource,
		NewAppResourceTypesDataSource,
		NewPoliciesDataSource,
		NewRiskLevelsDataSource,
		NewUsersDataSource,
		NewRequestCatalogDataSource,
		NewRequestCatalogsDataSource,
		NewComplianceFrameworksDataSource,
		NewWebhooksDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &ConductoroneProvider{
			version: version,
		}
	}
}
