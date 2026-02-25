package uhttp

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"net/http"
	"net/http/httputil"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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

const (
	defaultMaxRetries = 5
	maxRetryWait      = 30 * time.Second
)

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
	maxRetries      int
}

func newTransport() *Transport {
	return &Transport{
		tlsClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
		maxRetries: defaultMaxRetries,
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
	transportID := uuid.New().String()
	if !uat.debug {
		return uat.next.RoundTrip(req)
	}

	ctx := tflog.NewSubsystem(req.Context(), subsystem)
	requestBytes, err := httputil.DumpRequest(req, true)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("[ERROR] Conductor One API Request error: %#v", err))
		return nil, err
	} else {
		tflog.Debug(ctx, fmt.Sprintf("[DEBUG] "+logReqMsg, transportID, prettyPrintJsonLines(requestBytes)))
	}

	resp, err := uat.next.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	responseBytes, err := httputil.DumpResponse(resp, true)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("[ERROR] Conductor One API Response error: %#v", err))
		return nil, err
	} else {
		tflog.Debug(ctx, fmt.Sprintf("[DEBUG] "+logRespMsg, transportID, prettyPrintJsonLines(responseBytes)))
	}

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

// retryTripper retries requests that receive HTTP 429 (Too Many Requests).
// It respects the Retry-After header if present, falling back to exponential
// backoff. Context cancellation is honored during the wait.
type retryTripper struct {
	next       http.RoundTripper
	maxRetries int
}

func (rt *retryTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	for attempt := 0; ; attempt++ {
		// Re-create the request body for retries.
		if attempt > 0 && req.GetBody != nil {
			body, err := req.GetBody()
			if err != nil {
				return nil, err
			}
			req.Body = body
		}

		resp, err := rt.next.RoundTrip(req)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != http.StatusTooManyRequests || attempt >= rt.maxRetries {
			return resp, nil
		}

		wait := retryAfterDuration(resp)
		if wait <= 0 {
			// Exponential backoff: 1s, 2s, 4s, 8s, 16s, capped at 30s.
			wait = time.Second << uint(attempt)
			if wait > maxRetryWait {
				wait = maxRetryWait
			}
		}

		resp.Body.Close()

		select {
		case <-req.Context().Done():
			return nil, req.Context().Err()
		case <-time.After(wait):
		}
	}
}

// retryAfterDuration parses the Retry-After header from an HTTP response.
// Returns 0 if the header is missing or unparseable.
func retryAfterDuration(resp *http.Response) time.Duration {
	val := resp.Header.Get("Retry-After")
	if val == "" {
		return 0
	}

	// Try as seconds first.
	if seconds, err := strconv.ParseInt(val, 10, 64); err == nil {
		if seconds > 0 {
			return time.Duration(seconds) * time.Second
		}
		return 0
	}

	// Try as HTTP-date (RFC 1123).
	if t, err := time.Parse(time.RFC1123, val); err == nil {
		delta := time.Until(t)
		if delta > 0 {
			return delta
		}
	}

	return 0
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
	rv = &retryTripper{next: rv, maxRetries: t.maxRetries}
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

// prettyPrintJsonLines processes each line of the input.
// If the line is JSON, it will pretty-print it.
// If the line contains an Authorization header, it will redact its value.
func prettyPrintJsonLines(b []byte) string {
	parts := strings.Split(string(b), "\n")
	for i, p := range parts {
		if strings.HasPrefix(strings.ToLower(p), "authorization:") {
			parts[i] = "Authorization: (sensitive)"
			continue
		}

		if json.Valid([]byte(p)) {
			var out bytes.Buffer
			if err := json.Indent(&out, []byte(p), "", "  "); err == nil {
				parts[i] = out.String()
			}
		}
	}
	return strings.Join(parts, "\n")
}

const logReqMsg = `[%s] Conductor One API Request Details:
---[ REQUEST ]---------------------------------------
%s
-----------------------------------------------------`

const logRespMsg = `[%s] Conductor One API Response Details:
---[ RESPONSE ]--------------------------------------
%s
-----------------------------------------------------`
