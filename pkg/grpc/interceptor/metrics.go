package interceptor

import (
	"google.golang.org/grpc"

	"github.com/go-dev-frame/sponge/pkg/grpc/metrics"
)

// UnaryClientMetrics client-side metrics unary interceptor
func UnaryClientMetrics() grpc.UnaryClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryClientInterceptor)
}

// StreamClientMetrics client-side metrics stream interceptor
func StreamClientMetrics() grpc.StreamClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamClientInterceptor)
}

// UnaryServerMetrics server-side metrics unary interceptor
func UnaryServerMetrics(opts ...metrics.Option) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

// StreamServerMetrics server-side metrics stream interceptor
func StreamServerMetrics(opts ...metrics.Option) grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}
