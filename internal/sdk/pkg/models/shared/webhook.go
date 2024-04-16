// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// The Webhook message.
type Webhook struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	// The description field.
	Description *string `json:"description,omitempty"`
	// The displayName field.
	DisplayName *string `json:"displayName,omitempty"`
	// The id field.
	ID        *string    `json:"id,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// The url field.
	URL *string `json:"url,omitempty"`
}

func (o *Webhook) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Webhook) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *Webhook) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Webhook) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *Webhook) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Webhook) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *Webhook) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}