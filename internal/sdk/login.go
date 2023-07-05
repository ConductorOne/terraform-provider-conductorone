package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"terraform/internal/sdk/pkg/utils"
)

type ClientCredentials struct {
	ClientID     string
	ClientSecret string
}

type DeviceCodeResponse struct {
	DeviceCode      string `json:"device_code"`
	UserCode        string `json:"user_code"`
	VerificationURI string `json:"verification_uri_complete"`
	ExpiresIn       int64  `json:"expires_in"`
	Interval        int64  `json:"interval"`
}

func LoginFlow(
	ctx context.Context,
	tenantName string,
	clientID string,
	personalClientCredentialDisplayName string,
	cb func(validateUrl string) error,
) (*ClientCredentials, error) {
	tenantURLOption, err := WithTenant(tenantName)
	if err != nil {
		return nil, err
	}
	client := New(tenantURLOption)

	codeResp, responseUrl, err := getDeviceCode(ctx, client, clientID)
	if err != nil {
		return nil, fmt.Errorf("error getting device code: %w", err)
	}

	if responseUrl == "" {
		return nil, errors.New("invalid verification url")
	}

	if err := cb(responseUrl); err != nil {
		return nil, err
	}

	tokenResp, err := doTokenRequest(ctx, client, clientID, codeResp)
	if err != nil {
		return nil, fmt.Errorf("error doing token request: %w", err)
	}

	clientCredential, err := doClientCredentialRequest(ctx, client, tokenResp, personalClientCredentialDisplayName)
	if err != nil {
		return nil, fmt.Errorf("error doing client credential request: %w", err)
	}

	return &ClientCredentials{
		ClientID:     clientCredential.Client.ClientID,
		ClientSecret: clientCredential.ClientSecret,
	}, nil
}

func newReq(ctx context.Context, method, url string, reader io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, reader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "conductorone-sdk-go")
	return req, nil
}

func getDeviceCode(ctx context.Context, client *ConductoroneAPI, clientID string) (*DeviceCodeResponse, string, error) {
	httpClient := client.sdkConfiguration.DefaultClient
	baseURL := utils.ReplaceParameters(client.sdkConfiguration.GetServerDetails())

	deviceCodeURL, err := url.JoinPath(baseURL, "/auth/v1/device_authorization")
	if err != nil {
		return nil, "", err
	}
	vals := url.Values{}
	vals.Add("client_id", clientID)

	req, err := newReq(ctx, "POST", deviceCodeURL, strings.NewReader(vals.Encode()))
	if err != nil {
		return nil, "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("error getting device code: %s", string(data))
	}

	codeResp := &DeviceCodeResponse{}
	err = json.Unmarshal(data, codeResp)
	if err != nil {
		return nil, "", fmt.Errorf("error unmarshalling response data: %w - %s", err, data)
	}

	return codeResp, codeResp.VerificationURI, nil
}

func doTokenRequest(ctx context.Context, client *ConductoroneAPI, clientID string, deviceCodeResp *DeviceCodeResponse) (*tokenResponse, error) {
	httpClient := client.sdkConfiguration.DefaultClient
	baseURL := utils.ReplaceParameters(client.sdkConfiguration.GetServerDetails())

	tokenURL, err := url.JoinPath(baseURL, "/auth/v1/token")
	if err != nil {
		return nil, err
	}
	vals := url.Values{}
	vals.Add("client_id", clientID)
	vals.Add("device_code", deviceCodeResp.DeviceCode)
	vals.Add("grant_type", "urn:ietf:params:oauth:grant-type:device_code")

	intervalSeconds := deviceCodeResp.Interval
	startLoop := time.Now()
	for {
		select {
		case <-time.After(time.Duration(intervalSeconds) * time.Second):
			if time.Now().After(startLoop.Add(time.Duration(deviceCodeResp.ExpiresIn) * time.Second)) {
				return nil, errors.New("timeout")
			}
		case <-ctx.Done():
			return nil, ctx.Err()
		}

		req, err := newReq(ctx, "POST", tokenURL, strings.NewReader(vals.Encode()))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp, err := httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != http.StatusOK {
			errResp := &oauth2Error{}
			err = json.Unmarshal(body, errResp)
			if err != nil {
				return nil, err
			}

			if errResp.ErrorType == "authorization_pending" {
				continue
			}

			if errResp.ErrorType == "slow_down" {
				// Interval should be increased by 5 seconds when slow_down is received
				// https://datatracker.ietf.org/doc/html/rfc8628#section-3.5
				intervalSeconds += 5
			}

			return nil, errors.New(errResp.ErrorDescription)
		}

		tokenResp := &tokenResponse{}
		err = json.Unmarshal(body, tokenResp)
		if err != nil {
			return nil, err
		}

		return tokenResp, nil
	}
}

type clientDisplayName struct {
	DisplayName string `json:"display_name"`
}

func doClientCredentialRequest(ctx context.Context, client *ConductoroneAPI, tokenResp *tokenResponse, personalClientCredentialDisplayName string) (*clientResp, error) {
	httpClient := client.sdkConfiguration.DefaultClient
	baseURL := utils.ReplaceParameters(client.sdkConfiguration.GetServerDetails())

	pccURL, err := url.JoinPath(baseURL, "/api/v1/iam/personal_clients")
	if err != nil {
		return nil, err
	}

	displayNameArg, err := json.Marshal(clientDisplayName{DisplayName: personalClientCredentialDisplayName})
	if err != nil {
		return nil, err
	}
	personalClientReq, err := newReq(ctx, "POST", pccURL, bytes.NewReader(displayNameArg))
	if err != nil {
		return nil, err
	}
	personalClientReq.Header.Set("Content-Type", "application/json")
	personalClientReq.Header.Set("Authorization", "Bearer "+tokenResp.AccessToken)
	resp, err := httpClient.Do(personalClientReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error getting client credential code: %s", string(body))
	}

	clientRespJSON := &clientResp{}
	err = json.Unmarshal(body, clientRespJSON)
	if err != nil {
		return nil, err
	}

	return clientRespJSON, nil
}

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

type oauth2Error struct {
	ErrorType        string          `json:"error"`
	ErrorDescription string          `json:"error_description"`
	ErrorDetails     json.RawMessage `json:"error_details,omitempty"`
	State            string          `json:"state,omitempty"`
}

type clientStats struct {
	ClientID string `json:"clientId"`
}

type clientResp struct {
	Client       clientStats `json:"client"`
	ClientSecret string      `json:"clientSecret"`
}
