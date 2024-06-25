package sdk

import (
	"context"
	"crypto/tls"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"go.uber.org/zap"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/uhttp"
)

const c1TenantDomain = ".conductor.one"
const ClientIdGolangSDK = "2RCzHlak5q7CY14SdBc8HoZEJRf"

func WithTenant(input string) (SDKOption, error) {
	resp, err := NormalizeTenant(input)
	if err != nil {
		return nil, err
	}

	if resp.UseWithTenant() {
		return WithTenantDomain(resp.Tenant()), nil
	}

	if resp.UseWithServer() {
		return WithServerURL(resp.ServerURL()), nil
	}

	return func(api *ConductoroneAPI) {}, nil
}

type CustomSDKOption func(*CustomOptions)

func WithTenantCustom(input string) (CustomSDKOption, error) {
	resp, err := NormalizeTenant(input)
	if err != nil {
		return nil, err
	}

	return func(sdk *CustomOptions) {
		sdk.ClientConfig = resp
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

type ClientConfig struct {
	// These are mutually exclusive
	serverURL string
	tenant    string
}

func (c *ClientConfig) UseWithServer() bool {
	return c.serverURL != ""
}

func (c *ClientConfig) UseWithTenant() bool {
	return c.tenant != ""
}

func (c *ClientConfig) SetTenant(tenant string) error {
	if c.UseWithServer() {
		return errors.New("cannot set tenant, tenant and serverURL are mutually exclusive")
	}
	c.tenant = tenant
	return nil
}

func (c *ClientConfig) SetServerURL(serverURL string) error {
	if c.UseWithTenant() {
		return errors.New("cannot set serverURL, tenant and serverURL are mutually exclusive")
	}
	c.serverURL = serverURL
	return nil
}

func (c *ClientConfig) Tenant() string {
	return c.tenant
}

// ServerURL returns the server URL.
func (c *ClientConfig) ServerURL() string {
	return c.serverURL
}

// GetServerURL returns the server URL. If serverURL is empty (""), it constructs the server URL using the tenant. However, if the tenant is also empty, then it will return an empty string.
func (c *ClientConfig) GetServerURL() string {
	if c.UseWithServer() {
		return c.serverURL
	}
	if c.UseWithTenant() {
		u := &url.URL{}
		tenant := strings.ToLower(c.Tenant())
		u.Host = tenant + c1TenantDomain
		u.Scheme = "https"
		return u.String()
	}
	return ""
}

type CustomOptions struct {
	*ClientConfig

	// nolint:unused
	withClient *http.Client
	logger     *zap.Logger
	tlsConfig  *tls.Config

	userAgent string
}

func NewWithCredentials(ctx context.Context, cred *ClientCredentials, opts ...CustomSDKOption) (*ConductoroneAPI, error) {
	options := &CustomOptions{}
	for _, opt := range opts {
		opt(options)
	}
	if options.GetServerURL() == "" {
		resp, err := ParseClientID(cred.ClientID)
		if err != nil {
			return nil, err
		}
		options.ClientConfig = resp
	}

	tokenSource, err := NewTokenSource(ctx, cred.ClientID, cred.ClientSecret, options.GetServerURL())
	if err != nil {
		return nil, err
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
	if options.UseWithServer() {
		sdkOpts = append(sdkOpts, WithServerURL(options.ServerURL()))
	}
	if options.UseWithTenant() {
		sdkOpts = append(sdkOpts, WithTenantDomain(options.Tenant()))
	}
	sdkOpts = append(sdkOpts, WithClient(uclient))

	return New(sdkOpts...), nil
}

func NormalizeTenant(input string) (*ClientConfig, error) {
	input = strings.ToLower(input)

	var err error
	u := &url.URL{}
	if !strings.Contains(input, "://") {
		if !strings.Contains(input, ".") {
			input += c1TenantDomain
		}
		u.Host = input
	} else {
		u, err = url.Parse(input)
		if err != nil {
			return nil, err
		}
	}

	normalize := &ClientConfig{}

	parts := strings.Split(u.Host, ".")
	if len(parts) == 3 && parts[1] == "conductor" && parts[2] == "one" {
		err := normalize.SetTenant(parts[0])
		if err != nil {
			return nil, err
		}
		return normalize, nil
	}

	u.Scheme = "https"
	err = normalize.SetServerURL(u.String())
	if err != nil {
		return nil, err
	}
	return normalize, nil
}

func ParseClientID(input string) (*ClientConfig, error) {
	// split the input into 2 parts by @
	items := strings.SplitN(input, "@", 2)
	if len(items) != 2 {
		return nil, ErrInvalidClientID
	}

	// split the right part into 2 parts by /
	items = strings.SplitN(items[1], "/", 2)
	if len(items) != 2 {
		return nil, ErrInvalidClientID
	}

	resp, err := NormalizeTenant(items[0])
	if err != nil {
		return nil, err
	}

	if resp.GetServerURL() == "" {
		return nil, ErrInvalidClientID
	}

	return resp, nil
}
