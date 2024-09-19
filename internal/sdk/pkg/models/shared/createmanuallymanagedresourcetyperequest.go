// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ResourceType - The resourceType field.
type ResourceType string

const (
	ResourceTypeRole    ResourceType = "ROLE"
	ResourceTypeGroup   ResourceType = "GROUP"
	ResourceTypeLicense ResourceType = "LICENSE"
	ResourceTypeProject ResourceType = "PROJECT"
	ResourceTypeCatalog ResourceType = "CATALOG"
	ResourceTypeCustom  ResourceType = "CUSTOM"
)

func (e ResourceType) ToPointer() *ResourceType {
	return &e
}
func (e *ResourceType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ROLE":
		fallthrough
	case "GROUP":
		fallthrough
	case "LICENSE":
		fallthrough
	case "PROJECT":
		fallthrough
	case "CATALOG":
		fallthrough
	case "CUSTOM":
		*e = ResourceType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ResourceType: %v", v)
	}
}

// The CreateManuallyManagedResourceTypeRequest message.
type CreateManuallyManagedResourceTypeRequest struct {
	// The name field.
	Name *string `json:"name,omitempty"`
	// The resourceType field.
	ResourceType *ResourceType `json:"resourceType,omitempty"`
}

func (o *CreateManuallyManagedResourceTypeRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CreateManuallyManagedResourceTypeRequest) GetResourceType() *ResourceType {
	if o == nil {
		return nil
	}
	return o.ResourceType
}
