package tracer

import (
	"go.opentelemetry.io/otel/sdk/resource"
)

// alias, for other structs, the following code does not need to change the names of the resourceOptions
type resourceOptions = resourceConfig

// ResourceOption modifying struct field values by means of an interface
type ResourceOption interface {
	apply(*resourceOptions)
}

type resourceOptionFunc func(*resourceOptions)

func (o resourceOptionFunc) apply(cfg *resourceOptions) {
	_ = "STUB: not implemented"

	// set obj fields value
	return
}

func apply(obj *resourceOptions, opts ...ResourceOption) { _ = "STUB: not implemented"; return }

// WithServiceName set service name
func WithServiceName(name string) ResourceOption {
	_ = "STUB: not implemented"
	return *new(ResourceOption)
}

// WithServiceVersion set service version
func WithServiceVersion(version string) ResourceOption {
	_ = "STUB: not implemented"
	return *new(ResourceOption)
}

// WithEnvironment set service environment
func WithEnvironment(environment string) ResourceOption {
	_ = "STUB: not implemented"
	return *new(ResourceOption)
}

// WithAttributes set service attributes
func WithAttributes(attributes map[string]string) ResourceOption {
	_ = "STUB: not implemented"
	return *new(ResourceOption)
}

type resourceConfig struct {
	serviceName    string
	serviceVersion string
	environment    string

	attributes map[string]string
}

// NewResource returns a resource describing this application.
func NewResource(opts ...ResourceOption) *resource.Resource {
	_ = "STUB: not implemented"
	// default values
	return nil
}
