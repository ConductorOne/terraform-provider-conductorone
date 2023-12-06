// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/utils"
	"time"
)

type DurationUnset struct {
}

// AppEntitlement - The app entitlement represents one permission in a downstream App (SAAS) that can be granted. For example, GitHub Read vs GitHub Write.
//
// This message contains a oneof named max_grant_duration. Only a single field of the following list may be set at a time:
//   - durationUnset
//   - durationGrant
type AppEntitlement struct {
	// The alias of the app entitlement used by Cone. Also exact-match queryable.
	Alias *string `json:"alias,omitempty"`
	// The ID of the app that is associated with the app entitlement.
	AppID *string `json:"appId,omitempty"`
	// The ID of the app resource that is associated with the app entitlement
	AppResourceID *string `json:"appResourceId,omitempty"`
	// The ID of the app resource type that is associated with the app entitlement
	AppResourceTypeID *string `json:"appResourceTypeId,omitempty"`
	// The ID of the policy that will be used for certify tickets related to the app entitlement.
	CertifyPolicyID *string `json:"certifyPolicyId,omitempty"`
	// The IDs of different compliance frameworks associated with this app entitlement ex (SOX, HIPAA, PCI, etc.)
	ComplianceFrameworkValueIds []string   `json:"complianceFrameworkValueIds,omitempty"`
	CreatedAt                   *time.Time `json:"createdAt,omitempty"`
	DeletedAt                   *time.Time `json:"deletedAt,omitempty"`
	// The description of the app entitlement.
	Description *string `json:"description,omitempty"`
	// The display name of the app entitlement.
	DisplayName   *string        `json:"displayName,omitempty"`
	DurationGrant *string        `json:"durationGrant,omitempty"`
	DurationUnset *DurationUnset `json:"durationUnset,omitempty"`
	// This enables tasks to be created in an emergency and use a selected emergency access policy.
	EmergencyGrantEnabled *bool `json:"emergencyGrantEnabled,omitempty"`
	// The ID of the policy that will be used for emergency access grant tasks.
	EmergencyGrantPolicyID *string `json:"emergencyGrantPolicyId,omitempty"`
	// The amount of grants open for this entitlement
	GrantCount *string `json:"grantCount,omitempty"`
	// The ID of the policy that will be used for grant tickets related to the app entitlement.
	GrantPolicyID *string `json:"grantPolicyId,omitempty"`
	// The unique ID for the App Entitlement.
	ID *string `json:"id,omitempty"`
	// ProvisionPolicy is a oneOf that indicates how a provision step should be processed.
	//
	// This message contains a oneof named typ. Only a single field of the following list may be set at a time:
	//   - connector
	//   - manual
	//   - delegated
	//
	ProvisionPolicy *ProvisionPolicy `json:"provisionerPolicy,omitempty"`
	// The ID of the policy that will be used for revoke tickets related to the app entitlement
	RevokePolicyID *string `json:"revokePolicyId,omitempty"`
	// The riskLevelValueId field.
	RiskLevelValueID *string `json:"riskLevelValueId,omitempty"`
	// The slug is displayed as an oval next to the name in the frontend of C1, it tells you what permission the entitlement grants. See https://www.conductorone.com/docs/product/manage-access/entitlements/
	Slug *string `json:"slug,omitempty"`
	// This field indicates if this is a system builtin entitlement.
	SystemBuiltin  *bool      `json:"systemBuiltin,omitempty"`
	UpdatedAt      *time.Time `json:"updatedAt,omitempty"`
	UserEditedMask *string    `json:"userEditedMask,omitempty"`
}

func (a AppEntitlement) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AppEntitlement) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AppEntitlement) GetAlias() *string {
	if o == nil {
		return nil
	}
	return o.Alias
}

func (o *AppEntitlement) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *AppEntitlement) GetAppResourceID() *string {
	if o == nil {
		return nil
	}
	return o.AppResourceID
}

func (o *AppEntitlement) GetAppResourceTypeID() *string {
	if o == nil {
		return nil
	}
	return o.AppResourceTypeID
}

func (o *AppEntitlement) GetCertifyPolicyID() *string {
	if o == nil {
		return nil
	}
	return o.CertifyPolicyID
}

func (o *AppEntitlement) GetComplianceFrameworkValueIds() []string {
	if o == nil {
		return nil
	}
	return o.ComplianceFrameworkValueIds
}

func (o *AppEntitlement) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AppEntitlement) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *AppEntitlement) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AppEntitlement) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *AppEntitlement) GetDurationGrant() *string {
	if o == nil {
		return nil
	}
	return o.DurationGrant
}

func (o *AppEntitlement) GetDurationUnset() *DurationUnset {
	if o == nil {
		return nil
	}
	return o.DurationUnset
}

func (o *AppEntitlement) GetEmergencyGrantEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.EmergencyGrantEnabled
}

func (o *AppEntitlement) GetEmergencyGrantPolicyID() *string {
	if o == nil {
		return nil
	}
	return o.EmergencyGrantPolicyID
}

func (o *AppEntitlement) GetGrantCount() *string {
	if o == nil {
		return nil
	}
	return o.GrantCount
}

func (o *AppEntitlement) GetGrantPolicyID() *string {
	if o == nil {
		return nil
	}
	return o.GrantPolicyID
}

func (o *AppEntitlement) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AppEntitlement) GetProvisionPolicy() *ProvisionPolicy {
	if o == nil {
		return nil
	}
	return o.ProvisionPolicy
}

func (o *AppEntitlement) GetRevokePolicyID() *string {
	if o == nil {
		return nil
	}
	return o.RevokePolicyID
}

func (o *AppEntitlement) GetRiskLevelValueID() *string {
	if o == nil {
		return nil
	}
	return o.RiskLevelValueID
}

func (o *AppEntitlement) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *AppEntitlement) GetSystemBuiltin() *bool {
	if o == nil {
		return nil
	}
	return o.SystemBuiltin
}

func (o *AppEntitlement) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *AppEntitlement) GetUserEditedMask() *string {
	if o == nil {
		return nil
	}
	return o.UserEditedMask
}

// AppEntitlementInput - The app entitlement represents one permission in a downstream App (SAAS) that can be granted. For example, GitHub Read vs GitHub Write.
//
// This message contains a oneof named max_grant_duration. Only a single field of the following list may be set at a time:
//   - durationUnset
//   - durationGrant
type AppEntitlementInput struct {
	// The ID of the app that is associated with the app entitlement.
	AppID *string `json:"appId,omitempty"`
	// The ID of the app resource that is associated with the app entitlement
	AppResourceID *string `json:"appResourceId,omitempty"`
	// The ID of the app resource type that is associated with the app entitlement
	AppResourceTypeID *string `json:"appResourceTypeId,omitempty"`
	// The ID of the policy that will be used for certify tickets related to the app entitlement.
	CertifyPolicyID *string `json:"certifyPolicyId,omitempty"`
	// The IDs of different compliance frameworks associated with this app entitlement ex (SOX, HIPAA, PCI, etc.)
	ComplianceFrameworkValueIds []string `json:"complianceFrameworkValueIds,omitempty"`
	// The description of the app entitlement.
	Description *string `json:"description,omitempty"`
	// The display name of the app entitlement.
	DisplayName   *string        `json:"displayName,omitempty"`
	DurationGrant *string        `json:"durationGrant,omitempty"`
	DurationUnset *DurationUnset `json:"durationUnset,omitempty"`
	// This enables tasks to be created in an emergency and use a selected emergency access policy.
	EmergencyGrantEnabled *bool `json:"emergencyGrantEnabled,omitempty"`
	// The ID of the policy that will be used for emergency access grant tasks.
	EmergencyGrantPolicyID *string `json:"emergencyGrantPolicyId,omitempty"`
	// The ID of the policy that will be used for grant tickets related to the app entitlement.
	GrantPolicyID *string `json:"grantPolicyId,omitempty"`
	// ProvisionPolicy is a oneOf that indicates how a provision step should be processed.
	//
	// This message contains a oneof named typ. Only a single field of the following list may be set at a time:
	//   - connector
	//   - manual
	//   - delegated
	//
	ProvisionPolicy *ProvisionPolicy `json:"provisionerPolicy,omitempty"`
	// The ID of the policy that will be used for revoke tickets related to the app entitlement
	RevokePolicyID *string `json:"revokePolicyId,omitempty"`
	// The riskLevelValueId field.
	RiskLevelValueID *string `json:"riskLevelValueId,omitempty"`
	// The slug is displayed as an oval next to the name in the frontend of C1, it tells you what permission the entitlement grants. See https://www.conductorone.com/docs/product/manage-access/entitlements/
	Slug           *string `json:"slug,omitempty"`
	UserEditedMask *string `json:"userEditedMask,omitempty"`
}

func (o *AppEntitlementInput) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *AppEntitlementInput) GetAppResourceID() *string {
	if o == nil {
		return nil
	}
	return o.AppResourceID
}

func (o *AppEntitlementInput) GetAppResourceTypeID() *string {
	if o == nil {
		return nil
	}
	return o.AppResourceTypeID
}

func (o *AppEntitlementInput) GetCertifyPolicyID() *string {
	if o == nil {
		return nil
	}
	return o.CertifyPolicyID
}

func (o *AppEntitlementInput) GetComplianceFrameworkValueIds() []string {
	if o == nil {
		return nil
	}
	return o.ComplianceFrameworkValueIds
}

func (o *AppEntitlementInput) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AppEntitlementInput) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *AppEntitlementInput) GetDurationGrant() *string {
	if o == nil {
		return nil
	}
	return o.DurationGrant
}

func (o *AppEntitlementInput) GetDurationUnset() *DurationUnset {
	if o == nil {
		return nil
	}
	return o.DurationUnset
}

func (o *AppEntitlementInput) GetEmergencyGrantEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.EmergencyGrantEnabled
}

func (o *AppEntitlementInput) GetEmergencyGrantPolicyID() *string {
	if o == nil {
		return nil
	}
	return o.EmergencyGrantPolicyID
}

func (o *AppEntitlementInput) GetGrantPolicyID() *string {
	if o == nil {
		return nil
	}
	return o.GrantPolicyID
}

func (o *AppEntitlementInput) GetProvisionPolicy() *ProvisionPolicy {
	if o == nil {
		return nil
	}
	return o.ProvisionPolicy
}

func (o *AppEntitlementInput) GetRevokePolicyID() *string {
	if o == nil {
		return nil
	}
	return o.RevokePolicyID
}

func (o *AppEntitlementInput) GetRiskLevelValueID() *string {
	if o == nil {
		return nil
	}
	return o.RiskLevelValueID
}

func (o *AppEntitlementInput) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *AppEntitlementInput) GetUserEditedMask() *string {
	if o == nil {
		return nil
	}
	return o.UserEditedMask
}
