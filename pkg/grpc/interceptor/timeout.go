package interceptor

import (
	"time"

	"google.golang.org/grpc"
)

// ---------------------------------- client interceptor ----------------------------------

var timeoutVal = time.Second * 10

// UnaryClientTimeout client-side timeout unary interceptor
func UnaryClientTimeout(d time.Duration) grpc.UnaryClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryClientInterceptor)
}

//nolint

// StreamClientTimeout server-side timeout  interceptor
func StreamClientTimeout(d time.Duration) grpc.StreamClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamClientInterceptor)
}

//nolint
