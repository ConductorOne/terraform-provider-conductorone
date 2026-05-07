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
//   - (a) tags that aren't valid snake_case
//   - (b) tags equal to a reserved Terraform meta-argument name
//
// (a) catches camelCase leaking into customer-visible HCL attribute names.
// terraform-plugin-framework doesn't enforce snake_case on attribute names,
// but customers expect it (every other attribute in the provider follows it),
// and inconsistencies are confusing to write and read.
//
// (b) prevents collisions with Terraform's meta-arguments. Using a reserved
// name as a resource attribute produces opaque parse errors at apply time —
// e.g. an attribute called `count` competes with the `count` meta-argument.
//
// TODO: a stricter check could verify each `tfsdk:"foo_bar"` tag aligns with
// the JSON tag of the field it ultimately maps to in the wrapped SDK shared
// type — i.e. the customer-visible HCL name and the wire-level JSON name
// agree (modulo snake-vs-camel). Useful for catching cases where Speakeasy
// inherits a renamed schema name into a parent's nested attribute name and
// the TF attribute drifts from the underlying API field. Not worth the AST
// navigation complexity (resource model -> wrapped SDK type -> field tags)
// until we hit a regression (a)+(b) miss.
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
	fmt.Fprintln(w, "         Snake-case violations expose camelCase in customer-visible HCL attribute")
	fmt.Fprintln(w, "         names; reserved-keyword violations collide with Terraform meta-arguments")
	fmt.Fprintln(w, "         and produce opaque parse errors at customer apply time.")
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
