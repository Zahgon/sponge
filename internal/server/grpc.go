// Package server is a package that holds the http or grpc service.
package server

import (
	"net"
	"net/http"

	"google.golang.org/grpc"

	"github.com/go-dev-frame/sponge/pkg/app"
	"github.com/go-dev-frame/sponge/pkg/servicerd/registry"
)

var _ app.IServer = (*grpcServer)(nil)

var (
	defaultTokenAppID  = "grpc"
	defaultTokenAppKey = "mko09ijn"
)

type grpcServer struct {
	addr   string
	server *grpc.Server
	listen net.Listener

	mux                             *http.ServeMux
	httpServer                      *http.Server
	registerMetricsMuxAndMethodFunc func() error

	iRegistry registry.Registry
	instance  *registry.ServiceInstance
}

// Start grpc service
func (s *grpcServer) Start() error {
	_ = "STUB: not implemented"
	// registration Services
	return nil
}

//nolint

// if either pprof or metrics is enabled, the http service will be started

// block

// Stop grpc service
func (s *grpcServer) Stop() error { _ = "STUB: not implemented"; return nil }

//nolint

// String comment
func (s *grpcServer) String() string { _ = "STUB: not implemented"; return "" }

// secure option
func (s *grpcServer) secureServerOption() grpc.ServerOption {
	_ = "STUB: not implemented"
	return *new(grpc.ServerOption)
}

// server side certification

// both client and server side certification

// setting up unary server interceptors
func (s *grpcServer) unaryServerOptions() grpc.ServerOption {
	_ = "STUB: not implemented"
	return *new(grpc.ServerOption)
}

// logger interceptor, to print simple messages, replace interceptor.UnaryServerLog with interceptor.UnaryServerSimpleLog

// token interceptor

// todo the defaultTokenAppID and defaultTokenAppKey are usually retrieved from the cache or database

// jwt token interceptor
//unaryServerInterceptors = append(unaryServerInterceptors, interceptor.UnaryServerJwtAuth(
// // choose a verification method as needed
//interceptor.WithStandardVerify(standardVerifyFn), // standard verify (default), you can set standardVerifyFn to nil if you don't need it
//interceptor.WithCustomVerify(customVerifyFn), // custom verify
// // specify the grpc API to ignore token verification(full path)
//interceptor.WithAuthIgnoreMethods("/api.user.v1.User/Register", "/api.user.v1.User/Login"),
//))

// metrics interceptor

// limit interceptor

//interceptor.WithWindow(time.Second*5), // default 10s
//interceptor.WithBucket(200),           // default 100
//interceptor.WithCPUThreshold(900),     // default 800

// circuit breaker interceptor

//interceptor.WithBreakerOption(
//circuitbreaker.WithSuccess(75),           // default 60
//circuitbreaker.WithRequest(100),          // default 100
//circuitbreaker.WithBucket(10),            // default 10
//circuitbreaker.WithWindow(time.Second*3), // default 3s
//),
//interceptor.WithUnaryServerDegradeHandler(handler), // Add degradation processing
// Add timeout error codes can also trigger circuit breaking

// trace interceptor

// setting up stream server interceptors
func (s *grpcServer) streamServerOptions() grpc.ServerOption {
	_ = "STUB: not implemented"
	return *new(grpc.ServerOption)
}

//interceptor.StreamServerRequestID(),

// logger interceptor, to print simple messages, replace interceptor.StreamServerLog with interceptor.StreamServerSimpleLog

// token interceptor

// todo the defaultTokenAppID and defaultTokenAppKey are usually retrieved from the cache or database

// jwt token interceptor
//streamServerInterceptors = append(streamServerInterceptors, interceptor.StreamServerJwtAuth(
// // choose a verification method as needed
//interceptor.WithStandardVerify(standardVerifyFn), // standard verify (default), you can set standardVerifyFn to nil if you don't need it
//interceptor.WithCustomVerify(customVerifyFn), // custom verify
// // specify the grpc API to ignore token verification(full path)
//	interceptor.WithAuthIgnoreMethods("/api.user.v1.User/Register", "/api.user.v1.User/Login"),
//))

// metrics interceptor

// limit interceptor

// circuit breaker interceptor

// set rpc code for circuit breaker, default already includes codes.Internal and codes.Unavailable

// trace interceptor

func (s *grpcServer) setOptions() []grpc.ServerOption { _ = "STUB: not implemented"; return nil }

func (s *grpcServer) registerMetricsMuxAndMethod() func() error {
	_ = "STUB: not implemented"
	return nil
}

func (s *grpcServer) registerProfMux() { _ = "STUB: not implemented"; return }

func (s *grpcServer) addHTTPRouter() { _ = "STUB: not implemented"; return }

// error codes router

// config router

// NewGRPCServer creates a new grpc server
func NewGRPCServer(addr string, opts ...GrpcOption) app.IServer {
	_ = "STUB: not implemented"
	return *new(app.IServer)
}

// register for all services
