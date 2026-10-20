package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// oneofArm is one arm of a policy oneof: its Terraform attribute name, whether
// it is set, and how to clear it.
//
// The API models each provision and deprovisioner policy as a protobuf oneof and
// rejects a body carrying two arms. The generated models make every arm an
// independent optional field, and merge() loads prior state before overlaying the
// plan, so an arm being replaced survives next to its replacement. Each resource
// lists its arms and uses these helpers to keep only the one in configuration.
type oneofArm struct {
	name  string
	isSet func() bool
	clear func()
}

// selectedOneofArm returns the name of the one arm set, or "" when none is set.
// More than one is an error: choosing among them would apply a policy the
// practitioner did not write.
func selectedOneofArm(attribute string, arms []oneofArm) (string, diag.Diagnostics) {
	var diags diag.Diagnostics

	var set []string
	for _, arm := range arms {
		if arm.isSet() {
			set = append(set, arm.name)
		}
	}

	switch len(set) {
	case 0:
		return "", diags
	case 1:
		return set[0], diags
	}

	diags.AddError(
		fmt.Sprintf("Conflicting %s configuration", attribute),
		fmt.Sprintf(
			"Only one arm may be configured, but the configuration sets %s. The API models this as a oneof and rejects a body carrying more than one arm.",
			strings.Join(set, ", "),
		),
	)
	return "", diags
}

// keepOnlyOneofArm clears every arm except keep. An empty keep leaves the arms
// untouched, so a configuration naming no arm keeps the stored one.
func keepOnlyOneofArm(arms []oneofArm, keep string) {
	if keep == "" {
		return
	}
	for _, arm := range arms {
		if arm.name != keep {
			arm.clear()
		}
	}
}
