package metrics

// Option set the metrics options.
type Option func(*options)

type options struct {
	metricsPath          string
	ignoreStatusCodes    map[int]struct{}
	ignoreRequestPaths   map[string]struct{}
	ignoreRequestMethods map[string]struct{}
}

// defaultOptions default value
func defaultOptions() *options { _ = "STUB: not implemented"; return nil }

func (o *options) apply(opts ...Option) { _ = "STUB: not implemented"; return }

// WithMetricsPath set metrics path
func WithMetricsPath(metricsPath string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithIgnoreStatusCodes ignore status codes
func WithIgnoreStatusCodes(statusCodes ...int) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithIgnoreRequestPaths ignore request paths
func WithIgnoreRequestPaths(paths ...string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithIgnoreRequestMethods ignore request methods
func WithIgnoreRequestMethods(methods ...string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func (o *options) isIgnoreCodeStatus(statusCode int) bool { _ = "STUB: not implemented"; return false }

func (o *options) isIgnorePath(path string) bool { _ = "STUB: not implemented"; return false }

func (o *options) checkIgnoreMethod(method string) bool { _ = "STUB: not implemented"; return false }
