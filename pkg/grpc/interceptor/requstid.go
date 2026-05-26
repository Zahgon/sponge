package interceptor

import (
	"context"
	"sync"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	// ContextRequestIDKey request id key for context
	ContextRequestIDKey = "request_id"
	once                sync.Once
)

// SetContextRequestIDKey set context request id key
func SetContextRequestIDKey(key string) { _ = "STUB: not implemented"; return }

// CtxKeyString for context.WithValue key type
type CtxKeyString string

// RequestIDKey request_id
var RequestIDKey = CtxKeyString(ContextRequestIDKey)

// ---------------------------------- client interceptor ----------------------------------

// CtxRequestIDField get request id field from context.Context
func CtxRequestIDField(ctx context.Context) zap.Field {
	_ = "STUB: not implemented"
	return *new(zap.Field)
}

// ClientCtxRequestID get request id from rpc client context.Context
func ClientCtxRequestID(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

// ClientCtxRequestIDField get request id field from rpc client context.Context
func ClientCtxRequestIDField(ctx context.Context) zap.Field {
	_ = "STUB: not implemented"
	return *new(zap.Field)
}

// UnaryClientRequestID client-side request_id unary interceptor
func UnaryClientRequestID() grpc.UnaryClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryClientInterceptor)
}

// StreamClientRequestID client request id stream interceptor
func StreamClientRequestID() grpc.StreamClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamClientInterceptor)
}

// ---------------------------------- server interceptor ----------------------------------

// KV key value
type KV struct {
	Key string
	Val interface{}
}

// WrapServerCtx wrap context, used in grpc server-side
func WrapServerCtx(ctx context.Context, kvs ...KV) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

//nolint

//nolint

// ServerCtxRequestID get request id from rpc server context.Context
func ServerCtxRequestID(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

// ServerCtxRequestIDField get request id field from rpc server context.Context
func ServerCtxRequestIDField(ctx context.Context) zap.Field {
	_ = "STUB: not implemented"
	return *new(zap.Field)
}

// UnaryServerRequestID server-side request_id unary interceptor
func UnaryServerRequestID() grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

// StreamServerRequestID server-side request id stream interceptor
func StreamServerRequestID() grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	// todo
	return *new(grpc.StreamServerInterceptor)
}

//ctx := stream.Context()
//requestID := ServerCtxRequestID(ctx)
//if requestID == "" {
//	requestID = krand.String(krand.R_All, 10)
//	ctx = grpc_metadata.ExtractIncoming(ctx).Add(ContextRequestIDKey, requestID).ToIncoming(ctx)
//}
