// Package server is generic grpc server-side.
package server

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/go-dev-frame/sponge/pkg/grpc/metrics"
)

// RegisterFn register object
type RegisterFn func(srv *grpc.Server)

// ServiceRegisterFn used to register service address to Consul/ETCD/Nacos/Zookeeper...
type ServiceRegisterFn func() error

// Option set server option
type Option func(*options)

type options struct {
	credentials        credentials.TransportCredentials
	unaryInterceptors  []grpc.UnaryServerInterceptor
	streamInterceptors []grpc.StreamServerInterceptor
	serviceRegisterFn  ServiceRegisterFn

	isShowConnections bool
	connectionOptions []metrics.ConnectionOption
}

func defaultServerOptions() *options { _ = "STUB: not implemented"; return nil }

func (o *options) apply(opts ...Option) { _ = "STUB: not implemented"; return }

// WithSecure set secure
func WithSecure(credential credentials.TransportCredentials) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithUnaryInterceptor set unary interceptor
func WithUnaryInterceptor(interceptors ...grpc.UnaryServerInterceptor) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithStreamInterceptor set stream interceptor
func WithStreamInterceptor(interceptors ...grpc.StreamServerInterceptor) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithServiceRegister set service register
func WithServiceRegister(fn ServiceRegisterFn) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithStatConnections enable stat connections
func WithStatConnections(opts ...metrics.ConnectionOption) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func customInterceptorOptions(o *options) []grpc.ServerOption {
	_ = "STUB: not implemented"
	return nil
}

// Run grpc server with options, registerFn is the function to register object to the server
func Run(port int, registerFn RegisterFn, options ...Option) (*grpc.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// listening on TCP port

// create a grpc server where interceptors can be injected

// register object to the server

// register service address to Consul/ETCD/Nacos/Zookeeper...

// run the server
