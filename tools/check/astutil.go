package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// pln and pf wrap Fprintln/Fprintf and intentionally drop the error return.
// Subcommands write warnings to os.Stderr; if stderr writes fail, the user
// has bigger problems than this audit script. Wrapping centralizes the
// errcheck-suppression in one place rather than scattering `_, _ = ...`
// across every emit site.
func pln(w io.Writer, args ...any) {
	_, _ = fmt.Fprintln(w, args...)
}

func pf(w io.Writer, format string, args ...any) {
	_, _ = fmt.Fprintf(w, format, args...)
}

// providerGlob is the pattern for files that hold resource / data-source
// constructors and model structs.
const providerGlob = "internal/provider/*.go"

// registrationFiles list the files whose New* identifier references count as
// "registered". Anything referenced from these files is treated as wired into
// the live provider.
var registrationFiles = []string{
	"internal/provider/provider.go",
	"internal/provider/integrations.go",
}

// constructor describes a `func New<Name>() <pkg>.<Type>` declaration found
// during the AST walk.
type constructor struct {
	name string // e.g. "NewAppEntitlementResource"
	kind string // e.g. "resource.Resource" — pkg + "." + type
	file string // source file the declaration lives in
}

// collectConstructors walks providerGlob and returns every `func New*()
// <pkg>.<Type>` declaration that looks like a terraform-plugin-framework
// constructor.
func collectConstructors() ([]constructor, error) {
	files, err := filepath.Glob(providerGlob)
	if err != nil {
		return nil, err
	}
	fset := token.NewFileSet()
	var out []constructor
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
			kind, ok := returnKind(fn)
			if !ok {
				continue
			}
			out = append(out, constructor{
				name: fn.Name.Name,
				kind: kind,
				file: f,
			})
		}
	}
	return out, nil
}

// returnKind extracts the "pkg.Type" return signature from a constructor like
// `func NewFoo() resource.Resource`. Returns ("", false) if the shape doesn't
// match (e.g. multiple return values, anonymous return type).
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

// collectReferences returns the set of identifiers starting with "New" that
// appear anywhere in registrationFiles.
//
// We treat any New* identifier reference (slice entry, function call,
// variable, anything) as "registered". Coarser than parsing the specific
// Resources()/DataSources()/etc. methods, but resilient to refactors — if a
// constructor is referenced anywhere in these files, it's reachable.
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

// resourceTFName converts a constructor name + kind into the customer-visible
// TF type name (without the provider prefix). For example:
//
//	NewAppEntitlementResource, "resource.Resource"          -> "app_entitlement"
//	NewAppEntitlementsDataSource, "datasource.DataSource"   -> "app_entitlements"
//
// Returns ("", false) if the constructor doesn't follow the New<Name><Suffix>
// shape that matches the kind.
func resourceTFName(name, kind string) (string, bool) {
	suffix := kindSuffix(kind)
	if suffix == "" {
		return "", false
	}
	if !strings.HasPrefix(name, "New") || !strings.HasSuffix(name, suffix) {
		return "", false
	}
	stem := strings.TrimSuffix(strings.TrimPrefix(name, "New"), suffix)
	if stem == "" {
		return "", false
	}
	return camelToSnake(stem), true
}

// kindSuffix returns the conventional suffix on constructor names for a given
// terraform-plugin-framework kind: "resource.Resource" -> "Resource", etc.
// Returns "" for unrecognized kinds.
func kindSuffix(kind string) string {
	_, suffix, ok := strings.Cut(kind, ".")
	if !ok {
		return ""
	}
	return suffix
}

// camelToSnake converts CamelCase (with digits) to snake_case using the
// rules Speakeasy's TF generator applies, so output matches the directory
// names under examples/ and docs/ for each constructor:
//
//	"AppEntitlement"             -> "app_entitlement"
//	"HTTPServer"                 -> "http_server"      (upper-run end before lower)
//	"AppID"                      -> "app_id"           (upper-run preserved)
//	"Microsoft365"               -> "microsoft_365"    (digit after lower-letter)
//	"IntegrationDocusignV2"      -> "integration_docusign_v2"  (V2 at stem end stays joined)
//	"IntegrationGithubV2Test"    -> "integration_github_v_2_test"  (V2 mid-stem splits)
//
// Boundaries that insert an underscore:
//   - uppercase letter following a lowercase letter or digit
//   - uppercase letter ending an upper-run when the next rune is lowercase
//   - digit following a lowercase letter
//   - digit following an uppercase letter ONLY when a later rune is also
//     uppercase (so `V2Test` splits but `V2` at stem end doesn't)
func camelToSnake(s string) string {
	var b strings.Builder
	b.Grow(len(s) + 4)
	runes := []rune(s)
	for i, r := range runes {
		isUpper := r >= 'A' && r <= 'Z'
		isDigit := r >= '0' && r <= '9'
		insertUnderscore := false
		if i > 0 {
			prev := runes[i-1]
			prevUpper := prev >= 'A' && prev <= 'Z'
			prevLower := prev >= 'a' && prev <= 'z'
			prevDigit := prev >= '0' && prev <= '9'
			switch {
			case isUpper && (prevLower || prevDigit):
				insertUnderscore = true
			case isUpper && prevUpper && i+1 < len(runes):
				next := runes[i+1]
				if next >= 'a' && next <= 'z' {
					insertUnderscore = true
				}
			case isDigit && prevLower:
				insertUnderscore = true
			case isDigit && prevUpper && hasLaterUpper(runes, i):
				insertUnderscore = true
			}
		}
		if insertUnderscore {
			b.WriteByte('_')
		}
		if isUpper {
			b.WriteRune(r + ('a' - 'A'))
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// hasLaterUpper reports whether any rune at index > i in runes is uppercase.
// Used by camelToSnake to decide whether an upper-letter -> digit boundary
// should split. If no later uppercase exists, the digit is at the end of
// the identifier (e.g. `V2` at stem end after `Resource` is stripped) and
// the convention is to keep it joined.
func hasLaterUpper(runes []rune, i int) bool {
	for j := i + 1; j < len(runes); j++ {
		if runes[j] >= 'A' && runes[j] <= 'Z' {
			return true
		}
	}
	return false
}

// kindDirName maps a kind to the directory name used under examples/ and
// docs/: "resource.Resource" -> "resources", etc.
func kindDirName(kind string) string {
	switch kind {
	case "resource.Resource":
		return "resources"
	case "datasource.DataSource":
		return "data-sources"
	case "ephemeral.EphemeralResource":
		return "ephemeral-resources"
	case "function.Function":
		return "functions"
	case "list.ListResource":
		return "list-resources"
	case "action.Action":
		return "actions"
	}
	return ""
}

// kindDisplay maps a kind to the plural human-readable label used in warning
// output: "resource.Resource" -> "Resources", etc.
func kindDisplay(kind string) string {
	switch kind {
	case "resource.Resource":
		return "Resources"
	case "datasource.DataSource":
		return "DataSources"
	case "ephemeral.EphemeralResource":
		return "EphemeralResources"
	case "function.Function":
		return "Functions"
	case "list.ListResource":
		return "ListResources"
	case "action.Action":
		return "Actions"
	}
	return kind
}

// allKinds returns every kind we know about, in a stable order matching
// (*ConductoroneProvider) method order in provider.go.
func allKinds() []string {
	return []string{
		"resource.Resource",
		"datasource.DataSource",
		"ephemeral.EphemeralResource",
		"function.Function",
		"list.ListResource",
		"action.Action",
	}
}
