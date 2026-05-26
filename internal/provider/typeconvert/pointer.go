package typeconvert

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// BoolPointerOrFalse converts a *bool to a types.Bool, returning false instead
// of null when the pointer is nil. This prevents perpetual Terraform drift on
// optional boolean fields whose zero value (false) is omitted by protobuf APIs.
func BoolPointerOrFalse(value *bool) types.Bool {
	if value == nil {
		return types.BoolValue(false)
	}
	return types.BoolValue(*value)
}

// StringPointerOrEmpty converts a *string to a types.String, returning ""
// instead of null when the pointer is nil. This prevents perpetual Terraform
// drift on optional string fields whose zero value ("") is omitted by protobuf
// APIs.
func StringPointerOrEmpty(value *string) types.String {
	if value == nil {
		return types.StringValue("")
	}
	return types.StringValue(*value)
}
