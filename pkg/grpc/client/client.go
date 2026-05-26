// Package client is generic grpc client-side.
package client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/resolver"

	"github.com/go-dev-frame/sponge/pkg/servicerd/registry"
)

// Option client option func
type Option func(*options)

type options struct {
	builders   []resolver.Builder
	iDiscovery registry.Discovery
	isInsecure bool

	isLoadBalance      bool
	credentials        credentials.TransportCredentials
	unaryInterceptors  []grpc.UnaryClientInterceptor
	streamInterceptors []grpc.StreamClientInterceptor
	dialOptions        []grpc.DialOption
}

func defaultOptions() *options { _ = "STUB: not implemented"; return nil }

func (o *options) apply(opts ...Option) { _ = "STUB: not implemented"; return }

// WithServiceDiscover set service discover
func WithServiceDiscover(d registry.Discovery, isInsecure bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithServiceDiscoverBuilder set service discover builder
func WithServiceDiscoverBuilder(builder ...resolver.Builder) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithLoadBalance set load balance
func WithLoadBalance() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSecure set secure
func WithSecure(credential credentials.TransportCredentials) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithUnaryInterceptor set unary interceptor
func WithUnaryInterceptor(interceptors ...grpc.UnaryClientInterceptor) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithStreamInterceptor set stream interceptor
func WithStreamInterceptor(interceptors ...grpc.StreamClientInterceptor) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithDialOption set DialOption
func WithDialOption(dialOptions ...grpc.DialOption) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// NewClient create a new grpc client
func NewClient(endpoint string, opts ...Option) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// service discovery

// higher priority

// load balance option

// secure option

// custom dial option

// custom unary interceptor option

// custom stream interceptor option

// Dial to grpc server
func Dial(_ context.Context, endpoint string, opts ...Option) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
