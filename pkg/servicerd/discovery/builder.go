package discovery

import (
	"time"

	"google.golang.org/grpc/resolver"

	"github.com/go-dev-frame/sponge/pkg/servicerd/registry"
)

const name = "discovery"

// Option is builder option.
type Option func(o *builder)

// WithTimeout with timeout option.
func WithTimeout(timeout time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithInsecure with isSecure option.
func WithInsecure(insecure bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// DisableDebugLog disables update instances log.
func DisableDebugLog() Option { _ = "STUB: not implemented"; return *new(Option) }

type builder struct {
	discoverer       registry.Discovery
	timeout          time.Duration
	insecure         bool
	debugLogDisabled bool
}

// NewBuilder creates a builder which is used to factory registry resolvers.
func NewBuilder(d registry.Discovery, opts ...Option) resolver.Builder {
	_ = "STUB: not implemented"
	return *new(resolver.Builder)
}

func (b *builder) Build(target resolver.Target, cc resolver.ClientConn, _ resolver.BuildOptions) (resolver.Resolver, error) {
	_ = "STUB: not implemented"
	return *new(resolver.Resolver), nil
}

// Scheme return scheme of discovery
func (*builder) Scheme() string { _ = "STUB: not implemented"; return "" }
