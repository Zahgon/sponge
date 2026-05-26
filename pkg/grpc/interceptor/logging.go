package interceptor

import (
	"encoding/json"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/go-dev-frame/sponge/pkg/errcode"
)

var contentMark = []byte(" ...... ")

// ---------------------------------- client interceptor ----------------------------------

// UnaryClientLog client log unary interceptor
func UnaryClientLog(logger *zap.Logger, opts ...LogOption) grpc.UnaryClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryClientInterceptor)
}

// ignore printing of the specified method

// StreamClientLog client log stream interceptor
func StreamClientLog(logger *zap.Logger, opts ...LogOption) grpc.StreamClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamClientInterceptor)
}

// ignore printing of the specified method

// ---------------------------------- server interceptor ----------------------------------

var defaultMaxLength = 300 // max length of response data to print
var defaultMarshalFn = func(reply interface{}) []byte {
	data, _ := json.Marshal(reply)
	return data
}
var ignoreLogMethods = map[string]bool{ // ignore printing methods
	"/grpc.health.v1.Health/Check": true,
}
var printErrorBySpecifiedCodes = map[codes.Code]bool{
	codes.Internal:                           true,
	codes.Unavailable:                        true,
	errcode.StatusInternalServerError.Code(): true,
	errcode.StatusServiceUnavailable.Code():  true,
}

// LogOption log settings
type LogOption func(*logOptions)

type logOptions struct {
	maxLength           int
	isReplaceGRPCLogger bool
	marshalFn           func(reply interface{}) []byte // default json.Marshal
}

func defaultLogOptions() *logOptions { _ = "STUB: not implemented"; return nil }

func (o *logOptions) apply(opts ...LogOption) { _ = "STUB: not implemented"; return }

// WithMaxLen logger content max length
func WithMaxLen(maxLen int) LogOption { _ = "STUB: not implemented"; return *new(LogOption) }

// WithReplaceGRPCLogger replace grpc logger v2
func WithReplaceGRPCLogger() LogOption { _ = "STUB: not implemented"; return *new(LogOption) }

// WithPrintErrorByCodes set print error by grpc codes
func WithPrintErrorByCodes(code ...codes.Code) LogOption {
	_ = "STUB: not implemented"
	return *new(LogOption)
}

// WithMarshalFn custom response data marshal function
func WithMarshalFn(fn func(reply interface{}) []byte) LogOption {
	_ = "STUB: not implemented"
	return *new(LogOption)
}

// WithLogIgnoreMethods ignore printing methods
// fullMethodName format: /packageName.serviceName/methodName,
// example /api.userExample.v1.userExampleService/GetByID
func WithLogIgnoreMethods(fullMethodNames ...string) LogOption {
	_ = "STUB: not implemented"
	return *new(LogOption)
}

// UnaryServerLog server-side log unary interceptor
func UnaryServerLog(logger *zap.Logger, opts ...LogOption) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

// ignore printing of the specified method

// UnaryServerSimpleLog server-side log unary interceptor, only print response
func UnaryServerSimpleLog(logger *zap.Logger, opts ...LogOption) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

// ignore printing of the specified method

// StreamServerLog Server-side log stream interceptor
func StreamServerLog(logger *zap.Logger, opts ...LogOption) grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}

// ignore printing of the specified method

// StreamServerSimpleLog Server-side log stream interceptor, only print response
func StreamServerSimpleLog(logger *zap.Logger, opts ...LogOption) grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}

// ignore printing of the specified method
