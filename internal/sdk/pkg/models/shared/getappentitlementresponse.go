// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/ConductorOne/terraform-provider-conductorone/internal/sdk/pkg/utils"
)

// GetAppEntitlementResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type GetAppEntitlementResponseExpanded struct {
	// The type of the serialized message.
	AtType               *string     `json:"@type,omitempty"`
	AdditionalProperties interface{} `additionalProperties:"true" json:"-"`
}

func (g GetAppEntitlementResponseExpanded) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetAppEntitlementResponseExpanded) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetAppEntitlementResponseExpanded) GetAtType() *string {
	if o == nil {
		return nil
	}
	return o.AtType
}

func (o *GetAppEntitlementResponseExpanded) GetAdditionalProperties() interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

// GetAppEntitlementResponse - The get app entitlement response returns an entitlement view containing paths in the expanded array for the objects expanded as indicated by the expand mask in the request.
type GetAppEntitlementResponse struct {
	// The app entitlement view contains the serialized app entitlement and paths to objects referenced by the app entitlement.
	AppEntitlementView *AppEntitlementView `json:"appEntitlementView,omitempty"`
	// List of serialized related objects.
	Expanded []GetAppEntitlementResponseExpanded `json:"expanded,omitempty"`
}

func (o *GetAppEntitlementResponse) GetAppEntitlementView() *AppEntitlementView {
	if o == nil {
		return nil
	}
	return o.AppEntitlementView
}

func (o *GetAppEntitlementResponse) GetExpanded() []GetAppEntitlementResponseExpanded {
	if o == nil {
		return nil
	}
	return o.Expanded
}
