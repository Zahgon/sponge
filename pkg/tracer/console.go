package tracer

import (
	"io"
	"os"

	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
)

// NewConsoleExporter output to console
func NewConsoleExporter() (sdkTrace.SpanExporter, error) {
	_ = "STUB: not implemented"
	return *new(sdkTrace.SpanExporter), nil
}

// NewFileExporter output to file, note: close the file before ending
func NewFileExporter(filename string) (sdkTrace.SpanExporter, *os.File, error) {
	_ = "STUB: not implemented"
	return *new(sdkTrace.SpanExporter), nil, nil
}

// Write telemetry data to a file.

// newExporter returns a console exporter.
func newExporter(w io.Writer) (sdkTrace.SpanExporter, error) {
	_ = "STUB: not implemented"
	return *new(sdkTrace.SpanExporter), nil
}

// output to console.

// do not print timestamps for the demo.
