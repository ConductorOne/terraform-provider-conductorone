// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ConductorOne/terraform-provider-conductorone/internal/sdk/pkg/models/operations"
	"github.com/ConductorOne/terraform-provider-conductorone/internal/sdk/pkg/models/sdkerrors"
	"github.com/ConductorOne/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"github.com/ConductorOne/terraform-provider-conductorone/internal/sdk/pkg/utils"
	"io"
	"net/http"
	"strings"
)

type RequestCatalogSearch struct {
	sdkConfiguration sdkConfiguration
}

func newRequestCatalogSearch(sdkConfig sdkConfiguration) *RequestCatalogSearch {
	return &RequestCatalogSearch{
		sdkConfiguration: sdkConfig,
	}
}

// SearchEntitlements - Search Entitlements
// Search request catalogs based on filters specified in the request body.
func (s *RequestCatalogSearch) SearchEntitlements(ctx context.Context, request *shared.RequestCatalogSearchServiceSearchEntitlementsRequest) (*operations.C1APIRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/api/v1/search/request_catalog/entitlements"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, true, "Request", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	debugBody := bytes.NewBuffer([]byte{})
	debugReader := io.TeeReader(bodyReader, debugBody)

	req, err := http.NewRequestWithContext(ctx, "POST", url, debugReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	req.Header.Set("Content-Type", reqContentType)

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Request.Body = io.NopCloser(debugBody)
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.C1APIRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.RequestCatalogSearchServiceSearchEntitlementsResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.RequestCatalogSearchServiceSearchEntitlementsResponse = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}
