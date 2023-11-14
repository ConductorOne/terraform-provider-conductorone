package provider

import (
	"conductorone/internal/sdk/pkg/models/shared"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *AppEntitlementDataSourceModel) RefreshFromGetResponse(resp *shared.AppEntitlement) {
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
	r.MaxGrantDuration = &MaxGrantDuration{}
	if resp.DurationGrant != nil {
		r.MaxGrantDuration.DurationGrant = types.StringValue(*resp.DurationGrant)
	}
	if resp.DurationUnset != nil {
		r.MaxGrantDuration.DurationUnset = &AppEntitlementDurationUnset{}
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
	if r.ProvisionPolicy == nil {
		r.ProvisionPolicy = &ProvisionPolicy{}
	}
	if resp.ProvisionPolicy == nil {
		r.ProvisionPolicy = nil
	} else {
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
