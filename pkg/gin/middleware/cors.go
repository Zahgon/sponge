package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type CoresConfig = cors.Config

// CoresOption set coresOptions.
type CoresOption func(*coresOptions)

type coresOptions struct {
	newCoresConfig *CoresConfig // if nil, use default config under fields.

	allowOrigins     []string
	allowMethods     []string
	allowHeaders     []string
	exposeHeaders    []string
	maxAge           time.Duration
	allowWildcard    bool
	allowCredentials bool
}

func (o *coresOptions) apply(opts ...CoresOption) { _ = "STUB: not implemented"; return }

func defaultCoreOptions() *coresOptions { _ = "STUB: not implemented"; return nil }

// WithNewConfig set cors config, if nil, use default config under fields.
func WithNewConfig(config *CoresConfig) CoresOption {
	_ = "STUB: not implemented"
	return *new(CoresOption)
}

// WithAllowOrigins set allowOrigins, e.g. "https://yourdomain.com", "https://*.subdomain.com"
func WithAllowOrigins(allowOrigins ...string) CoresOption {
	_ = "STUB: not implemented"
	return *new(CoresOption)
}

// WithAllowMethods set allowMethods, e.g. "GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"
func WithAllowMethods(allowMethods ...string) CoresOption {
	_ = "STUB: not implemented"
	return *new(CoresOption)
}

// WithAllowHeaders set allowHeaders, e.g. "Origin", "Authorization", "Content-Type", "Accept"
func WithAllowHeaders(allowHeaders ...string) CoresOption {
	_ = "STUB: not implemented"
	return *new(CoresOption)
}

// WithExposeHeaders set exposeHeaders
func WithExposeHeaders(exposeHeaders ...string) CoresOption {
	_ = "STUB: not implemented"
	return *new(CoresOption)
}

// WithMaxAge set maxAge
func WithMaxAge(maxAge time.Duration) CoresOption {
	_ = "STUB: not implemented"
	return *new(CoresOption)
}

// WithAllowCredentials set allowCredentials
func WithAllowCredentials(allowCredentials bool) CoresOption {
	_ = "STUB: not implemented"
	return *new(CoresOption)
}

// WithAllowWildcard set allowWildcard
func WithAllowWildcard(allowWildcard bool) CoresOption {
	_ = "STUB: not implemented"
	return *new(CoresOption)
}

// Cors cross domain
func Cors(opts ...CoresOption) gin.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(gin.HandlerFunc)
}
