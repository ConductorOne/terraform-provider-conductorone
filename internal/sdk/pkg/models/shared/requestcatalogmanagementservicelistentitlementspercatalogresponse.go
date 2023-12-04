// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/ConductorOne/terraform-provider-conductorone/internal/sdk/pkg/utils"
)

// RequestCatalogManagementServiceListEntitlementsPerCatalogResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type RequestCatalogManagementServiceListEntitlementsPerCatalogResponseExpanded struct {
	// The type of the serialized message.
	AtType               *string     `json:"@type,omitempty"`
	AdditionalProperties interface{} `additionalProperties:"true" json:"-"`
}

func (r RequestCatalogManagementServiceListEntitlementsPerCatalogResponseExpanded) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RequestCatalogManagementServiceListEntitlementsPerCatalogResponseExpanded) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RequestCatalogManagementServiceListEntitlementsPerCatalogResponseExpanded) GetAtType() *string {
	if o == nil {
		return nil
	}
	return o.AtType
}

func (o *RequestCatalogManagementServiceListEntitlementsPerCatalogResponseExpanded) GetAdditionalProperties() interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

// The RequestCatalogManagementServiceListEntitlementsPerCatalogResponse message contains a list of results and a nextPageToken if applicable.
type RequestCatalogManagementServiceListEntitlementsPerCatalogResponse struct {
	// List of serialized related objects.
	Expanded []RequestCatalogManagementServiceListEntitlementsPerCatalogResponseExpanded `json:"expanded,omitempty"`
	// The list of results containing up to X results, where X is the page size defined in the request.
	List []AppEntitlementView `json:"list,omitempty"`
	// The nextPageToken is shown for the next page if the number of results is larger than the max page size.
	//  The server returns one page of results and the nextPageToken until all results are retreived.
	//  To retrieve the next page, use the same request and append a pageToken field with the value of nextPageToken shown on the previous page.
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

func (o *RequestCatalogManagementServiceListEntitlementsPerCatalogResponse) GetExpanded() []RequestCatalogManagementServiceListEntitlementsPerCatalogResponseExpanded {
	if o == nil {
		return nil
	}
	return o.Expanded
}

func (o *RequestCatalogManagementServiceListEntitlementsPerCatalogResponse) GetList() []AppEntitlementView {
	if o == nil {
		return nil
	}
	return o.List
}

func (o *RequestCatalogManagementServiceListEntitlementsPerCatalogResponse) GetNextPageToken() *string {
	if o == nil {
		return nil
	}
	return o.NextPageToken
}
