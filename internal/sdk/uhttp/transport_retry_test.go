package uhttp

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
)

type mockTripper struct {
	responses []*http.Response
	calls     int
}

func (m *mockTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if m.calls >= len(m.responses) {
		return m.responses[len(m.responses)-1], nil
	}
	resp := m.responses[m.calls]
	m.calls++
	return resp, nil
}

func resp429(retryAfter string) *http.Response {
	h := http.Header{}
	if retryAfter != "" {
		h.Set("Retry-After", retryAfter)
	}
	return &http.Response{
		StatusCode: http.StatusTooManyRequests,
		Header:     h,
		Body:       io.NopCloser(strings.NewReader("")),
	}
}

func resp200() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{},
		Body:       io.NopCloser(strings.NewReader("ok")),
	}
}

type errorThenSuccessTripper struct {
	calls int
}

func (m *errorThenSuccessTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	m.calls++
	if m.calls == 1 {
		return nil, errors.New("connection reset")
	}
	return resp200(), nil
}

type bodyRecordingTripper struct {
	calls  int
	bodies []string
}

func (m *bodyRecordingTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	m.calls++
	m.bodies = append(m.bodies, string(body))
	if m.calls == 1 {
		return &http.Response{
			StatusCode: http.StatusServiceUnavailable,
			Header:     http.Header{},
			Body:       io.NopCloser(strings.NewReader("")),
		}, nil
	}
	return resp200(), nil
}

type unexpectedEOFBody struct{}

func (unexpectedEOFBody) Read([]byte) (int, error) {
	return 0, io.ErrUnexpectedEOF
}

func (unexpectedEOFBody) Close() error {
	return nil
}

func TestRetryTripperRetriesOn429(t *testing.T) {
	mock := &mockTripper{
		responses: []*http.Response{resp429("0"), resp200()},
	}
	rt := &retryTripper{next: mock, maxRetries: 3, rateLimitBudget: time.Minute}

	req, _ := http.NewRequestWithContext(context.Background(), "GET", "http://example.com", nil)
	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	if mock.calls != 2 {
		t.Errorf("expected 2 calls, got %d", mock.calls)
	}
}

func resp503() *http.Response {
	return &http.Response{
		StatusCode: http.StatusServiceUnavailable,
		Header:     http.Header{},
		Body:       io.NopCloser(strings.NewReader("")),
	}
}

func TestRetryTripperStopsAtMaxRetries(t *testing.T) {
	mock := &mockTripper{
		responses: []*http.Response{resp503(), resp503(), resp503()},
	}
	rt := &retryTripper{next: mock, maxRetries: 1, rateLimitBudget: time.Minute}

	req, _ := http.NewRequestWithContext(context.Background(), "GET", "http://example.com", nil)
	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Errorf("expected 503 after max retries, got %d", resp.StatusCode)
	}
	if mock.calls != 2 {
		t.Errorf("expected 2 calls, got %d", mock.calls)
	}
}

// A throttled Terraform run can need more 429 retries than maxRetries allows
// for server errors; 429s are bounded by rateLimitBudget instead.
func TestRetryTripper429IsNotLimitedByMaxRetries(t *testing.T) {
	mock := &mockTripper{
		responses: []*http.Response{resp429("1"), resp200()},
	}
	rt := &retryTripper{next: mock, maxRetries: 0, rateLimitBudget: time.Minute}

	req, _ := http.NewRequestWithContext(context.Background(), "GET", "http://example.com", nil)
	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	if mock.calls != 2 {
		t.Errorf("expected 2 calls, got %d", mock.calls)
	}
}

func TestRetryTripperStops429AtBudget(t *testing.T) {
	mock := &mockTripper{
		responses: []*http.Response{resp429("1"), resp429("1"), resp429("1"), resp200()},
	}
	// The first 1s wait (plus up to 0.5s jitter) fits; a second would not.
	rt := &retryTripper{next: mock, maxRetries: 3, rateLimitBudget: 1500 * time.Millisecond}

	req, _ := http.NewRequestWithContext(context.Background(), "GET", "http://example.com", nil)
	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusTooManyRequests {
		t.Errorf("expected 429 once the budget is spent, got %d", resp.StatusCode)
	}
	if mock.calls != 2 {
		t.Errorf("expected 2 calls, got %d", mock.calls)
	}
}

func TestRetryTripperReturns429WithoutWaitingPastLimits(t *testing.T) {
	tests := []struct {
		name    string
		budget  time.Duration
		timeout time.Duration
	}{
		{"retry-after exceeds budget", time.Minute, 0},
		{"retry-after exceeds context deadline", time.Hour, 10 * time.Second},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockTripper{
				responses: []*http.Response{resp429("120"), resp200()},
			}
			rt := &retryTripper{next: mock, maxRetries: 3, rateLimitBudget: tt.budget}

			ctx := context.Background()
			if tt.timeout > 0 {
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(ctx, tt.timeout)
				defer cancel()
			}
			req, _ := http.NewRequestWithContext(ctx, "GET", "http://example.com", nil)
			start := time.Now()
			resp, err := rt.RoundTrip(req)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if resp.StatusCode != http.StatusTooManyRequests {
				t.Errorf("expected 429, got %d", resp.StatusCode)
			}
			if mock.calls != 1 {
				t.Errorf("expected 1 call, got %d", mock.calls)
			}
			if elapsed := time.Since(start); elapsed > time.Second {
				t.Errorf("expected an immediate return, took %v", elapsed)
			}
		})
	}
}

// The budget and deadline must bound the wait actually slept, jitter
// included. These run on synctest's virtual clock, so elapsed time is exact.
func TestRetryTripperJitteredWaitStaysWithinLimits(t *testing.T) {
	t.Run("budget", func(t *testing.T) {
		for range 50 {
			synctest.Test(t, func(t *testing.T) {
				mock := &mockTripper{responses: []*http.Response{resp429("1")}}
				rt := &retryTripper{next: mock, maxRetries: 3, rateLimitBudget: 1200 * time.Millisecond}

				req, _ := http.NewRequestWithContext(context.Background(), "GET", "http://example.com", nil)
				start := time.Now()
				resp, err := rt.RoundTrip(req)
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if resp.StatusCode != http.StatusTooManyRequests {
					t.Fatalf("expected 429, got %d", resp.StatusCode)
				}
				if elapsed := time.Since(start); elapsed > rt.rateLimitBudget {
					t.Fatalf("waited %v, past the %v budget", elapsed, rt.rateLimitBudget)
				}
			})
		}
	})
	t.Run("context deadline", func(t *testing.T) {
		for range 50 {
			synctest.Test(t, func(t *testing.T) {
				mock := &mockTripper{responses: []*http.Response{resp429("1")}}
				rt := &retryTripper{next: mock, maxRetries: 3, rateLimitBudget: time.Hour}

				ctx, cancel := context.WithTimeout(context.Background(), 1100*time.Millisecond)
				defer cancel()
				req, _ := http.NewRequestWithContext(ctx, "GET", "http://example.com", nil)
				resp, err := rt.RoundTrip(req)
				if err != nil {
					t.Fatalf("expected the 429 rather than a context error, got %v", err)
				}
				if resp.StatusCode != http.StatusTooManyRequests {
					t.Fatalf("expected 429, got %d", resp.StatusCode)
				}
			})
		}
	})
}

func TestBackoff(t *testing.T) {
	want := []time.Duration{time.Second, 2 * time.Second, 4 * time.Second, 8 * time.Second, 16 * time.Second, maxRetryWait, maxRetryWait}
	for attempt, w := range want {
		if got := backoff(attempt); got != w {
			t.Errorf("backoff(%d) = %v, want %v", attempt, got, w)
		}
	}
	if got := backoff(100); got != maxRetryWait {
		t.Errorf("backoff(100) = %v, want %v", got, maxRetryWait)
	}
}

func TestRetryTripperRetriesOn5xx(t *testing.T) {
	mock := &mockTripper{
		responses: []*http.Response{
			{
				StatusCode: http.StatusServiceUnavailable,
				Header:     http.Header{},
				Body:       io.NopCloser(strings.NewReader("")),
			},
			resp200(),
		},
	}
	rt := &retryTripper{next: mock, maxRetries: 3}

	req, _ := http.NewRequestWithContext(context.Background(), "GET", "http://example.com", nil)
	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	if mock.calls != 2 {
		t.Errorf("expected 2 calls, got %d", mock.calls)
	}
}

func TestRetryTripperRetriesWhenResponseDrainFails(t *testing.T) {
	mock := &mockTripper{
		responses: []*http.Response{
			{
				StatusCode: http.StatusServiceUnavailable,
				Header:     http.Header{},
				Body:       unexpectedEOFBody{},
			},
			resp200(),
		},
	}
	rt := &retryTripper{next: mock, maxRetries: 3}

	req, _ := http.NewRequestWithContext(context.Background(), "GET", "http://example.com", nil)
	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	if mock.calls != 2 {
		t.Errorf("expected 2 calls, got %d", mock.calls)
	}
}

func TestRetryTripperPassesThroughNonRetryableStatus(t *testing.T) {
	mock := &mockTripper{
		responses: []*http.Response{{
			StatusCode: http.StatusNotFound,
			Header:     http.Header{},
			Body:       io.NopCloser(strings.NewReader("")),
		}},
	}
	rt := &retryTripper{next: mock, maxRetries: 3}

	req, _ := http.NewRequestWithContext(context.Background(), "GET", "http://example.com", nil)
	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("expected 404, got %d", resp.StatusCode)
	}
	if mock.calls != 1 {
		t.Errorf("expected 1 call for 404, got %d", mock.calls)
	}
}

func TestRetryTripperRetriesTransportErrors(t *testing.T) {
	mock := &errorThenSuccessTripper{}
	rt := &retryTripper{next: mock, maxRetries: 3}

	req, _ := http.NewRequestWithContext(context.Background(), "GET", "http://example.com", nil)
	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	if mock.calls != 2 {
		t.Errorf("expected 2 calls, got %d", mock.calls)
	}
}

func TestRetryTripperRetriesNonReplayableRequestBody(t *testing.T) {
	mock := &bodyRecordingTripper{}
	rt := &retryTripper{next: mock, maxRetries: 3}

	req, _ := http.NewRequestWithContext(
		context.Background(),
		http.MethodPut,
		"http://example.com",
		io.NopCloser(strings.NewReader("payload")),
	)
	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	if req.GetBody == nil {
		t.Fatal("expected request body to be replayable")
	}
	if got, want := strings.Join(mock.bodies, ","), "payload,payload"; got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestRetryTripperDoesNotRetryUnsafeRequests(t *testing.T) {
	mock := &mockTripper{
		responses: []*http.Response{
			{
				StatusCode: http.StatusServiceUnavailable,
				Header:     http.Header{},
				Body:       io.NopCloser(strings.NewReader("")),
			},
			resp200(),
		},
	}
	rt := &retryTripper{next: mock, maxRetries: 3}

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "http://example.com", nil)
	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Errorf("expected 503, got %d", resp.StatusCode)
	}
	if mock.calls != 1 {
		t.Errorf("expected 1 call for POST, got %d", mock.calls)
	}
}

func TestRetryTripperRetriesUnsafeRequestsOn429(t *testing.T) {
	for _, method := range []string{http.MethodPost, http.MethodPatch} {
		t.Run(method, func(t *testing.T) {
			var bodies []string
			mock := roundTripFunc(func(req *http.Request) (*http.Response, error) {
				body, err := io.ReadAll(req.Body)
				if err != nil {
					return nil, err
				}
				bodies = append(bodies, string(body))
				if len(bodies) == 1 {
					return resp429("1"), nil
				}
				return resp200(), nil
			})
			rt := &retryTripper{next: mock, maxRetries: 3, rateLimitBudget: time.Minute}

			req, _ := http.NewRequestWithContext(context.Background(), method, "http://example.com", nil)
			req.Body = io.NopCloser(strings.NewReader(`{"query":"x"}`))
			resp, err := rt.RoundTrip(req)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if resp.StatusCode != http.StatusOK {
				t.Errorf("expected 200, got %d", resp.StatusCode)
			}
			if len(bodies) != 2 || bodies[0] != bodies[1] || bodies[1] != `{"query":"x"}` {
				t.Errorf("expected the same body on both attempts, got %q", bodies)
			}
		})
	}
}

// The provider wraps the transport in Terraform's logging transport, whose
// httputil.DumpRequestOut call replaces a nil GET body with http.NoBody and
// leaves GetBody unset.
func TestRetryTripperRetriesGetBehindLoggingTransport(t *testing.T) {
	var hits int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.AddInt32(&hits, 1) == 1 {
			w.Header().Set("Retry-After", "1")
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	client := &http.Client{
		Transport: logging.NewSubsystemLoggingHTTPTransport(subsystem, &retryTripper{next: http.DefaultTransport, maxRetries: 3, rateLimitBudget: time.Minute}),
	}
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, srv.URL, nil)
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	if got := atomic.LoadInt32(&hits); got != 2 {
		t.Errorf("expected 2 requests, got %d", got)
	}
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestRetryTripperRespectsContextCancellation(t *testing.T) {
	mock := &mockTripper{
		responses: []*http.Response{resp429("10"), resp200()},
	}
	rt := &retryTripper{next: mock, maxRetries: 3, rateLimitBudget: time.Minute}

	ctx, cancel := context.WithCancel(context.Background())
	req, _ := http.NewRequestWithContext(ctx, "GET", "http://example.com", nil)

	// Cancel immediately so the wait is interrupted.
	cancel()

	_, err := rt.RoundTrip(req)
	if err == nil {
		t.Fatal("expected context cancellation error, got nil")
	}
}

type closeRecordingBody struct {
	io.Reader
	closed bool
}

func (b *closeRecordingBody) Close() error {
	b.closed = true
	return nil
}

// A retryable response discarded because the request was cancelled must still
// be closed, or its connection leaks.
func TestRetryTripperClosesResponseWhenContextCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	body := &closeRecordingBody{Reader: strings.NewReader("")}
	next := roundTripFunc(func(*http.Request) (*http.Response, error) {
		cancel()
		return &http.Response{StatusCode: http.StatusTooManyRequests, Header: http.Header{}, Body: body}, nil
	})
	rt := &retryTripper{next: next, maxRetries: 3, rateLimitBudget: time.Minute}

	req, _ := http.NewRequestWithContext(ctx, "GET", "http://example.com", nil)
	resp, err := rt.RoundTrip(req)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context.Canceled, got resp=%v err=%v", resp, err)
	}
	if !body.closed {
		t.Error("expected the discarded response body to be closed")
	}
}

func TestWithJitterStaysWithinBounds(t *testing.T) {
	for _, wait := range []time.Duration{0, time.Nanosecond, time.Second, maxRetryWait} {
		for range 1000 {
			got := withJitter(wait)
			if got < wait || got > wait+wait/2 {
				t.Fatalf("withJitter(%v) = %v, want within [%v, %v]", wait, got, wait, wait+wait/2)
			}
		}
	}
}

func TestRetryAfterDuration(t *testing.T) {
	tests := []struct {
		name     string
		header   string
		wantZero bool
	}{
		{"empty header", "", true},
		{"seconds", "2", false},
		{"zero seconds", "0", true},
		{"negative seconds", "-1", true},
		{"garbage", "not-a-number", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{Header: http.Header{}}
			if tt.header != "" {
				resp.Header.Set("Retry-After", tt.header)
			}
			d := retryAfterDuration(resp)
			if tt.wantZero && d != 0 {
				t.Errorf("expected 0, got %v", d)
			}
			if !tt.wantZero && d == 0 {
				t.Errorf("expected non-zero duration for header %q", tt.header)
			}
		})
	}
}

func TestRetryAfterDurationSeconds(t *testing.T) {
	resp := &http.Response{Header: http.Header{}}
	resp.Header.Set("Retry-After", strconv.Itoa(5))
	d := retryAfterDuration(resp)
	if d != 5*time.Second {
		t.Errorf("expected 5s, got %v", d)
	}
}
