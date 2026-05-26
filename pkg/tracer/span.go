package tracer

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

var traceName = "unknown"

// SetTraceName each service corresponds to a traceName
func SetTraceName(name string) { _ = "STUB: not implemented"; return }

// NewSpan create a span, to end a span you must call span.End()
func NewSpan(ctx context.Context, spanName string, tags map[string]interface{}) (context.Context, trace.Span) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(trace.Span)
}

//nolint
