package registry

// Option service instance  options
type Option func(*options)

type options struct {
	version  string
	metadata map[string]string
}

func defaultOptions() *options { _ = "STUB: not implemented"; return nil }

func (o *options) apply(opts ...Option) { _ = "STUB: not implemented"; return }

// WithVersion set server version
func WithVersion(version string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMetadata set metadata
func WithMetadata(metadata map[string]string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
