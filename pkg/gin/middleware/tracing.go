package middleware

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/propagation"
	oteltrace "go.opentelemetry.io/otel/trace"
)

const (
	tracerKey  = "otel-tracer"
	tracerName = "otelgin"
)

type traceConfig struct {
	TracerProvider oteltrace.TracerProvider
	Propagators    propagation.TextMapPropagator
}

// TraceOption specifies instrumentation configuration options.
type TraceOption func(*traceConfig)

// WithPropagators specifies propagators to use for extracting
// information from the HTTP requests. If none are specified, global
// ones will be used.
func WithPropagators(propagators propagation.TextMapPropagator) TraceOption {
	_ = "STUB: not implemented"
	return *new(TraceOption)
}

// WithTracerProvider specifies a tracer provider to use for creating a tracer.
// If none is specified, the global provider is used.
func WithTracerProvider(provider oteltrace.TracerProvider) TraceOption {
	_ = "STUB: not implemented"
	return *new(TraceOption)
}

// Tracing returns interceptor that will trace incoming requests.
// The service parameter should describe the name of the (virtual)
// server handling the request.
func Tracing(serviceName string, opts ...TraceOption) gin.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(gin.HandlerFunc)
}

// pass the span through the request context

// serve the request to the next interceptor
