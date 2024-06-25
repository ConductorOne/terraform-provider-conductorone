package provider

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/operations"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"github.com/conductorone/terraform-provider-conductorone/internal/validators"
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func NewWebhookDataSource() datasource.DataSource {
	return &WebhookDataSource{}
}

// WebhookResource defines the data source implementation.
type WebhookDataSource struct {
	client *sdk.ConductoroneAPI
}

// WebhookResourceModel describes the resource data model.
type WebhookDataSourceModel struct {
	CreatedAt   types.String `tfsdk:"created_at"`
	DeletedAt   types.String `tfsdk:"deleted_at"`
	Description types.String `tfsdk:"description"`
	DisplayName types.String `tfsdk:"display_name"`
	ID          types.String `tfsdk:"id"`
	UpdatedAt   types.String `tfsdk:"updated_at"`
	URL         types.String `tfsdk:"url"`
}

func (r *WebhookDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_webhook"
}

func (r *WebhookDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Webhook Data Source",

		Attributes: map[string]schema.Attribute{
			"created_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"deleted_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: `The description field.`,
			},
			"display_name": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The displayName field.`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The id field.`,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"url": schema.StringAttribute{
				Computed:    true,
				Description: `The url field.`,
			},
		},
	}
}

func (r *WebhookDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.ConductoroneAPI)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *WebhookDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *WebhookResourceModel
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

	id := data.ID.ValueString()
	displayName := data.DisplayName.ValueStringPointer()
	if id == "" && (displayName == nil || *displayName == "") {
		resp.Diagnostics.AddError("either id or displayName must be set", "")
		return
	}

	// If we have an ID, we can use the Get API to fetch the attribute
	if id != "" {
		request := operations.C1APIWebhooksV1WebhooksServiceGetRequest{
			ID: id,
		}
		res, err := r.client.Webhooks.Get(ctx, request)
		if err != nil {
			resp.Diagnostics.AddError("failure to invoke API", err.Error())
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
		if res.WebhooksServiceGetResponse == nil {
			resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
			return
		}
		data.RefreshFromSharedWebhook(res.WebhooksServiceGetResponse.Webhook)

		// Save updated data into Terraform state
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		return
	}

	// If we don't have an ID but we do have a value we can use the Search API to fetch the latest data.
	request := &shared.WebhooksSearchRequest{
		DisplayName: displayName,
	}
	res, err := r.client.WebhooksSearch.Search(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
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
	if res.WebhooksSearchResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}

	if len(res.WebhooksSearchResponse.List) == 0 {
		resp.Diagnostics.AddError("unexpected response from API. Webhook was not found", debugResponse(res.RawResponse))
		return
	}

	if len(res.WebhooksSearchResponse.List) > 1 {
		resp.Diagnostics.AddError("unexpected response from API. More than 1 Webhook was found", debugResponse(res.RawResponse))
		return
	}

	data.RefreshFromSharedWebhook(&res.WebhooksSearchResponse.List[0])

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
