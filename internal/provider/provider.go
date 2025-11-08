package provider

import (
	"context"
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
}

func (p *ConductoroneProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "conductorone"
	resp.Version = p.version
}

func (p *ConductoroneProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: `ConductorOne API: The ConductorOne API is a HTTP API for managing ConductorOne resources.`,
		Attributes: map[string]schema.Attribute{
			"server_url": schema.StringAttribute{
				MarkdownDescription: "Server URL (defaults to https://{tenantDomain}.conductor.one)",
				Optional:            true,
			},
			"client_id": schema.StringAttribute{
				MarkdownDescription: "ClientId for Auth",
				Optional:            true,
			},
			"client_secret": schema.StringAttribute{
				MarkdownDescription: "ClientSecret for Auth",
				Optional:            true,
			},
			"tenant_domain": schema.StringAttribute{
				MarkdownDescription: "Tenant Domain (derived from client_id if not provided)",
				Optional:            true,
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

	ServerURL := data.ServerURL.ValueString()
	if ServerURL == "" {
		ServerURL = os.Getenv("CONDUCTORONE_SERVER_URL")
	}
	ClientID := data.ClientID.ValueString()
	if ClientID == "" {
		ClientID = os.Getenv("CONDUCTORONE_CLIENT_ID")
	}
	ClientSecret := data.ClientSecret.ValueString()
	if ClientSecret == "" {
		ClientSecret = os.Getenv("CONDUCTORONE_CLIENT_SECRET")
	}
	TenantDomain := data.TenantDomain.ValueString()
	if TenantDomain == "" {
		TenantDomain = os.Getenv("CONDUCTORONE_TENANT_DOMAIN")
	}

	optStr := ""
	if TenantDomain != "" {
		optStr = TenantDomain
	} else if ServerURL != "" {
		optStr = ServerURL
	}

	var opts []sdk.CustomSDKOption
	if optStr != "" {
		opt, err := sdk.WithTenantCustom(optStr)
		if err != nil {
			return
		}
		opts = append(opts, opt)
	}

	client, err := sdk.NewWithCredentials(ctx, &sdk.ClientCredentials{
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
	}, opts...)
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
	}

	resources = append(resources, getIntegrationResources()...)

	return resources
}

func (p *ConductoroneProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
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
