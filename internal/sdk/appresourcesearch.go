// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"bytes"
	"context"
	"fmt"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/operations"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/utils"
	"io"
	"net/http"
	"strings"
)

type appResourceSearch struct {
	sdkConfiguration sdkConfiguration
}

func newAppResourceSearch(sdkConfig sdkConfiguration) *appResourceSearch {
	return &appResourceSearch{
		sdkConfiguration: sdkConfig,
	}
}

// SearchAppResourceTypes - Invokes the c1.api.app.v1.AppResourceSearch.SearchAppResourceTypes method.
func (s *appResourceSearch) SearchAppResourceTypes(ctx context.Context, request shared.SearchAppResourceTypesRequest) (*operations.C1APIAppV1AppResourceSearchSearchAppResourceTypesResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/api/v1/search/app_resource_types"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "Request", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	req.Header.Set("Content-Type", reqContentType)

	client := s.sdkConfiguration.DefaultClient

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
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.C1APIAppV1AppResourceSearchSearchAppResourceTypesResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.SearchAppResourceTypesResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.SearchAppResourceTypesResponse = out
		}
	}

	return res, nil
}
