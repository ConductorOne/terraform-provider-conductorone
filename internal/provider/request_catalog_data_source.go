package provider

import (
	"context"
	"fmt"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &RequestCatalogDataSource{}
var _ datasource.DataSourceWithConfigure = &RequestCatalogDataSource{}

func NewRequestCatalogDataSource() datasource.DataSource {
	return &RequestCatalogDataSource{}
}

// RequestCatalogDataSource is the data source implementation.
type RequestCatalogDataSource struct {
	// Provider configured SDK client.
	client *sdk.ConductoroneAPI
}

// RequestCatalogDataSourceModel describes the data model.
type RequestCatalogDataSourceModel struct {
	CreatedAt                       types.String `tfsdk:"created_at"`
	CreatedByUserID                 types.String `tfsdk:"created_by_user_id"`
	DeletedAt                       types.String `tfsdk:"deleted_at"`
	Description                     types.String `tfsdk:"description"`
	DisplayName                     types.String `tfsdk:"display_name"`
	EnrollmentBehavior              types.String `tfsdk:"enrollment_behavior"`
	ID                              types.String `tfsdk:"id"`
	Published                       types.Bool   `tfsdk:"published"`
	RequestBundle                   types.Bool   `tfsdk:"request_bundle"`
	UnenrollmentBehavior            types.String `tfsdk:"unenrollment_behavior"`
	UnenrollmentEntitlementBehavior types.String `tfsdk:"unenrollment_entitlement_behavior"`
	UpdatedAt                       types.String `tfsdk:"updated_at"`
	VisibleToEveryone               types.Bool   `tfsdk:"visible_to_everyone"`
}

// Metadata returns the data source type name.
func (r *RequestCatalogDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_request_catalog"
}

// Schema defines the schema for the data source.
func (r *RequestCatalogDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "RequestCatalog DataSource",

		Attributes: map[string]schema.Attribute{
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"created_by_user_id": schema.StringAttribute{
				Computed:    true,
				Description: `The id of the user this request catalog was created by.`,
			},
			"deleted_at": schema.StringAttribute{
				Computed: true,
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: `The description of the request catalog.`,
			},
			"display_name": schema.StringAttribute{
				Computed:    true,
				Description: `The display name of the request catalog.`,
			},
			"enrollment_behavior": schema.StringAttribute{
				Computed:    true,
				Description: `Defines how to handle the request policies of the entitlements in the catalog during enrollment.`,
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `The id of the request catalog.`,
			},
			"published": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether or not this catalog is published.`,
			},
			"request_bundle": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether all the entitlements in the catalog can be requested at once. Your tenant must have the bundles feature to use this.`,
			},
			"unenrollment_behavior": schema.StringAttribute{
				Computed:    true,
				Description: `Defines how to handle the revocation of the entitlements in the catalog during unenrollment.`,
			},
			"unenrollment_entitlement_behavior": schema.StringAttribute{
				Computed:    true,
				Description: `Defines how to handle the revoke policies of the entitlements in the catalog during unenrollment.`,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
			"visible_to_everyone": schema.BoolAttribute{
				Computed:    true,
				Description: `If this is true, the access entitlement requirement is ignored.`,
			},
		},
	}
}

func (r *RequestCatalogDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *RequestCatalogDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *RequestCatalogDataSourceModel
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

	request, requestDiags := data.ToOperationsC1APIRequestcatalogV1RequestCatalogManagementServiceGetRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.RequestCatalogManagement.Get(ctx, *request)
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
	if !(res.RequestCatalogManagementServiceGetResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedRequestCatalogManagementServiceGetResponse(ctx, res.RequestCatalogManagementServiceGetResponse)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
