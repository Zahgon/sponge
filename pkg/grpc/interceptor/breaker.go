// Package interceptor provides commonly used grpc client-side and server-side interceptors.
package interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/go-dev-frame/sponge/pkg/container/group"
	"github.com/go-dev-frame/sponge/pkg/shield/circuitbreaker"
)

// ErrNotAllowed error not allowed.
var ErrNotAllowed = circuitbreaker.ErrNotAllowed

// CircuitBreakerOption set the circuit breaker circuitBreakerOptions.
type CircuitBreakerOption func(*circuitBreakerOptions)

type circuitBreakerOptions struct {
	group *group.Group
	// rpc code for circuit breaker, default already includes codes.Internal and codes.Unavailable
	validCodes map[codes.Code]struct{}

	// degrade handler for unary server
	unaryServerDegradeHandler func(ctx context.Context, req interface{}) (reply interface{}, err error)
}

func defaultCircuitBreakerOptions() *circuitBreakerOptions { _ = "STUB: not implemented"; return nil }

func (o *circuitBreakerOptions) apply(opts ...CircuitBreakerOption) {
	_ = "STUB: not implemented"
	return
}

// WithGroup with circuit breaker group.
// Deprecated: use WithBreakerOption instead
func WithGroup(g *group.Group) CircuitBreakerOption {
	_ = "STUB: not implemented"
	return *new(CircuitBreakerOption)
}

// WithBreakerOption set the circuit breaker options.
func WithBreakerOption(opts ...circuitbreaker.Option) CircuitBreakerOption {
	_ = "STUB: not implemented"
	return *new(CircuitBreakerOption)
}

// WithValidCode rpc code to mark failed
func WithValidCode(code ...codes.Code) CircuitBreakerOption {
	_ = "STUB: not implemented"
	return *new(CircuitBreakerOption)
}

// WithUnaryServerDegradeHandler unary server degrade handler function
func WithUnaryServerDegradeHandler(handler func(ctx context.Context, req interface{}) (reply interface{}, err error)) CircuitBreakerOption {
	_ = "STUB: not implemented"
	return *new(CircuitBreakerOption)
}

// UnaryClientCircuitBreaker client-side unary circuit breaker interceptor
func UnaryClientCircuitBreaker(opts ...CircuitBreakerOption) grpc.UnaryClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryClientInterceptor)
}

// NOTE: when client reject request locally, keep adding counter let the drop ratio higher.

// NOTE: need to check internal and service unavailable error

// StreamClientCircuitBreaker client-side stream circuit breaker interceptor
func StreamClientCircuitBreaker(opts ...CircuitBreakerOption) grpc.StreamClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamClientInterceptor)
}

// NOTE: when client reject request locally, keep adding counter let the drop ratio higher.

// NOTE: need to check internal and service unavailable error

// UnaryServerCircuitBreaker server-side unary circuit breaker interceptor
func UnaryServerCircuitBreaker(opts ...CircuitBreakerOption) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

// NOTE: when client reject request locally, keep adding let the drop ratio higher.

// NOTE: need to check internal and service unavailable error

// StreamServerCircuitBreaker server-side stream circuit breaker interceptor
func StreamServerCircuitBreaker(opts ...CircuitBreakerOption) grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}

// NOTE: when client reject request locally, keep adding counter let the drop ratio higher.

// NOTE: need to check internal and service unavailable error
