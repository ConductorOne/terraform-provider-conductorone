package provider

import (
	"context"
	"slices"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const (
	providerConfig = `
provider "conductorone" {
  server_url    = ""
  client_id     = ""
  client_secret = ""
}
`
)

var (
	testAccProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"conductorone": providerserver.NewProtocol6WithError(New("test")()),
	}
)

func TestRegistersAccessReviewSetupResources(t *testing.T) {
	t.Parallel()

	names := registeredResourceNames(t)

	for _, want := range []string{
		"conductorone_access_review_setup",
		"conductorone_access_review_template_setup",
	} {
		if !slices.Contains(names, want) {
			t.Fatalf("resource %q is not registered; got %v", want, names)
		}
	}
}

func TestRegistersAccessReviewSetupDataSources(t *testing.T) {
	t.Parallel()

	names := registeredDataSourceNames(t)

	for _, want := range []string{
		"conductorone_access_review_setup",
		"conductorone_access_review_template_setup",
	} {
		if !slices.Contains(names, want) {
			t.Fatalf("data source %q is not registered; got %v", want, names)
		}
	}
}

func registeredResourceNames(t *testing.T) []string {
	ctx := context.Background()
	p := mustTestProvider(t)
	names := make([]string, 0, len(p.Resources(ctx)))

	for _, factory := range p.Resources(ctx) {
		r := factory()
		var resp resource.MetadataResponse
		r.Metadata(ctx, resource.MetadataRequest{ProviderTypeName: "conductorone"}, &resp)
		names = append(names, resp.TypeName)
	}

	return names
}

func registeredDataSourceNames(t *testing.T) []string {
	ctx := context.Background()
	p := mustTestProvider(t)
	names := make([]string, 0, len(p.DataSources(ctx)))

	for _, factory := range p.DataSources(ctx) {
		ds := factory()
		var resp datasource.MetadataResponse
		ds.Metadata(ctx, datasource.MetadataRequest{ProviderTypeName: "conductorone"}, &resp)
		names = append(names, resp.TypeName)
	}

	return names
}

func mustTestProvider(t *testing.T) *ConductoroneProvider {
	t.Helper()

	candidate := New("test")()
	p, ok := candidate.(*ConductoroneProvider)
	if !ok {
		t.Fatalf("test provider has unexpected type %T", candidate)
	}

	return p
}
