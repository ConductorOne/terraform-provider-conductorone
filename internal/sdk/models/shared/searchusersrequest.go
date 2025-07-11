// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ExcludeTypes string

const (
	ExcludeTypesUserTypeUnspecified ExcludeTypes = "USER_TYPE_UNSPECIFIED"
	ExcludeTypesUserTypeSystem      ExcludeTypes = "USER_TYPE_SYSTEM"
	ExcludeTypesUserTypeHuman       ExcludeTypes = "USER_TYPE_HUMAN"
	ExcludeTypesUserTypeService     ExcludeTypes = "USER_TYPE_SERVICE"
	ExcludeTypesUserTypeAgent       ExcludeTypes = "USER_TYPE_AGENT"
)

func (e ExcludeTypes) ToPointer() *ExcludeTypes {
	return &e
}
func (e *ExcludeTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "USER_TYPE_UNSPECIFIED":
		fallthrough
	case "USER_TYPE_SYSTEM":
		fallthrough
	case "USER_TYPE_HUMAN":
		fallthrough
	case "USER_TYPE_SERVICE":
		fallthrough
	case "USER_TYPE_AGENT":
		*e = ExcludeTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ExcludeTypes: %v", v)
	}
}

type UserStatuses string

const (
	UserStatusesUnknown  UserStatuses = "UNKNOWN"
	UserStatusesEnabled  UserStatuses = "ENABLED"
	UserStatusesDisabled UserStatuses = "DISABLED"
	UserStatusesDeleted  UserStatuses = "DELETED"
)

func (e UserStatuses) ToPointer() *UserStatuses {
	return &e
}
func (e *UserStatuses) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "UNKNOWN":
		fallthrough
	case "ENABLED":
		fallthrough
	case "DISABLED":
		fallthrough
	case "DELETED":
		*e = UserStatuses(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UserStatuses: %v", v)
	}
}

// SearchUsersRequest - Search for users based on some filters.
type SearchUsersRequest struct {
	// Search for users based on their email (exact match).
	Email *string `json:"email,omitempty"`
	// An array of users IDs to exclude from the results.
	ExcludeIds []string `json:"excludeIds,omitempty"`
	// An array of types to exclude from the results.
	ExcludeTypes []ExcludeTypes `json:"excludeTypes,omitempty"`
	// Deprecated. Use refs array instead.
	Ids []string `json:"ids,omitempty"`
	// The pageSize where 0 <= pageSize <= 100. Values < 10 will be set to 10. A value of 0 returns the default page size (currently 25)
	PageSize *int `json:"pageSize,omitempty"`
	// The pageToken field.
	PageToken *string `json:"pageToken,omitempty"`
	// Query the apps with a fuzzy search on display name and emails.
	Query *string `json:"query,omitempty"`
	// An array of user refs to restrict the return values to by ID.
	Refs []UserRef `json:"refs,omitempty"`
	// Search for users that have any of the role IDs on this list.
	RoleIds []string `json:"roleIds,omitempty"`
	// Search for users that have any of the statuses on this list. This can only be ENABLED, DISABLED, and DELETED
	UserStatuses []UserStatuses `json:"userStatuses,omitempty"`
}

func (o *SearchUsersRequest) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *SearchUsersRequest) GetExcludeIds() []string {
	if o == nil {
		return nil
	}
	return o.ExcludeIds
}

func (o *SearchUsersRequest) GetExcludeTypes() []ExcludeTypes {
	if o == nil {
		return nil
	}
	return o.ExcludeTypes
}

func (o *SearchUsersRequest) GetIds() []string {
	if o == nil {
		return nil
	}
	return o.Ids
}

func (o *SearchUsersRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *SearchUsersRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

func (o *SearchUsersRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *SearchUsersRequest) GetRefs() []UserRef {
	if o == nil {
		return nil
	}
	return o.Refs
}

func (o *SearchUsersRequest) GetRoleIds() []string {
	if o == nil {
		return nil
	}
	return o.RoleIds
}

func (o *SearchUsersRequest) GetUserStatuses() []UserStatuses {
	if o == nil {
		return nil
	}
	return o.UserStatuses
}
