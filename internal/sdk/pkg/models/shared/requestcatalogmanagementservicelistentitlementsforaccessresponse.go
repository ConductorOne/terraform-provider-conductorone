// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
)

// RequestCatalogManagementServiceListEntitlementsForAccessResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type RequestCatalogManagementServiceListEntitlementsForAccessResponseExpanded struct {
	// The type of the serialized message.
	AtType *string `json:"@type,omitempty"`

	AdditionalProperties interface{} `json:"-"`
}
type _RequestCatalogManagementServiceListEntitlementsForAccessResponseExpanded RequestCatalogManagementServiceListEntitlementsForAccessResponseExpanded

func (c *RequestCatalogManagementServiceListEntitlementsForAccessResponseExpanded) UnmarshalJSON(bs []byte) error {
	data := _RequestCatalogManagementServiceListEntitlementsForAccessResponseExpanded{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = RequestCatalogManagementServiceListEntitlementsForAccessResponseExpanded(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "@type")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c RequestCatalogManagementServiceListEntitlementsForAccessResponseExpanded) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_RequestCatalogManagementServiceListEntitlementsForAccessResponseExpanded(c))
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

	return json.Marshal(out)
}

// RequestCatalogManagementServiceListEntitlementsForAccessResponse - The RequestCatalogManagementServiceListEntitlementsForAccessResponse message.
type RequestCatalogManagementServiceListEntitlementsForAccessResponse struct {
	// The expanded field.
	Expanded []RequestCatalogManagementServiceListEntitlementsForAccessResponseExpanded `json:"expanded,omitempty"`
	// The list field.
	List []AppEntitlementView `json:"list,omitempty"`
	// The nextPageToken field.
	NextPageToken *string `json:"nextPageToken,omitempty"`
}