package provider

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
)

func TestAppEntitlementDataSourceEmptySearchResult(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("request method = %s, want %s", r.Method, http.MethodPost)
		}
		if r.URL.Path != "/api/v1/search/entitlements" {
			t.Errorf("request path = %s, want /api/v1/search/entitlements", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"list":[],"expanded":[],"nextPageToken":"","facets":null}`))
	}))
	defer server.Close()

	client := sdk.New(sdk.WithServerURL(server.URL))
	result, err := client.AppEntitlementSearch.Search(context.Background(), &shared.AppEntitlementSearchServiceSearchRequest{})
	if err != nil {
		t.Fatalf("search: %v", err)
	}
	if result.StatusCode != http.StatusOK {
		t.Fatalf("search status = %d, want %d", result.StatusCode, http.StatusOK)
	}

	data := &AppEntitlementDataSourceModel{}
	diagnostics := data.RefreshFromSharedAppEntitlementSearchServiceSearchResponse(context.Background(), result.AppEntitlementSearchServiceSearchResponse)
	if diagnostics.HasError() {
		t.Fatalf("empty search response produced diagnostics: %v", diagnostics)
	}
	if data.NextPageToken.ValueString() != "" {
		t.Fatalf("next_page_token = %q, want empty string", data.NextPageToken.ValueString())
	}
}
