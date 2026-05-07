package main

import (
	"os"
	"sort"
)

// runUnregistered diffs constructors defined in internal/provider/*.go against
// constructors referenced from provider.go + integrations.go. Anything
// defined-but-unreferenced is unreachable from the live provider — customers
// can't use it and tfplugindocs won't render a doc page for it until someone
// hand-registers it.
//
// The root cause is that provider.go is in .genignore (preserves the custom
// Configure() auth chain plus the getIntegrationResources() append), so
// Speakeasy never updates its registration slices. See IGA-1741 for the
// redesign tracking issue.
func runUnregistered() (int, error) {
	defined, err := collectConstructors()
	if err != nil {
		return 0, err
	}
	referenced, err := collectReferences()
	if err != nil {
		return 0, err
	}

	// Bucket missing constructors by kind so output groups stay stable.
	missing := map[string][]string{}
	for _, c := range defined {
		if !referenced[c.name] {
			missing[c.kind] = append(missing[c.kind], c.name)
		}
	}

	total := 0
	for _, names := range missing {
		sort.Strings(names)
		total += len(names)
	}

	if total == 0 {
		return 0, nil
	}

	w := os.Stderr
	pln(w)
	pf(w, "WARNING [unregistered]: %d Speakeasy-generated constructor(s) are defined but unregistered.\n", total)
	pln(w, "         provider.go is in .genignore, so Speakeasy can't update its slices.")
	pln(w, "         Hand-add new entries to provider.go's registration methods")
	pln(w, "         (or integrations.go's getIntegrationResources() for integration_*).")
	pln(w, "         See IGA-1741 for the redesign tracking issue.")
	for _, kind := range allKinds() {
		names, ok := missing[kind]
		if !ok {
			continue
		}
		pf(w, "\n  %s (%d):\n", kindDisplay(kind), len(names))
		for _, n := range names {
			pf(w, "    %s\n", n)
		}
	}
	pln(w)
	return total, nil
}
