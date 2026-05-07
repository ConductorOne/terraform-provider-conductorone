package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
)

// reservedTFNames are Terraform meta-arguments and other reserved attribute
// names that should never appear as a resource attribute. Using them produces
// confusing parse errors at apply time. Sourced from
// https://developer.hashicorp.com/terraform/language/meta-arguments and the
// resource block reference.
var reservedTFNames = map[string]bool{
	"count":      true,
	"depends_on": true,
	"for_each":   true,
	"lifecycle":  true,
	"locals":     true,
	"provider":   true,
	"providers":  true,
	"source":     true,
	"terraform":  true,
}

// runTagParity scans every `tfsdk:"..."` tag on resource / data-source model
// structs and flags two classes of issue:
//
//   - tags that aren't valid snake_case (catches camelCase leakage from
//     Phase-2-class collision-resolution renames)
//   - tags equal to a reserved Terraform meta-argument name
//
// TODO: extend with the stricter "JSON-tag <-> TF-attribute parity" check
// from the original brainstorm — for each TF attribute, walk to the wrapped
// SDK shared type and verify the json: tag matches the snake_case form of
// the tfsdk: tag. That would catch Phase-2-style renames where Speakeasy's
// generator inherited a renamed schema name into the parent's nested
// attribute name (the speakeasy-feature-request-property-override.md issue).
// Not worth the AST-walking complexity until we hit a regression that
// (a) and (b) fail to catch.
func runTagParity() (int, error) {
	files, err := filepath.Glob(providerGlob)
	if err != nil {
		return 0, err
	}

	type finding struct {
		tag     string // the tfsdk tag value
		field   string // Go field name carrying the tag
		struct_ string // enclosing struct
		file    string
		why     string // "non-snake-case" or "reserved-keyword"
	}
	var findings []finding

	fset := token.NewFileSet()
	for _, f := range files {
		if strings.HasSuffix(f, "_test.go") {
			continue
		}
		af, err := parser.ParseFile(fset, f, nil, parser.SkipObjectResolution)
		if err != nil {
			continue
		}
		ast.Inspect(af, func(n ast.Node) bool {
			ts, ok := n.(*ast.TypeSpec)
			if !ok {
				return true
			}
			st, ok := ts.Type.(*ast.StructType)
			if !ok {
				return true
			}
			for _, field := range st.Fields.List {
				if field.Tag == nil {
					continue
				}
				tag := reflect.StructTag(strings.Trim(field.Tag.Value, "`")).Get("tfsdk")
				if tag == "" || tag == "-" {
					continue
				}
				fieldName := ""
				if len(field.Names) > 0 {
					fieldName = field.Names[0].Name
				}
				if !isSnakeCase(tag) {
					findings = append(findings, finding{
						tag: tag, field: fieldName, struct_: ts.Name.Name, file: f,
						why: "non-snake-case",
					})
				}
				if reservedTFNames[tag] {
					findings = append(findings, finding{
						tag: tag, field: fieldName, struct_: ts.Name.Name, file: f,
						why: "reserved-keyword",
					})
				}
			}
			return true
		})
	}

	if len(findings) == 0 {
		return 0, nil
	}

	// Stable order: file, then struct, then field.
	sort.Slice(findings, func(i, j int) bool {
		if findings[i].file != findings[j].file {
			return findings[i].file < findings[j].file
		}
		if findings[i].struct_ != findings[j].struct_ {
			return findings[i].struct_ < findings[j].struct_
		}
		return findings[i].field < findings[j].field
	})

	w := os.Stderr
	fmt.Fprintln(w)
	fmt.Fprintf(w, "WARNING [tag-parity]: %d tfsdk: tag(s) violate naming rules.\n", len(findings))
	fmt.Fprintln(w, "         Snake-case violations leak camelCase from Phase-2-class collision renames.")
	fmt.Fprintln(w, "         Reserved-keyword violations conflict with Terraform meta-arguments and")
	fmt.Fprintln(w, "         produce confusing parse errors at customer apply time.")
	fmt.Fprintln(w)
	for _, x := range findings {
		fmt.Fprintf(w, "    [%s] %s.%s tfsdk:%q  (%s)\n", x.why, x.struct_, x.field, x.tag, x.file)
	}
	fmt.Fprintln(w)
	return len(findings), nil
}

// isSnakeCase reports whether s is valid snake_case: lowercase ASCII letters
// or digits, with single underscores between segments, no leading/trailing
// underscore, no consecutive underscores.
func isSnakeCase(s string) bool {
	if s == "" {
		return false
	}
	if s[0] == '_' || s[len(s)-1] == '_' {
		return false
	}
	prevUnderscore := false
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z':
			prevUnderscore = false
		case r >= '0' && r <= '9':
			prevUnderscore = false
		case r == '_':
			if prevUnderscore {
				return false
			}
			prevUnderscore = true
		default:
			return false
		}
	}
	return true
}
