// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	tfTypes "github.com/conductorone/terraform-provider-conductorone/internal/provider/types"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &AppEntitlementProxyBindingDataSource{}
var _ datasource.DataSourceWithConfigure = &AppEntitlementProxyBindingDataSource{}

func NewAppEntitlementProxyBindingDataSource() datasource.DataSource {
	return &AppEntitlementProxyBindingDataSource{}
}

// AppEntitlementProxyBindingDataSource is the data source implementation.
type AppEntitlementProxyBindingDataSource struct {
	client *sdk.ConductoroneAPI
}

// AppEntitlementProxyBindingDataSourceModel describes the data model.
type AppEntitlementProxyBindingDataSourceModel struct {
	CreatedAt           types.String                                     `tfsdk:"created_at"`
	DeletedAt           types.String                                     `tfsdk:"deleted_at"`
	DstAppEntitlementID types.String                                     `tfsdk:"dst_app_entitlement_id"`
	DstAppID            types.String                                     `tfsdk:"dst_app_id"`
	Expanded            []tfTypes.GetAppEntitlementProxyResponseExpanded `tfsdk:"expanded"`
	SrcAppEntitlementID types.String                                     `tfsdk:"src_app_entitlement_id"`
	SrcAppID            types.String                                     `tfsdk:"src_app_id"`
	SystemBuiltin       types.Bool                                       `tfsdk:"system_builtin"`
	UpdatedAt           types.String                                     `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *AppEntitlementProxyBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_app_entitlement_proxy_binding"
}

// Schema defines the schema for the data source.
func (r *AppEntitlementProxyBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "AppEntitlementProxyBinding DataSource",

		Attributes: map[string]schema.Attribute{
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"deleted_at": schema.StringAttribute{
				Computed: true,
			},
			"dst_app_entitlement_id": schema.StringAttribute{
				Required: true,
			},
			"dst_app_id": schema.StringAttribute{
				Required: true,
			},
			"expanded": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{},
				},
				Description: `The expanded field.`,
			},
			"src_app_entitlement_id": schema.StringAttribute{
				Required: true,
			},
			"src_app_id": schema.StringAttribute{
				Required: true,
			},
			"system_builtin": schema.BoolAttribute{
				Computed:    true,
				Description: `The systemBuiltin field.`,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *AppEntitlementProxyBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.ConductoroneAPI)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.ConductoroneAPI, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *AppEntitlementProxyBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *AppEntitlementProxyBindingDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request, requestDiags := data.ToOperationsC1APIAppV1AppEntitlementsProxyGetRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.AppEntitlementsProxy.Get(ctx, *request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.GetAppEntitlementProxyResponse != nil && res.GetAppEntitlementProxyResponse.AppEntitlementProxyView != nil && res.GetAppEntitlementProxyResponse.AppEntitlementProxyView.AppEntitlementProxy != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedAppEntitlementProxy(ctx, res.GetAppEntitlementProxyResponse.AppEntitlementProxyView.AppEntitlementProxy)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
