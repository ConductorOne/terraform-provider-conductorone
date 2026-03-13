package uhttp

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
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

func TestRetryTripperRetriesOn429(t *testing.T) {
	mock := &mockTripper{
		responses: []*http.Response{resp429("0"), resp200()},
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

func TestRetryTripperStopsAtMaxRetries(t *testing.T) {
	mock := &mockTripper{
		responses: []*http.Response{resp429("0"), resp429("0"), resp429("0")},
	}
	rt := &retryTripper{next: mock, maxRetries: 2}

	req, _ := http.NewRequestWithContext(context.Background(), "GET", "http://example.com", nil)
	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// After 2 retries (3 total attempts), should return the 429.
	if resp.StatusCode != http.StatusTooManyRequests {
		t.Errorf("expected 429 after max retries, got %d", resp.StatusCode)
	}
}

func TestRetryTripperPassesThroughNon429(t *testing.T) {
	mock := &mockTripper{
		responses: []*http.Response{{
			StatusCode: http.StatusInternalServerError,
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
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", resp.StatusCode)
	}
	if mock.calls != 1 {
		t.Errorf("expected 1 call (no retry for 500), got %d", mock.calls)
	}
}

func TestRetryTripperRespectsContextCancellation(t *testing.T) {
	mock := &mockTripper{
		responses: []*http.Response{resp429("10"), resp200()},
	}
	rt := &retryTripper{next: mock, maxRetries: 3}

	ctx, cancel := context.WithCancel(context.Background())
	req, _ := http.NewRequestWithContext(ctx, "GET", "http://example.com", nil)

	// Cancel immediately so the wait is interrupted.
	cancel()

	_, err := rt.RoundTrip(req)
	if err == nil {
		t.Fatal("expected context cancellation error, got nil")
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
