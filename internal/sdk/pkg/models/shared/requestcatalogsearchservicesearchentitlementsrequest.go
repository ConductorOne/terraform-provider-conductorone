// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatus - The grantedStatus field.
type RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatus string

const (
	RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatusUnspecified RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatus = "UNSPECIFIED"
	RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatusAll         RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatus = "ALL"
	RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatusGranted     RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatus = "GRANTED"
	RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatusNotGranted  RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatus = "NOT_GRANTED"
)

func (e RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatus) ToPointer() *RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatus {
	return &e
}

func (e *RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "UNSPECIFIED":
		fallthrough
	case "ALL":
		fallthrough
	case "GRANTED":
		fallthrough
	case "NOT_GRANTED":
		*e = RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatus: %v", v)
	}
}

// RequestCatalogSearchServiceSearchEntitlementsRequest - The RequestCatalogSearchServiceSearchEntitlementsRequest message.
type RequestCatalogSearchServiceSearchEntitlementsRequest struct {
	// The AppEntitlementExpandMask message.
	AppEntitlementExpandMask *AppEntitlementExpandMask `json:"expandMask,omitempty"`
	// The entitlementAlias field.
	EntitlementAlias *string `json:"entitlementAlias,omitempty"`
	// The grantedStatus field.
	GrantedStatus *RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatus `json:"grantedStatus,omitempty"`
	// The pageSize field.
	PageSize *float64 `json:"pageSize,omitempty"`
	// The pageToken field.
	PageToken *string `json:"pageToken,omitempty"`
	// The query field.
	Query *string `json:"query,omitempty"`
}