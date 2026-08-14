// specnorm rewrites nullable object properties in an OpenAPI document from
//
//	prop:
//	    oneOf:
//	        - $ref: '#/components/schemas/Foo'
//	        - type: "null"
//
// back to the bare reference
//
//	prop:
//	    $ref: '#/components/schemas/Foo'
//
// Speakeasy's Terraform generator names an attribute after the *referenced
// schema* for a bare $ref (Foo -> foo_bar_provision) and after the *property*
// for a oneOf-wrapped ref (prop -> prop). The insulator spec switched every
// nullable object property from the first shape to the second between provider
// 1.4.0 and 1.5.0, which renames ~250 customer-visible Terraform attributes
// across ~30 resources — with no product change behind it, and no way to
// express an attribute rename in `terraform state mv`. Collapsing the shape
// before generation keeps 1.4.x attribute names while retaining every additive
// change in the spec.
//
// The correct long-term fix is upstream, in whatever emits the public spec.
// Until then this runs inside `make gen`, between `speakeasy overlay apply` and
// `speakeasy generate sdk`.
//
// Usage:
//
//	specnorm -in combined.yaml -out combined.normalized.yaml
//
// The rewrite is confined to mappings whose ONLY key is `oneOf`, whose value is
// a two-element sequence of exactly one `$ref` mapping and one `type: "null"`
// mapping. Anything else — a three-way oneOf, a oneOf carrying a description, a
// discriminated union — is left untouched.
package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	in := flag.String("in", "", "path to the input OpenAPI document (YAML)")
	out := flag.String("out", "", "path to write the normalized document")
	flag.Parse()

	if *in == "" || *out == "" {
		fmt.Fprintln(os.Stderr, "specnorm: both -in and -out are required")
		os.Exit(2)
	}

	if err := run(*in, *out); err != nil {
		fmt.Fprintf(os.Stderr, "specnorm: %v\n", err)
		os.Exit(1)
	}
}

func run(in, out string) error {
	src, err := os.ReadFile(in)
	if err != nil {
		return err
	}

	var doc yaml.Node
	if err := yaml.Unmarshal(src, &doc); err != nil {
		return fmt.Errorf("parsing %s: %w", in, err)
	}

	n := collapse(&doc)

	buf, err := yaml.Marshal(&doc)
	if err != nil {
		return fmt.Errorf("serializing: %w", err)
	}
	if err := os.WriteFile(out, buf, 0o644); err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "specnorm: collapsed %d nullable oneOf-of-$ref properties to a bare $ref\n", n)
	return nil
}

// collapse rewrites every nullable oneOf-of-$ref mapping in the tree in place
// and reports how many it rewrote.
func collapse(n *yaml.Node) int {
	if n == nil {
		return 0
	}

	count := 0
	if ref := nullableRef(n); ref != nil {
		*n = yaml.Node{
			Kind: yaml.MappingNode,
			Tag:  "!!map",
			Content: []*yaml.Node{
				{Kind: yaml.ScalarNode, Tag: "!!str", Value: "$ref"},
				{Kind: yaml.ScalarNode, Tag: "!!str", Value: ref.Value, Style: ref.Style},
			},
			Line:   n.Line,
			Column: n.Column,
		}
		count++
	}

	for _, c := range n.Content {
		count += collapse(c)
	}
	return count
}

// nullableRef returns the $ref value node when n is a mapping whose only key is
// `oneOf` holding exactly one `$ref` member and one `type: "null"` member, and
// nil otherwise.
func nullableRef(n *yaml.Node) *yaml.Node {
	if n.Kind != yaml.MappingNode || len(n.Content) != 2 {
		return nil
	}
	if n.Content[0].Value != "oneOf" {
		return nil
	}
	seq := n.Content[1]
	if seq.Kind != yaml.SequenceNode || len(seq.Content) != 2 {
		return nil
	}

	var ref, null *yaml.Node
	for _, m := range seq.Content {
		switch {
		case soleEntry(m, "$ref") != nil:
			ref = soleEntry(m, "$ref")
		case soleEntry(m, "type") != nil && soleEntry(m, "type").Value == "null":
			null = m
		}
	}
	if ref == nil || null == nil {
		return nil
	}
	return ref
}

// soleEntry returns the value node for key when n is a mapping with that single
// key, and nil otherwise.
func soleEntry(n *yaml.Node, key string) *yaml.Node {
	if n.Kind != yaml.MappingNode || len(n.Content) != 2 || n.Content[0].Value != key {
		return nil
	}
	return n.Content[1]
}
