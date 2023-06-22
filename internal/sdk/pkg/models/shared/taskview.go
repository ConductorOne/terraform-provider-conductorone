// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TaskView - The TaskView message.
type TaskView struct {
	// The accessReviewPath field.
	AccessReviewPath *string `json:"accessReviewPath,omitempty"`
	// The appPath field.
	AppPath *string `json:"appPath,omitempty"`
	// The appUserPath field.
	AppUserPath *string `json:"appUserPath,omitempty"`
	// The createdByUserPath field.
	CreatedByUserPath *string `json:"createdByUserPath,omitempty"`
	// The entitlementsPath field.
	EntitlementsPath *string `json:"entitlementsPath,omitempty"`
	// The identityUserPath field.
	IdentityUserPath *string `json:"identityUserPath,omitempty"`
	// The stepApproversPath field.
	StepApproversPath *string `json:"stepApproversPath,omitempty"`
	// The Task message.
	Task *Task `json:"task,omitempty"`
	// The userPath field.
	UserPath *string `json:"userPath,omitempty"`
}
