package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// providerName is the prefix for example directory names. Hardcoded rather
// than read from somewhere because it never changes for this provider.
const providerName = "conductorone"

// runDocsCoverage walks every constructor that IS registered (the inverse of
// runUnregistered's filter) and verifies it has both:
//
//   - an example HCL file under examples/<kind-dir>/<provider>_<name>/
//   - a rendered doc page at docs/<kind-dir>/<name>.md
//
// Either gap means a customer hitting the registry sees a barebones page
// without the HCL the rest of the provider's docs include. We caught this
// when the tfplugindocs Makefile target was misconfigured — flag now so the
// next regression surfaces immediately rather than on next regen.
//
// Note: the related "no-diff after running tfplugindocs" invariant lives in
// CI/Makefile, not here. It requires shelling out to tfplugindocs, which we
// keep out of this static-analysis binary.
func runDocsCoverage() (int, error) {
	defined, err := collectConstructors()
	if err != nil {
		return 0, err
	}
	referenced, err := collectReferences()
	if err != nil {
		return 0, err
	}

	type miss struct {
		constructor string
		examplePath string // expected dir under examples/, or "" if found
		docPath     string // expected file under docs/, or "" if found
	}
	missesByKind := map[string][]miss{}
	totalMissing := 0

	for _, c := range defined {
		if !referenced[c.name] {
			// Out of scope here — flagged by runUnregistered.
			continue
		}
		tfName, ok := resourceTFName(c.name, c.kind)
		if !ok {
			continue
		}
		dir := kindDirName(c.kind)
		if dir == "" {
			continue
		}
		// Example dir contains at least one .tf file (resource.tf,
		// data-source.tf, import.sh, etc.). We treat the dir's existence
		// + non-emptiness as "covered".
		exampleDir := filepath.Join("examples", dir, providerName+"_"+tfName)
		hasExample, err := dirHasFiles(exampleDir)
		if err != nil {
			return 0, err
		}
		docPage := filepath.Join("docs", dir, tfName+".md")
		hasDoc := fileExists(docPage)

		if hasExample && hasDoc {
			continue
		}
		m := miss{constructor: c.name}
		if !hasExample {
			m.examplePath = exampleDir
		}
		if !hasDoc {
			m.docPath = docPage
		}
		missesByKind[c.kind] = append(missesByKind[c.kind], m)
		totalMissing++
	}

	if totalMissing == 0 {
		return 0, nil
	}

	w := os.Stderr
	fmt.Fprintln(w)
	fmt.Fprintf(w, "WARNING [docs-coverage]: %d registered constructor(s) lack examples/ or docs/ pages.\n", totalMissing)
	fmt.Fprintln(w, "         For each constructor: confirm an example HCL file exists under")
	fmt.Fprintln(w, "         examples/<kind>/<provider>_<name>/, then rerun `make generate`")
	fmt.Fprintln(w, "         to render docs/<kind>/<name>.md.")
	for _, kind := range allKinds() {
		entries, ok := missesByKind[kind]
		if !ok {
			continue
		}
		sort.Slice(entries, func(i, j int) bool { return entries[i].constructor < entries[j].constructor })
		fmt.Fprintf(w, "\n  %s (%d):\n", kindDisplay(kind), len(entries))
		for _, m := range entries {
			fmt.Fprintf(w, "    %s\n", m.constructor)
			if m.examplePath != "" {
				fmt.Fprintf(w, "      missing example: %s/\n", m.examplePath)
			}
			if m.docPath != "" {
				fmt.Fprintf(w, "      missing doc:     %s\n", m.docPath)
			}
		}
	}
	fmt.Fprintln(w)
	return totalMissing, nil
}

// dirHasFiles reports whether dir exists and contains at least one regular
// file. Used to treat example dirs as covered iff they have any HCL/sh/etc.
func dirHasFiles(dir string) (bool, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	for _, e := range entries {
		if e.Type().IsRegular() {
			return true, nil
		}
	}
	return false, nil
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}
