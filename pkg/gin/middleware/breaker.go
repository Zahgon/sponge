package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/go-dev-frame/sponge/pkg/container/group"
	"github.com/go-dev-frame/sponge/pkg/shield/circuitbreaker"
)

// ErrNotAllowed error not allowed.
var ErrNotAllowed = circuitbreaker.ErrNotAllowed

// CircuitBreakerOption set the circuit breaker circuitBreakerOptions.
type CircuitBreakerOption func(*circuitBreakerOptions)

type circuitBreakerOptions struct {
	group *group.Group
	// http code for circuit breaker, default already includes 500 and 503
	validCodes map[int]struct{}
	// degrade func
	degradeHandler func(c *gin.Context)
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

// WithValidCode http code to mark failed
func WithValidCode(code ...int) CircuitBreakerOption {
	_ = "STUB: not implemented"
	return *new(CircuitBreakerOption)
}

// WithDegradeHandler set degrade handler function
func WithDegradeHandler(handler func(c *gin.Context)) CircuitBreakerOption {
	_ = "STUB: not implemented"
	return *new(CircuitBreakerOption)
}

// CircuitBreaker a circuit breaker middleware
func CircuitBreaker(opts ...CircuitBreakerOption) gin.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(gin.HandlerFunc)
}

// NOTE: when client reject request locally, keep adding counter let the drop ratio higher.

// NOTE: need to check internal and service unavailable error
