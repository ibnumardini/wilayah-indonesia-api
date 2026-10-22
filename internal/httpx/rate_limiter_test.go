package httpx

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func noContentHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
}

func TestRateLimiterLimitsRequestsPerClient(t *testing.T) {
	limiter := NewRateLimiter(60, 2)
	handler := limiter.Handler(noContentHandler())

	for requestNumber := 1; requestNumber <= 2; requestNumber++ {
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.RemoteAddr = "192.0.2.1:1234"
		recorder := httptest.NewRecorder()

		handler.ServeHTTP(recorder, request)

		if recorder.Code != http.StatusNoContent {
			t.Fatalf("request %d returned status %d, want %d", requestNumber, recorder.Code, http.StatusNoContent)
		}
	}

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	request.RemoteAddr = "192.0.2.1:5678"
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusTooManyRequests {
		t.Fatalf("limited request returned status %d, want %d", recorder.Code, http.StatusTooManyRequests)
	}
	if recorder.Header().Get("Retry-After") != "1" {
		t.Fatalf("Retry-After = %q, want 1", recorder.Header().Get("Retry-After"))
	}
}

func TestRateLimiterSeparatesClients(t *testing.T) {
	limiter := NewRateLimiter(60, 1)
	handler := limiter.Handler(noContentHandler())

	for _, ip := range []string{"192.0.2.1:1234", "192.0.2.2:1234"} {
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.RemoteAddr = ip
		recorder := httptest.NewRecorder()

		handler.ServeHTTP(recorder, request)

		if recorder.Code != http.StatusNoContent {
			t.Fatalf("client %s returned status %d, want %d", ip, recorder.Code, http.StatusNoContent)
		}
	}
}

func TestRateLimiterUsesForwardedIPFromTrustedProxy(t *testing.T) {
	limiter := NewRateLimiter(60, 1, "10.0.0.0/8")
	handler := limiter.Handler(noContentHandler())

	forwarded := httptest.NewRequest(http.MethodGet, "/", nil)
	forwarded.RemoteAddr = "10.0.0.5:1234"
	forwarded.Header.Set("X-Forwarded-For", "192.0.2.1")
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, forwarded)
	if recorder.Code != http.StatusNoContent {
		t.Fatalf("trusted proxy request returned status %d, want %d", recorder.Code, http.StatusNoContent)
	}

	sameClient := httptest.NewRequest(http.MethodGet, "/", nil)
	sameClient.RemoteAddr = "10.0.0.6:1234"
	sameClient.Header.Set("X-Forwarded-For", "192.0.2.1")
	recorder = httptest.NewRecorder()
	handler.ServeHTTP(recorder, sameClient)
	if recorder.Code != http.StatusTooManyRequests {
		t.Fatalf("forwarded client request returned status %d, want %d", recorder.Code, http.StatusTooManyRequests)
	}
}

func TestRateLimiterIgnoresForwardedIPFromUntrustedClient(t *testing.T) {
	limiter := NewRateLimiter(60, 1)
	handler := limiter.Handler(noContentHandler())

	first := httptest.NewRequest(http.MethodGet, "/", nil)
	first.RemoteAddr = "192.0.2.1:1234"
	first.Header.Set("X-Forwarded-For", "198.51.100.1")
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, first)

	second := httptest.NewRequest(http.MethodGet, "/", nil)
	second.RemoteAddr = "192.0.2.1:5678"
	second.Header.Set("X-Forwarded-For", "198.51.100.1")
	recorder = httptest.NewRecorder()
	handler.ServeHTTP(recorder, second)
	if recorder.Code != http.StatusTooManyRequests {
		t.Fatalf("untrusted forwarded header request returned status %d, want %d", recorder.Code, http.StatusTooManyRequests)
	}
}

func TestRateLimiterSkipsSwaggerRoutes(t *testing.T) {
	limiter := NewRateLimiter(60, 1)
	handler := limiter.Handler(noContentHandler())

	for requestNumber := 1; requestNumber <= 2; requestNumber++ {
		request := httptest.NewRequest(http.MethodGet, "/swagger/index.html", nil)
		request.RemoteAddr = "192.0.2.1:1234"
		recorder := httptest.NewRecorder()

		handler.ServeHTTP(recorder, request)

		if recorder.Code != http.StatusNoContent {
			t.Fatalf("Swagger request %d returned status %d, want %d", requestNumber, recorder.Code, http.StatusNoContent)
		}
	}
}
