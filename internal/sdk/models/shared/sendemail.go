// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The SendEmail message.
type SendEmail struct {
	// The body field.
	Body *string `json:"body,omitempty"`
	// The subject field.
	Subject *string `json:"subject,omitempty"`
	// The title field.
	Title *string `json:"title,omitempty"`
	// If true, the step will use the subject user of the automation as the subject.
	UseSubjectUser *bool `json:"useSubjectUser,omitempty"`
	// The userIdsCel field.
	UserIdsCel *string `json:"userIdsCel,omitempty"`
	// The userRefs field.
	UserRefs []UserRef `json:"userRefs,omitempty"`
}

func (o *SendEmail) GetBody() *string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *SendEmail) GetSubject() *string {
	if o == nil {
		return nil
	}
	return o.Subject
}

func (o *SendEmail) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *SendEmail) GetUseSubjectUser() *bool {
	if o == nil {
		return nil
	}
	return o.UseSubjectUser
}

func (o *SendEmail) GetUserIdsCel() *string {
	if o == nil {
		return nil
	}
	return o.UserIdsCel
}

func (o *SendEmail) GetUserRefs() []UserRef {
	if o == nil {
		return nil
	}
	return o.UserRefs
}
