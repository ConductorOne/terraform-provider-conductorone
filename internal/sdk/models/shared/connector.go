// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/internal/utils"
	"time"
)

// Config - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type Config struct {
	// The type of the serialized message.
	AtType               *string `json:"@type,omitempty"`
	AdditionalProperties any     `additionalProperties:"true" json:"-"`
}

func (c Config) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *Config) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Config) GetAtType() *string {
	if o == nil {
		return nil
	}
	return o.AtType
}

func (o *Config) GetAdditionalProperties() any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

// A Connector is used to sync objects into Apps
type Connector struct {
	// The id of the app the connector is associated with.
	AppID *string `json:"appId,omitempty"`
	// The catalogId describes which catalog entry this connector is an instance of. For example, every Okta connector will have the same catalogId indicating it is an Okta connector.
	CatalogID *string `json:"catalogId,omitempty"`
	// Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
	Config    *Config    `json:"config,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	// The description of the connector.
	Description *string `json:"description,omitempty"`
	// The display name of the connector.
	DisplayName *string `json:"displayName,omitempty"`
	// The downloadUrl for a spreadsheet if the connector was created from uploading a file.
	DownloadURL *string `json:"downloadUrl,omitempty"`
	// The id of the connector.
	ID *string `json:"id,omitempty"`
	// OAuth2AuthorizedAs tracks the user that OAuthed with the connector.
	OAuth2AuthorizedAs *OAuth2AuthorizedAs `json:"oauthAuthorizedAs,omitempty"`
	// List of profile attributes to sync, when set only these attributes will be synced
	ProfileAllowList []string `json:"profileAllowList,omitempty"`
	// List of profile attributes to ignore (not sync), when set other attributes will be synced, but these will not.
	ProfileIgnoreList []string `json:"profileIgnoreList,omitempty"`
	// The status field on the connector is used to track the status of the connectors sync, and when syncing last started, completed, or caused the connector to update.
	ConnectorStatus *ConnectorStatus `json:"status,omitempty"`
	SyncDisabledAt  *time.Time       `json:"syncDisabledAt,omitempty"`
	// The category of the connector sync that was disabled.
	SyncDisabledCategory *string `json:"syncDisabledCategory,omitempty"`
	// The reason the connector sync was disabled.
	SyncDisabledReason *string    `json:"syncDisabledReason,omitempty"`
	UpdatedAt          *time.Time `json:"updatedAt,omitempty"`
	// The userIds field is used to define the integration owners of the connector.
	UserIds []string `json:"userIds,omitempty"`
}

func (c Connector) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *Connector) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Connector) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *Connector) GetCatalogID() *string {
	if o == nil {
		return nil
	}
	return o.CatalogID
}

func (o *Connector) GetConfig() *Config {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *Connector) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Connector) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *Connector) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Connector) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *Connector) GetDownloadURL() *string {
	if o == nil {
		return nil
	}
	return o.DownloadURL
}

func (o *Connector) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Connector) GetOAuth2AuthorizedAs() *OAuth2AuthorizedAs {
	if o == nil {
		return nil
	}
	return o.OAuth2AuthorizedAs
}

func (o *Connector) GetProfileAllowList() []string {
	if o == nil {
		return nil
	}
	return o.ProfileAllowList
}

func (o *Connector) GetProfileIgnoreList() []string {
	if o == nil {
		return nil
	}
	return o.ProfileIgnoreList
}

func (o *Connector) GetConnectorStatus() *ConnectorStatus {
	if o == nil {
		return nil
	}
	return o.ConnectorStatus
}

func (o *Connector) GetSyncDisabledAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.SyncDisabledAt
}

func (o *Connector) GetSyncDisabledCategory() *string {
	if o == nil {
		return nil
	}
	return o.SyncDisabledCategory
}

func (o *Connector) GetSyncDisabledReason() *string {
	if o == nil {
		return nil
	}
	return o.SyncDisabledReason
}

func (o *Connector) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *Connector) GetUserIds() []string {
	if o == nil {
		return nil
	}
	return o.UserIds
}

// ConnectorInput - A Connector is used to sync objects into Apps
type ConnectorInput struct {
	// The id of the app the connector is associated with.
	AppID *string `json:"appId,omitempty"`
	// The catalogId describes which catalog entry this connector is an instance of. For example, every Okta connector will have the same catalogId indicating it is an Okta connector.
	CatalogID *string `json:"catalogId,omitempty"`
	// Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
	Config *Config `json:"config,omitempty"`
	// The description of the connector.
	Description *string `json:"description,omitempty"`
	// The display name of the connector.
	DisplayName *string `json:"displayName,omitempty"`
	// The id of the connector.
	ID *string `json:"id,omitempty"`
	// OAuth2AuthorizedAs tracks the user that OAuthed with the connector.
	OAuth2AuthorizedAs *OAuth2AuthorizedAsInput `json:"oauthAuthorizedAs,omitempty"`
	// List of profile attributes to sync, when set only these attributes will be synced
	ProfileAllowList []string `json:"profileAllowList,omitempty"`
	// List of profile attributes to ignore (not sync), when set other attributes will be synced, but these will not.
	ProfileIgnoreList []string `json:"profileIgnoreList,omitempty"`
	// The status field on the connector is used to track the status of the connectors sync, and when syncing last started, completed, or caused the connector to update.
	ConnectorStatus *ConnectorStatus `json:"status,omitempty"`
	// The category of the connector sync that was disabled.
	SyncDisabledCategory *string `json:"syncDisabledCategory,omitempty"`
	// The reason the connector sync was disabled.
	SyncDisabledReason *string `json:"syncDisabledReason,omitempty"`
	// The userIds field is used to define the integration owners of the connector.
	UserIds []string `json:"userIds,omitempty"`
}

func (o *ConnectorInput) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *ConnectorInput) GetCatalogID() *string {
	if o == nil {
		return nil
	}
	return o.CatalogID
}

func (o *ConnectorInput) GetConfig() *Config {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *ConnectorInput) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ConnectorInput) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *ConnectorInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ConnectorInput) GetOAuth2AuthorizedAs() *OAuth2AuthorizedAsInput {
	if o == nil {
		return nil
	}
	return o.OAuth2AuthorizedAs
}

func (o *ConnectorInput) GetProfileAllowList() []string {
	if o == nil {
		return nil
	}
	return o.ProfileAllowList
}

func (o *ConnectorInput) GetProfileIgnoreList() []string {
	if o == nil {
		return nil
	}
	return o.ProfileIgnoreList
}

func (o *ConnectorInput) GetConnectorStatus() *ConnectorStatus {
	if o == nil {
		return nil
	}
	return o.ConnectorStatus
}

func (o *ConnectorInput) GetSyncDisabledCategory() *string {
	if o == nil {
		return nil
	}
	return o.SyncDisabledCategory
}

func (o *ConnectorInput) GetSyncDisabledReason() *string {
	if o == nil {
		return nil
	}
	return o.SyncDisabledReason
}

func (o *ConnectorInput) GetUserIds() []string {
	if o == nil {
		return nil
	}
	return o.UserIds
}
