// Package grpccli is grpc client with support for service discovery, logging, load balancing, trace, metrics, retries, circuit breaker.
package grpccli

import (
	"context"

	"google.golang.org/grpc"
)

// NewClient creates a new grpc client
func NewClient(endpoint string, opts ...Option) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// service discovery

// load balance option

// secure option

// token option

// unary options

// stream options

// custom options

// Dial to grpc server
// Deprecated: use NewClient instead
func Dial(_ context.Context, endpoint string, opts ...Option) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func secureOption(o *options) (grpc.DialOption, error) {
	_ = "STUB: not implemented"
	return *new(grpc.DialOption), nil
}

// server side certification

// both client and server side certification

func unaryClientOptions(o *options) grpc.DialOption {
	_ = "STUB: not implemented"
	return *new(grpc.DialOption)
}

// request id

// logging

// metrics

// circuit breaker
//if o.enableCircuitBreaker {
//	unaryClientInterceptors = append(unaryClientInterceptors, interceptor.UnaryClientCircuitBreaker(
//	// set rpc code for circuit breaker, default already includes codes.Internal and codes.Unavailable
//	//interceptor.WithValidCode(codes.PermissionDenied),
//	))
//}

// retry

// trace

// custom unary interceptors

func streamClientOptions(o *options) grpc.DialOption {
	_ = "STUB: not implemented"
	return *new(grpc.DialOption)
}

// request id

// logging

// metrics

// circuit breaker
//if o.enableCircuitBreaker {
//	streamClientInterceptors = append(streamClientInterceptors, interceptor.StreamClientCircuitBreaker(
//	// set rpc code for circuit breaker, default already includes codes.Internal and codes.Unavailable
//	//interceptor.WithValidCode(codes.PermissionDenied),
//	))
//}

// retry

// trace

// custom stream interceptors
