// check-unregistered lists Speakeasy-generated constructors (resources, data
// sources, ephemeral resources, functions, list resources, actions) that are
// defined in internal/provider/*.go but not referenced from provider.go's
// registration slices or integrations.go's getIntegrationResources().
//
// provider.go is in .genignore (preserves the custom Configure() auth chain
// and the getIntegrationResources() append), so Speakeasy never updates its
// registration slices. Every regen that adds a new constructor produces a
// function that's defined-but-unreachable from the live provider until
// someone hand-adds it. This tool surfaces the delta so the regen-PR
// reviewer can register new entries explicitly.
//
// All six terraform-plugin-framework constructor kinds are checked; today
// only Resource and DataSource have any defined constructors, but the others
// are watched so that future expansion of gen.yaml's additional* slots
// stays caught.
//
// Output is informational. Exits 0 even when entries are missing — failing
// the build would force every regen to clear the historical backlog before
// shipping. Tracking the cleanup separately (see IGA-1741).
package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// kind describes one terraform-plugin-framework constructor type.
//
// pkg + typ identify the return type ("resource.Resource", etc.). display is
// the human-readable label used in warning output.
type kind struct {
	pkg, typ, display string
}

// All terraform-plugin-framework constructor kinds. Add a new entry here
// when the framework adds another type. Order determines warning section
// order; we keep it consistent with provider.go's method order.
var kinds = []kind{
	{"resource", "Resource", "Resources"},
	{"datasource", "DataSource", "DataSources"},
	{"ephemeral", "EphemeralResource", "EphemeralResources"},
	{"function", "Function", "Functions"},
	{"list", "ListResource", "ListResources"},
	{"action", "Action", "Actions"},
}

// Files whose identifier references count as "registered". Anything
// referenced from these files is treated as wired into the live provider.
var registrationFiles = []string{
	"internal/provider/provider.go",
	"internal/provider/integrations.go",
}

// Glob for files to scan for constructor definitions. Test files are
// excluded inline.
const definitionGlob = "internal/provider/*.go"

func main() {
	quiet := flag.Bool("quiet", false, "suppress success message when nothing is missing")
	flag.Parse()

	if err := chdirToRepoRoot(); err != nil {
		fail("locating repo root: %v", err)
	}

	defined, err := collectDefinitions()
	if err != nil {
		fail("scanning definitions: %v", err)
	}

	referenced, err := collectReferences()
	if err != nil {
		fail("scanning references: %v", err)
	}

	missing := diff(defined, referenced)
	total := 0
	for _, names := range missing {
		total += len(names)
	}

	if total == 0 {
		if !*quiet {
			fmt.Println("All defined New*() constructors are registered.")
		}
		return
	}

	emitWarning(total, missing)
}

// collectDefinitions walks the provider package and returns a map of
// kind ("pkg.Type") -> sorted constructor names.
func collectDefinitions() (map[string][]string, error) {
	files, err := filepath.Glob(definitionGlob)
	if err != nil {
		return nil, err
	}
	fset := token.NewFileSet()
	out := map[string][]string{}
	for _, f := range files {
		if strings.HasSuffix(f, "_test.go") {
			continue
		}
		af, err := parser.ParseFile(fset, f, nil, parser.SkipObjectResolution)
		if err != nil {
			// Skip unparseable files rather than failing the whole audit; a
			// regen-in-progress can leave a transiently-broken file behind.
			continue
		}
		for _, decl := range af.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Recv != nil {
				continue
			}
			if !strings.HasPrefix(fn.Name.Name, "New") {
				continue
			}
			k, ok := returnKind(fn)
			if !ok {
				continue
			}
			out[k] = append(out[k], fn.Name.Name)
		}
	}
	for k := range out {
		sort.Strings(out[k])
	}
	return out, nil
}

// returnKind extracts the "pkg.Type" return signature from a constructor
// like `func NewFoo() resource.Resource`. Returns ("", false) if the shape
// doesn't match.
func returnKind(fn *ast.FuncDecl) (string, bool) {
	if fn.Type.Results == nil || len(fn.Type.Results.List) != 1 {
		return "", false
	}
	sel, ok := fn.Type.Results.List[0].Type.(*ast.SelectorExpr)
	if !ok {
		return "", false
	}
	pkgIdent, ok := sel.X.(*ast.Ident)
	if !ok {
		return "", false
	}
	return pkgIdent.Name + "." + sel.Sel.Name, true
}

// collectReferences walks every registrationFiles entry and returns the
// set of identifiers starting with "New".
//
// We treat any New* identifier reference (slice entry, function call,
// variable, anything) as "registered". Coarser than parsing the specific
// Resources()/DataSources()/etc. methods, but resilient to refactors —
// if a constructor is referenced anywhere in these files, it's reachable.
func collectReferences() (map[string]bool, error) {
	out := map[string]bool{}
	fset := token.NewFileSet()
	for _, f := range registrationFiles {
		af, err := parser.ParseFile(fset, f, nil, parser.SkipObjectResolution)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return nil, err
		}
		ast.Inspect(af, func(n ast.Node) bool {
			id, ok := n.(*ast.Ident)
			if ok && strings.HasPrefix(id.Name, "New") {
				out[id.Name] = true
			}
			return true
		})
	}
	return out, nil
}

// diff returns a map of kind -> sorted constructor names that are defined
// but not referenced. Kinds with zero missing entries are omitted.
func diff(defined map[string][]string, referenced map[string]bool) map[string][]string {
	out := map[string][]string{}
	for k, names := range defined {
		var missing []string
		for _, n := range names {
			if !referenced[n] {
				missing = append(missing, n)
			}
		}
		if len(missing) > 0 {
			out[k] = missing
		}
	}
	return out
}

// emitWarning prints the human-readable warning to stderr in the same
// section order as kinds (predictable diffs in CI logs).
func emitWarning(total int, missing map[string][]string) {
	w := os.Stderr
	fmt.Fprintln(w)
	fmt.Fprintf(w, "WARNING: %d Speakeasy-generated constructor(s) are defined but unregistered.\n", total)
	fmt.Fprintln(w, "         provider.go is in .genignore, so Speakeasy can't update its slices.")
	fmt.Fprintln(w, "         Hand-add new entries to provider.go's registration methods")
	fmt.Fprintln(w, "         (or integrations.go's getIntegrationResources() for integration_*).")
	fmt.Fprintln(w, "         See IGA-1741 for the redesign tracking issue.")
	for _, k := range kinds {
		key := k.pkg + "." + k.typ
		names, ok := missing[key]
		if !ok {
			continue
		}
		fmt.Fprintf(w, "\n  %s (%d):\n", k.display, len(names))
		for _, n := range names {
			fmt.Fprintf(w, "    %s\n", n)
		}
	}
	fmt.Fprintln(w)
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
	fmt.Fprintf(os.Stderr, "check-unregistered: "+format+"\n", args...)
	os.Exit(2)
}
