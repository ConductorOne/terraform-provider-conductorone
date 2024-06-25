// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"context"
	"fmt"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/operations"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"github.com/conductorone/terraform-provider-conductorone/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func NewRiskLevelDataSource() datasource.DataSource {
	return &RiskLevelDataSource{}
}

// RiskLevel defines the resource implementation.
type RiskLevelDataSource struct {
	client *sdk.ConductoroneAPI
}

// RiskLevelModel describes the resource data model.
type RiskLevelDataSourceModel struct {
	AttributeTypeID types.String `tfsdk:"attribute_type_id"`
	CreatedAt       types.String `tfsdk:"created_at"`
	DeletedAt       types.String `tfsdk:"deleted_at"`
	ID              types.String `tfsdk:"id"`
	UpdatedAt       types.String `tfsdk:"updated_at"`
	Value           types.String `tfsdk:"value"`
}

func (r *RiskLevelDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_risk_level"
}

func (r *RiskLevelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Risk Level Resource",

		Attributes: map[string]schema.Attribute{
			"attribute_type_id": schema.StringAttribute{
				Computed:    true,
				Description: `The attributeTypeId field.`,
			},
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
			"value": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The value field is the value of the attribute. In this case it is the name of the Risk Level (e.g. High, Low, etc.).`,
			},
		},
	}
}

func (r *RiskLevelDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *RiskLevelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *RiskLevelModel
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
	value := data.Value.ValueStringPointer()
	if id == "" && (value == nil || *value == "") {
		resp.Diagnostics.AddError("either id or value must be set", "")
		return
	}

	// If we have an ID, we can use the Get API to fetch the attribute
	if id != "" {
		request := operations.C1APIAttributeV1AttributesGetAttributeValueRequest{
			ID: id,
		}
		res, err := r.client.Attributes.GetAttributeValue(ctx, request)
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
		if res.GetAttributeValueResponse.AttributeValue == nil {
			resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
			return
		}
		data.RefreshFromGetResponse(res.GetAttributeValueResponse.AttributeValue)

		// Save updated data into Terraform state
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		return
	}

	// If we don't have an ID but we do have a value we can use the Search API to fetch the latest data.
	request := shared.SearchAttributeValuesRequest{
		Value: value,
		AttributeTypeIds: []string{
			riskLevelAttributeTypeID,
		},
	}
	res, err := r.client.AttributeSearch.SearchAttributeValues(ctx, request)
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
	if res.SearchAttributeValuesResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}

	if len(res.SearchAttributeValuesResponse.List) == 0 {
		resp.Diagnostics.AddError("unexpected response from API. RiskLevel was not found", debugResponse(res.RawResponse))
		return
	}

	if len(res.SearchAttributeValuesResponse.List) > 1 {
		resp.Diagnostics.AddError("unexpected response from API. More than 1 RiskLevel was found", debugResponse(res.RawResponse))
		return
	}

	data.RefreshFromGetResponse(&res.SearchAttributeValuesResponse.List[0])

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
