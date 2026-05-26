// Package tracer is a library wrapped in go.opentelemetry.io/otel.
package tracer

import (
	"context"

	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
)

var tp *trace.TracerProvider

// Init Initialize tracer, parameter fraction is fraction, default is 1.0, value >= 1.0 means all links are sampled,
// value <= 0 means all are not sampled, 0 < value < 1 only samples percentage
func Init(exporter trace.SpanExporter, res *resource.Resource, fractions ...float64) {
	_ = "STUB: not implemented"
	return
}

// sampling rate

// register the TracerProvider as global so that any future imports of package go.opentelemetry.io/otel/trace will use it by default.

// propagation of context across processes

// Close tracer
func Close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// InitWithConfig Initialize tracer according to configuration, fraction is fraction, default is 1.0, value >= 1.0 means all links are sampled,
// value <= 0 means all are not sampled, 0 < value < 1 only samples percentage
func InitWithConfig(appName string, appEnv string, appVersion string,
	jaegerAgentHost string, jaegerAgentPort string, jaegerSamplingRate float64) {
	_ = "STUB: not implemented"
	return
}

// initializing tracing

// GetProvider get tracer provider
func GetProvider() *trace.TracerProvider { _ = "STUB: not implemented"; return nil }
