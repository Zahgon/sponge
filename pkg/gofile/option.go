package gofile

const (
	prefix  = "prefix"
	suffix  = "suffix"
	contain = "contain"
)

var (
	defaultFilterType = "" // with prefix, suffix, contain, no filter by default
)

type options struct {
	filter string
	name   string

	noAbsolutePath bool
}

func defaultOptions() *options { _ = "STUB: not implemented"; return nil }

// Option set the file options.
type Option func(*options)

func (o *options) apply(opts ...Option) { _ = "STUB: not implemented"; return }

// WithSuffix set suffix matching
func WithSuffix(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPrefix set prefix matching
func WithPrefix(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithContain set contain matching
func WithContain(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithNoAbsolutePath set no absolute path
func WithNoAbsolutePath() Option { _ = "STUB: not implemented"; return *new(Option) }
