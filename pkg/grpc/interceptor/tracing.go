package interceptor

import (
	"google.golang.org/grpc"
)

// UnaryClientTracing client-side tracing unary interceptor
func UnaryClientTracing() grpc.UnaryClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryClientInterceptor)
}

//nolint

// StreamClientTracing client-side tracing stream interceptor
func StreamClientTracing() grpc.StreamClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamClientInterceptor)
}

//nolint

// UnaryServerTracing server-side tracing unary interceptor
func UnaryServerTracing() grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

//nolint

// StreamServerTracing server-side tracing stream interceptor
func StreamServerTracing() grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}

//nolint

// ClientOptionTracing client-side tracing interceptor
func ClientOptionTracing() grpc.DialOption { _ = "STUB: not implemented"; return *new(grpc.DialOption) }

// ServerOptionTracing server-side tracing interceptor
func ServerOptionTracing() grpc.ServerOption {
	_ = "STUB: not implemented"
	return *new(grpc.ServerOption)
}
