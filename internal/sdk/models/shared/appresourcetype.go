// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/internal/utils"
	"time"
)

// The AppResourceType is referenced by an app entitlement defining its resource types. Commonly things like Group or Role.
type AppResourceType struct {
	// The ID of the app that is associated with the app resource type
	AppID     *string    `json:"appId,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	// The display name of the app resource type.
	DisplayName *string `json:"displayName,omitempty"`
	// The unique ID for the app resource type.
	ID *string `json:"id,omitempty"`
	// Associated trait ids
	TraitIds  []string   `json:"traitIds,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

func (a AppResourceType) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AppResourceType) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AppResourceType) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *AppResourceType) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AppResourceType) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *AppResourceType) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *AppResourceType) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AppResourceType) GetTraitIds() []string {
	if o == nil {
		return nil
	}
	return o.TraitIds
}

func (o *AppResourceType) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
