package provider

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"math/big"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *AppDataSourceModel) RefreshFromGetResponse(resp *shared.App) {
	if resp.AppAccountID != nil {
		r.AppAccountID = types.StringValue(*resp.AppAccountID)
	} else {
		r.AppAccountID = types.StringNull()
	}
	if resp.AppAccountName != nil {
		r.AppAccountName = types.StringValue(*resp.AppAccountName)
	} else {
		r.AppAccountName = types.StringNull()
	}
	if resp.CertifyPolicyID != nil {
		r.CertifyPolicyID = types.StringValue(*resp.CertifyPolicyID)
	} else {
		r.CertifyPolicyID = types.StringNull()
	}
	if resp.CreatedAt != nil {
		r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339))
	} else {
		r.CreatedAt = types.StringNull()
	}
	if resp.DeletedAt != nil {
		r.DeletedAt = types.StringValue(resp.DeletedAt.Format(time.RFC3339))
	} else {
		r.DeletedAt = types.StringNull()
	}
	if resp.Description != nil {
		r.Description = types.StringValue(*resp.Description)
	} else {
		r.Description = types.StringNull()
	}
	if resp.DisplayName != nil {
		r.DisplayName = types.StringValue(*resp.DisplayName)
	} else {
		r.DisplayName = types.StringNull()
	}
	if resp.FieldMask != nil {
		r.FieldMask = types.StringValue(*resp.FieldMask)
	} else {
		r.FieldMask = types.StringNull()
	}
	if resp.GrantPolicyID != nil {
		r.GrantPolicyID = types.StringValue(*resp.GrantPolicyID)
	} else {
		r.GrantPolicyID = types.StringNull()
	}
	if resp.IconURL != nil {
		r.IconURL = types.StringValue(*resp.IconURL)
	} else {
		r.IconURL = types.StringNull()
	}
	if resp.ID != nil {
		r.ID = types.StringValue(*resp.ID)
	} else {
		r.ID = types.StringNull()
	}
	if resp.LogoURI != nil {
		r.LogoURI = types.StringValue(*resp.LogoURI)
	} else {
		r.LogoURI = types.StringNull()
	}
	if resp.MonthlyCostUsd != nil {
		r.MonthlyCostUsd = types.NumberValue(big.NewFloat(*resp.MonthlyCostUsd))
	} else {
		r.MonthlyCostUsd = types.NumberNull()
	}
	if resp.ParentAppID != nil {
		r.ParentAppID = types.StringValue(*resp.ParentAppID)
	} else {
		r.ParentAppID = types.StringNull()
	}
	if resp.RevokePolicyID != nil {
		r.RevokePolicyID = types.StringValue(*resp.RevokePolicyID)
	} else {
		r.RevokePolicyID = types.StringNull()
	}
	if resp.UpdatedAt != nil {
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339))
	} else {
		r.UpdatedAt = types.StringNull()
	}
	if resp.UserCount != nil {
		r.UserCount = types.StringValue(*resp.UserCount)
	} else {
		r.UserCount = types.StringNull()
	}
}
