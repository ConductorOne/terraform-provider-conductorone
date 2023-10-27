package provider

import (
	"conductorone/internal/sdk"
	"conductorone/internal/sdk/pkg/models/operations"
	"conductorone/internal/sdk/pkg/models/shared"
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppOwnerResource{}
var _ resource.ResourceWithImportState = &AppOwnerResource{}

func NewAppOwnerResource() resource.Resource {
	return &AppOwnerResource{}
}

// AppOwnerResource defines the resource implementation.
type AppOwnerResource struct {
	client *sdk.ConductoroneAPI
}

// AppOwnerResourceModel describes the resource data model.
type AppOwnerResourceModel struct {
	AppID   types.String   `tfsdk:"app_id"`
	UserIds []types.String `tfsdk:"user_ids"`
}

func (r *AppOwnerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_app_owner"
}

func (r *AppOwnerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "AppOwner Resource",

		Attributes: map[string]schema.Attribute{
			"app_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Required: true,
			},
			"user_ids": schema.ListAttribute{
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				Required:    true,
				ElementType: types.StringType,
				Description: `The user_ids field for the users to set as an owner of the app.`,
			},
		},
	}
}

func (r *AppOwnerResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *AppOwnerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *AppOwnerResourceModel
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

	setAppOwnersRequest := data.ToCreateSDKType()
	appID := data.AppID.ValueString()
	request := operations.C1APIAppV1AppOwnersSetRequest{
		SetAppOwnersRequest: setAppOwnersRequest,
		AppID:               appID,
	}
	res, err := r.client.AppOwners.C1APIAppV1AppOwnersSet(ctx, request)
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
	if res.SetAppOwnersResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.SetAppOwnersResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppOwnerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *AppOwnerResourceModel
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

	var owners []shared.User
	pageToken := ""
	for {
		listResp, err := r.client.AppOwners.C1APIAppV1AppOwnersList(ctx, operations.C1APIAppV1AppOwnersListRequest{
			AppID:     data.AppID.ValueString(),
			PageToken: &pageToken,
		})
		if err != nil {
			resp.Diagnostics.AddError("failure to invoke list API", err.Error())
			return
		}
		if listResp == nil {
			resp.Diagnostics.AddError("unexpected response from list API", fmt.Sprintf("%v", listResp))
			return
		}
		if listResp.ListAppOwnersResponse == nil || listResp.ListAppOwnersResponse.List == nil {
			resp.Diagnostics.AddError("unexpected response from list API", fmt.Sprintf("%v", listResp))
			return
		}

		owners = append(owners, listResp.ListAppOwnersResponse.List...)

		if listResp.ListAppOwnersResponse.NextPageToken == nil || *listResp.ListAppOwnersResponse.NextPageToken == "" {
			break
		}
		pageToken = *listResp.ListAppOwnersResponse.NextPageToken
	}

	data.RefreshFromReadResponse(owners)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppOwnerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *AppOwnerResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; all attributes marked as RequiresReplace

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppOwnerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *AppOwnerResourceModel
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

	// Not Implemented; entity does not have a configured DELETE operation
}

func (r *AppOwnerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource app_owner.")
}
