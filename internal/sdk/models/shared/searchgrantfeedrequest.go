// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/internal/utils"
	"time"
)

// The SearchGrantFeedRequest message.
type SearchGrantFeedRequest struct {
	After *time.Time `json:"after,omitempty"`
	// The list of app entitlements to limit the search to.
	AppEntitlementRefs []AppEntitlementRef `json:"appEntitlementRefs,omitempty"`
	// The list of apps to limit the search to.
	AppRefs []AppRef `json:"appRefs,omitempty"`
	// The list of app users to limit the search to.
	AppUserRefs []AppUserRef `json:"appUserRefs,omitempty"`
	Before      *time.Time   `json:"before,omitempty"`
	// The AppEntitlementUserBindingExpandHistoryMask message.
	AppEntitlementUserBindingExpandHistoryMask *AppEntitlementUserBindingExpandHistoryMask `json:"expandMask,omitempty"`
	// The pageSize where 10 <= pageSize <= 100, default 25.
	PageSize *int `json:"pageSize,omitempty"`
	// The page_token field for pagination.
	PageToken *string `json:"pageToken,omitempty"`
	// The list of C1 users to limit the search to.
	UserRefs []UserRef `json:"userRefs,omitempty"`
}

func (s SearchGrantFeedRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SearchGrantFeedRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SearchGrantFeedRequest) GetAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *SearchGrantFeedRequest) GetAppEntitlementRefs() []AppEntitlementRef {
	if o == nil {
		return nil
	}
	return o.AppEntitlementRefs
}

func (o *SearchGrantFeedRequest) GetAppRefs() []AppRef {
	if o == nil {
		return nil
	}
	return o.AppRefs
}

func (o *SearchGrantFeedRequest) GetAppUserRefs() []AppUserRef {
	if o == nil {
		return nil
	}
	return o.AppUserRefs
}

func (o *SearchGrantFeedRequest) GetBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.Before
}

func (o *SearchGrantFeedRequest) GetAppEntitlementUserBindingExpandHistoryMask() *AppEntitlementUserBindingExpandHistoryMask {
	if o == nil {
		return nil
	}
	return o.AppEntitlementUserBindingExpandHistoryMask
}

func (o *SearchGrantFeedRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *SearchGrantFeedRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

func (o *SearchGrantFeedRequest) GetUserRefs() []UserRef {
	if o == nil {
		return nil
	}
	return o.UserRefs
}
