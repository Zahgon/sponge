// Package routers is a package dedicated to registering routes, and supports both
// manual route registration and automatic route registration.
package routers

import (
	"github.com/gin-gonic/gin"
)

var (
	apiV1RouterFns []func(r *gin.RouterGroup) // group router functions
	// if you have other group routes you can define them here
	// example:
	//     apiV2RouterFns []func(r *gin.RouterGroup)
)

// NewRouter create a new router
func NewRouter() *gin.Engine { _ = "STUB: not implemented"; return nil }

// if you need more fine-grained control over your routes, set the timeout in your routes, unsetting the timeout globally here.

// request id middleware

// logger middleware, to print simple messages, replace middleware.Logging with middleware.SimpleLog

// ignore path

// metrics middleware

//metrics.WithMetricsPath("/metrics"),                // default is /metrics
// ignore 404 status codes

// limit middleware

//middleware.WithWindow(time.Second*5), // default 10s
//middleware.WithBucket(1000), // default 100
//middleware.WithCPUThreshold(750), // default 800

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

// register swagger routes, generate code via swag init

// access path /swagger/index.html

// register routers, middleware support

// if you have other group routes you can add them here
// example:
//    registerRouters(r, "/api/v2", apiV2RouteFns, middleware.Auth())

func registerRouters(r *gin.Engine, groupPath string, routerFns []func(*gin.RouterGroup), handlers ...gin.HandlerFunc) {
	_ = "STUB: not implemented"
	return
}
