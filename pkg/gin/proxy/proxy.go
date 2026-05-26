package proxy

import (
	"github.com/gin-gonic/gin"

	"github.com/go-dev-frame/sponge/pkg/proxykit"
)

// Proxy is a proxy server.
type Proxy struct {
	r       *gin.Engine
	manager *proxykit.RouteManager
}

// New creates a new Proxy instance.
func New(r *gin.Engine, opts ...Option) *Proxy { _ = "STUB: not implemented"; return nil }

// setup manager endpoints routes

// Pass registers proxy endpoints to gin engine.
func (p *Proxy) Pass(prefixPath string, endpoints []string, opts ...PassOption) error {
	_ = "STUB: not implemented"
	return nil
}

// setup proxy endpoints routes
// /prefixPath/*path
