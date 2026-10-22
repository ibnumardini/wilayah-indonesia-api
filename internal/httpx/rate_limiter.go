package httpx

import (
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

const (
	requestsPerMinute = 120
	burstSize         = 30
)

type clientLimiter struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

type RateLimiter struct {
	mu      sync.Mutex
	clients map[string]*clientLimiter
	rate    rate.Limit
	burst   int
	proxies []*net.IPNet
}

func NewRateLimiter(requestsPerMinute, burst int, trustedProxyCIDRs ...string) *RateLimiter {
	limiter := &RateLimiter{
		clients: make(map[string]*clientLimiter),
		rate:    rate.Limit(float64(requestsPerMinute) / time.Minute.Seconds()),
		burst:   burst,
	}
	for _, value := range trustedProxyCIDRs {
		if _, network, err := net.ParseCIDR(value); err == nil {
			limiter.proxies = append(limiter.proxies, network)
		}
	}
	return limiter
}

func DefaultRateLimiter() *RateLimiter {
	return NewRateLimiter(requestsPerMinute, burstSize, trustedProxyCIDRsFromEnv()...)
}

func (rl *RateLimiter) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/swagger") {
			next.ServeHTTP(w, r)
			return
		}

		client := rl.clientIP(r)
		if !rl.allow(client) {
			w.Header().Set("Retry-After", "1")
			ResponseError(w, Response{
				Status:  http.StatusTooManyRequests,
				Message: "rate limit exceeded",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (rl *RateLimiter) allow(client string) bool {
	now := time.Now()

	rl.mu.Lock()
	defer rl.mu.Unlock()

	entry, exists := rl.clients[client]
	if !exists {
		entry = &clientLimiter{
			limiter: rate.NewLimiter(rl.rate, rl.burst),
		}
		rl.clients[client] = entry
	}
	entry.lastSeen = now

	return entry.limiter.Allow()
}

func (rl *RateLimiter) clientIP(r *http.Request) string {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		host = r.RemoteAddr
	}
	remoteIP := net.ParseIP(host)
	if remoteIP == nil || !rl.isTrustedProxy(remoteIP) {
		return host
	}

	if forwardedIP := firstForwardedIP(r.Header.Get("X-Forwarded-For")); forwardedIP != "" {
		return forwardedIP
	}
	if realIP := net.ParseIP(strings.TrimSpace(r.Header.Get("X-Real-IP"))); realIP != nil {
		return realIP.String()
	}
	return host
}

func (rl *RateLimiter) isTrustedProxy(ip net.IP) bool {
	for _, network := range rl.proxies {
		if network.Contains(ip) {
			return true
		}
	}
	return false
}

func firstForwardedIP(value string) string {
	for _, part := range strings.Split(value, ",") {
		if ip := net.ParseIP(strings.TrimSpace(part)); ip != nil {
			return ip.String()
		}
	}
	return ""
}

func trustedProxyCIDRsFromEnv() []string {
	value := os.Getenv("TRUSTED_PROXY_CIDRS")
	if value == "" {
		return nil
	}
	return strings.Split(value, ",")
}
