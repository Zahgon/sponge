package interceptor

import (
	"time"

	"google.golang.org/grpc"

	rl "github.com/go-dev-frame/sponge/pkg/shield/ratelimit"
)

// ---------------------------------- server interceptor ----------------------------------

// ErrLimitExceed is returned when the rate limiter is
// triggered and the request is rejected due to limit exceeded.
var ErrLimitExceed = rl.ErrLimitExceed

// RatelimitOption set the rate limits ratelimitOptions.
type RatelimitOption func(*ratelimitOptions)

type ratelimitOptions struct {
	window       time.Duration
	bucket       int
	cpuThreshold int64
	cpuQuota     float64
}

func defaultRatelimitOptions() *ratelimitOptions { _ = "STUB: not implemented"; return nil }

func (o *ratelimitOptions) apply(opts ...RatelimitOption) { _ = "STUB: not implemented"; return }

// WithWindow with window size.
func WithWindow(d time.Duration) RatelimitOption {
	_ = "STUB: not implemented"
	return *new(RatelimitOption)
}

// WithBucket with bucket size.
func WithBucket(b int) RatelimitOption { _ = "STUB: not implemented"; return *new(RatelimitOption) }

// WithCPUThreshold with cpu threshold
func WithCPUThreshold(threshold int64) RatelimitOption {
	_ = "STUB: not implemented"
	return *new(RatelimitOption)
}

// WithCPUQuota with real cpu quota(if it can not collect from process correct);
func WithCPUQuota(quota float64) RatelimitOption {
	_ = "STUB: not implemented"
	return *new(RatelimitOption)
}

// UnaryServerRateLimit server-side unary circuit breaker interceptor
func UnaryServerRateLimit(opts ...RatelimitOption) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

// StreamServerRateLimit server-side stream circuit breaker interceptor
func StreamServerRateLimit(opts ...RatelimitOption) grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}
