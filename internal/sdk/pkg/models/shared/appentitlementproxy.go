// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/utils"
	"time"
)

// The AppEntitlementProxy message.
//
// This message contains a oneof named _implicit. Only a single field of the following list may be set at a time:
//   - implicit
type AppEntitlementProxy struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	// The dstAppEntitlementId field.
	DstAppEntitlementID *string `json:"dstAppEntitlementId,omitempty"`
	// The dstAppId field.
	DstAppID *string `json:"dstAppId,omitempty"`
	// If true, the binding doesn't not exist yet and is from the list of the entitlements from the parent app.
	//  typically the IdP that handles provisioning for the app instead of C1s connector.
	// This field is part of the `_implicit` oneof.
	// See the documentation for `c1.api.app.v1.AppEntitlementProxy` for more details.
	Implicit *bool `json:"implicit,omitempty"`
	// The srcAppEntitlementId field.
	SrcAppEntitlementID *string `json:"srcAppEntitlementId,omitempty"`
	// The srcAppId field.
	SrcAppID *string `json:"srcAppId,omitempty"`
	// The systemBuiltin field.
	SystemBuiltin *bool      `json:"systemBuiltin,omitempty"`
	UpdatedAt     *time.Time `json:"updatedAt,omitempty"`
}

func (a AppEntitlementProxy) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AppEntitlementProxy) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AppEntitlementProxy) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AppEntitlementProxy) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *AppEntitlementProxy) GetDstAppEntitlementID() *string {
	if o == nil {
		return nil
	}
	return o.DstAppEntitlementID
}

func (o *AppEntitlementProxy) GetDstAppID() *string {
	if o == nil {
		return nil
	}
	return o.DstAppID
}

func (o *AppEntitlementProxy) GetImplicit() *bool {
	if o == nil {
		return nil
	}
	return o.Implicit
}

func (o *AppEntitlementProxy) GetSrcAppEntitlementID() *string {
	if o == nil {
		return nil
	}
	return o.SrcAppEntitlementID
}

func (o *AppEntitlementProxy) GetSrcAppID() *string {
	if o == nil {
		return nil
	}
	return o.SrcAppID
}

func (o *AppEntitlementProxy) GetSystemBuiltin() *bool {
	if o == nil {
		return nil
	}
	return o.SystemBuiltin
}

func (o *AppEntitlementProxy) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
