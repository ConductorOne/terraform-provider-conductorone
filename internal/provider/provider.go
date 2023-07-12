// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"conductorone/internal/sdk"
	"context"

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
	ServerURL types.String `tfsdk:"server_url"`
	ClientID     types.String `tfsdk:"client_id"`
	ClientSecret types.String `tfsdk:"client_secret"`
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
				Required:            false,
			},
			"client_id": schema.StringAttribute{
				MarkdownDescription: "ClientId for Auth",
				Optional:            false,
				Required:            true,
			},
			"client_secret": schema.StringAttribute{
				MarkdownDescription: "ClientSecret for Auth",
				Optional:            false,
				Required:            true,
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
	ClientID := data.ClientID.ValueString()
	ClientSecret := data.ClientSecret.ValueString()

	if ServerURL == "" {
		ServerURL = "https://{tenantDomain}.conductor.one"
	}

	opt, err := sdk.WithTenantCustom(ServerURL)
	if err != nil {
		return
	}

	opts := []sdk.CustomSDKOption{
		opt,
	}
	client, err := sdk.NewWithCredentials(ctx, &sdk.ClientCredentials{
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
	}, opts...)
	if err != nil {
		panic(err)
	}

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *ConductoroneProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewPolicyResource,
	}
}

func (p *ConductoroneProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &ConductoroneProvider{
			version: version,
		}
	}
}
