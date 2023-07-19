// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// ConnectorConfig - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type ConnectorConfig struct {
	// The type of the serialized message.
	AtType *string `json:"@type,omitempty"`

	AdditionalProperties interface{} `json:"-"`
}
type _ConnectorConfig ConnectorConfig

func (c *ConnectorConfig) UnmarshalJSON(bs []byte) error {
	data := _ConnectorConfig{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = ConnectorConfig(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "@type")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c ConnectorConfig) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_ConnectorConfig(c))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	bs, err = json.Marshal(c.AdditionalProperties)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	o, err := json.Marshal(out)
	if err != nil {
		return nil, err
	}
	fmt.Println("********* CONFIG*******", string(o))

	return json.Marshal(out)
}

// Connector - The Connector message.
type Connector struct {
	// The ConnectorStatus message.
	ConnectorStatus *ConnectorStatus `json:"status,omitempty"`
	// The OAuth2AuthorizedAs message.
	OAuth2AuthorizedAs *OAuth2AuthorizedAs `json:"oauthAuthorizedAs,omitempty"`
	// The appId field.
	AppID *string `json:"appId,omitempty"`
	// The catalogId field.
	CatalogID *string `json:"catalogId,omitempty"`
	// Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
	Config    *ConnectorConfig `json:"config,omitempty"`
	CreatedAt *time.Time       `json:"createdAt,omitempty"`
	DeletedAt *time.Time       `json:"deletedAt,omitempty"`
	// The description field.
	Description *string `json:"description,omitempty"`
	// The displayName field.
	DisplayName *string `json:"displayName,omitempty"`
	// The downloadUrl field.
	DownloadURL *string `json:"downloadUrl,omitempty"`
	// The id field.
	ID        *string    `json:"id,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// The userIds field.
	UserIds []string `json:"userIds,omitempty"`
}