// check runs static audits over the generated provider code, surfacing
// classes of regen drift that are easy to miss at PR review time:
//
//   unregistered    — Speakeasy-generated constructors not wired into
//                     provider.go's registration slices or
//                     integrations.go's getIntegrationResources().
//
//   docs-coverage   — registered constructors that lack an example HCL
//                     block in examples/<kind>/conductorone_<name>/ or a
//                     rendered doc page in docs/<kind>/<name>.md.
//
//   tag-parity      — `tfsdk:` tags on resource model structs that violate
//                     either snake_case style or reserved Terraform keyword
//                     rules. (Stricter "JSON tag <-> TF attribute name
//                     parity" — see TODO in parity.go — deferred.)
//
//   todo-markers    — `// TODO` comments inside Speakeasy-generated files.
//                     Generated code shouldn't ship with author-stub
//                     placeholders; flagging them surfaces regen output
//                     that the upstream generator emitted as incomplete.
//
//   all             — runs every subcommand above. Aggregate exit code is
//                     the max of the subcommand exit codes (so a single
//                     failure propagates).
//
// All subcommands exit 0 by default and emit warnings to stderr. Failing
// the build would force every regen to clear the historical backlog before
// shipping; warnings let the reviewer judge per-PR. Pass -strict to upgrade
// any non-empty finding into exit code 1 (suitable for CI gating once the
// backlog is clean).
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// runFunc is the signature each subcommand exposes. It writes its findings
// to stderr and returns the count of items found (0 = clean). The dispatcher
// translates count + -strict into an exit code.
type runFunc func() (count int, err error)

var subcommands = map[string]runFunc{
	"unregistered":  runUnregistered,
	"docs-coverage": runDocsCoverage,
	"tag-parity":    runTagParity,
	"todo-markers":  runTodoMarkers,
}

func main() {
	strict := flag.Bool("strict", false, "exit 1 on any finding (default: warning only, exit 0)")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [-strict] <subcommand>\n", filepath.Base(os.Args[0]))
		fmt.Fprintln(os.Stderr, "subcommands:")
		fmt.Fprintln(os.Stderr, "  unregistered    Speakeasy-generated constructors not wired into provider.go")
		fmt.Fprintln(os.Stderr, "  docs-coverage   registered constructors lacking examples/ or docs/ pages")
		fmt.Fprintln(os.Stderr, "  tag-parity      tfsdk: tags violating snake_case / reserved keyword rules")
		fmt.Fprintln(os.Stderr, "  todo-markers    // TODO comments inside Speakeasy-generated files")
		fmt.Fprintln(os.Stderr, "  all             run every subcommand above")
	}
	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(2)
	}
	sub := flag.Arg(0)

	if err := chdirToRepoRoot(); err != nil {
		fail("locating repo root: %v", err)
	}

	if sub == "all" {
		exit := 0
		for _, name := range []string{"unregistered", "docs-coverage", "tag-parity", "todo-markers"} {
			count, err := subcommands[name]()
			if err != nil {
				fail("%s: %v", name, err)
			}
			if count > 0 && *strict {
				exit = 1
			}
		}
		os.Exit(exit)
	}

	fn, ok := subcommands[sub]
	if !ok {
		flag.Usage()
		os.Exit(2)
	}
	count, err := fn()
	if err != nil {
		fail("%s: %v", sub, err)
	}
	if count > 0 && *strict {
		os.Exit(1)
	}
}

// chdirToRepoRoot walks up from the current working directory until it
// finds a go.mod, then chdirs there. This lets the binary be invoked from
// any subdirectory without flag-threading.
func chdirToRepoRoot() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	for {
		if _, err := os.Stat(filepath.Join(wd, "go.mod")); err == nil {
			return os.Chdir(wd)
		}
		parent := filepath.Dir(wd)
		if parent == wd {
			return fmt.Errorf("go.mod not found above %q", wd)
		}
		wd = parent
	}
}

func fail(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "check: "+format+"\n", args...)
	os.Exit(2)
}
