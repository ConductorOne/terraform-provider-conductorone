// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"context"
	"fmt"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/utils"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	// The ConductorOne API server for the current tenant.
	"https://{tenantDomain}.conductor.one",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	ServerDefaults    []map[string]string
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], c.ServerDefaults[c.ServerIndex]
}

// SDK - ConductorOne API: The ConductorOne API is a HTTP API for managing ConductorOne resources.
type SDK struct {
	Apps                      *Apps
	Connector                 *Connector
	AppEntitlements           *AppEntitlements
	AppEntitlementUserBinding *AppEntitlementUserBinding
	AppEntitlementOwners      *AppEntitlementOwners
	AppOwners                 *AppOwners
	AppReport                 *AppReport
	AppReportAction           *AppReportAction
	AppResourceType           *AppResourceType
	AppResource               *AppResource
	AppResourceOwners         *AppResourceOwners
	AppUsageControls          *AppUsageControls
	AppUser                   *AppUser
	Attributes                *Attributes
	Auth                      *Auth
	RequestCatalogManagement  *RequestCatalogManagement
	Directory                 *Directory
	PersonalClient            *PersonalClient
	Roles                     *Roles
	Policies                  *Policies
	PolicyValidate            *PolicyValidate
	AppResourceSearch         *AppResourceSearch
	AppSearch                 *AppSearch
	AttributeSearch           *AttributeSearch
	AppEntitlementSearch      *AppEntitlementSearch
	PolicySearch              *PolicySearch
	RequestCatalogSearch      *RequestCatalogSearch
	TaskSearch                *TaskSearch
	UserSearch                *UserSearch
	Task                      *Task
	TaskActions               *TaskActions
	User                      *User

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*SDK)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *SDK) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *SDK) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithTenantDomain allows setting the tenantDomain variable for url substitution
func WithTenantDomain(tenantDomain string) SDKOption {
	return func(sdk *SDK) {
		for idx := range sdk.sdkConfiguration.ServerDefaults {
			if _, ok := sdk.sdkConfiguration.ServerDefaults[idx]["tenantDomain"]; !ok {
				continue
			}

			sdk.sdkConfiguration.ServerDefaults[idx]["tenantDomain"] = fmt.Sprintf("%v", tenantDomain)
		}
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

func withSecurity(security interface{}) func(context.Context) (interface{}, error) {
	return func(context.Context) (interface{}, error) {
		return &security, nil
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.Security = withSecurity(security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *SDK {
	sdk := &SDK{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "0.1.0-alpha",
			SDKVersion:        "0.0.1",
			GenVersion:        "2.210.6",
			UserAgent:         "speakeasy-sdk/go 0.0.1 2.210.6 0.1.0-alpha terraform",
			ServerDefaults: []map[string]string{
				{
					"tenantDomain": "example",
				},
			},
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.Apps = newApps(sdk.sdkConfiguration)

	sdk.Connector = newConnector(sdk.sdkConfiguration)

	sdk.AppEntitlements = newAppEntitlements(sdk.sdkConfiguration)

	sdk.AppEntitlementUserBinding = newAppEntitlementUserBinding(sdk.sdkConfiguration)

	sdk.AppEntitlementOwners = newAppEntitlementOwners(sdk.sdkConfiguration)

	sdk.AppOwners = newAppOwners(sdk.sdkConfiguration)

	sdk.AppReport = newAppReport(sdk.sdkConfiguration)

	sdk.AppReportAction = newAppReportAction(sdk.sdkConfiguration)

	sdk.AppResourceType = newAppResourceType(sdk.sdkConfiguration)

	sdk.AppResource = newAppResource(sdk.sdkConfiguration)

	sdk.AppResourceOwners = newAppResourceOwners(sdk.sdkConfiguration)

	sdk.AppUsageControls = newAppUsageControls(sdk.sdkConfiguration)

	sdk.AppUser = newAppUser(sdk.sdkConfiguration)

	sdk.Attributes = newAttributes(sdk.sdkConfiguration)

	sdk.Auth = newAuth(sdk.sdkConfiguration)

	sdk.RequestCatalogManagement = newRequestCatalogManagement(sdk.sdkConfiguration)

	sdk.Directory = newDirectory(sdk.sdkConfiguration)

	sdk.PersonalClient = newPersonalClient(sdk.sdkConfiguration)

	sdk.Roles = newRoles(sdk.sdkConfiguration)

	sdk.Policies = newPolicies(sdk.sdkConfiguration)

	sdk.PolicyValidate = newPolicyValidate(sdk.sdkConfiguration)

	sdk.AppResourceSearch = newAppResourceSearch(sdk.sdkConfiguration)

	sdk.AppSearch = newAppSearch(sdk.sdkConfiguration)

	sdk.AttributeSearch = newAttributeSearch(sdk.sdkConfiguration)

	sdk.AppEntitlementSearch = newAppEntitlementSearch(sdk.sdkConfiguration)

	sdk.PolicySearch = newPolicySearch(sdk.sdkConfiguration)

	sdk.RequestCatalogSearch = newRequestCatalogSearch(sdk.sdkConfiguration)

	sdk.TaskSearch = newTaskSearch(sdk.sdkConfiguration)

	sdk.UserSearch = newUserSearch(sdk.sdkConfiguration)

	sdk.Task = newTask(sdk.sdkConfiguration)

	sdk.TaskActions = newTaskActions(sdk.sdkConfiguration)

	sdk.User = newUser(sdk.sdkConfiguration)

	return sdk
}