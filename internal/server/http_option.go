package server

import (
	"github.com/go-dev-frame/sponge/pkg/servicerd/registry"

	"github.com/go-dev-frame/sponge/internal/config"
)

// HTTPOption setting up http
type HTTPOption func(*httpOptions)

type httpOptions struct {
	isProd    bool
	instance  *registry.ServiceInstance
	iRegistry registry.Registry
	tls       config.TLS
}

func defaultHTTPOptions() *httpOptions { _ = "STUB: not implemented"; return nil }

func (o *httpOptions) apply(opts ...HTTPOption) { _ = "STUB: not implemented"; return }

// WithHTTPIsProd setting up production environment markers
func WithHTTPIsProd(isProd bool) HTTPOption { _ = "STUB: not implemented"; return *new(HTTPOption) }

// WithHTTPRegistry registration services
func WithHTTPRegistry(iRegistry registry.Registry, instance *registry.ServiceInstance) HTTPOption {
	_ = "STUB: not implemented"
	return *new(HTTPOption)
}

// WithHTTPTLS setting up tls
func WithHTTPTLS(tls config.TLS) HTTPOption { _ = "STUB: not implemented"; return *new(HTTPOption) }
