package uhttp

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"sync"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/oauth2"
)

// NewTransport creates a new Transport, applies the options, and then cycles the transport.
func NewTransport(ctx context.Context, options ...Option) (*Transport, error) {
	t := newTransport()
	for _, opt := range options {
		opt.Apply(t)
	}
	_, err := t.cycle(ctx)
	if err != nil {
		return nil, err
	}
	return t, nil
}

type Transport struct {
	userAgent       string
	tokenSource     oauth2.TokenSource
	tlsClientConfig *tls.Config
	roundTripper    http.RoundTripper
	logger          *zap.Logger
	log             bool
	nextCycle       time.Time
	mtx             sync.RWMutex
	debug           bool
}

func newTransport() *Transport {
	return &Transport{
		tlsClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}
}

func (t *Transport) cycle(ctx context.Context) (http.RoundTripper, error) {
	n := time.Now()
	t.mtx.RLock()
	if t.nextCycle.After(n) && t.roundTripper != nil {
		defer t.mtx.RUnlock()
		return t.roundTripper, nil
	}
	t.mtx.RUnlock()

	t.mtx.Lock()
	defer t.mtx.Unlock()
	n = time.Now()
	// other goroutine changed it under us
	if t.nextCycle.After(n) && t.roundTripper != nil {
		return t.roundTripper, nil
	}
	var err error
	t.roundTripper, err = t.make(ctx)
	if err != nil {
		return nil, err
	}
	t.nextCycle = n.Add(time.Minute * 5)
	return t.roundTripper, nil
}

type userAgentTripper struct {
	next      http.RoundTripper
	userAgent string
}

func (uat *userAgentTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Header.Get("User-Agent") == "" {
		req.Header.Set("User-Agent", uat.userAgent)
	}
	return uat.next.RoundTrip(req)
}

type debugTripper struct {
	next  http.RoundTripper
	debug bool
}

func (uat *debugTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if !uat.debug {
		return uat.next.RoundTrip(req)
	}

	requestBytes, err := httputil.DumpRequest(req, true)
	if err != nil {
		return nil, err
	}
	//nolint:forbidigo
	fmt.Println(string(requestBytes))

	resp, err := uat.next.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	responseBytes, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return nil, err
	}
	//nolint:forbidigo
	fmt.Println(string(responseBytes))

	return resp, nil
}

type tokenSourceTripper struct {
	next        http.RoundTripper
	tokenSource oauth2.TokenSource
}

func (uts *tokenSourceTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if uts.tokenSource == nil {
		return uts.next.RoundTrip(req)
	}
	token, err := uts.tokenSource.Token()
	if err != nil {
		return nil, err
	}
	token.SetAuthHeader(req)
	return uts.next.RoundTrip(req)
}

func (t *Transport) make(ctx context.Context) (http.RoundTripper, error) {
	// based on http.DefaultTransport
	baseTransport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       t.tlsClientConfig,
	}
	err := http2.ConfigureTransport(baseTransport)
	if err != nil {
		return nil, err
	}
	var rv http.RoundTripper = baseTransport
	t.userAgent = fmt.Sprintf("%s cone", t.userAgent)
	rv = &debugTripper{next: rv, debug: t.debug}
	rv = &userAgentTripper{next: rv, userAgent: t.userAgent}
	rv = &tokenSourceTripper{next: rv, tokenSource: t.tokenSource}
	return rv, nil
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()
	rt, err := t.cycle(ctx)
	if err != nil {
		return nil, fmt.Errorf("uhttp: cycle failed: %w", err)
	}
	if t.log {
		t.l(ctx).Debug("Request started",
			zap.String("http.method", req.Method),
			zap.String("http.url_details.host", req.URL.Host),
			zap.String("http.url_details.path", req.URL.Path),
		)
	}
	resp, err := rt.RoundTrip(req)
	if t.log {
		fields := []zap.Field{zap.String("http.method", req.Method),
			zap.String("http.url_details.host", req.URL.Host),
			zap.String("http.url_details.path", req.URL.Path),
		}

		if err != nil {
			fields = append(fields, zap.Error(err))
		}

		if resp != nil {
			fields = append(fields, zap.Int("http.status_code", resp.StatusCode))
		}

		t.l(ctx).Debug("Request complete", fields...)
	}
	return resp, err
}

func (t *Transport) l(ctx context.Context) *zap.Logger {
	if t.logger != nil {
		return t.logger
	}
	return ctxzap.Extract(ctx)
}
