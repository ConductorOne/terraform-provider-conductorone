// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/operations"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/speakeasy/terraform-provider-terraform/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ConnectorCredentialResource{}
var _ resource.ResourceWithImportState = &ConnectorCredentialResource{}

func NewConnectorCredentialResource() resource.Resource {
	return &ConnectorCredentialResource{}
}

// ConnectorCredentialResource defines the resource implementation.
type ConnectorCredentialResource struct {
	client *sdk.SDK
}

// ConnectorCredentialResourceModel describes the resource data model.
type ConnectorCredentialResourceModel struct {
	AppID                                   types.String   `tfsdk:"app_id"`
	ClientID                                types.String   `tfsdk:"client_id"`
	ClientSecret                            types.String   `tfsdk:"client_secret"`
	ConnectorID                             types.String   `tfsdk:"connector_id"`
	ConnectorServiceRotateCredentialRequest *DurationUnset `tfsdk:"connector_service_rotate_credential_request"`
	CreatedAt                               types.String   `tfsdk:"created_at"`
	DeletedAt                               types.String   `tfsdk:"deleted_at"`
	DisplayName                             types.String   `tfsdk:"display_name"`
	ExpiresTime                             types.String   `tfsdk:"expires_time"`
	ID                                      types.String   `tfsdk:"id"`
	LastUsedAt                              types.String   `tfsdk:"last_used_at"`
	UpdatedAt                               types.String   `tfsdk:"updated_at"`
}

func (r *ConnectorCredentialResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_connector_credential"
}

func (r *ConnectorCredentialResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "ConnectorCredential Resource",

		Attributes: map[string]schema.Attribute{
			"app_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Optional:    true,
				Description: `The appId of the app the connector is attached to.`,
			},
			"client_id": schema.StringAttribute{
				Computed:    true,
				Description: `The client id of the ConnectorCredential.`,
			},
			"client_secret": schema.StringAttribute{
				Computed:    true,
				Description: `The new clientSecret returned after rotating the connector credential.`,
			},
			"connector_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Optional:    true,
				Description: `The connectorId of the connector the credential is associated with.`,
			},
			"connector_service_rotate_credential_request": schema.SingleNestedAttribute{
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplace(),
				},
				Optional:    true,
				Attributes:  map[string]schema.Attribute{},
				Description: `ConnectorServiceRotateCredentialRequest is a request for rotating connector credentials. It uses URL values for input.`,
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
			"display_name": schema.StringAttribute{
				Computed:    true,
				Description: `The display name of the ConnectorCredential.`,
			},
			"expires_time": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `The id of the ConnectorCredential.`,
			},
			"last_used_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
		},
	}
}

func (r *ConnectorCredentialResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *ConnectorCredentialResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ConnectorCredentialResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
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

	appID := data.AppID.ValueString()
	connectorID := data.ConnectorID.ValueString()
	connectorServiceRotateCredentialRequest := data.ToCreateSDKType()
	request := operations.C1APIAppV1ConnectorServiceRotateCredentialRequest{
		AppID:                                   appID,
		ConnectorID:                             connectorID,
		ConnectorServiceRotateCredentialRequest: connectorServiceRotateCredentialRequest,
	}
	res, err := r.client.Connector.RotateCredential(ctx, request)
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
	if res.ConnectorServiceRotateCredentialResponse == nil || res.ConnectorServiceRotateCredentialResponse.ConnectorCredential == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.ConnectorServiceRotateCredentialResponse.ConnectorCredential)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ConnectorCredentialResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ConnectorCredentialResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	appID := data.AppID.ValueString()
	connectorID := data.ConnectorID.ValueString()
	id := data.ID.ValueString()
	request := operations.C1APIAppV1ConnectorServiceGetCredentialsRequest{
		AppID:       appID,
		ConnectorID: connectorID,
		ID:          id,
	}
	res, err := r.client.Connector.GetCredentials(ctx, request)
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
	if res.ConnectorServiceGetCredentialsResponse == nil || res.ConnectorServiceGetCredentialsResponse.ConnectorCredential == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.ConnectorServiceGetCredentialsResponse.ConnectorCredential)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ConnectorCredentialResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ConnectorCredentialResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; all attributes marked as RequiresReplace

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ConnectorCredentialResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ConnectorCredentialResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	appID := data.AppID.ValueString()
	connectorID := data.ConnectorID.ValueString()
	id := data.ID.ValueString()
	connectorServiceRevokeCredentialRequest := data.ToDeleteSDKType()
	request := operations.C1APIAppV1ConnectorServiceRevokeCredentialRequest{
		AppID:                                   appID,
		ConnectorID:                             connectorID,
		ID:                                      id,
		ConnectorServiceRevokeCredentialRequest: connectorServiceRevokeCredentialRequest,
	}
	res, err := r.client.Connector.RevokeCredential(ctx, request)
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

}

func (r *ConnectorCredentialResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource connector_credential. Reason: composite imports strings not supported.")
}
