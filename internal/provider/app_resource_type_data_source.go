// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/operations"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func NewAppResourceTypeDataSource() datasource.DataSource {
	return &AppResourceTypeDataSource{}
}

// AppResourceTypeDataSource is the data source implementation.
type AppResourceTypeDataSource struct {
	client *sdk.ConductoroneAPI
}

// AppResourceTypeDataSourceModel describes the data model.
type AppResourceTypeDataSourceModel struct {
	AppID       types.String `tfsdk:"app_id"`
	CreatedAt   types.String `tfsdk:"created_at"`
	DeletedAt   types.String `tfsdk:"deleted_at"`
	DisplayName types.String `tfsdk:"display_name"`
	ID          types.String `tfsdk:"id"`
	UpdatedAt   types.String `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *AppResourceTypeDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_app_resource_type"
}

// Schema defines the schema for the data source.
func (r *AppResourceTypeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "AppResourceType DataSource",

		Attributes: map[string]schema.Attribute{
			"app_id": schema.StringAttribute{
				Required: true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"deleted_at": schema.StringAttribute{
				Computed: true,
			},
			"display_name": schema.StringAttribute{
				Computed:    true,
				Description: `The display name of the app resource type.`,
			},
			"id": schema.StringAttribute{
				Required: true,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *AppResourceTypeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.ConductoroneAPI)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.Conductorone, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *AppResourceTypeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *AppResourceTypeDataSourceModel
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

	var appID string
	appID = data.AppID.ValueString()

	var id string
	id = data.ID.ValueString()

	request := operations.C1APIAppV1AppResourceTypeServiceGetRequest{
		AppID: appID,
		ID:    id,
	}
	res, err := r.client.AppResourceType.Get(ctx, request)
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
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.AppResourceTypeServiceGetResponse != nil && res.AppResourceTypeServiceGetResponse.AppResourceTypeView != nil && res.AppResourceTypeServiceGetResponse.AppResourceTypeView.AppResourceType != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAppResourceType(res.AppResourceTypeServiceGetResponse.AppResourceTypeView.AppResourceType)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}