package proxy

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Option set options.
type Option func(*options)

type options struct {
	managerPrefixPath  string // default "/endpoints"
	managerMiddlewares []gin.HandlerFunc
	zapLogger          *zap.Logger
}

func (o *options) apply(opts ...Option) { _ = "STUB: not implemented"; return }

func defaultOptions() *options { _ = "STUB: not implemented"; return nil }

// WithManagerEndpoints sets manager prefix path and middlewares, managerPrefixPath default "/endpoints".
func WithManagerEndpoints(managerPrefixPath string, middlewares ...gin.HandlerFunc) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithLogger sets logger.
func WithLogger(logger *zap.Logger) Option { _ = "STUB: not implemented"; return *new(Option) }

// -------------------------------------------------------------------------------------------

var (
	BalancerRoundRobin = "round_robin"
	BalancerLeastConn  = "least_conn"
	BalancerIPHash     = "ip_hash"
)

// PassOption set passOptions.
type PassOption func(*passOptions)

type passOptions struct {
	healthCheckInterval time.Duration // default 5s
	healthCheckTimeout  time.Duration // default 3s
	balancerType        string        // supported values: "round_robin", "least_conn", "ip_hash", default "round_robin"
	passMiddlewares     []gin.HandlerFunc
}

func (o *passOptions) apply(opts ...PassOption) { _ = "STUB: not implemented"; return }

func defaultPassOptions() *passOptions { _ = "STUB: not implemented"; return nil }

// WithPassBalancer sets balancer type.
func WithPassBalancer(balancerType string) PassOption {
	_ = "STUB: not implemented"
	return *new(PassOption)
}

// WithPassHealthCheck sets health check interval and timeout.
func WithPassHealthCheck(interval time.Duration, timeout time.Duration) PassOption {
	_ = "STUB: not implemented"
	return *new(PassOption)
}

// WithPassMiddlewares sets proxy middlewares.
func WithPassMiddlewares(middlewares ...gin.HandlerFunc) PassOption {
	_ = "STUB: not implemented"
	return *new(PassOption)
}
