package tracer

import (
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
)

// JaegerOption set fields
type JaegerOption func(*jaegerOptions)

type jaegerOptions struct {
	username string
	password string
}

func (o *jaegerOptions) apply(opts ...JaegerOption) { _ = "STUB: not implemented"; return }

// default setting
func defaultJaegerOptions() *jaegerOptions { _ = "STUB: not implemented"; return nil }

// WithUsername set username
func WithUsername(username string) JaegerOption {
	_ = "STUB: not implemented"
	return *new(JaegerOption)
}

// WithPassword set password
func WithPassword(password string) JaegerOption {
	_ = "STUB: not implemented"
	return *new(JaegerOption)
}

// NewJaegerExporter use jaeger collector as exporter, e.g. default url=http://localhost:14268/api/traces
func NewJaegerExporter(url string, opts ...JaegerOption) (sdkTrace.SpanExporter, error) {
	_ = "STUB: not implemented"
	return *new(sdkTrace.SpanExporter), nil
}

// NewJaegerAgentExporter use jaeger agent as exporter, e.g. host=localhost port=6831
func NewJaegerAgentExporter(host string, port string) (sdkTrace.SpanExporter, error) {
	_ = "STUB: not implemented"
	return *new(sdkTrace.SpanExporter), nil
}
