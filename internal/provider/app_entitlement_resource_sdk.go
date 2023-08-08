package provider

import (
	"conductorone/internal/sdk/pkg/models/shared"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *AppEntitlementResourceModel) ToUpdateSDKType() *shared.AppEntitlement {
	var provisionPolicy *shared.ProvisionPolicy
	if r.ProvisionPolicy != nil {
		var connectorProvision *shared.ConnectorProvision
		if r.ProvisionPolicy.ConnectorProvision != nil {
			connectorProvision = &shared.ConnectorProvision{}
		}
		var delegatedProvision *shared.DelegatedProvision
		if r.ProvisionPolicy.DelegatedProvision != nil {
			appID := new(string)
			if !r.ProvisionPolicy.DelegatedProvision.AppID.IsUnknown() && !r.ProvisionPolicy.DelegatedProvision.AppID.IsNull() {
				*appID = r.ProvisionPolicy.DelegatedProvision.AppID.ValueString()
			} else {
				appID = nil
			}
			entitlementID := new(string)
			if !r.ProvisionPolicy.DelegatedProvision.EntitlementID.IsUnknown() && !r.ProvisionPolicy.DelegatedProvision.EntitlementID.IsNull() {
				*entitlementID = r.ProvisionPolicy.DelegatedProvision.EntitlementID.ValueString()
			} else {
				entitlementID = nil
			}
			delegatedProvision = &shared.DelegatedProvision{
				AppID:         appID,
				EntitlementID: entitlementID,
			}
		}
		var manualProvision *shared.ManualProvision
		if r.ProvisionPolicy.ManualProvision != nil {
			instructions := new(string)
			if !r.ProvisionPolicy.ManualProvision.Instructions.IsUnknown() && !r.ProvisionPolicy.ManualProvision.Instructions.IsNull() {
				*instructions = r.ProvisionPolicy.ManualProvision.Instructions.ValueString()
			} else {
				instructions = nil
			}
			userIds := make([]string, 0)
			for _, userIdsItem := range r.ProvisionPolicy.ManualProvision.UserIds {
				userIds = append(userIds, userIdsItem.ValueString())
			}
			manualProvision = &shared.ManualProvision{
				Instructions: instructions,
				UserIds:      userIds,
			}
		}
		provisionPolicy = &shared.ProvisionPolicy{
			ConnectorProvision: connectorProvision,
			DelegatedProvision: delegatedProvision,
			ManualProvision:    manualProvision,
		}
	}
	alias := new(string)
	if !r.Alias.IsUnknown() && !r.Alias.IsNull() {
		*alias = r.Alias.ValueString()
	} else {
		alias = nil
	}
	appId1 := new(string)
	if !r.AppID.IsUnknown() && !r.AppID.IsNull() {
		*appId1 = r.AppID.ValueString()
	} else {
		appId1 = nil
	}
	appResourceID := new(string)
	if !r.AppResourceID.IsUnknown() && !r.AppResourceID.IsNull() {
		*appResourceID = r.AppResourceID.ValueString()
	} else {
		appResourceID = nil
	}
	appResourceTypeID := new(string)
	if !r.AppResourceTypeID.IsUnknown() && !r.AppResourceTypeID.IsNull() {
		*appResourceTypeID = r.AppResourceTypeID.ValueString()
	} else {
		appResourceTypeID = nil
	}
	certifyPolicyID := new(string)
	if !r.CertifyPolicyID.IsUnknown() && !r.CertifyPolicyID.IsNull() {
		*certifyPolicyID = r.CertifyPolicyID.ValueString()
	} else {
		certifyPolicyID = nil
	}
	complianceFrameworkValueIds := make([]string, 0)
	for _, complianceFrameworkValueIdsItem := range r.ComplianceFrameworkValueIds {
		complianceFrameworkValueIds = append(complianceFrameworkValueIds, complianceFrameworkValueIdsItem.ValueString())
	}
	createdAt := new(time.Time)
	if !r.CreatedAt.IsUnknown() && !r.CreatedAt.IsNull() {
		*createdAt, _ = time.Parse(time.RFC3339Nano, r.CreatedAt.ValueString())
	} else {
		createdAt = nil
	}
	deletedAt := new(time.Time)
	if !r.DeletedAt.IsUnknown() && !r.DeletedAt.IsNull() {
		*deletedAt, _ = time.Parse(time.RFC3339Nano, r.DeletedAt.ValueString())
	} else {
		deletedAt = nil
	}
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	displayName := new(string)
	if !r.DisplayName.IsUnknown() && !r.DisplayName.IsNull() {
		*displayName = r.DisplayName.ValueString()
	} else {
		displayName = nil
	}
	var durationUnset *shared.AppEntitlementDurationUnset
	durationGrant := new(string)
	if r.MaxGrantDuration != nil {
		if r.MaxGrantDuration.DurationUnset != nil {
			durationUnset = &shared.AppEntitlementDurationUnset{}
		}
		if !r.MaxGrantDuration.DurationGrant.IsUnknown() && !r.MaxGrantDuration.DurationGrant.IsNull() {
			*durationGrant = r.MaxGrantDuration.DurationGrant.ValueString()
		} else {
			durationGrant = nil
		}
	} else {
		durationGrant = nil
		durationUnset = &shared.AppEntitlementDurationUnset{}
	}

	emergencyGrantEnabled := new(bool)
	if !r.EmergencyGrantEnabled.IsUnknown() && !r.EmergencyGrantEnabled.IsNull() {
		*emergencyGrantEnabled = r.EmergencyGrantEnabled.ValueBool()
	} else {
		emergencyGrantEnabled = nil
	}
	emergencyGrantPolicyID := new(string)
	if !r.EmergencyGrantPolicyID.IsUnknown() && !r.EmergencyGrantPolicyID.IsNull() {
		*emergencyGrantPolicyID = r.EmergencyGrantPolicyID.ValueString()
	} else {
		emergencyGrantPolicyID = nil
	}
	grantPolicyID := new(string)
	if !r.GrantPolicyID.IsUnknown() && !r.GrantPolicyID.IsNull() {
		*grantPolicyID = r.GrantPolicyID.ValueString()
	} else {
		grantPolicyID = nil
	}
	id := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id = r.ID.ValueString()
	} else {
		id = nil
	}
	revokePolicyID := new(string)
	if !r.RevokePolicyID.IsUnknown() && !r.RevokePolicyID.IsNull() {
		*revokePolicyID = r.RevokePolicyID.ValueString()
	} else {
		revokePolicyID = nil
	}
	riskLevelValueID := new(string)
	if !r.RiskLevelValueID.IsUnknown() && !r.RiskLevelValueID.IsNull() {
		*riskLevelValueID = r.RiskLevelValueID.ValueString()
	} else {
		riskLevelValueID = nil
	}
	slug := new(string)
	if !r.Slug.IsUnknown() && !r.Slug.IsNull() {
		*slug = r.Slug.ValueString()
	} else {
		slug = nil
	}

	out := shared.AppEntitlement{
		ProvisionPolicy:             provisionPolicy,
		Alias:                       alias,
		AppID:                       appId1,
		AppResourceID:               appResourceID,
		AppResourceTypeID:           appResourceTypeID,
		CertifyPolicyID:             certifyPolicyID,
		ComplianceFrameworkValueIds: complianceFrameworkValueIds,
		CreatedAt:                   createdAt,
		DeletedAt:                   deletedAt,
		Description:                 description,
		DisplayName:                 displayName,
		DurationGrant:               durationGrant,
		DurationUnset:               durationUnset,
		EmergencyGrantEnabled:       emergencyGrantEnabled,
		EmergencyGrantPolicyID:      emergencyGrantPolicyID,
		GrantPolicyID:               grantPolicyID,
		ID:                          id,
		RevokePolicyID:              revokePolicyID,
		RiskLevelValueID:            riskLevelValueID,
		Slug:                        slug,
	}
	return &out
}

func (r *AppEntitlementResourceModel) RefreshFromGetResponse(resp *shared.AppEntitlement) {
	if resp.Alias != nil {
		r.Alias = types.StringValue(*resp.Alias)
	} else {
		r.Alias = types.StringNull()
	}
	if resp.AppID != nil {
		r.AppID = types.StringValue(*resp.AppID)
	} else {
		r.AppID = types.StringNull()
	}
	if resp.AppResourceID != nil {
		r.AppResourceID = types.StringValue(*resp.AppResourceID)
	} else {
		r.AppResourceID = types.StringNull()
	}
	if resp.AppResourceTypeID != nil {
		r.AppResourceTypeID = types.StringValue(*resp.AppResourceTypeID)
	} else {
		r.AppResourceTypeID = types.StringNull()
	}
	if resp.CertifyPolicyID != nil {
		r.CertifyPolicyID = types.StringValue(*resp.CertifyPolicyID)
	} else {
		r.CertifyPolicyID = types.StringNull()
	}
	r.ComplianceFrameworkValueIds = nil
	for _, v := range resp.ComplianceFrameworkValueIds {
		r.ComplianceFrameworkValueIds = append(r.ComplianceFrameworkValueIds, types.StringValue(v))
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
	if resp.DurationGrant != nil || resp.DurationUnset != nil {
		if r.MaxGrantDuration != nil {
			r.MaxGrantDuration = &MaxGrantDuration{}
			if resp.DurationGrant != nil {
				r.MaxGrantDuration.DurationGrant = types.StringValue(*resp.DurationGrant)
			} else {
				r.MaxGrantDuration.DurationGrant = types.StringNull()
			}
			if r.MaxGrantDuration.DurationUnset == nil {
				r.MaxGrantDuration.DurationUnset = &AppEntitlementDurationUnset{}
			}
			if resp.DurationUnset == nil {
				r.MaxGrantDuration.DurationUnset = nil
			} else {
				r.MaxGrantDuration.DurationUnset = &AppEntitlementDurationUnset{}
			}
		}
	}
	if resp.EmergencyGrantEnabled != nil {
		r.EmergencyGrantEnabled = types.BoolValue(*resp.EmergencyGrantEnabled)
	} else {
		r.EmergencyGrantEnabled = types.BoolNull()
	}
	if resp.EmergencyGrantPolicyID != nil {
		r.EmergencyGrantPolicyID = types.StringValue(*resp.EmergencyGrantPolicyID)
	} else {
		r.EmergencyGrantPolicyID = types.StringNull()
	}
	if resp.GrantPolicyID != nil {
		r.GrantPolicyID = types.StringValue(*resp.GrantPolicyID)
	} else {
		r.GrantPolicyID = types.StringNull()
	}
	if resp.ID != nil {
		r.ID = types.StringValue(*resp.ID)
	} else {
		r.ID = types.StringNull()
	}

	if resp.ProvisionPolicy != nil {
		if r.ProvisionPolicy != nil {
			r.ProvisionPolicy = &ProvisionPolicy{}
			if r.ProvisionPolicy.ConnectorProvision == nil {
				r.ProvisionPolicy.ConnectorProvision = &ConnectorProvision{}
			}
			if resp.ProvisionPolicy.ConnectorProvision == nil {
				r.ProvisionPolicy.ConnectorProvision = nil
			} else {
				r.ProvisionPolicy.ConnectorProvision = &ConnectorProvision{}
			}
			if r.ProvisionPolicy.DelegatedProvision == nil {
				r.ProvisionPolicy.DelegatedProvision = &DelegatedProvision{}
			}
			if resp.ProvisionPolicy.DelegatedProvision == nil {
				r.ProvisionPolicy.DelegatedProvision = nil
			} else {
				r.ProvisionPolicy.DelegatedProvision = &DelegatedProvision{}
				if resp.ProvisionPolicy.DelegatedProvision.AppID != nil {
					r.ProvisionPolicy.DelegatedProvision.AppID = types.StringValue(*resp.ProvisionPolicy.DelegatedProvision.AppID)
				} else {
					r.ProvisionPolicy.DelegatedProvision.AppID = types.StringNull()
				}
				if resp.ProvisionPolicy.DelegatedProvision.EntitlementID != nil {
					r.ProvisionPolicy.DelegatedProvision.EntitlementID = types.StringValue(*resp.ProvisionPolicy.DelegatedProvision.EntitlementID)
				} else {
					r.ProvisionPolicy.DelegatedProvision.EntitlementID = types.StringNull()
				}
			}
			if r.ProvisionPolicy.ManualProvision == nil {
				r.ProvisionPolicy.ManualProvision = &ManualProvision{}
			}
			if resp.ProvisionPolicy.ManualProvision == nil {
				r.ProvisionPolicy.ManualProvision = nil
			} else {
				r.ProvisionPolicy.ManualProvision = &ManualProvision{}
				if resp.ProvisionPolicy.ManualProvision.Instructions != nil {
					r.ProvisionPolicy.ManualProvision.Instructions = types.StringValue(*resp.ProvisionPolicy.ManualProvision.Instructions)
				} else {
					r.ProvisionPolicy.ManualProvision.Instructions = types.StringNull()
				}
				r.ProvisionPolicy.ManualProvision.UserIds = nil
				for _, v := range resp.ProvisionPolicy.ManualProvision.UserIds {
					r.ProvisionPolicy.ManualProvision.UserIds = append(r.ProvisionPolicy.ManualProvision.UserIds, types.StringValue(v))
				}
			}
		}
	}
	if resp.RevokePolicyID != nil {
		r.RevokePolicyID = types.StringValue(*resp.RevokePolicyID)
	} else {
		r.RevokePolicyID = types.StringNull()
	}
	if resp.RiskLevelValueID != nil {
		r.RiskLevelValueID = types.StringValue(*resp.RiskLevelValueID)
	} else {
		r.RiskLevelValueID = types.StringNull()
	}
	if resp.Slug != nil {
		r.Slug = types.StringValue(*resp.Slug)
	} else {
		r.Slug = types.StringNull()
	}
	if resp.UpdatedAt != nil {
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339))
	} else {
		r.UpdatedAt = types.StringNull()
	}
}

func (r *AppEntitlementResourceModel) RefreshFromCreateResponse(resp *shared.AppEntitlement) {
	r.RefreshFromGetResponse(resp)
}

func (r *AppEntitlementResourceModel) RefreshFromUpdateResponse(resp *shared.AppEntitlement) {
	r.RefreshFromGetResponse(resp)
}
