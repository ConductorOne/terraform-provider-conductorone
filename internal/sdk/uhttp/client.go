package uhttp

import (
	"context"
	"crypto/tls"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

const subsystem = "terraform-provider-conductorone"

type tlsClientConfigOption struct {
	config *tls.Config
}

func (o tlsClientConfigOption) Apply(c *Transport) {
	c.tlsClientConfig = o.config
}

// WithTLSClientConfig returns an Option that sets the TLS client configuration.
// `tlsConfig` is a structure that is used to configure a TLS client or server.
func WithTLSClientConfig(tlsConfig *tls.Config) Option {
	return tlsClientConfigOption{config: tlsConfig}
}

type loggerOption struct {
	log    bool
	logger *zap.Logger
}

func (o loggerOption) Apply(c *Transport) {
	c.log = o.log
	c.logger = o.logger
}

// WithLogger sets a logger options to the transport layer.
func WithLogger(log bool, logger *zap.Logger) Option {
	return loggerOption{
		log:    log,
		logger: logger,
	}
}

type tokenSourceOption struct {
	tokenSource oauth2.TokenSource
}

func (t tokenSourceOption) Apply(c *Transport) {
	c.tokenSource = t.tokenSource
}

// WithTokenSource sets a token source option to the transport layer.
func WithTokenSource(tokenSource oauth2.TokenSource) Option {
	return tokenSourceOption{
		tokenSource: tokenSource,
	}
}

type userAgentOption struct {
	userAgent string
}

func (o userAgentOption) Apply(c *Transport) {
	c.userAgent = o.userAgent
}

// WithUserAgent sets a user agent option to the transport layer.
func WithUserAgent(userAgent string) Option {
	return userAgentOption{
		userAgent: userAgent,
	}
}

type debugOption struct {
	debug bool
}

func (o debugOption) Apply(c *Transport) {
	c.debug = o.debug
}

func WithDebug(debug bool) Option {
	return debugOption{
		debug: debug,
	}
}

type Option interface {
	Apply(*Transport)
}

// NewClient creates a new HTTP client that uses the given context and options to create a new transport layer.
func NewClient(ctx context.Context, options ...Option) (*http.Client, error) {
	t, err := NewTransport(ctx, options...)
	if err != nil {
		return nil, err
	}
	return &http.Client{
		Transport: logging.NewSubsystemLoggingHTTPTransport(subsystem, t),
	}, nil
}
