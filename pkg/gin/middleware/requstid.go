package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	// ContextRequestIDKey request id for context
	ContextRequestIDKey = "request_id"

	// HeaderXRequestIDKey header request id key
	HeaderXRequestIDKey = "X-Request-Id"
)

// RequestIDOption set the request id  options.
type RequestIDOption func(*requestIDOptions)

type requestIDOptions struct {
	contextRequestIDKey string
	headerXRequestIDKey string
}

func defaultRequestIDOptions() *requestIDOptions { _ = "STUB: not implemented"; return nil }

func (o *requestIDOptions) apply(opts ...RequestIDOption) { _ = "STUB: not implemented"; return }

func (o *requestIDOptions) setRequestIDKey() { _ = "STUB: not implemented"; return }

// WithContextRequestIDKey set context request id key, minimum length of 4
func WithContextRequestIDKey(key string) RequestIDOption {
	_ = "STUB: not implemented"
	return *new(RequestIDOption)
}

// WithHeaderRequestIDKey set header request id key, minimum length of 4
func WithHeaderRequestIDKey(key string) RequestIDOption {
	_ = "STUB: not implemented"
	return *new(RequestIDOption)
}

// CtxKeyString for context.WithValue key type
type CtxKeyString string

// RequestIDKey request_id
var RequestIDKey = CtxKeyString(ContextRequestIDKey)

// -------------------------------------------------------------------------------------------

// RequestID is an interceptor that injects a 'request id' into the context and request/response header of each request.
func RequestID(opts ...RequestIDOption) gin.HandlerFunc {
	_ = "STUB: not implemented"
	// customized request id key
	return *new(gin.HandlerFunc)
}

// Check for incoming header, use it if exists

// Create request id

// Expose it for use in the application

// Set X-Request-Id header

// GCtxRequestID get request id from gin.Context
func GCtxRequestID(c *gin.Context) string { _ = "STUB: not implemented"; return "" }

// GCtxRequestIDField get request id field from gin.Context
func GCtxRequestIDField(c *gin.Context) zap.Field {
	_ = "STUB: not implemented"
	return *new(zap.Field)
}

// HeaderRequestID get request id from the header
func HeaderRequestID(c *gin.Context) string { _ = "STUB: not implemented"; return "" }

// HeaderRequestIDField get request id field from header
func HeaderRequestIDField(c *gin.Context) zap.Field {
	_ = "STUB: not implemented"
	return *new(zap.Field)
}

// -------------------------------------------------------------------------------------------

// RequestHeaderKey request header key
var RequestHeaderKey = "request_header_key"

// WrapCtx wrap context, put the Keys and Header of gin.Context into context
func WrapCtx(c *gin.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

//nolint
//nolint

// AdaptCtx adapt context, if ctx is gin.Context, return gin.Context and context of the transformation
func AdaptCtx(ctx context.Context) (*gin.Context, context.Context) {
	_ = "STUB: not implemented"
	return nil, *new(context.Context)
}

// GetFromCtx get value from context
func GetFromCtx(ctx context.Context, key string) interface{} { _ = "STUB: not implemented"; return nil }

// CtxRequestID get request id from context.Context
func CtxRequestID(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

// CtxRequestIDField get request id field from context.Context
func CtxRequestIDField(ctx context.Context) zap.Field {
	_ = "STUB: not implemented"
	return *new(zap.Field)
}

// GetFromHeader get value from header
func GetFromHeader(ctx context.Context, key string) string { _ = "STUB: not implemented"; return "" }

// GetFromHeaders get values from header
func GetFromHeaders(ctx context.Context, key string) []string {
	_ = "STUB: not implemented"
	return nil
}
