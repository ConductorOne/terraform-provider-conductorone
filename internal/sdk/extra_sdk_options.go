package sdk

import (
	"context"
	"crypto/tls"
	"net/http"
	"net/url"
	"strings"

	"go.uber.org/zap"

	"terraform/internal/sdk/uhttp"
)

const ClientIdGolangSDK = "2RCzHlak5q7CY14SdBc8HoZEJRf"

type normalizeTenantResp struct {
	useWithServer bool
	useWithTenant bool

	ServerURL string
	Tenant    string
}

func normalizeTenant(input string) (*normalizeTenantResp, error) {
	input = strings.ToLower(input)

	var err error
	u := &url.URL{}
	if !strings.Contains(input, "//") {
		if !strings.Contains(input, ".") {
			input += ".conductor.one"
		}
		u.Host = input
	} else {
		u, err = url.Parse(input)
		if err != nil {
			return nil, err
		}
	}

	normalize := &normalizeTenantResp{}

	parts := strings.Split(u.Host, ".")
	if len(parts) == 3 && parts[1] == "conductor" && parts[2] == "one" {
		normalize.useWithTenant = true
		normalize.Tenant = parts[0]

		return normalize, nil
	}

	u.Scheme = "https"
	normalize.useWithServer = true
	normalize.ServerURL = u.String()
	return normalize, nil
}

func WithTenant(input string) (SDKOption, error) {
	resp, err := normalizeTenant(input)
	if err != nil {
		return nil, err
	}

	if resp.useWithTenant {
		return WithTenantDomain(resp.Tenant), nil
	}

	if resp.useWithServer {
		return WithServerURL(resp.ServerURL), nil
	}

	return func(api *ConductoroneAPI) {}, nil
}

type CustomSDKOption func(*CustomOptions)

//func WithClient(client HTTPClient) SDKOption {
//	return func(sdk *ConductoroneAPI) {
//		sdk.sdkConfiguration.DefaultClient = client
//	}
//}

func WithTenantCustom(input string) (CustomSDKOption, error) {
	resp, err := normalizeTenant(input)
	if err != nil {
		return nil, err
	}

	return func(sdk *CustomOptions) {
		sdk.useWithServer = resp.useWithServer
		sdk.useWithTenant = resp.useWithTenant
		sdk.ServerURL = resp.ServerURL
		sdk.Tenant = resp.Tenant
	}, nil
}

func WithLog(logger *zap.Logger) CustomSDKOption {
	return func(sdk *CustomOptions) {
		sdk.logger = logger
	}
}

func WithUserAgent(userAgent string) CustomSDKOption {
	return func(sdk *CustomOptions) {
		sdk.userAgent = userAgent
	}
}
func WithTLSConfig(tlsConfig *tls.Config) CustomSDKOption {
	return func(sdk *CustomOptions) {
		sdk.tlsConfig = tlsConfig
	}
}

type CustomOptions struct {
	useWithServer bool
	useWithTenant bool

	ServerURL string
	Tenant    string

	withClient *http.Client
	logger     *zap.Logger
	tlsConfig  *tls.Config

	userAgent string
}

func NewWithCredentials(ctx context.Context, cred *ClientCredentials, opts ...CustomSDKOption) (*ConductoroneAPI, error) {
	tokenSource, err := NewC1TokenSource(ctx, cred.ClientID, cred.ClientSecret)
	if err != nil {
		return nil, err
	}

	options := &CustomOptions{}
	for _, opt := range opts {
		opt(options)
	}

	if options.userAgent == "" {
		options.userAgent = "conductorone-sdk-go"
	}

	uclient, err := uhttp.NewClient(
		ctx,
		uhttp.WithTokenSource(tokenSource),
		uhttp.WithLogger(options.logger != nil, options.logger),
		uhttp.WithUserAgent(options.userAgent),
		uhttp.WithTLSClientConfig(options.tlsConfig),
	)
	if err != nil {
		return nil, err
	}

	sdkOpts := []SDKOption{}
	if options.useWithServer {
		sdkOpts = append(sdkOpts, WithServerURL(options.ServerURL))
	}
	if options.useWithTenant {
		sdkOpts = append(sdkOpts, WithTenantDomain(options.Tenant))
	}
	sdkOpts = append(sdkOpts, WithClient(uclient))

	return New(sdkOpts...), nil
}
