package middleware

import (
	"time"

	"github.com/gin-gonic/gin"

	rl "github.com/go-dev-frame/sponge/pkg/shield/ratelimit"
)

// ErrLimitExceed is returned when the rate limiter is
// triggered and the request is rejected due to limit exceeded.
var ErrLimitExceed = rl.ErrLimitExceed

// RateLimitOption set the rate limits rateLimitOptions.
type RateLimitOption func(*rateLimitOptions)

type rateLimitOptions struct {
	window       time.Duration
	bucket       int
	cpuThreshold int64
	cpuQuota     float64
}

func defaultRatelimitOptions() *rateLimitOptions { _ = "STUB: not implemented"; return nil }

func (o *rateLimitOptions) apply(opts ...RateLimitOption) { _ = "STUB: not implemented"; return }

// WithWindow with window size.
func WithWindow(d time.Duration) RateLimitOption {
	_ = "STUB: not implemented"
	return *new(RateLimitOption)
}

// WithBucket with bucket size.
func WithBucket(b int) RateLimitOption { _ = "STUB: not implemented"; return *new(RateLimitOption) }

// WithCPUThreshold with cpu threshold
func WithCPUThreshold(threshold int64) RateLimitOption {
	_ = "STUB: not implemented"
	return *new(RateLimitOption)
}

// WithCPUQuota with real cpu quota(if it can not collect from process correct);
func WithCPUQuota(quota float64) RateLimitOption {
	_ = "STUB: not implemented"
	return *new(RateLimitOption)
}

// RateLimit an adaptive rate limiter middleware
func RateLimit(opts ...RateLimitOption) gin.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(gin.HandlerFunc)
}

// Timeout request time out
func Timeout(d time.Duration) gin.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(gin.HandlerFunc)
}

//nolint
