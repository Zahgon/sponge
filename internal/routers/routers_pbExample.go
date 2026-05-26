package routers

import (
	"github.com/gin-gonic/gin"
)

type routeFns = []func(r *gin.Engine, groupPathMiddlewares map[string][]gin.HandlerFunc, singlePathMiddlewares map[string][]gin.HandlerFunc)

var (
	// all route functions
	allRouteFns = make(routeFns, 0)
	// all middleware functions
	allMiddlewareFns = []func(c *middlewareConfig){}
)

// NewRouter_pbExample create a new router
func NewRouter_pbExample() *gin.Engine {
	_ = "STUB: not implemented" //nolint
	return nil
}

// if you need more fine-grained control over your routes, set the timeout in your routes, unsetting the timeout globally here.

// request id middleware

// logger middleware, to print simple messages, replace middleware.Logging with middleware.SimpleLog

// ignore path

// metrics middleware

//metrics.WithMetricsPath("/metrics"),                // default is /metrics
// ignore 404 status codes

// limit middleware

//middleware.WithWindow(time.Second*5), // default 10s
//middleware.WithBucket(200), // default 100
//middleware.WithCPUThreshold(900), // default 800

// circuit breaker middleware

//middleware.WithBreakerOption(
//circuitbreaker.WithSuccess(75),           // default 60
//circuitbreaker.WithRequest(100),          // default 100
//circuitbreaker.WithBucket(20),            // default 10
//circuitbreaker.WithWindow(time.Second*3), // default 3s
//),
//middleware.WithDegradeHandler(handler),              // Add degradation processing
// Add error codes to trigger circuit breaking

// trace middleware

// profile performance analysis

// access path /apis/swagger/index.html

// set up all middlewares

// register all routes

type middlewareConfig struct {
	groupPathMiddlewares  map[string][]gin.HandlerFunc // middleware functions corresponding to route group
	singlePathMiddlewares map[string][]gin.HandlerFunc // middleware functions corresponding to a single route
}

func newMiddlewareConfig() *middlewareConfig { _ = "STUB: not implemented"; return nil }

func (c *middlewareConfig) setGroupPath(groupPath string, handlers ...gin.HandlerFunc) {
	_ = "STUB: not implemented" //nolint
	return
}

func (c *middlewareConfig) setSinglePath(method string, singlePath string, handlers ...gin.HandlerFunc) {
	_ = "STUB: not implemented" //nolint
	return
}

func getSinglePathKey(method string, singlePath string) string {
	_ = "STUB: not implemented" //nolint
	return ""
}
