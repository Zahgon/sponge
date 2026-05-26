package interceptor

import (
	"google.golang.org/grpc"
)

// ---------------------------------- client interceptor ----------------------------------

// UnaryClientRecovery client-side unary recovery
func UnaryClientRecovery() grpc.UnaryClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryClientInterceptor)
}

// StreamClientRecovery client-side recovery stream interceptor
func StreamClientRecovery() grpc.StreamClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamClientInterceptor)
}

// ---------------------------------- server interceptor ----------------------------------

// UnaryServerRecovery recovery unary interceptor
func UnaryServerRecovery() grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

// StreamServerRecovery recovery stream interceptor
func StreamServerRecovery() grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}
